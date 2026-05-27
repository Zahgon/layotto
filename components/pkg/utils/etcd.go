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
	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	defaultKeyPrefix   = "/layotto/"
	defaultDialTimeout = 5
	prefixKey          = "keyPrefixPath"
	usernameKey        = "username"
	passwordKey        = "password"
	dialTimeoutKey     = "dialTimeout"
	endpointsKey       = "endpoints"
	tlsCertPathKey     = "tlsCert"
	tlsCertKeyPathKey  = "tlsCertKey"
	tlsCaPathKey       = "tlsCa"
)

func ParseEtcdMetadata(properties map[string]string) (EtcdMetadata, error) {
	_ = "STUB: not implemented"
	return *new(EtcdMetadata), nil
}

type EtcdMetadata struct {
	KeyPrefix   string
	DialTimeout int
	Endpoints   []string
	Username    string
	Password    string

	TlsCa      string
	TlsCert    string
	TlsCertKey string
}

func addPathSeparator(p string) string { _ = "STUB: not implemented"; return "" }

func NewEtcdClient(meta EtcdMetadata) (*clientv3.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//enable tls

//ping
