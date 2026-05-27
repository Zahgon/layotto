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

package in_memory

import (
	"sync"

	"mosn.io/layotto/components/pkg/actuators"
	"mosn.io/layotto/components/sequencer"
)

var (
	once               sync.Once
	readinessIndicator *actuators.HealthIndicator
	livenessIndicator  *actuators.HealthIndicator
)

const componentName = "in-memory"

func init() {
	readinessIndicator = actuators.NewHealthIndicator()
	livenessIndicator = actuators.NewHealthIndicator()
}

type InMemorySequencer struct {
	data *sync.Map
}

func registerActuator() { _ = "STUB: not implemented"; return }

func NewInMemorySequencer() *InMemorySequencer { _ = "STUB: not implemented"; return nil }

func (s *InMemorySequencer) Init(_ sequencer.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *InMemorySequencer) GetNextId(req *sequencer.GetNextIdRequest) (*sequencer.GetNextIdResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *InMemorySequencer) GetSegment(req *sequencer.GetSegmentRequest) (bool, *sequencer.GetSegmentResponse, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}
