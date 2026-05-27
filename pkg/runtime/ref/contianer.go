/*
 * Copyright 2021 Layotto Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ref

import (
	"github.com/dapr/components-contrib/secretstores"

	"mosn.io/layotto/components/configstores"
)

// RefContainer  hold all secret&config store
type RefContainer struct {
	SecretRef map[string]secretstores.SecretStore
	ConfigRef map[string]configstores.Store
}

// NewRefContainer return a new container
func NewRefContainer() *RefContainer { _ = "STUB: not implemented"; return nil }

func (r *RefContainer) getSecretStore(key string) secretstores.SecretStore {
	_ = "STUB: not implemented"
	return *new(secretstores.SecretStore)
}

func (r *RefContainer) getConfigStore(key string) configstores.Store {
	_ = "STUB: not implemented"
	return *new(configstores.Store)
}
