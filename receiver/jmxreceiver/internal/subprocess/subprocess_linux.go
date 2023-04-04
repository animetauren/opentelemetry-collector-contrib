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

<<<<<<< HEAD:receiver/jmxreceiver/internal/subprocess/subprocess_linux.go
//go:build linux
// +build linux

package subprocess // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jmxreceiver/internal/subprocess"

import (
	"os/exec"
	"syscall"
)

func applyOSSpecificCmdModifications(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		// This is Linux-specific and will cause the subprocess to be killed by the OS if
		// the collector dies
		Pdeathsig: syscall.SIGTERM,
	}
=======
//go:build linux || darwin
// +build linux darwin

package loggingexporter // import "go.opentelemetry.io/collector/exporter/loggingexporter"

import (
	"errors"
	"syscall"
)

var knownSyncErrors = []error{
	// sync /dev/stdout: invalid argument
	syscall.EINVAL,
	// sync /dev/stdout: not supported
	syscall.ENOTSUP,
	// sync /dev/stdout: inappropriate ioctl for device
	syscall.ENOTTY,
	// sync /dev/stdout: bad file descriptor
	syscall.EBADF,
}

// knownSyncError returns true if the given error is one of the known
// non-actionable errors returned by Sync on Linux and macOS.
func knownSyncError(err error) bool {
	for _, syncError := range knownSyncErrors {
		if errors.Is(err, syncError) {
			return true
		}
	}

	return false
>>>>>>> upstream/main:exporter/loggingexporter/known_sync_error.go
}
