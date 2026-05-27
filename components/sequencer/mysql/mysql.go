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
package mysql

import (
	"database/sql"
	"sync"

	"mosn.io/layotto/kit/logger"

	"mosn.io/layotto/components/pkg/actuators"
	"mosn.io/layotto/components/pkg/utils"
	"mosn.io/layotto/components/sequencer"
)

const (
	componentName = "sequencer-mysql"
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

type MySQLSequencer struct {
	metadata   utils.MySQLMetadata
	biggerThan map[string]int64
	logger     logger.Logger
	db         *sql.DB
}

func NewMySQLSequencer() *MySQLSequencer { _ = "STUB: not implemented"; return nil }

func (e *MySQLSequencer) OnLogLevelChanged(level logger.LogLevel) {
	_ = "STUB: not implemented"
	return
}

func (e *MySQLSequencer) Init(config sequencer.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *MySQLSequencer) GetNextId(req *sequencer.GetNextIdRequest) (*sequencer.GetNextIdResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *MySQLSequencer) GetSegment(req *sequencer.GetSegmentRequest) (support bool, result *sequencer.GetSegmentResponse, err error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

func (e *MySQLSequencer) Close(db *sql.DB) error { _ = "STUB: not implemented"; return nil }
