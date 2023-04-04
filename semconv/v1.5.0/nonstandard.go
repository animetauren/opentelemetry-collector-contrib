// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

<<<<<<<< HEAD:internal/metadataproviders/internal/goos.go
package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/internal"

// GOOSToOSType maps a runtime.GOOS-like value to os.type style.
func GOOSToOSType(goos string) string {
	switch goos {
	case "dragonfly":
		return "dragonflybsd"
	case "zos":
		return "z_os"
	}
	return goos
}
========
package semconv // import "go.opentelemetry.io/collector/semconv/v1.5.0"

const (
	InstrumentationLibraryName    = "otel.library.name"
	InstrumentationLibraryVersion = "otel.library.version"
)

const (
	OtelStatusCode        = "otel.status_code"
	OtelStatusDescription = "otel.status_description"
)
>>>>>>>> upstream/main:semconv/v1.5.0/nonstandard.go
