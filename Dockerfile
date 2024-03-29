FROM golang:latest AS build-stage
WORKDIR /go/src/app
COPY . .
RUN make build

FROM gcr.io/distroless/base-debian11 AS production-stage
WORKDIR /
COPY --from=build-stage /go/src/app/bin/cli /
COPY --from=build-stage /go/src/app/test/config/config.yml /
COPY --from=build-stage /go/src/app/test/data/runner.json /
COPY --from=build-stage /go/src/app/test/data/scheduler1.json /
COPY --from=build-stage /go/src/app/test/data/scheduler2.json /
COPY --from=build-stage /go/src/app/test/data/scheduler3.json /
COPY --from=build-stage /go/src/app/test/data/scheduler4.json /
USER nonroot:nonroot
CMD ["/cli", "--config-file=/config.yml", "--runner-file=/runner.json", "--scheduler-file=/scheduler1.json"]
