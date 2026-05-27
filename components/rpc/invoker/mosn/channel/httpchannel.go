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
	"net"

	"mosn.io/pkg/buffer"

	"github.com/valyala/fasthttp"
	// bridge to mosn
	_ "mosn.io/mosn/pkg/stream/http"

	"mosn.io/layotto/components/rpc"
)

// init is regist http channel
func init() {
	RegistChannel("http", newHttpChannel)
}

// hstate is a pipe for readloop goroutine to communicate with request goroutine
type hstate struct {
	// request goroutine will read data from it
	reader net.Conn
	// readloop goroutine will write data to it
	writer net.Conn
}

func (h *hstate) onData(b buffer.IoBuffer) error { _ = "STUB: not implemented"; return nil }

func (h *hstate) close() { _ = "STUB: not implemented"; return }

// httpChannel is Channel implement
type httpChannel struct {
	pool *connPool
}

// newHttpChannel is used to create rpc.Channel according to ChannelConfig
func newHttpChannel(config ChannelConfig) (rpc.Channel, error) {
	_ = "STUB: not implemented"
	return *new(rpc.Channel), nil
}

// dialFunc

// the goroutine model is:
// request goroutine --->  localTcpConn ---> 	mosn
//		^											|
//		|											|
//		|											|
// 		hstate(net.Pipe) <-- readloop goroutine <---

// stateFunc

// hstate is a pipe for readloop goroutine to communicate with request goroutine

// Do is used to handle RPCRequest and return RPCResponse
func (h *httpChannel) Do(req *rpc.RPCRequest) (*rpc.RPCResponse, error) {
	_ = "STUB: not implemented"
	// 1. context.WithTimeout
	return nil, nil
}

// 2. get a fake connection with mosn
// The pool will start a readloop gorountine,
// which aims to read data from mosn and then write data to the hstate.writer

// 3. set deadline before write data to this connection

// 4. write data to this fake connection

// 5. read response data and parse it into fasthttp.Response

// 6. convert result to rpc.RPCResponse,which is the response of rpc invoker

// constructReq is handle rpc.RPCRequest to fasthttp.Request
func (h *httpChannel) constructReq(req *rpc.RPCRequest) *fasthttp.Request {
	_ = "STUB: not implemented"
	return nil
}

func (h *httpChannel) onData(conn *wrapConn) error { _ = "STUB: not implemented"; return nil }

func (h *httpChannel) cleanup(conn *wrapConn, err error) { _ = "STUB: not implemented"; return }
