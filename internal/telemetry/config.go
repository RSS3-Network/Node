package telemetry

type Config struct {
	OpenTelemetry *OpenTelemetryConfig `yaml:"opentelemetry" validate:"required"`
}

type OpenTelemetryConfig struct {
	Metrics *OpenTelemetryMetricsConfig `yaml:"metrics" validate:"required"`
	Traces  *OpenTelemetryTracesConfig  `yaml:"traces" validate:"required"`
}

type OpenTelemetryMetricsConfig struct {
	Enable   bool   `yaml:"enable" default:"false"`
	Endpoint string `yaml:"endpoint" default:"0.0.0.0:9090"`
}

type OpenTelemetryTracesConfig struct {
	Enable   bool   `yaml:"enable" default:"false"`
	Insecure bool   `yaml:"insecure" default:"false"`
	Endpoint string `yaml:"endpoint" validate:"required" default:"0.0.0.0:4318"`
}
