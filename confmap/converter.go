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

package confmap // import "go.opentelemetry.io/collector/confmap"

import (
<<<<<<< HEAD:cmd/configschema/docsgen/main.go
	"path/filepath"

	"github.com/open-telemetry/opentelemetry-collector-contrib/cmd/configschema"
	"github.com/open-telemetry/opentelemetry-collector-contrib/cmd/configschema/docsgen/docsgen"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/components"
)

func main() {
	c, err := components.Components()
	if err != nil {
		panic(err)
	}
	dr := configschema.NewDirResolver(filepath.Join("..", ".."), configschema.DefaultModule)
	docsgen.CLI(c, dr)
=======
	"context"
)

// Converter is a converter interface for the confmap.Conf that allows distributions
// (in the future components as well) to build backwards compatible config converters.
type Converter interface {
	// Convert applies the conversion logic to the given "conf".
	Convert(ctx context.Context, conf *Conf) error
>>>>>>> upstream/main:confmap/converter.go
}
