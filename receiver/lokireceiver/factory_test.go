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

package lokireceiver

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/component/componenttest"
<<<<<<< HEAD:receiver/lokireceiver/factory_test.go
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/confignet"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/receiver/receivertest"
=======
	"go.opentelemetry.io/collector/processor/processortest"
>>>>>>> upstream/main:processor/batchprocessor/factory_test.go
)

func TestCreateDefaultConfig(t *testing.T) {
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()
	assert.NotNil(t, cfg, "failed to create default config")
	assert.NoError(t, componenttest.CheckConfigStruct(cfg))
}

func TestCreateReceiver(t *testing.T) {
	factory := NewFactory()
	cfg := factory.CreateDefaultConfig()
<<<<<<< HEAD:receiver/lokireceiver/factory_test.go
	cfg.(*Config).Protocols.GRPC = &configgrpc.GRPCServerSettings{
		NetAddr: confignet.NetAddr{
			Endpoint:  defaultGRPCBindEndpoint,
			Transport: "tcp",
		},
	}
	set := receivertest.NewNopCreateSettings()
	receiver, err := factory.CreateLogsReceiver(context.Background(), set, cfg, consumertest.NewNop())
	assert.NoError(t, err, "receiver creation failed")
	assert.NotNil(t, receiver, "receiver creation failed")
=======
	creationSet := processortest.NewNopCreateSettings()
	tp, err := factory.CreateTracesProcessor(context.Background(), creationSet, cfg, nil)
	assert.NotNil(t, tp)
	assert.NoError(t, err, "cannot create trace processor")

	mp, err := factory.CreateMetricsProcessor(context.Background(), creationSet, cfg, nil)
	assert.NotNil(t, mp)
	assert.NoError(t, err, "cannot create metric processor")

	lp, err := factory.CreateLogsProcessor(context.Background(), creationSet, cfg, nil)
	assert.NotNil(t, lp)
	assert.NoError(t, err, "cannot create logs processor")
>>>>>>> upstream/main:processor/batchprocessor/factory_test.go
}
