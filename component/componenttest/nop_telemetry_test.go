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

<<<<<<<< HEAD:receiver/sqlqueryreceiver/metrics_test.go
package sqlqueryreceiver
========
package componenttest
>>>>>>>> upstream/main:component/componenttest/nop_telemetry_test.go

import (
	"testing"

	"github.com/stretchr/testify/assert"
<<<<<<<< HEAD:receiver/sqlqueryreceiver/metrics_test.go
	"go.opentelemetry.io/collector/pdata/pmetric"
)

func TestSetDataPointValue(t *testing.T) {
	err := setDataPointValue(MetricCfg{
		ValueType:   MetricValueTypeInt,
		ValueColumn: "some-col",
	}, "", pmetric.NewNumberDataPoint())
	assert.EqualError(t, err, `setDataPointValue: col "some-col": error converting to integer: strconv.Atoi: parsing "": invalid syntax`)
========

	"go.opentelemetry.io/collector/config/configtelemetry"
)

func TestNewNopTelemetrySettings(t *testing.T) {
	nts := NewNopTelemetrySettings()
	assert.NotNil(t, nts.Logger)
	assert.NotNil(t, nts.TracerProvider)
	assert.NotPanics(t, func() {
		nts.TracerProvider.Tracer("test")
	})
	assert.NotNil(t, nts.MeterProvider)
	assert.NotPanics(t, func() {
		nts.MeterProvider.Meter("test")
	})
	assert.Equal(t, configtelemetry.LevelNone, nts.MetricsLevel)
>>>>>>>> upstream/main:component/componenttest/nop_telemetry_test.go
}
