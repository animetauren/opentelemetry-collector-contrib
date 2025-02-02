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

package hostmetadata

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
)

<<<<<<<< HEAD:exporter/datadogexporter/internal/hostmetadata/host_test.go
func TestHost(t *testing.T) {
	p, err := GetSourceProvider(componenttest.NewNopTelemetrySettings(), "test-host")
	require.NoError(t, err)
	src, err := p.Source(context.Background())
	require.NoError(t, err)
	assert.Equal(t, src.Identifier, "test-host")
========
func TestExampleReceiver(t *testing.T) {
	rcv := &ExampleReceiver{}
	host := componenttest.NewNopHost()
	assert.False(t, rcv.Started())
	assert.NoError(t, rcv.Start(context.Background(), host))
	assert.True(t, rcv.Started())

	assert.False(t, rcv.Stopped())
	assert.NoError(t, rcv.Shutdown(context.Background()))
	assert.True(t, rcv.Stopped())
>>>>>>>> upstream/main:service/internal/testcomponents/example_receiver_test.go
}
