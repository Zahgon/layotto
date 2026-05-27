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

// init mosn bolt or boltv2 protocol
func init() {
	RegistProtocol("bolt", newBoltProtocol())
	RegistProtocol("boltv2", newBoltV2Protocol())
}

type boltCommon struct {
	className string
	fromFrame
}

// Init is init boltCommon info
func (b *boltCommon) Init(conf map[string]interface{}) error { _ = "STUB: not implemented"; return nil }

// FromFrame is boltProtocol transform
func (b *boltCommon) FromFrame(resp api.XRespFrame) (*rpc.RPCResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// newBoltProtocol is create boltProtocol
func newBoltProtocol() TransportProtocol { _ = "STUB: not implemented"; return *new(TransportProtocol) }

// boltProtocol is one of TransportProtocol
type boltProtocol struct {
	boltCommon
	api.XProtocol
}

// ToFrame is boltProtocol transform
func (b *boltProtocol) ToFrame(req *rpc.RPCRequest) api.XFrame {
	_ = "STUB: not implemented"
	return *new(api.XFrame)
}

// newBoltV2Protocol is create boltV2Protocol
func newBoltV2Protocol() TransportProtocol {
	_ = "STUB: not implemented"
	return *new(TransportProtocol)
}

// boltv2Protocol is one of TransportProtocol
type boltv2Protocol struct {
	boltCommon
	api.XProtocol
}

// ToFrame is boltv2Protocol transform
func (b *boltv2Protocol) ToFrame(req *rpc.RPCRequest) api.XFrame {
	_ = "STUB: not implemented"
	return *new(api.XFrame)
}

// s2b is convert string to byte slice
func s2b(s string) []byte { _ = "STUB: not implemented"; return nil }
