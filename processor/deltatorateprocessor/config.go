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

<<<<<<< HEAD:processor/deltatorateprocessor/config.go
package deltatorateprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatorateprocessor"

import (
	"fmt"
)

// Config defines the configuration for the processor.
type Config struct {

	// List of delta sum metrics to convert to rates
	Metrics []string `mapstructure:"metrics"`
}

// Validate checks whether the input configuration has all of the required fields for the processor.
// An error is returned if there are any invalid inputs.
func (config *Config) Validate() error {
	if len(config.Metrics) == 0 {
		return fmt.Errorf("metric names are missing")
	}
	return nil
=======
package consumertest // import "go.opentelemetry.io/collector/consumer/consumertest"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

// NewNop returns a Consumer that just drops all received data and returns no error.
func NewNop() Consumer {
	return &baseConsumer{
		ConsumeTracesFunc:  func(ctx context.Context, td ptrace.Traces) error { return nil },
		ConsumeMetricsFunc: func(ctx context.Context, md pmetric.Metrics) error { return nil },
		ConsumeLogsFunc:    func(ctx context.Context, ld plog.Logs) error { return nil },
	}
>>>>>>> upstream/main:consumer/consumertest/nop.go
}
