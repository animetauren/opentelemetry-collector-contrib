// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

<<<<<<< HEAD:exporter/fileexporter/doc.go
// Package fileexporter exports data to files.
package fileexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/fileexporter"
=======
package plogotlp

import (
	"encoding/json"
)

var _ json.Unmarshaler = ExportResponse{}
var _ json.Marshaler = ExportResponse{}
>>>>>>> upstream/main:pdata/plog/plogotlp/response_test.go
