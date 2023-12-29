package telemetry

import (
	"fmt"

	"github.com/naturalselectionlabs/rss3-node/internal/constant"
	"github.com/naturalselectionlabs/rss3-node/internal/telemetry"
	meterx "github.com/naturalselectionlabs/rss3-node/internal/telemetry/meter"
	tracerx "github.com/naturalselectionlabs/rss3-node/internal/telemetry/tracer"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric"
	meter "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	tracer "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
)

func OpenTelemetryTracer(config *telemetry.OpenTelemetryConfig) (trace.TracerProvider, error) {
	var (
		tracerDriver   = tracerx.DriverOpenTelemetryProtocol
		tracerInsecure = config.Traces.Insecure
		tracerEndpoint = config.Traces.Endpoint
	)

	spanExporter, err := tracerx.Open(tracerx.Driver(tracerDriver), tracerEndpoint, tracerInsecure)
	if err != nil {
		return nil, fmt.Errorf("open tracer: %w", err)
	}

	var (
		serviceName    = constant.Name
		serviceVersion = constant.BuildVersion()
	)

	tracerProvider := tracer.NewTracerProvider(
		tracer.WithBatcher(spanExporter),
		tracer.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
			semconv.ServiceVersionKey.String(serviceVersion),
		)),
	)

	return tracerProvider, nil
}

func OpenTelemetryMeter() (metric.MeterProvider, error) {
	exporter, err := prometheus.New()
	if err != nil {
		return nil, fmt.Errorf("create prometheus exporter: %w", err)
	}

	meterProvider := meter.NewMeterProvider(meter.WithReader(exporter))

	return meterProvider, nil
}

func OpenTelemetryMeterServer() (meterx.Server, error) {
	return meterx.New(meterx.Driver(meterx.DriverPrometheus))
}
