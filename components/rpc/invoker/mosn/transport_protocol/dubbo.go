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

package transport_protocol

import (
	"mosn.io/api"

	"mosn.io/layotto/components/rpc"
)

// init dubbo protocol
func init() {
	RegistProtocol("dubbo", newDubboProtocol())
}

// newDubboProtocol is create dubbo TransportProtocol
func newDubboProtocol() TransportProtocol {
	_ = "STUB: not implemented"
	return *new(TransportProtocol)
}

type dubboProtocol struct {
	fromFrame
	api.XProtocol
}

func (d *dubboProtocol) Init(map[string]interface{}) error {
	_ = "STUB: not implemented"

	// ToFrame is dubboProtocol transform
	return nil
}

func (d *dubboProtocol) ToFrame(req *rpc.RPCRequest) api.XFrame {
	_ = "STUB: not implemented"
	return *new(api.XFrame)
}

// FromFrame is dubboProtocol transform
func (d *dubboProtocol) FromFrame(resp api.XRespFrame) (*rpc.RPCResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
