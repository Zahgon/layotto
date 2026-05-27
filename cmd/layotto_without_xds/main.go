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
	"encoding/json"
	_ "net/http/pprof"
	"os"

	layottoLogger "mosn.io/layotto/kit/logger"
	actuatorLogger "mosn.io/layotto/pkg/actuator/logger"

	// Hello

	// Configuration

	// Pub/Sub

	"github.com/dapr/kit/logger"

	// RPC

	// State Stores

	// Lock

	// Sequencer

	// Actuator
	_ "mosn.io/layotto/pkg/actuator"
	"mosn.io/layotto/pkg/actuator/health"
	actuatorInfo "mosn.io/layotto/pkg/actuator/info"
	_ "mosn.io/layotto/pkg/filter/stream/actuator/http"
	"mosn.io/layotto/pkg/integrate/actuator"

	"github.com/urfave/cli"
	"google.golang.org/grpc"
	_ "mosn.io/mosn/pkg/filter/network/grpc"
	mgrpc "mosn.io/mosn/pkg/filter/network/grpc"
	_ "mosn.io/mosn/pkg/filter/network/proxy"
	_ "mosn.io/mosn/pkg/filter/stream/flowcontrol"
	_ "mosn.io/mosn/pkg/filter/stream/grpcmetric"
	_ "mosn.io/mosn/pkg/metrics/sink"
	_ "mosn.io/mosn/pkg/metrics/sink/prometheus"
	_ "mosn.io/mosn/pkg/network"
	_ "mosn.io/mosn/pkg/stream/http"
	_ "mosn.io/mosn/pkg/wasm/runtime/wasmer"
	_ "mosn.io/pkg/buffer"

	_ "mosn.io/layotto/pkg/filter/network/tcpcopy"
	_ "mosn.io/layotto/pkg/wasm"

	_ "mosn.io/mosn/pkg/filter/listener/originaldst"
	_ "mosn.io/mosn/pkg/filter/network/connectionmanager"
	_ "mosn.io/mosn/pkg/filter/network/streamproxy"
	_ "mosn.io/mosn/pkg/filter/network/tunnel"
	_ "mosn.io/mosn/pkg/filter/stream/dsl"
	_ "mosn.io/mosn/pkg/filter/stream/dubbo"
	_ "mosn.io/mosn/pkg/filter/stream/faultinject"
	_ "mosn.io/mosn/pkg/filter/stream/faulttolerance"
	_ "mosn.io/mosn/pkg/filter/stream/gzip"
	_ "mosn.io/mosn/pkg/filter/stream/headertometadata"
	_ "mosn.io/mosn/pkg/filter/stream/ipaccess"
	_ "mosn.io/mosn/pkg/filter/stream/mirror"
	_ "mosn.io/mosn/pkg/filter/stream/payloadlimit"
	_ "mosn.io/mosn/pkg/filter/stream/proxywasm"
	_ "mosn.io/mosn/pkg/filter/stream/seata"
	_ "mosn.io/mosn/pkg/filter/stream/transcoder/http2bolt"
	_ "mosn.io/mosn/pkg/filter/stream/transcoder/httpconv"
	_ "mosn.io/mosn/pkg/protocol"
	_ "mosn.io/mosn/pkg/protocol/xprotocol"
	_ "mosn.io/mosn/pkg/router"
	_ "mosn.io/mosn/pkg/server/keeper"
	_ "mosn.io/mosn/pkg/stream/http2"
	_ "mosn.io/mosn/pkg/stream/xprotocol"
	_ "mosn.io/mosn/pkg/trace/jaeger"
	_ "mosn.io/mosn/pkg/trace/skywalking"
	_ "mosn.io/mosn/pkg/trace/skywalking/http"
	_ "mosn.io/mosn/pkg/trace/sofa/http"
	_ "mosn.io/mosn/pkg/trace/sofa/xprotocol"
	_ "mosn.io/mosn/pkg/trace/sofa/xprotocol/bolt"
	_ "mosn.io/mosn/pkg/upstream/healthcheck"
	_ "mosn.io/mosn/pkg/upstream/servicediscovery/dubbod"

	_ "mosn.io/layotto/diagnostics/exporter_iml"
)

// loggerForDaprComp is constructed for reusing dapr's components.
var loggerForDaprComp = logger.NewLogger("reuse.dapr.component")

// loggerForLayotto is constructed for layotto.
var loggerForLayotto = layottoLogger.NewLayottoLogger("layotto")

// GitVersion mosn version is specified by latest tag
var GitVersion = ""

func init() {
	mgrpc.RegisterServerHandler("runtime", NewRuntimeGrpcServer)
	// Register default actuator implementations
	actuatorInfo.AddInfoContributor("app", actuator.GetAppContributor())
	actuatorLogger.NewEndpoint()
	health.AddReadinessIndicator("runtime_startup", actuator.GetRuntimeReadinessIndicator())
	health.AddLivenessIndicator("runtime_startup", actuator.GetRuntimeLivenessIndicator())
}

func NewRuntimeGrpcServer(data json.RawMessage, opts ...grpc.ServerOption) (mgrpc.RegisteredServer, error) {
	_ = "STUB: not implemented"
	return *new(mgrpc.RegisteredServer), nil
}

// fail fast if error occurs during startup.
// The reason we panic in a new goroutine is to prevent mosn from recovering.

// 1. parse config

// 2. new instance

// 3. run

// wrap the grpc server with actuator

// register your gRPC API here

// Hello

// Configuration

// RPC

// File

// PubSub

// State

// Lock

// bindings

//OSS

// Cryption

// Sms

// Sequencer

// secretstores

func main() {
	app := newRuntimeApp(&cmdStart)
	registerAppInfo(app)
	_ = app.Run(os.Args)
}

func registerAppInfo(app *cli.App) { _ = "STUB: not implemented"; return }

func newRuntimeApp(startCmd *cli.Command) *cli.App { _ = "STUB: not implemented"; return nil }

// commands

// action
