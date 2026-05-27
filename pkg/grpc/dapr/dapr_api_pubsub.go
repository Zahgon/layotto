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

package dapr

import (
	"context"

	"github.com/dapr/components-contrib/pubsub"
	"google.golang.org/protobuf/types/known/emptypb"
	"mosn.io/pkg/log"

	dapr_v1pb "mosn.io/layotto/pkg/grpc/dapr/proto/runtime/v1"
)

const (
	Metadata_key_pubsubName = "pubsubName"
)

type Details struct {
	metadata map[string]string
}

type TopicSubscriptions struct {
	topic2Details map[string]Details
}

func (d *daprGrpcAPI) PublishEvent(ctx context.Context, in *dapr_v1pb.PublishEventRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	// 1. validate
	return nil, nil
}

// doPublishEvent is a protocal irrelevant function to do event publishing.
// It's easy to add APIs for other protocals(e.g. for http api). Just move this func to a separate layer if you need.
func (d *daprGrpcAPI) doPublishEvent(ctx context.Context, pubsubName string, topic string, data []byte, contentType string, metadata map[string]string) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	// 1. validate
	return nil, nil
}

// 2. get component

// 3. new cloudevent request

// 4. publish

// TODO limit topic scope

func (d *daprGrpcAPI) startSubscribing() error {
	_ = "STUB: not implemented"
	// 1. check if there is no need to do it
	return nil
}

// 2. list topics

// return if no need to dosubscription

// 3. loop subscribe

func (d *daprGrpcAPI) getInterestedTopics() (map[string]TopicSubscriptions, error) {
	_ = "STUB: not implemented"
	// 1. check
	return nil, nil
}

// 2. handle app subscriptions

// TODO handle declarative subscriptions

// 3. prepare result

// 4. log

// 5. cache the result

func (d *daprGrpcAPI) beginPubSub(pubsubName string, ps pubsub.PubSub, topicRoutes map[string]TopicSubscriptions) error {
	_ = "STUB: not implemented"
	// 1. call app to find topic topic2Details.
	return nil
}

// 2. loop subscribing every <topic, route>

// TODO limit topic scope

// ask component to subscribe

func (d *daprGrpcAPI) publishMessageGRPC(ctx context.Context, msg *pubsub.NewMessage) error {
	_ = "STUB: not implemented"
	// 1. unmarshal to cloudEvent model
	return nil
}

// 2. drop msg if the current cloud event has expired

// 3. convert request

// set data field

// 4. call appcallback

// 5. check result

func retryStrategy(err error, res *dapr_v1pb.TopicEventResponse, cloudEvent map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// DROP

// on error from application, return error for redelivery of event

// on uninitialized status, this is the case it defaults to as an uninitialized status defaults to 0 which is
// success from protobuf definition

// Consider unknown status field as error and retry

func listTopicSubscriptions(client dapr_v1pb.AppCallbackClient, log log.ErrorLogger) []*dapr_v1pb.TopicSubscription {
	_ = "STUB: not implemented"
	return nil
}
