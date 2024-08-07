package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/alecthomas/kingpin/v2"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"

	"github.com/pipego/cli/config"
	"github.com/pipego/cli/dag"
	"github.com/pipego/cli/pipeline"
	"github.com/pipego/cli/runner"
	"github.com/pipego/cli/scheduler"
	_runner "github.com/pipego/dag/runner"
)

var (
	app           = kingpin.New("cli", "pipego cli").Version(config.Version + "-build-" + config.Build)
	configFile    = app.Flag("config-file", "Config file (.yml)").Required().String()
	runnerFile    = app.Flag("runner-file", "Runner file (.json)").Required().String()
	schedulerFile = app.Flag("scheduler-file", "Scheduler file (.json)").Required().String()
)

func Run(ctx context.Context) error {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	cfg, err := initConfig(ctx, *configFile)
	if err != nil {
		return errors.Wrap(err, "failed to init config")
	}

	d, err := initDag(ctx, cfg)
	if err != nil {
		return errors.Wrap(err, "failed to init dag")
	}

	t, g, m, c, err := initRunner(ctx, cfg, *runnerFile, d)
	if err != nil {
		return errors.Wrap(err, "failed to init runner")
	}

	s, err := initScheduler(ctx, cfg, *schedulerFile)
	if err != nil {
		return errors.Wrap(err, "failed to init scheduler")
	}

	p, err := initPipeline(ctx, cfg, t, s)
	if err != nil {
		return errors.Wrap(err, "failed to init pipeline")
	}

	if err := runPipeline(ctx, t, p); err != nil {
		return errors.Wrap(err, "failed to run pipeline")
	}

	if err := runGlance(ctx, g); err != nil {
		return errors.Wrap(err, "failed to run glance")
	}

	if err := runMaint(ctx, m); err != nil {
		return errors.Wrap(err, "failed to run maint")
	}

	if err := runConfig(ctx, c); err != nil {
		return errors.Wrap(err, "failed to run config")
	}

	return nil
}

func initConfig(_ context.Context, name string) (*config.Config, error) {
	c := config.New()

	fi, err := os.Open(name)
	if err != nil {
		return c, errors.Wrap(err, "failed to open")
	}

	defer func() {
		_ = fi.Close()
	}()

	buf, _ := io.ReadAll(fi)

	if err := yaml.Unmarshal(buf, c); err != nil {
		return c, errors.Wrap(err, "failed to unmarshal")
	}

	return c, nil
}

func loadFile(name string) ([]byte, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open file")
	}

	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	buf, err := io.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read file")
	}

	return buf, nil
}

func initDag(ctx context.Context, cfg *config.Config) (dag.DAG, error) {
	c := dag.DefaultConfig()
	if c == nil {
		return nil, errors.New("failed to config")
	}

	c.Config = *cfg

	return dag.New(ctx, c), nil
}

// nolint: funlen,gocyclo
func initRunner(ctx context.Context, cfg *config.Config, name string, d dag.DAG) (runner.Tasker, runner.Glancer,
	runner.Mainter, runner.Configer, error) {
	tasker := func() (*runner.TaskerConfig, error) {
		t := runner.TaskerDefaultConfig()
		if t == nil {
			return nil, errors.New("failed to config")
		}
		t.Config = *cfg
		t.Dag = d
		buf, err := loadFile(name)
		if err != nil {
			return nil, errors.Wrap(err, "failed to load")
		}
		if err := json.Unmarshal(buf, &t.Data); err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal")
		}
		return t, nil
	}

	glancer := func() (*runner.GlancerConfig, error) {
		g := runner.GlancerDefaultConfig()
		if g == nil {
			return nil, errors.New("failed to config")
		}
		g.Config = *cfg
		buf, err := loadFile(name)
		if err != nil {
			return nil, errors.Wrap(err, "failed to load")
		}
		if err := json.Unmarshal(buf, &g.Data); err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal")
		}
		return g, nil
	}

	mainter := func() (*runner.MainterConfig, error) {
		m := runner.MainterDefaultConfig()
		if m == nil {
			return nil, errors.New("failed to config")
		}
		m.Config = *cfg
		buf, err := loadFile(name)
		if err != nil {
			return nil, errors.Wrap(err, "failed to load")
		}
		if err := json.Unmarshal(buf, &m.Data); err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal")
		}
		return m, nil
	}

	configer := func() (*runner.ConfigerConfig, error) {
		m := runner.ConfigerDefaultConfig()
		if m == nil {
			return nil, errors.New("failed to config")
		}
		m.Config = *cfg
		buf, err := loadFile(name)
		if err != nil {
			return nil, errors.Wrap(err, "failed to load")
		}
		if err := json.Unmarshal(buf, &m.Data); err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal")
		}
		return m, nil
	}

	t, err := tasker()
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "failed to init tasker")
	}

	g, err := glancer()
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "failed to init glancer")
	}

	m, err := mainter()
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "failed to init mainter")
	}

	c, err := configer()
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "failed to init configer")
	}

	return runner.TaskerNew(ctx, t), runner.GlancerNew(ctx, g), runner.MainterNew(ctx, m), runner.ConfigerNew(ctx, c), nil
}

func initScheduler(ctx context.Context, cfg *config.Config, name string) (scheduler.Scheduler, error) {
	c := scheduler.DefaultConfig()
	if c == nil {
		return nil, errors.New("failed to config")
	}

	c.Config = *cfg

	buf, err := loadFile(name)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load")
	}

	if err := json.Unmarshal(buf, &c.Data); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal")
	}

	return scheduler.New(ctx, c), nil
}

func initPipeline(ctx context.Context, cfg *config.Config, tasker runner.Tasker, sched scheduler.Scheduler) (pipeline.Pipeline, error) {
	c := pipeline.DefaultConfig()
	if c == nil {
		return nil, errors.New("failed to config")
	}

	c.Config = *cfg
	c.Tasker = tasker
	c.Scheduler = sched

	return pipeline.New(ctx, c), nil
}

// nolint: gosec
func runPipeline(ctx context.Context, tasker runner.Tasker, pipe pipeline.Pipeline) error {
	if err := pipe.Init(ctx); err != nil {
		return errors.Wrap(err, "failed to init")
	}

	s, l, err := pipe.Run(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to run")
	}

	fmt.Println("    Run: scheduler")
	fmt.Println("   Name:", s.Name)
	fmt.Println("  Error:", s.Error)

	fmt.Println()
	fmt.Println("    Run: runner.tasker")

	done := make(chan bool, 1)
	go printer(ctx, tasker, l, done)

L:
	for {
		select {
		case <-done:
			break L
		case <-ctx.Done():
			break L
		}
	}

	_ = pipe.Deinit(ctx)

	return nil
}

func printer(ctx context.Context, tasker runner.Tasker, log _runner.Log, done chan<- bool) {
	tasks := tasker.Tasks(ctx)

	for range tasks {
		for {
			line := <-log.Line
			fmt.Println("    Pos:", line.Pos)
			fmt.Println("   Time:", line.Time)
			fmt.Println("Message:", line.Message)
			if line.Message == "EOF" {
				break
			}
		}
	}

	done <- true
}

func runGlance(ctx context.Context, glancer runner.Glancer) error {
	if err := glancer.Init(ctx); err != nil {
		return errors.Wrap(err, "failed to init")
	}

	out, err := glancer.Run(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to run")
	}

	buf, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return errors.Wrap(err, "failed to marshal")
	}

	fmt.Println("    Run: runner.glancer")
	fmt.Println(" Output:", string(buf))

	_ = glancer.Deinit(ctx)

	return nil
}

func runMaint(ctx context.Context, mainter runner.Mainter) error {
	if err := mainter.Init(ctx); err != nil {
		return errors.Wrap(err, "failed to init")
	}

	out, err := mainter.Run(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to run")
	}

	buf, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return errors.Wrap(err, "failed to marshal")
	}

	fmt.Println("    Run: runner.mainter")
	fmt.Println(" Output:", string(buf))

	_ = mainter.Deinit(ctx)

	return nil
}

func runConfig(ctx context.Context, configer runner.Configer) error {
	if err := configer.Init(ctx); err != nil {
		return errors.Wrap(err, "failed to init")
	}

	out, err := configer.Run(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to run")
	}

	buf, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return errors.Wrap(err, "failed to marshal")
	}

	fmt.Println("    Run: runner.configer")
	fmt.Println(" Output:", string(buf))

	_ = configer.Deinit(ctx)

	return nil
}
