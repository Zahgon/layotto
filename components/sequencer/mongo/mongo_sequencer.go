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
package mongo

import (
	"context"
	"sync"

	"mosn.io/layotto/kit/logger"

	"mosn.io/layotto/components/pkg/actuators"
	"mosn.io/layotto/components/pkg/utils"
	"mosn.io/layotto/components/sequencer"
)

const (
	componentName = "sequencer-mongo"
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

type MongoSequencer struct {
	factory utils.MongoFactory

	client     utils.MongoClient
	session    utils.MongoSession
	collection utils.MongoCollection
	singResult utils.MongoSingleResult
	metadata   utils.MongoMetadata
	biggerThan map[string]int64

	logger logger.Logger

	ctx    context.Context
	cancel context.CancelFunc
}

type SequencerDocument struct {
	Id              string `bson:"_id"`
	Sequencer_value int64  `bson:"sequencer_value"`
}

// MongoSequencer returns a new mongo sequencer
func NewMongoSequencer() *MongoSequencer { _ = "STUB: not implemented"; return nil }

func (e *MongoSequencer) OnLogLevelChanged(level logger.LogLevel) {
	_ = "STUB: not implemented"
	return
}

func (e *MongoSequencer) Init(config sequencer.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

// 1.parse config

// 2. construct client

// Connections Collection

// find key of biggerThan

// check biggerThan's value

func (e *MongoSequencer) GetNextId(req *sequencer.GetNextIdRequest) (*sequencer.GetNextIdResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// create mongo session

// check session

// close mongo session

// rollback

// commit

func (e *MongoSequencer) GetSegment(req *sequencer.GetSegmentRequest) (support bool, result *sequencer.GetSegmentResponse, err error) {
	_ = "STUB: not implemented"
	return false,

		// size=0 only check support
		nil, nil
}

// create mongo session

// check session

// close mongo session

// find and upsert

// rollback

// commit

func (e *MongoSequencer) Close() error { _ = "STUB: not implemented"; return nil }
