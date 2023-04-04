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

<<<<<<<< HEAD:internal/coreinternal/testdata/resource.go
package testdata
========
package components
>>>>>>>> upstream/main:service/internal/components/host_wrapper_test.go

import "go.opentelemetry.io/collector/pdata/pcommon"

<<<<<<<< HEAD:internal/coreinternal/testdata/resource.go
func initResource1(r pcommon.Resource) {
	r.Attributes().PutStr("resource-attr", "resource-attr-val-1")
}

func initResource2(r pcommon.Resource) {
	r.Attributes().PutStr("resource-attr", "resource-attr-val-2")
========
	"go.uber.org/zap"

	"go.opentelemetry.io/collector/component/componenttest"
)

func Test_newHostWrapper(_ *testing.T) {
	hw := NewHostWrapper(componenttest.NewNopHost(), zap.NewNop())
	hw.ReportFatalError(errors.New("test error"))
>>>>>>>> upstream/main:service/internal/components/host_wrapper_test.go
}
