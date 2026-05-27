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

	"github.com/dapr/components-contrib/pubsub"
	"google.golang.org/protobuf/types/known/emptypb"

	"mosn.io/pkg/log"

	runtimev1pb "mosn.io/layotto/spec/proto/runtime/v1"
)

// Publishes events to the specific topic.
func (a *api) PublishEvent(ctx context.Context, in *runtimev1pb.PublishEventRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *api) startSubscribing() error {
	_ = "STUB: not implemented"
	// 1. check if there is no need to do it
	return nil
}

// 2. list topics

// return if no need to dosubscription

// 3. loop subscribe

func (a *api) beginPubSub(pubsubName string, ps pubsub.PubSub, topicRoutes map[string]TopicSubscriptions) error {
	_ = "STUB: not implemented"
	// 1. call app to find topic topic2Details.
	return nil
}

// 2. loop subscribing every <topic, route>

// TODO limit topic scope

// ask component to subscribe

type Details struct {
	metadata map[string]string
}

type TopicSubscriptions struct {
	topic2Details map[string]Details
}

func (a *api) getInterestedTopics() (map[string]TopicSubscriptions, error) {
	_ = "STUB: not implemented"
	// 1. check
	return nil, nil
}

// 2. handle app subscriptions

// TODO handle declarative subscriptions

// 3. prepare result

// 4. log

// 5. cache the result

func (a *api) publishMessageGRPC(ctx context.Context, msg *pubsub.NewMessage) error {
	_ = "STUB: not implemented"

	// TODO tracing
	return nil
}

// Call appcallback

// Check result

// retryStrategy returns error when the message should be redelivered
func retryStrategy(err error, res *runtimev1pb.TopicEventResponse, cloudEvent map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// DROP

// on error from application, return error for redelivery of event

// on uninitialized status, this is the case it defaults to as an uninitialized status defaults to 0 which is
// success from protobuf definition

// Consider unknown status field as error and retry

func listTopicSubscriptions(client runtimev1pb.AppCallbackClient, log log.ErrorLogger) []*runtimev1pb.TopicSubscription {
	_ = "STUB: not implemented"
	return nil
}

func (a *api) envelopeFromSubscriptionMessage(ctx context.Context, msg *pubsub.NewMessage) (*runtimev1pb.TopicEventRequest, map[string]interface{}, error) {
	_ = "STUB: not implemented"
	// 1. Unmarshal to cloudEvent model
	return nil, nil, nil
}

// 2. Drop msg if the current cloud event has expired

// 3. Convert to proto domain struct

// set data field
