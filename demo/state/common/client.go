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

package main

import (
	"context"
	"flag"
	"fmt"

	client "github.com/layotto/go-sdk/client"
)

const (
	key1 = "key1"
	key2 = "key2"
	key3 = "key3"
	key4 = "key4"
	key5 = "key5"
)

var storeName string

func init() {
	flag.StringVar(&storeName, "s", "", "set `storeName`")
}

func main() {
	// parse command arguments
	flag.Parse()
	if storeName == "" {
		panic("storeName is empty.")
	}

	// create a layotto client
	cli, err := client.NewClient()
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	ctx := context.Background()
	value := []byte("hello world")
	fmt.Printf("Start testing %v\n", storeName)

	// Belows are CRUD examples.
	// save state
	testSave(ctx, cli, storeName, key1, value)

	// get state
	testGet(ctx, cli, storeName, key1)

	// SaveBulkState with options and metadata
	testSaveBulkState(ctx, cli, storeName, key1, value, key2)

	keyTostate := testGetBulkState(ctx, cli, storeName, key1, key2)

	// delete state
	testDelete(ctx, cli, storeName, key1, keyTostate[key1].Etag)
	testDelete(ctx, cli, storeName, key2, keyTostate[key2].Etag)
}

func testGetBulkState(ctx context.Context, cli client.Client, store string, key1 string, key2 string) map[string]*client.BulkStateItem {
	_ = "STUB: not implemented"
	return nil
}

func testDelete(ctx context.Context, cli client.Client, store string, key string, etag string) {
	_ = "STUB: not implemented"
	return
}

func testSaveBulkState(ctx context.Context, cli client.Client, store string, key string, value []byte, key2 string) {
	_ = "STUB: not implemented"
	return
}

// etag is used to implement Optimistic Concurrency Control (OCC)
//	see https://docs.dapr.io/developing-applications/building-blocks/state-management/state-management-overview/#concurrency

func testGet(ctx context.Context, cli client.Client, store string, key string) {
	_ = "STUB: not implemented"
	return
}

func testSave(ctx context.Context, cli client.Client, store string, key string, value []byte) {
	_ = "STUB: not implemented"
	return
}
