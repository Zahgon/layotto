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
	"time"

	client "github.com/layotto/go-sdk/client"
)

const (
	appid      = "testApplication_yang"
	group      = "application"
	writeTimes = 4
)

var (
	storeName string
	mode      string
)

func init() {
	flag.StringVar(&storeName, "s", "", "set `storeName`")
	flag.StringVar(&mode, "mode", "raw", "set `mode`")
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

	// Belows are CRUD examples
	// 1. set
	testSet(ctx, cli)

	// 2. get after set
	// Since configuration data might be cached and eventual-consistent,we need to sleep a while before querying new data
	time.Sleep(time.Second * 2)
	testGet(ctx, cli)

	// 3. delete
	testDelete(ctx, cli)

	// 4. get after delete
	//sleep a while before querying deleted data
	time.Sleep(time.Second * 2)
	testGet(ctx, cli)

	// 5. show how to use subscribe API
	// with sdk
	if mode == "sdk" {
		testSubscribeWithSDK(ctx, cli)
	} else {
		// besides sdk,u can also call layotto with grpc
		testSubscribeWithGrpc(ctx)
	}

}

func testSubscribeWithSDK(ctx context.Context, cli client.Client) {
	_ = "STUB: not implemented"
	return
}

// 1. subscribe

// 2. client loop receiving changes in another gorountine

// 3. loop setting configuration in another gorountine

func testSubscribeWithGrpc(ctx context.Context) {
	_ = "STUB: not implemented"
	// 1. connect with grpc
	return
}

// get client for subscribe

// client receive changes

// if it's the last item, break and end this demo

// client send subscribe request

// loop write in another gorountine

func testSet(ctx context.Context, cli client.Client) { _ = "STUB: not implemented"; return }

func testGet(ctx context.Context, cli client.Client) { _ = "STUB: not implemented"; return }

//validate

func testDelete(ctx context.Context, cli client.Client) { _ = "STUB: not implemented"; return }
