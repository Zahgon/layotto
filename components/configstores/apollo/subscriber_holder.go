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

package apollo

import (
	"sync"

	"mosn.io/layotto/components/configstores"
)

// Holding subscribers' chan and ctx.
type subscriberHolder struct {
	sync.RWMutex
	chanMap map[subscriberKey][]*subscriber
}

func (h *subscriberHolder) findByTopic(namespace string, keyWithLabel string) []*subscriber {
	_ = "STUB: not implemented"
	return nil
}

func (h *subscriberHolder) addByTopic(namespace string, keyWithLabel string, respChan chan *configstores.SubscribeResp) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *subscriberHolder) remove(s *subscriber) {
	_ = "STUB: not implemented"
	// check
	return
}

// find related slice

// find and remove the subscriber

//	remove

func (h *subscriberHolder) reset() { _ = "STUB: not implemented"; return }

type subscriberKey struct {
	//appId        string
	group        string
	keyWithLabel string
}

type subscriber struct {
	respChan      chan *configstores.SubscribeResp
	group         string
	subscriberKey *subscriberKey
	//	TODO add "context" field for canceling
}

func newSubscriberHolder() *subscriberHolder { _ = "STUB: not implemented"; return nil }
