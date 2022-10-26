// Package otlphttp provides an otelcol.exporter.otlphttp component.
package otlphttp

import (
	"errors"
	"time"

	"github.com/grafana/agent/component"
	"github.com/grafana/agent/component/otelcol"
	"github.com/grafana/agent/component/otelcol/exporter"
	"github.com/grafana/agent/pkg/river"
	otelcomponent "go.opentelemetry.io/collector/component"
	otelconfig "go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/exporter/otlphttpexporter"
)

func init() {
	component.Register(component.Registration{
		Name:    "otelcol.exporter.otlphttp",
		Args:    Arguments{},
		Exports: otelcol.ConsumerExports{},

		Build: func(opts component.Options, args component.Arguments) (component.Component, error) {
			fact := otlphttpexporter.NewFactory()
			return exporter.New(opts, fact, args.(Arguments))
		},
	})
}

// Arguments configures the otelcol.exporter.otlphttp component.
type Arguments struct {
	Client HTTPClientArguments    `river:"client,block"`
	Queue  otelcol.QueueArguments `river:"sending_queue,block,optional"`
	Retry  otelcol.RetryArguments `river:"retry_on_failure,block,optional"`

	// The URLs to send metrics/logs/traces to. If omitted the exporter will
	// use Client.Endpoint by appending "/v1/metrics", "/v1/logs" or
	// "/v1/traces", respectively. If set, these settings override
	// Client.Endpoint for the corresponding signal.
	TracesEndpoint  string `river:"traces_endpoint,attr,optional"`
	MetricsEndpoint string `river:"metrics_endpoint,attr,optional"`
	LogsEndpoint    string `river:"logs_endpoint,attr,optional"`
}

var (
	_ river.Unmarshaler  = (*Arguments)(nil)
	_ river.Unmarshaler  = (*HTTPClientArguments)(nil)
	_ exporter.Arguments = Arguments{}
)

// DefaultArguments holds default values for Arguments.
var DefaultArguments = Arguments{
	Queue:  otelcol.DefaultQueueArguments,
	Retry:  otelcol.DefaultRetryArguments,
	Client: DefaultHTTPClientArguments,
}

// UnmarshalRiver implements river.Unmarshaler.
func (args *Arguments) UnmarshalRiver(f func(interface{}) error) error {
	*args = DefaultArguments
	type arguments Arguments
	err := f((*arguments)(args))
	if err != nil {
		return err
	}
	return args.Validate()
}

// Convert implements exporter.Arguments.
func (args Arguments) Convert() otelconfig.Exporter {
	return &otlphttpexporter.Config{
		ExporterSettings:   otelconfig.NewExporterSettings(otelconfig.NewComponentID("otlp")),
		HTTPClientSettings: *(*otelcol.HTTPClientArguments)(&args.Client).Convert(),
		QueueSettings:      *args.Queue.Convert(),
		RetrySettings:      *args.Retry.Convert(),
	}
}

// Extensions implements exporter.Arguments.
func (args Arguments) Extensions() map[otelconfig.ComponentID]otelcomponent.Extension {
	return (*otelcol.HTTPClientArguments)(&args.Client).Extensions()
}

// Exporters implements exporter.Arguments.
func (args Arguments) Exporters() map[otelconfig.DataType]map[otelconfig.ComponentID]otelcomponent.Exporter {
	return nil
}

// Validate returns an error if the configuration is invalid.
func (args *Arguments) Validate() error {
	if args.Client.Endpoint == "" && args.TracesEndpoint == "" && args.MetricsEndpoint == "" && args.LogsEndpoint == "" {
		return errors.New("at least one endpoint must be specified")
	}
	return nil
}

// HTTPClientArguments is used to configure otelcol.exporter.otlphttp with
// component-specific defaults.
type HTTPClientArguments otelcol.HTTPClientArguments

// Default server settings.
var (
	DefaultMaxIddleConns       = 100
	DefaultIdleConnTimeout     = 90 * time.Second
	DefaultHTTPClientArguments = HTTPClientArguments{
		MaxIdleConns:    &DefaultMaxIddleConns,
		IdleConnTimeout: &DefaultIdleConnTimeout,

		Timeout:         30 * time.Second,
		Headers:         map[string]string{},
		Compression:     otelcol.CompressionTypeGzip,
		ReadBufferSize:  0,
		WriteBufferSize: 512 * 1024,
	}
)

// UnmarshalRiver implements river.Unmarshaler and supplies defaults.
func (args *HTTPClientArguments) UnmarshalRiver(f func(interface{}) error) error {
	*args = DefaultHTTPClientArguments
	type arguments HTTPClientArguments
	return f((*arguments)(args))
}
