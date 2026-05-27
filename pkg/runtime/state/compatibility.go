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
// CODE ATTRIBUTION: https://github.com/dapr/dapr
// We copied these code here to make our runtime compatible with dapr's component.
package state

const (
	strategyKey = "keyPrefix"

	strategyAppid     = "appid"
	strategyStoreName = "name"
	strategyNone      = "none"
	strategyDefault   = strategyNone

	daprSeparator = "||"
)

var statesConfiguration = map[string]*StoreConfiguration{}

type StoreConfiguration struct {
	keyPrefixStrategy string
}

// Save StateConfiguration by storeName
func SaveStateConfiguration(storeName string, metadata map[string]string) error {
	_ = "STUB: not implemented"
	// convert
	return nil
}

// Change strategy to lowercase

//if strategy is "",use default values("none")

// Check if the secret key is legitimate

// convert

func GetModifiedStateKey(key, storeName, appID string) (string, error) {
	_ = "STUB: not implemented"
	// Check if the secret key is legitimate
	return "", nil
}

// Get stateConfiguration by storeName

// Determine the keyPrefixStrategy type

func GetOriginalStateKey(modifiedStateKey string) string {
	_ = "STUB: not implemented"
	// Split modifiedStateKey by daprSeparator("||")
	return ""
}

func getStateConfiguration(storeName string) *StoreConfiguration {
	_ = "STUB: not implemented"
	// Get statesConfiguration by storeName
	return nil
}

// If statesConfiguration is empty, strategyDefault("none") is provided

func checkKeyIllegal(key string) error {
	_ = "STUB: not implemented"
	// Determine if the key contains daprSeparator
	return nil
}
