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

<<<<<<< HEAD:exporter/skywalkingexporter/config.go
package skywalkingexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/skywalkingexporter"

import (
	"errors"
=======
package otlpexporter // import "go.opentelemetry.io/collector/exporter/otlpexporter"

import (
	"fmt"
>>>>>>> upstream/main:exporter/otlpexporter/config.go

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

// Config defines configuration for SkyWalking exporter.
type Config struct {
<<<<<<< HEAD:exporter/skywalkingexporter/config.go
	configgrpc.GRPCClientSettings  `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct.
=======
	exporterhelper.TimeoutSettings `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct.
>>>>>>> upstream/main:exporter/otlpexporter/config.go
	exporterhelper.QueueSettings   `mapstructure:"sending_queue"`
	exporterhelper.RetrySettings   `mapstructure:"retry_on_failure"`
	exporterhelper.TimeoutSettings `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct.

	// The number of grpc streams that send the gRPC requests.
	NumStreams int `mapstructure:"num_streams"`
}

var _ component.Config = (*Config)(nil)

// Validate checks if the exporter configuration is valid
func (cfg *Config) Validate() error {
<<<<<<< HEAD:exporter/skywalkingexporter/config.go
	if cfg.Endpoint == "" {
		return errors.New("Skywalking exporter cfg requires an Endpoint")
	}

	if cfg.NumStreams <= 0 {
		return errors.New("Skywalking exporter cfg requires at least one stream")
	}
=======
	if err := cfg.QueueSettings.Validate(); err != nil {
		return fmt.Errorf("queue settings has invalid configuration: %w", err)
	}

>>>>>>> upstream/main:exporter/otlpexporter/config.go
	return nil
}
