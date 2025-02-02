// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

<<<<<<<< HEAD:extension/headerssetterextension/internal/source/static_test.go
package source

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStaticSource(t *testing.T) {
	ts := &StaticSource{Value: "acme"}
	tenant, err := ts.Get(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, "acme", tenant)
========
package main

import (
	"go.opentelemetry.io/collector/pdata/internal/cmd/pdatagen/internal"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	for _, fp := range internal.AllPackages {
		check(fp.GenerateFiles())
		check(fp.GenerateTestFiles())
		check(fp.GenerateInternalFiles())
	}
>>>>>>>> upstream/main:pdata/internal/cmd/pdatagen/main.go
}
