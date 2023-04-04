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

<<<<<<< HEAD:processor/tailsamplingprocessor/main_test.go
package tailsamplingprocessor
=======
package obsreport // import "go.opentelemetry.io/collector/obsreport"
>>>>>>> upstream/main:obsreport/obsreport.go

import (
	"testing"

	"go.uber.org/goleak"
)

<<<<<<< HEAD:processor/tailsamplingprocessor/main_test.go
func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m, goleak.IgnoreTopFunction("go.opencensus.io/stats/view.(*worker).start"))
=======
const (
	scopeName = "go.opentelemetry.io/collector/obsreport"

	nameSep = "/"
)

func recordError(span trace.Span, err error) {
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
	}
>>>>>>> upstream/main:obsreport/obsreport.go
}
