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

package otlp // import "go.opentelemetry.io/collector/pdata/internal/otlp"

<<<<<<<< HEAD:cmd/mdatagen/statusdata.go
var distros = map[string]string{
	"core":    "https://github.com/open-telemetry/opentelemetry-collector-releases/tree/main/distributions/otelcol",
	"contrib": "https://github.com/open-telemetry/opentelemetry-collector-releases/tree/main/distributions/otelcol-contrib",
}

type Status struct {
	Stability     string   `mapstructure:"stability"`
	Pipelines     []string `mapstructure:"pipelines"`
	Distributions []string `mapstructure:"distributions"`
	Type          string   `mapstructure:"type"`
	Warnings      []string `mapstructure:"warnings"`
========
import (
	otlptrace "go.opentelemetry.io/collector/pdata/internal/data/protogen/trace/v1"
)

// MigrateTraces implements any translation needed due to deprecation in OTLP traces protocol.
// Any ptrace.Unmarshaler implementation from OTLP (proto/json) MUST call this, and the gRPC Server implementation.
func MigrateTraces(rss []*otlptrace.ResourceSpans) {
	for _, rs := range rss {
		if len(rs.ScopeSpans) == 0 {
			rs.ScopeSpans = rs.DeprecatedScopeSpans
		}
		rs.DeprecatedScopeSpans = nil
	}
>>>>>>>> upstream/main:pdata/internal/otlp/traces.go
}
