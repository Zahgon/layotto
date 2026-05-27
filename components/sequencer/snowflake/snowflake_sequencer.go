// Copyright 2021 Layotto Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package snowflake

import (
	"context"
	"database/sql"
	"sync"

	"mosn.io/layotto/kit/logger"

	"mosn.io/layotto/components/pkg/actuators"
	"mosn.io/layotto/components/sequencer"
)

const (
	componentName = "sequencer-snowflake"
)

var (
	once               sync.Once
	readinessIndicator *actuators.HealthIndicator
	livenessIndicator  *actuators.HealthIndicator
)

func init() {
	readinessIndicator = actuators.NewHealthIndicator()
	livenessIndicator = actuators.NewHealthIndicator()
}

type SnowFlakeSequencer struct {
	metadata   SnowflakeMetadata
	workerId   int64
	db         *sql.DB
	mu         sync.Mutex
	smap       map[string]chan int64
	biggerThan map[string]int64
	logger     logger.Logger
	ctx        context.Context
	cancel     context.CancelFunc
}

func NewSnowFlakeSequencer() *SnowFlakeSequencer { _ = "STUB: not implemented"; return nil }

func (s *SnowFlakeSequencer) OnLogLevelChanged(level logger.LogLevel) {
	_ = "STUB: not implemented"
	return
}

func (s *SnowFlakeSequencer) Init(config sequencer.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

//for unit test

func (s *SnowFlakeSequencer) GetNextId(req *sequencer.GetNextIdRequest) (*sequencer.GetNextIdResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//If the key appears for the first time, start a new goroutine for it. If the key doesn't appear for a long time, close the goroutine

func (s *SnowFlakeSequencer) GetSegment(req *sequencer.GetSegmentRequest) (support bool, result *sequencer.GetSegmentResponse, err error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

func (s *SnowFlakeSequencer) producer(id, currentTimeStamp int64, ch chan int64, key string) {
	_ = "STUB: not implemented"
	return
}

//if timeout, remove key from map and record key, workerId, timestamp to mysql

func (s *SnowFlakeSequencer) Close() error { _ = "STUB: not implemented"; return nil }
