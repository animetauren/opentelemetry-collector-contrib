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

package network

import (
	"bytes"
	"encoding/binary"
	"testing"
	"unicode/utf16"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
<<<<<<< HEAD:receiver/dotnetdiagnosticsreceiver/network/writer_test.go
)

func TestWriteUTF16String(t *testing.T) {
	buf := &bytes.Buffer{}
	const msg = "hello"
	WriteUTF16String(buf, msg)
	var msgLen int32
	err := binary.Read(buf, ByteOrder, &msgLen)
	require.NoError(t, err)
	assert.Equal(t, len(msg)+1, int(msgLen))
	var p []uint16
	for i := 0; i < len(msg); i++ {
		var v uint16
		err := binary.Read(buf, ByteOrder, &v)
		require.NoError(t, err)
		p = append(p, v)
	}
	require.Equal(t, msg, string(utf16.Decode(p)))
=======

	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

func TestErr(t *testing.T) {
	err := errors.New("my error")
	ec := NewErr(err)
	require.NotNil(t, ec)
	assert.NotPanics(t, ec.unexported)
	assert.Equal(t, err, ec.ConsumeLogs(context.Background(), plog.NewLogs()))
	assert.Equal(t, err, ec.ConsumeMetrics(context.Background(), pmetric.NewMetrics()))
	assert.Equal(t, err, ec.ConsumeTraces(context.Background(), ptrace.NewTraces()))
>>>>>>> upstream/main:consumer/consumertest/err_test.go
}
