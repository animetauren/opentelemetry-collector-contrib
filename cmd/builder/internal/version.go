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

<<<<<<<< HEAD:testbed/correctnesstests/metrics/metric_supplier.go
package metrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/correctnesstests/metrics"

import (
	"go.opentelemetry.io/collector/pdata/pmetric"
)

type metricSupplier struct {
	pdms    []pmetric.Metrics
	currIdx int
}

func newMetricSupplier(pdms []pmetric.Metrics) *metricSupplier {
	return &metricSupplier{pdms: pdms}
}

func (p *metricSupplier) nextMetrics() (pdm pmetric.Metrics, done bool) {
	if p.currIdx == len(p.pdms) {
		return pmetric.Metrics{}, true
========
package internal // import "go.opentelemetry.io/collector/cmd/builder/internal"

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version = "dev"
	date    = "unknown"
)

func versionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Version of ocb",
		Long:  "Prints the version of the ocb binary",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println(fmt.Sprintf("%s version %s", cmd.Parent().Name(), version))
		},
>>>>>>>> upstream/main:cmd/builder/internal/version.go
	}
}
