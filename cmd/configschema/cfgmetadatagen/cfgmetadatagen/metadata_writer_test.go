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

// skipping windows to avoid this golang bug: https://github.com/golang/go/issues/51442
//go:build !windows

package cfgmetadatagen

import (
<<<<<<< HEAD:cmd/configschema/cfgmetadatagen/cfgmetadatagen/metadata_writer_test.go
	"io"
	"os"
	"path/filepath"
=======
	"errors"
>>>>>>> upstream/main:receiver/scrapererror/partialscrapeerror_test.go
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/open-telemetry/opentelemetry-collector-contrib/cmd/configschema"
)

<<<<<<< HEAD:cmd/configschema/cfgmetadatagen/cfgmetadatagen/metadata_writer_test.go
func TestMetadataFileWriter(t *testing.T) {
	tempDir := t.TempDir()
	w := newMetadataFileWriter(tempDir)
	err := w.write(configschema.CfgInfo{Group: "mygroup", Type: "mytype"}, []byte("hello"))
	require.NoError(t, err)
	file, err := os.Open(filepath.Join(tempDir, "mygroup", "mytype.yaml"))
	require.NoError(t, err)
	bytes, err := io.ReadAll(file)
	require.NoError(t, err)
	assert.EqualValues(t, "hello", bytes)
=======
func TestPartialScrapeError(t *testing.T) {
	failed := 2
	err := errors.New("some error")
	partialErr := NewPartialScrapeError(err, failed)
	assert.Equal(t, err.Error(), partialErr.Error())
	assert.Equal(t, failed, partialErr.Failed)
}

func TestIsPartialScrapeError(t *testing.T) {
	err := errors.New("testError")
	require.False(t, IsPartialScrapeError(err))

	err = NewPartialScrapeError(err, 2)
	require.True(t, IsPartialScrapeError(err))
>>>>>>> upstream/main:receiver/scrapererror/partialscrapeerror_test.go
}
