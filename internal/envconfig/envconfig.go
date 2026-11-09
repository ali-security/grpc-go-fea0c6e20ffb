/*
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package envconfig contains grpc settings configured by environment variables.
package envconfig

import (
	"os"
	"strconv"
	"strings"
)

const (
	prefix          = "GRPC_GO_"
	txtErrIgnoreStr = prefix + "IGNORE_TXT_ERRORS"
)

var (
	// TXTErrIgnore is set if TXT errors should be ignored ("GRPC_GO_IGNORE_TXT_ERRORS" is not "false").
	TXTErrIgnore = !strings.EqualFold(os.Getenv(txtErrIgnoreStr), "false")

	// DisableStrictPathChecking indicates whether strict path checking is
	// disabled. This feature can be disabled by setting the environment
	// variable GRPC_GO_EXPERIMENTAL_DISABLE_STRICT_PATH_CHECKING to "true".
	//
	// When strict path checking is enabled, gRPC will reject requests with
	// paths that do not conform to the gRPC over HTTP/2 specification found at
	// https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md.
	//
	// When disabled, gRPC will allow paths that do not contain a leading slash.
	// Enabling strict path checking is recommended for security reasons, as it
	// prevents potential path traversal vulnerabilities.
	//
	// A future release will remove this environment variable, enabling strict
	// path checking behavior unconditionally.
	DisableStrictPathChecking = boolFromEnv("GRPC_GO_EXPERIMENTAL_DISABLE_STRICT_PATH_CHECKING", false)

	// ControlBufferThrottleLimit is the maximum number of control frames that can
	// be queued in the control buffer before throttling is applied. The value
	// must be between 1 and 100,000, and is set to 10,000 by default.
	//
	// This environment variable serves as an escape hatch to increase the
	// throttling limit if unforeseen issues arise, and it will be removed in a
	// future release.
	ControlBufferThrottleLimit = uint64FromEnv("GRPC_GO_EXPERIMENTAL_CONTROL_BUFFER_THROTTLE_LIMIT", 10000, 1, 100000)

	// EnableReceiveBufferCompaction enables the compaction of data buffers
	// to reduce the number of buffers in the receive buffer.
	//
	// This environment variable serves as an escape hatch to disable the
	// feature if unforeseen issues arise, and it will be removed in a future
	// release.
	//
	EnableReceiveBufferCompaction = boolFromEnv("GRPC_GO_EXPERIMENTAL_ENABLE_RECEIVE_BUFFER_COMPACTION", true)
)

func boolFromEnv(envVar string, def bool) bool {
	v := os.Getenv(envVar)
	if v == "" {
		return def
	}
	return strings.EqualFold(v, "true")
}

func uint64FromEnv(envVar string, def, min, max uint64) uint64 {
	v, err := strconv.ParseUint(os.Getenv(envVar), 10, 64)
	if err != nil {
		return def
	}
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}
