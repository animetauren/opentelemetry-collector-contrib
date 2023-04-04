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

<<<<<<<< HEAD:extension/experimental/storage/filestorage/default_windows.go
//go:build windows
// +build windows

package filestorage // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/filestorage"

import (
	"os"
	"path/filepath"
)

func getDefaultDirectory() string {
	return filepath.Join(os.Getenv("ProgramData"), "Otelcol", "FileStorage")
}
========
package semconv // import "go.opentelemetry.io/collector/semconv/v1.5.0"

// SchemaURL is the schema URL that matches the version of the semantic conventions
// that this package defines. Conventions packages starting from v1.4.0 must declare
// non-empty schema URL in the form https://opentelemetry.io/schemas/<version>
const SchemaURL = "https://opentelemetry.io/schemas/1.5.0"
>>>>>>>> upstream/main:semconv/v1.5.0/schema.go
