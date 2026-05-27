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

package mock_state

import (
	"sync"

	"github.com/dapr/components-contrib/state"
	"github.com/dapr/kit/logger"
)

type inMemStateStoreItem struct {
	data []byte
	etag *string
}

type inMemoryStore struct {
	items map[string]*inMemStateStoreItem
	lock  *sync.RWMutex
	log   logger.Logger
}

func New(logger logger.Logger) state.Store { _ = "STUB: not implemented"; return *new(state.Store) }

func (store *inMemoryStore) newItem(data []byte, etagString *string) *inMemStateStoreItem {
	_ = "STUB: not implemented"
	return nil
}

func (store *inMemoryStore) Init(metadata state.Metadata) error {
	_ = "STUB: not implemented"
	return nil
}

func (store *inMemoryStore) Ping() error { _ = "STUB: not implemented"; return nil }

func (store *inMemoryStore) Features() []state.Feature { _ = "STUB: not implemented"; return nil }

func (store *inMemoryStore) Delete(req *state.DeleteRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (store *inMemoryStore) BulkDelete(req []state.DeleteRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (store *inMemoryStore) Get(req *state.GetRequest) (*state.GetResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (store *inMemoryStore) BulkGet(req []state.GetRequest) (bool, []state.BulkGetResponse, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

func (store *inMemoryStore) Set(req *state.SetRequest) error { _ = "STUB: not implemented"; return nil }

func (store *inMemoryStore) BulkSet(req []state.SetRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (store *inMemoryStore) Multi(request *state.TransactionalStateRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// First we check all eTags

// Now we can perform the operation.

func marshal(value interface{}) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func unmarshal(val interface{}) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
