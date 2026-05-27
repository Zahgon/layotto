/*
 * Copyright 2021 Layotto Authors
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
 */

package actuators

import (
	"sync"
)

const (
	reasonKey = "reason"
)

func NewHealthIndicator() *HealthIndicator { _ = "STUB: not implemented"; return nil }

type HealthIndicator struct {
	mu sync.Mutex

	started   bool
	isErr     bool
	errReason string
}

func (idc *HealthIndicator) Report() (status Status, details map[string]interface{}) {
	_ = "STUB: not implemented"
	return *new(Status), nil
}

func (idc *HealthIndicator) ReportError(reason string) { _ = "STUB: not implemented"; return }

func (idc *HealthIndicator) SetStarted() { _ = "STUB: not implemented"; return }
