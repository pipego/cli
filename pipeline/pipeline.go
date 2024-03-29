package pipeline

import (
	"context"

	"github.com/pkg/errors"

	"github.com/pipego/cli/config"
	"github.com/pipego/cli/runner"
	"github.com/pipego/cli/scheduler"
	_runner "github.com/pipego/dag/runner"
)

type Pipeline interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Run(context.Context) (scheduler.Result, _runner.Log, error)
}

type Config struct {
	Config    config.Config
	Tasker    runner.Tasker
	Scheduler scheduler.Scheduler
}

type pipeline struct {
	cfg *Config
}

func New(_ context.Context, cfg *Config) Pipeline {
	return &pipeline{
		cfg: cfg,
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (p *pipeline) Init(ctx context.Context) error {
	if err := p.cfg.Scheduler.Init(ctx); err != nil {
		return errors.Wrap(err, "failed to init scheduler")
	}

	if err := p.cfg.Tasker.Init(ctx); err != nil {
		return errors.Wrap(err, "failed to init runner")
	}

	return nil
}

func (p *pipeline) Deinit(ctx context.Context) error {
	_ = p.cfg.Tasker.Deinit(ctx)
	_ = p.cfg.Scheduler.Deinit(ctx)

	return nil
}

func (p *pipeline) Run(ctx context.Context) (s scheduler.Result, l _runner.Log, e error) {
	var err error

	if s, err = p.cfg.Scheduler.Run(ctx); err != nil {
		return scheduler.Result{}, _runner.Log{}, errors.Wrap(err, "failed to issuerail scheduler")
	}

	if err = p.cfg.Tasker.Run(ctx); err != nil {
		return scheduler.Result{}, _runner.Log{}, errors.Wrap(err, "failed to run runner")
	}

	l = p.cfg.Tasker.Tail(ctx)

	return s, l, nil
}
