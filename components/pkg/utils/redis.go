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
package utils

import (
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	db                     = "db"
	host                   = "redisHost"
	redisHosts             = "redisHosts"
	password               = "redisPassword"
	enableTLS              = "enableTLS"
	maxRetries             = "maxRetries"
	concurrency            = "concurrency"
	maxRetryBackoff        = "maxRetryBackoff"
	defaultBase            = 10
	defaultBitSize         = 0
	defaultDB              = 0
	defaultMaxRetries      = 3
	defaultMaxRetryBackoff = time.Second * 2
	defaultEnableTLS       = false
)

func NewRedisClient(m RedisMetadata) *redis.Client { _ = "STUB: not implemented"; return nil }

type RedisMetadata struct {
	Host            string
	Password        string
	MaxRetries      int
	MaxRetryBackoff time.Duration
	EnableTLS       bool
	DB              int
}

func ParseRedisMetadata(properties map[string]string) (RedisMetadata, error) {
	_ = "STUB: not implemented"
	return *new(RedisMetadata), nil
}

func NewClusterRedisClient(m RedisClusterMetadata) []*redis.Client {
	_ = "STUB: not implemented"
	return nil
}

type RedisClusterMetadata struct {
	Hosts           []string
	Concurrency     int
	Password        string
	MaxRetries      int
	MaxRetryBackoff time.Duration
	EnableTLS       bool
	DB              int
}

func ParseRedisClusterMetadata(properties map[string]string) (RedisClusterMetadata, error) {
	_ = "STUB: not implemented"
	return *new(RedisClusterMetadata), nil
}

func getConcurrency(properties map[string]string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func GetMiliTimestamp(i int64) int64 { _ = "STUB: not implemented"; return 0 }
