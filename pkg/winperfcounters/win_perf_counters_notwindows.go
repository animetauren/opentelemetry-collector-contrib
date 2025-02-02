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

<<<<<<<< HEAD:pkg/winperfcounters/win_perf_counters_notwindows.go
//go:build !windows
// +build !windows

package winperfcounters // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/winperfcounters"
========
package main

import (
	"github.com/spf13/cobra"

	"go.opentelemetry.io/collector/cmd/builder/internal"
)

func main() {
	cmd, err := internal.Command()
	cobra.CheckErr(err)
	cobra.CheckErr(cmd.Execute())
}
>>>>>>>> upstream/main:cmd/builder/main.go
