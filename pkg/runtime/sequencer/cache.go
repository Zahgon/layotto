// Copyright 2021 Layotto Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package sequencer

import (
	"context"
	"sync"
	"time"

	"mosn.io/layotto/components/sequencer"
)

const defaultSize = 10000
const defaultLimit = 1000
const defaultRetry = 5
const waitTime = time.Second * 2

// DoubleBuffer is double segment id buffer.
// There are two buffers in DoubleBuffer: inUseBuffer is in use, BackUpBuffer is a backup buffer.
// Their default capacity is 1000. When the inUseBuffer usage exceeds 30%, the BackUpBuffer will be initialized.
// When inUseBuffer is used up, swap them.
type DoubleBuffer struct {
	Key              string
	size             int
	inUseBuffer      *Buffer
	backUpBufferChan chan *Buffer
	lock             sync.Mutex
	Store            sequencer.Store
}

type Buffer struct {
	from int64
	to   int64
}

func NewDoubleBuffer(key string, store sequencer.Store) *DoubleBuffer {
	_ = "STUB: not implemented"
	return nil
}

// init double buffer
func (d *DoubleBuffer) init() error {

	buffer, err := d.getNewBuffer()
	if err != nil {
		return err
	}

	d.inUseBuffer = buffer

	return nil
}

// getId next id
func (d *DoubleBuffer) getId() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

//check swap

//when inUseBuffer id more than limit used, initialize BackUpBuffer.
//equal make sure only one thread enter

//quick retry

//slow retry

// swap inUseBuffer and BackUpBuffer, must be locked
func (d *DoubleBuffer) swap() error { _ = "STUB: not implemented"; return nil }

//timeout, return error

// getNewBuffer return a new segment
func (d *DoubleBuffer) getNewBuffer() (*Buffer, error) { _ = "STUB: not implemented"; return nil, nil }

// BufferCatch catch key and buffer
var BufferCatch = map[string]*DoubleBuffer{}

// read/write lock for BufferCatch
var rwLock sync.RWMutex

func GetNextIdFromCache(ctx context.Context, store sequencer.Store, req *sequencer.GetNextIdRequest) (bool, int64, error) {
	_ = "STUB: not implemented"

	// 1. check support
	return false, 0, nil
}

// return if not support

// 2. find the DoubleBuffer for this store and key

// 3. get the next id.
// The buffer should automatically load segment into cache if the cache is (nearly) empty

// get DoubleBuffer using write lock
func getDoubleBufferInWL(key string, store sequencer.Store) (*DoubleBuffer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//double check

// get DoubleBuffer using read lock
func getDoubleBufferInRL(key string) *DoubleBuffer { _ = "STUB: not implemented"; return nil }
