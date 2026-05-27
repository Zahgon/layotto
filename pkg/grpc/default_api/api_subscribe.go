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

package default_api

import (
	"context"
	"sync"

	"github.com/dapr/components-contrib/pubsub"

	runtimev1pb "mosn.io/layotto/spec/proto/runtime/v1"
)

type streamer struct {
	subscribers map[string]*conn
	lock        sync.RWMutex
}

type conn struct {
	lock             sync.RWMutex
	streamLock       sync.Mutex
	stream           runtimev1pb.Runtime_SubscribeTopicEventsServer
	publishResponses map[string]chan *runtimev1pb.SubscribeTopicEventsRequestProcessed
}

// SubscribeTopicEvents is called by the layotto runtime to ad hoc stream
// subscribe to topics. If gRPC API server closes, returns func early with nil
// to close stream.
func (a *api) SubscribeTopicEvents(stream runtimev1pb.Runtime_SubscribeTopicEventsServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *api) streamSubscribe(stream runtimev1pb.Runtime_SubscribeTopicEventsServer, subDone chan struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *streamer) Subscribe(stream runtimev1pb.Runtime_SubscribeTopicEventsServer, req *runtimev1pb.SubscribeTopicEventsRequestInitial) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *api) publishMessageForStream(ctx context.Context, msg *pubsub.NewMessage, pubsubName string) error {
	_ = "STUB: not implemented"
	return nil
}

// 5. Check result

func (c *conn) notifyPublishResponse(ctx context.Context, resp *runtimev1pb.SubscribeTopicEventsRequestProcessed) {
	_ = "STUB: not implemented"
	return
}

func (c *conn) registerPublishResponse(id string) (chan *runtimev1pb.SubscribeTopicEventsRequestProcessed, func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *streamer) StreamerKey(pubsub, topic string) string { _ = "STUB: not implemented"; return "" }
