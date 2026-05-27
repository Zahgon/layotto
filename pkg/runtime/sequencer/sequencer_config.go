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
package sequencer

const (
	strategyKey       = "keyPrefix"
	strategyAppid     = "appid"
	strategyStoreName = "name"
	strategyNone      = "none"
	strategyDefault   = strategyAppid
	apiPrefix         = "sequencer"
	apiSeparator      = "|||"
	separator         = "||"
)

var seqConfiguration = map[string]*StoreConfiguration{}

type StoreConfiguration struct {
	keyPrefixStrategy string
}

func SaveSeqConfiguration(storeName string, metadata map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func GetModifiedSeqKey(key, storeName, appID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getConfiguration(storeName string) *StoreConfiguration { _ = "STUB: not implemented"; return nil }

func checkKeyIllegal(key string) error { _ = "STUB: not implemented"; return nil }
