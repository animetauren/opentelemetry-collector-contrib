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

package pmetric

import (
	"testing"

	"github.com/stretchr/testify/assert"
<<<<<<< HEAD:receiver/hostmetricsreceiver/internal/scraper/memoryscraper/factory_test.go
	"go.opentelemetry.io/collector/receiver/receivertest"
=======
>>>>>>> upstream/main:pdata/pmetric/metric_data_point_flags_test.go
)

func TestLogRecordFlags(t *testing.T) {
	flags := DataPointFlags(1)
	assert.True(t, flags.NoRecordedValue())
	assert.EqualValues(t, uint32(1), flags)

	flags = flags.WithNoRecordedValue(false)
	assert.False(t, flags.NoRecordedValue())
	assert.EqualValues(t, uint32(0), flags)

	flags = flags.WithNoRecordedValue(true)
	assert.True(t, flags.NoRecordedValue())
	assert.EqualValues(t, uint32(1), flags)
}

<<<<<<< HEAD:receiver/hostmetricsreceiver/internal/scraper/memoryscraper/factory_test.go
func TestCreateMetricsScraper(t *testing.T) {
	factory := &Factory{}
	cfg := &Config{}

	scraper, err := factory.CreateMetricsScraper(context.Background(), receivertest.NewNopCreateSettings(), cfg)

	assert.NoError(t, err)
	assert.NotNil(t, scraper)
=======
func TestDefaultLogRecordFlags(t *testing.T) {
	flags := DefaultDataPointFlags
	assert.False(t, flags.NoRecordedValue())
	assert.EqualValues(t, uint32(0), flags)
>>>>>>> upstream/main:pdata/pmetric/metric_data_point_flags_test.go
}
