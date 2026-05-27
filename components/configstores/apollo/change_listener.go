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
	"time"

	"github.com/apolloconfig/agollo/v4/storage"

	"mosn.io/layotto/kit/logger"

	"mosn.io/layotto/components/configstores"
)

type changeListener struct {
	subscribers *subscriberHolder
	timeout     time.Duration
	store       RepoForListener
	logger      logger.Logger
}

type RepoForListener interface {
	splitKey(keyWithLabel string) (key string, label string)
	getAllTags(group string, keyWithLabel string) (tags map[string]string, err error)
	GetAppId() string
	GetStoreName() string
}

func newChangeListener(c RepoForListener, log logger.Logger) *changeListener {
	_ = "STUB: not implemented"
	return nil
}

func (lis *changeListener) OnChange(changeEvent *storage.ChangeEvent) {
	_ = "STUB: not implemented"
	// 1. find related subscribers
	return
}

// 2. notice

func (lis *changeListener) OnNewestChange(event *storage.FullChangeEvent) {
	_ = "STUB: not implemented"
	return
}

func (lis *changeListener) notify(s *subscriber, keyWithLabel string, change *storage.ConfigChange) {
	_ = "STUB: not implemented"
	return
}

// 1 recover panic caused when interacting with the chan

// make sure unused chan are all deleted

// 2 prepare response

// TODO add a removed flag in response struct.

//	log and ignore

// 3 write

// 4 close chan if timeout

// remove for gc

func (lis *changeListener) addByTopic(namespace string, keyWithLabel string, respChan chan *configstores.SubscribeResp) error {
	_ = "STUB: not implemented"
	return nil
}

func (lis *changeListener) reset() { _ = "STUB: not implemented"; return }
