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

<<<<<<< HEAD:receiver/fluentforwardreceiver/observ/metrics_test.go
package observ
=======
package configgrpc // import "go.opentelemetry.io/collector/config/configgrpc"
>>>>>>> upstream/main:config/configgrpc/gzip.go

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestViews(t *testing.T) {
	require.Equal(t, len(MetricViews()), 5)
}
