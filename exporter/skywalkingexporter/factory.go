// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

<<<<<<< HEAD:exporter/skywalkingexporter/factory.go
package skywalkingexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/skywalkingexporter"
=======
package otlpexporter // import "go.opentelemetry.io/collector/exporter/otlpexporter"
>>>>>>> upstream/main:exporter/otlpexporter/factory.go

import (
	"context"

	"go.opentelemetry.io/collector/component"
<<<<<<< HEAD:exporter/skywalkingexporter/factory.go
=======
	"go.opentelemetry.io/collector/config/configcompression"
>>>>>>> upstream/main:exporter/otlpexporter/factory.go
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

const (
	// The value of "type" key in configuration.
	typeStr = "skywalking"
	// The stability level of the exporter.
	stability = component.StabilityLevelBeta
)

<<<<<<< HEAD:exporter/skywalkingexporter/factory.go
// NewFactory creates a factory for Skywalking exporter.
=======
// NewFactory creates a factory for OTLP exporter.
>>>>>>> upstream/main:exporter/otlpexporter/factory.go
func NewFactory() exporter.Factory {
	return exporter.NewFactory(
		typeStr,
		createDefaultConfig,
<<<<<<< HEAD:exporter/skywalkingexporter/factory.go
		exporter.WithLogs(createLogsExporter, stability),
		exporter.WithMetrics(createMetricsExporter, stability))
=======
		exporter.WithTraces(createTracesExporter, component.StabilityLevelStable),
		exporter.WithMetrics(createMetricsExporter, component.StabilityLevelStable),
		exporter.WithLogs(createLogsExporter, component.StabilityLevelBeta),
	)
>>>>>>> upstream/main:exporter/otlpexporter/factory.go
}

func createDefaultConfig() component.Config {
	return &Config{
<<<<<<< HEAD:exporter/skywalkingexporter/factory.go
		RetrySettings:   exporterhelper.NewDefaultRetrySettings(),
		QueueSettings:   exporterhelper.NewDefaultQueueSettings(),
		TimeoutSettings: exporterhelper.NewDefaultTimeoutSettings(),
		GRPCClientSettings: configgrpc.GRPCClientSettings{
=======
		TimeoutSettings: exporterhelper.NewDefaultTimeoutSettings(),
		RetrySettings:   exporterhelper.NewDefaultRetrySettings(),
		QueueSettings:   exporterhelper.NewDefaultQueueSettings(),
		GRPCClientSettings: configgrpc.GRPCClientSettings{
			Headers: map[string]configopaque.String{},
			// Default to gzip compression
			Compression: configcompression.Gzip,
>>>>>>> upstream/main:exporter/otlpexporter/factory.go
			// We almost read 0 bytes, so no need to tune ReadBufferSize.
			WriteBufferSize: 512 * 1024,
		},
		NumStreams: 2,
	}
}

<<<<<<< HEAD:exporter/skywalkingexporter/factory.go
=======
func createTracesExporter(
	ctx context.Context,
	set exporter.CreateSettings,
	cfg component.Config,
) (exporter.Traces, error) {
	oce, err := newExporter(cfg, set)
	if err != nil {
		return nil, err
	}
	oCfg := cfg.(*Config)
	return exporterhelper.NewTracesExporter(ctx, set, cfg,
		oce.pushTraces,
		exporterhelper.WithCapabilities(consumer.Capabilities{MutatesData: false}),
		exporterhelper.WithTimeout(oCfg.TimeoutSettings),
		exporterhelper.WithRetry(oCfg.RetrySettings),
		exporterhelper.WithQueue(oCfg.QueueSettings),
		exporterhelper.WithStart(oce.start),
		exporterhelper.WithShutdown(oce.shutdown))
}

func createMetricsExporter(
	ctx context.Context,
	set exporter.CreateSettings,
	cfg component.Config,
) (exporter.Metrics, error) {
	oce, err := newExporter(cfg, set)
	if err != nil {
		return nil, err
	}
	oCfg := cfg.(*Config)
	return exporterhelper.NewMetricsExporter(ctx, set, cfg,
		oce.pushMetrics,
		exporterhelper.WithCapabilities(consumer.Capabilities{MutatesData: false}),
		exporterhelper.WithTimeout(oCfg.TimeoutSettings),
		exporterhelper.WithRetry(oCfg.RetrySettings),
		exporterhelper.WithQueue(oCfg.QueueSettings),
		exporterhelper.WithStart(oce.start),
		exporterhelper.WithShutdown(oce.shutdown),
	)
}

>>>>>>> upstream/main:exporter/otlpexporter/factory.go
func createLogsExporter(
	ctx context.Context,
	set exporter.CreateSettings,
	cfg component.Config,
) (exporter.Logs, error) {
<<<<<<< HEAD:exporter/skywalkingexporter/factory.go
	oCfg := cfg.(*Config)
	oce := newLogsExporter(ctx, oCfg, set.TelemetrySettings)
	return exporterhelper.NewLogsExporter(
		ctx,
		set,
		cfg,
=======
	oce, err := newExporter(cfg, set)
	if err != nil {
		return nil, err
	}
	oCfg := cfg.(*Config)
	return exporterhelper.NewLogsExporter(ctx, set, cfg,
>>>>>>> upstream/main:exporter/otlpexporter/factory.go
		oce.pushLogs,
		exporterhelper.WithCapabilities(consumer.Capabilities{MutatesData: false}),
		exporterhelper.WithRetry(oCfg.RetrySettings),
		exporterhelper.WithQueue(oCfg.QueueSettings),
		exporterhelper.WithTimeout(oCfg.TimeoutSettings),
		exporterhelper.WithStart(oce.start),
		exporterhelper.WithShutdown(oce.shutdown),
	)
}

func createMetricsExporter(ctx context.Context, set exporter.CreateSettings, cfg component.Config) (exporter.Metrics, error) {
	oCfg := cfg.(*Config)
	oce := newMetricsExporter(ctx, oCfg, set.TelemetrySettings)
	return exporterhelper.NewMetricsExporter(
		ctx,
		set,
		cfg,
		oce.pushMetrics,
		exporterhelper.WithCapabilities(consumer.Capabilities{MutatesData: false}),
		exporterhelper.WithRetry(oCfg.RetrySettings),
		exporterhelper.WithQueue(oCfg.QueueSettings),
		exporterhelper.WithTimeout(oCfg.TimeoutSettings),
		exporterhelper.WithStart(oce.start),
		exporterhelper.WithShutdown(oce.shutdown))
}
