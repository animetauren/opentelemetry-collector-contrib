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

<<<<<<<< HEAD:connector/servicegraphconnector/factory.go
package servicegraphconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/servicegraphconnector"

import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/servicegraphprocessor"

var NewFactory = servicegraphprocessor.NewConnectorFactory
========
//go:build !linux && !darwin && !windows
// +build !linux,!darwin,!windows

package loggingexporter // import "go.opentelemetry.io/collector/exporter/loggingexporter"

// knownSyncError returns true if the given error is one of the known
// non-actionable errors returned by Sync on Plan 9.
func knownSyncError(err error) bool {
	return false
}
>>>>>>>> upstream/main:exporter/loggingexporter/known_sync_error_other.go
