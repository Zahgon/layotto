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
	"container/list"
	"context"
	"errors"
	"net"
	"sync"

	"mosn.io/pkg/buffer"
)

const (
	defaultBufSize = 16 * 1024
	maxBufSize     = 512 * 1024
)

var (
	connpoolTimeout = errors.New("connection pool timeout")
)

// wrapConn is wrap connect
type wrapConn struct {
	net.Conn
	buf        buffer.IoBuffer
	state      interface{}
	closed     int32
	cancelCtx  context.Context
	cancelFunc context.CancelFunc
}

// isClose is checked wrapConn close or not
func (w *wrapConn) isClose() bool { _ = "STUB: not implemented"; return false }

// close is real close connect
func (w *wrapConn) close() error { _ = "STUB: not implemented"; return nil }

// newConnPool is reduced the overhead of creating connections and improve program performance
// im-memory fake conn pool
func newConnPool(
	// max active connected count
	maxActive int,
	// create new conn
	dialFunc func() (net.Conn, error),
	// state
	stateFunc func() interface{},
	// handle data
	onDataFunc func(*wrapConn) error,
	// clean connected
	cleanupFunc func(*wrapConn, error)) *connPool {
	_ = "STUB: not implemented"
	return nil
}

// connPool is connected pool
type connPool struct {
	maxActive   int
	dialFunc    func() (net.Conn, error)
	stateFunc   func() interface{}
	onDataFunc  func(*wrapConn) error
	cleanupFunc func(*wrapConn, error)

	sema chan struct{}
	mu   sync.Mutex
	free *list.List
}

// Get is get wrapConn by context.Context
func (p *connPool) Get(ctx context.Context) (*wrapConn, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// get free conn

// create new conn

// start a readloop gorountine to read and handle data

// Put when connected less than maxActive
func (p *connPool) Put(c *wrapConn, close bool) { _ = "STUB: not implemented"; return }

// readloop is loop to read connected then exec onDataFunc
func (p *connPool) readloop(c *wrapConn) { _ = "STUB: not implemented"; return }

// read data from connection

// handle data.
// it will delegate to hstate if it's constructed by httpchannel

func (p *connPool) waitTurn(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (p *connPool) freeTurn() { _ = "STUB: not implemented"; return }
