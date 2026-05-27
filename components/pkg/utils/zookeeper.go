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

	"github.com/go-zookeeper/zk"
)

const (
	zkHost                = "zookeeperHosts"
	zkPassword            = "zookeeperPassword"
	sessionTimeout        = "SessionTimeout"
	logInfo               = "LogInfo"
	defaultSessionTimeout = 5 * time.Second
)

type ConnectionFactory interface {
	NewConnection(expire time.Duration, meta ZookeeperMetadata) (ZKConnection, error)
}

type ConnectionFactoryImpl struct {
}

func (c *ConnectionFactoryImpl) NewConnection(expire time.Duration, meta ZookeeperMetadata) (ZKConnection, error) {
	_ = "STUB: not implemented"
	return *new(ZKConnection), nil
}

type ZKConnection interface {
	Get(path string) ([]byte, *zk.Stat, error)
	Set(path string, data []byte, version int32) (*zk.Stat, error)
	Delete(path string, version int32) error
	Create(path string, data []byte, flags int32, acl []zk.ACL) (string, error)
	Close()
}

type ZookeeperMetadata struct {
	Hosts          []string
	Password       string
	SessionTimeout time.Duration
	LogInfo        bool
}

func ParseZookeeperMetadata(properties map[string]string) (ZookeeperMetadata, error) {
	_ = "STUB: not implemented"
	return *new(ZookeeperMetadata), nil
}
