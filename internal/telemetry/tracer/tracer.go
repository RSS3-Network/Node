package tracer

import (
	"context"
	"fmt"
	"os"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/trace"
)

type Driver string

const (
	DriverLocal                 Driver = "local"
	DriverOpenTelemetryProtocol Driver = "otlp"
)

func Open(driver Driver, endpoint string, insecure bool) (trace.SpanExporter, error) {
	switch driver {
	case DriverOpenTelemetryProtocol:
		return openOpenTelemetryProtocolExporter(endpoint, insecure)
	case DriverLocal:
		fallthrough
	default:
		return openLocalExporter(endpoint, insecure)
	}
}

func openOpenTelemetryProtocolExporter(endpoint string, insecure bool) (exporter trace.SpanExporter, err error) {
	options := []otlptracehttp.Option{
		otlptracehttp.WithEndpoint(endpoint),
	}

	if insecure {
		options = append(options, otlptracehttp.WithInsecure())
	}

	return otlptrace.New(context.Background(), otlptracehttp.NewClient(options...))
}

func openLocalExporter(path string, _ bool) (exporter trace.SpanExporter, err error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("open local exporter file: %w", err)
	}

	return stdouttrace.New(
		stdouttrace.WithWriter(file),
		stdouttrace.WithPrettyPrint(),
		stdouttrace.WithoutTimestamps(),
	)
}
