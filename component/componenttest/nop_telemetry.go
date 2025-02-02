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

<<<<<<<< HEAD:processor/logstransformprocessor/config.go
package logstransformprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/logstransformprocessor"

import (
	"errors"

	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/adapter"
)

// Config defines configuration for Resource processor.
type Config struct {
	adapter.BaseConfig `mapstructure:",squash"`
}

var _ component.Config = (*Config)(nil)

// Validate checks if the processor configuration is valid
func (cfg *Config) Validate() error {
	if len(cfg.BaseConfig.Operators) == 0 {
		return errors.New("no operators were configured for this logs transform processor")
	}
	return nil
========
package componenttest // import "go.opentelemetry.io/collector/component/componenttest"

import (
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configtelemetry"
)

// NewNopTelemetrySettings returns a new nop telemetry settings for Create* functions.
func NewNopTelemetrySettings() component.TelemetrySettings {
	return component.TelemetrySettings{
		Logger:         zap.NewNop(),
		TracerProvider: trace.NewNoopTracerProvider(),
		MeterProvider:  metric.NewNoopMeterProvider(),
		MetricsLevel:   configtelemetry.LevelNone,
	}
>>>>>>>> upstream/main:component/componenttest/nop_telemetry.go
}
