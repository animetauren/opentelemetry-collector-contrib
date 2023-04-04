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

<<<<<<<< HEAD:receiver/fluentforwardreceiver/config_test.go
package fluentforwardreceiver
========
package extensiontest
>>>>>>>> upstream/main:extension/extensiontest/nop_extension_test.go

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
<<<<<<<< HEAD:receiver/fluentforwardreceiver/config_test.go
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap/confmaptest"
)

func TestLoadConfig(t *testing.T) {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)

	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()

	sub, err := cm.Sub(component.NewID("fluentforward").String())
	require.NoError(t, err)
	require.NoError(t, component.UnmarshalConfig(sub, cfg))

	assert.NoError(t, component.ValidateConfig(cfg))
	assert.Equal(t, factory.CreateDefaultConfig(), cfg)
========

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenttest"
)

func TestNewNopFactory(t *testing.T) {
	factory := NewNopFactory()
	require.NotNil(t, factory)
	assert.Equal(t, component.Type("nop"), factory.Type())
	cfg := factory.CreateDefaultConfig()
	assert.Equal(t, &nopConfig{}, cfg)

	traces, err := factory.CreateExtension(context.Background(), NewNopCreateSettings(), cfg)
	require.NoError(t, err)
	assert.NoError(t, traces.Start(context.Background(), componenttest.NewNopHost()))
	assert.NoError(t, traces.Shutdown(context.Background()))
>>>>>>>> upstream/main:extension/extensiontest/nop_extension_test.go
}

func TestNewNopBuilder(t *testing.T) {
	builder := NewNopBuilder()
	require.NotNil(t, builder)

	factory := NewNopFactory()
	cfg := factory.CreateDefaultConfig()
	set := NewNopCreateSettings()
	set.ID = component.NewID(typeStr)

	ext, err := factory.CreateExtension(context.Background(), set, cfg)
	require.NoError(t, err)
	bExt, err := builder.Create(context.Background(), set)
	require.NoError(t, err)
	assert.IsType(t, ext, bExt)
}
