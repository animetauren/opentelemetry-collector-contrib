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

<<<<<<<< HEAD:receiver/sqlqueryreceiver/factory.go
package sqlqueryreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sqlqueryreceiver"

import (
	"database/sql"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/receiver"
)

const (
	typeStr   = "sqlquery"
	stability = component.StabilityLevelUndefined
)

func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		typeStr,
		createDefaultConfig,
		receiver.WithMetrics(createReceiverFunc(sql.Open, newDbClient), stability),
	)
========
package testcomponents // import "go.opentelemetry.io/collector/service/internal/testcomponents"

import (
	"context"

	"go.opentelemetry.io/collector/component"
)

type componentState struct {
	started bool
	stopped bool
}

func (cs *componentState) Started() bool {
	return cs.started
}

func (cs *componentState) Stopped() bool {
	return cs.stopped
}

func (cs *componentState) Start(_ context.Context, _ component.Host) error {
	cs.started = true
	return nil
}

func (cs *componentState) Shutdown(_ context.Context) error {
	cs.stopped = true
	return nil
>>>>>>>> upstream/main:service/internal/testcomponents/stateful_component.go
}
