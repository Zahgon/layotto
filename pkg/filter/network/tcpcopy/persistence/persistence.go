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

package persistence

import (
	"os"
	"sync"

	"mosn.io/layotto/pkg/common"
	"mosn.io/layotto/pkg/filter/network/tcpcopy/model"

	rlog "mosn.io/pkg/log"
)

const (
	dumpBasePath         = "dump"
	tcpcopyDumpFile      = dumpBasePath + string(os.PathSeparator) + "dump_tcp_copy.log"
	memConfDumpFile      = dumpBasePath + string(os.PathSeparator) + "dump_mem_dump.log"
	staticConfDumpFile   = dumpBasePath + string(os.PathSeparator) + "dump_static_conf.log"
	portraitDataDumpFile = dumpBasePath + string(os.PathSeparator) + "dump_portrait_data.log"

	incrementLog = "no_change"
)

type GetLogPath func(fileName string) string

var (
	getLogPath GetLogPath = common.GetLogPath
	//Logger
	tcpcopyPersistence      rlog.ErrorLogger
	memPersistence          rlog.ErrorLogger
	staticConfPersistence   rlog.ErrorLogger
	portraitDataPersistence rlog.ErrorLogger
	//md5 for diff
	md5ValueOfMemDump string
	// md5ValueOfStaticConf string

	memConfDumpFilePath string

	initLoggerOnce sync.Once
)

func getMemConfDumpFilePath() string { _ = "STUB: not implemented"; return "" }

func GetTcpcopyLogger() rlog.ErrorLogger { _ = "STUB: not implemented"; return *new(rlog.ErrorLogger) }

func GetMemLogger() rlog.ErrorLogger { _ = "STUB: not implemented"; return *new(rlog.ErrorLogger) }

func GetStaticConfLogger() rlog.ErrorLogger {
	_ = "STUB: not implemented"
	return *new(rlog.ErrorLogger)
}

func GetPortraitDataLogger() rlog.ErrorLogger {
	_ = "STUB: not implemented"
	return *new(rlog.ErrorLogger)
}

func InitLogger() { _ = "STUB: not implemented"; return }

func doInitLogger() {
	_ = "STUB: not implemented"
	// local variable
	return
}

// write global variable

// init logger using these path variables.

func IsPersistence() bool {
	_ = "STUB: not implemented"
	// Determine the switch state
	return false
}

// Determine whether it is within the sampling period

// Determine whether it is fused

func persistence(config *model.DumpUploadDynamicConfig) {
	_ = "STUB: not implemented"
	// 1.Persist binary data
	return
}

// 2. Persistent user-defined data

// 3. Persistent memory configuration data, only make incremental changes

// 3.1. dump if the data has been changed
