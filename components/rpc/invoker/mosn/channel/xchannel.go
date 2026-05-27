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

package channel

import (
	"sync"

	"mosn.io/api"

	"mosn.io/layotto/components/rpc"
	"mosn.io/layotto/components/rpc/invoker/mosn/transport_protocol"
)

// init is regist bolt、boltv2、dubbo channel
func init() {
	RegistChannel("bolt", newXChannel)
	RegistChannel("boltv2", newXChannel)
	RegistChannel("dubbo", newXChannel)
}

// newXChannel is create rpc.Channel by ChannelConfig
func newXChannel(config ChannelConfig) (rpc.Channel, error) {
	_ = "STUB: not implemented"
	return *new(rpc.Channel), nil
}

// dialFunc

// the goroutine model is:
// request goroutine --->  localTcpConn ---> 	mosn
//		^											|
//		|											|
//		|											|
//		|											v
// 		xstate.calls[reqId](a channel) <-- readloop goroutine

// stateFunc

// xstate is record state
type xstate struct {
	reqid uint32
	mu    sync.Mutex
	calls map[uint32]chan call
}

type call struct {
	resp api.XRespFrame
	err  error
}

// xChannel is Channel implement
type xChannel struct {
	proto transport_protocol.TransportProtocol
	pool  *connPool
}

// InvokeWithTargetAddress send request to specific provider address
func (m *xChannel) InvokeWithTargetAddress(req *rpc.RPCRequest) (*rpc.RPCResponse, error) {
	_ = "STUB: not implemented"
	// 1. context.WithTimeout
	return nil, nil
}

// 2. get connection with specific address

// 3. encode request

// 4. set timeout

// 5. read package

// 6. write packet

func (m *xChannel) readResponse(wc *wrapConn, callChan chan<- call) {
	_ = "STUB: not implemented"
	return
}

// read data from connection

func (m *xChannel) Invoke(req *rpc.RPCRequest) (*rpc.RPCResponse, error) {
	_ = "STUB: not implemented"
	// 1. context.WithTimeout
	return nil, nil
}

// 2. get fake connection with mosn

// 3. encode request

// register response channel

// write packet

// if is new conn, send heart

// read response and decode it

// Do is handle RPCRequest to RPCResponse
func (m *xChannel) Do(req *rpc.RPCRequest) (*rpc.RPCResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// removeCall is delete xstate.calls by id
func (m *xChannel) removeCall(xstate *xstate, id uint32) { _ = "STUB: not implemented"; return }

// onData is handle xstate data
func (m *xChannel) onData(conn *wrapConn) error { _ = "STUB: not implemented"; return nil }

// cleanup is clean all xstate.calls
func (m *xChannel) cleanup(c *wrapConn, err error) { _ = "STUB: not implemented"; return }

// cleanup pending calls

func (m *xChannel) sendHeartbeat(c *wrapConn) { _ = "STUB: not implemented"; return }
