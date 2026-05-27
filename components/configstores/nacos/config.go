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

package nacos

// map keys
const (
	namespaceIdKey = "namespace_id"
	userNameKey    = "username"
	passwordKey    = "password"
	endPointKey    = "end_point"
	regionIdKey    = "region_id"
	accessKey      = "access_key"
	secretKey      = "secret_key"
	logDirKey      = "log_dir"
	cacheDirKey    = "cache_dir"
	logLevelKey    = "log_level" // support debug, info, warn, error
)

type Metadata struct {
	NameSpaceId string
	Username    string
	Password    string
	// ACM & KMS
	Endpoint  string
	RegionId  string
	AccessKey string
	SecretKey string
	OpenKMS   bool
	// log & cache files
	LogDir   string
	CacheDir string
	LogLevel string
}

func ParseNacosMetadata(properties map[string]string) (*Metadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// the namespace of config, not required

// ACM & KMS

// log & cache files
