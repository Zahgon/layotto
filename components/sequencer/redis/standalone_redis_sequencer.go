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
package redis

import (
	"context"
	"sync"

	"github.com/go-redis/redis/v8"

	"mosn.io/layotto/kit/logger"

	"mosn.io/layotto/components/pkg/actuators"
	"mosn.io/layotto/components/pkg/utils"
	"mosn.io/layotto/components/sequencer"
)

const (
	componentName = "sequencer-redis-standalone"
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

type StandaloneRedisSequencer struct {
	client     *redis.Client
	metadata   utils.RedisMetadata
	biggerThan map[string]int64

	logger logger.Logger

	ctx    context.Context
	cancel context.CancelFunc
}

// NewStandaloneRedisSequencer returns a new redis sequencer
func NewStandaloneRedisSequencer() *StandaloneRedisSequencer { _ = "STUB: not implemented"; return nil }

func (s *StandaloneRedisSequencer) OnLogLevelChanged(level logger.LogLevel) {
	_ = "STUB: not implemented"
	return
}

/*
1. exists and >= biggerThan, no operation required, return 0
2. not exists or < biggthan, reset val, return 1
3. lua script occur error, such as tonumer(string), return error
*/
const initScript = `
if  redis.call('exists', KEYS[1])==1 and tonumber(redis.call('get', KEYS[1])) >= tonumber(ARGV[1]) then
    return 0
else
     redis.call('set', KEYS[1],ARGV[1])
     return 1
end
`

func (s *StandaloneRedisSequencer) Init(config sequencer.Configuration) error {
	_ = "STUB: not implemented"
	return nil
}

//init

// construct client

//check biggerThan, initialize if not satisfied

//occur error,  such as value is string type

//As long as there is no error, the initialization is successful
//It may be a reset value or it may be satisfied before

func (s *StandaloneRedisSequencer) GetNextId(req *sequencer.GetNextIdRequest) (*sequencer.GetNextIdResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *StandaloneRedisSequencer) GetSegment(req *sequencer.GetSegmentRequest) (bool, *sequencer.GetSegmentResponse, error) {
	_ = "STUB: not implemented"

	// size=0 only check support
	return false, nil, nil
}

func (s *StandaloneRedisSequencer) Close() error { _ = "STUB: not implemented"; return nil }
