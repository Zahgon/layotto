package mock_appcallback

import (
	"context"

	"github.com/golang/mock/gomock"
	empty "google.golang.org/protobuf/types/known/emptypb"

	dapr_common_v1pb "mosn.io/layotto/pkg/grpc/dapr/proto/common/v1"
	dapr_v1pb "mosn.io/layotto/pkg/grpc/dapr/proto/runtime/v1"
)

type MockDaprAppCallbackServer struct {
	ctrl     *gomock.Controller
	recorder *MockDaprAppCallbackServerMockRecorder
}

type MockDaprAppCallbackServerMockRecorder struct {
	mock *MockDaprAppCallbackServer
}

func (m *MockDaprAppCallbackServer) OnInvoke(ctx context.Context, in *dapr_common_v1pb.InvokeRequest) (*dapr_common_v1pb.InvokeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MockDaprAppCallbackServer) ListInputBindings(ctx context.Context, in *empty.Empty) (*dapr_v1pb.ListInputBindingsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MockDaprAppCallbackServer) OnBindingEvent(ctx context.Context, in *dapr_v1pb.BindingEventRequest) (*dapr_v1pb.BindingEventResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MockDaprAppCallbackServer) ListTopicSubscriptions(arg0 context.Context, arg1 *empty.Empty) (*dapr_v1pb.ListTopicSubscriptionsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MockDaprAppCallbackServer) OnTopicEvent(ctx context.Context, in *dapr_v1pb.TopicEventRequest) (*dapr_v1pb.TopicEventResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mr *MockDaprAppCallbackServerMockRecorder) OnInvoke(arg0, arg1 interface{}) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (mr *MockDaprAppCallbackServerMockRecorder) ListInputBindings(arg0, arg1 interface{}) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (mr *MockDaprAppCallbackServerMockRecorder) OnBindingEvent(arg0, arg1 interface{}) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (mr *MockDaprAppCallbackServerMockRecorder) ListTopicSubscriptions(arg0, arg1 interface{}) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func (mr *MockDaprAppCallbackServerMockRecorder) OnTopicEvent(arg0, arg1 interface{}) *gomock.Call {
	_ = "STUB: not implemented"
	return nil
}

func NewMockDaprAppCallbackServer(ctrl *gomock.Controller) *MockDaprAppCallbackServer {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockDaprAppCallbackServer) EXPECT() *MockDaprAppCallbackServerMockRecorder {
	_ = "STUB: not implemented"
	return nil
}
