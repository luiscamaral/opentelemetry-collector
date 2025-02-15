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

package consumererror // import "go.opentelemetry.io/collector/consumer/consumererror"

import (
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type retryable[V ptrace.Traces | pmetric.Metrics | plog.Logs] struct {
	error
	data V
}

// Unwrap returns the wrapped error for functions Is and As in standard package errors.
func (err retryable[V]) Unwrap() error {
	return err.error
}

// Data returns the telemetry data that failed to be processed or sent.
func (err retryable[V]) Data() V {
	return err.data
}

// Traces is an error that may carry associated Trace data for a subset of received data
// that failed to be processed or sent.
type Traces struct {
	retryable[ptrace.Traces]
}

// NewTraces creates a Traces that can encapsulate received data that failed to be processed or sent.
func NewTraces(err error, data ptrace.Traces) error {
	return Traces{
		retryable: retryable[ptrace.Traces]{
			error: err,
			data:  data,
		},
	}
}

// Deprecated: [v0.73.0] Use `Data` instead.
func (err Traces) GetTraces() ptrace.Traces {
	return err.Data()
}

// Logs is an error that may carry associated Log data for a subset of received data
// that failed to be processed or sent.
type Logs struct {
	retryable[plog.Logs]
}

// NewLogs creates a Logs that can encapsulate received data that failed to be processed or sent.
func NewLogs(err error, data plog.Logs) error {
	return Logs{
		retryable: retryable[plog.Logs]{
			error: err,
			data:  data,
		},
	}
}

// Deprecated: [v0.73.0] Use `Data` instead.
func (err Logs) GetLogs() plog.Logs {
	return err.Data()
}

// Metrics is an error that may carry associated Metrics data for a subset of received data
// that failed to be processed or sent.
type Metrics struct {
	retryable[pmetric.Metrics]
}

// NewMetrics creates a Metrics that can encapsulate received data that failed to be processed or sent.
func NewMetrics(err error, data pmetric.Metrics) error {
	return Metrics{
		retryable: retryable[pmetric.Metrics]{
			error: err,
			data:  data,
		},
	}
}

// Deprecated: [v0.73.0] Use `Data` instead.
func (err Metrics) GetMetrics() pmetric.Metrics {
	return err.Data()
}
