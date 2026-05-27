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
	"mosn.io/layotto/components/ref"
)

type DefaultInjector struct {
	Container RefContainer
}

// NewDefaultInjector return a single Inject
func NewDefaultInjector(secretStores map[string]secretstores.SecretStore, configStores map[string]configstores.Store) *DefaultInjector {
	_ = "STUB: not implemented"
	return nil
}

// InjectSecretRef  inject secret to metaData
// TODO: permission control
func (i *DefaultInjector) InjectSecretRef(items []*ref.SecretRefConfig, metaData map[string]string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//avoid part of assign because of err

func (i *DefaultInjector) GetConfigStore(cf *ref.ComponentRefConfig) (configstores.Store, error) {
	_ = "STUB: not implemented"
	return *new(configstores.Store), nil
}

func (i *DefaultInjector) GetSecretStore(cf *ref.ComponentRefConfig) (secretstores.SecretStore, error) {
	_ = "STUB: not implemented"
	return *new(secretstores.SecretStore), nil
}
