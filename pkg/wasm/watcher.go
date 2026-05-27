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
package wasm

import (
	"mosn.io/pkg/utils"

	"mosn.io/mosn/pkg/log"

	"github.com/fsnotify/fsnotify"
)

var (
	watcher *fsnotify.Watcher
	// map[wasm-file-path]config
	configs = make(map[string]*filterConfigItem)
	// map[wasm-file-path]Factory
	factories = make(map[string]*FilterConfigFactory)
)

// Init watcher
func init() {
	var err error
	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		log.DefaultLogger.Errorf("[proxywasm] [watcher] init fail to create watcher: %v", err)
		return
	}
	utils.GoWithRecover(runWatcher, nil)
}

// Watching wasm
func runWatcher() { _ = "STUB: not implemented"; return }

// rewatch the file if it exists
// remove this file then nename other file to this name will cause this case

// Add watching file
func addWatchFile(cfg *filterConfigItem, factory *FilterConfigFactory) {
	_ = "STUB: not implemented"
	return

	// Add starts watching the named file or directory (non-recursively).
}

// remove watching file
func removeWatchFile(cfg *filterConfigItem) { _ = "STUB: not implemented"; return }

// Add starts watching the named file or directory (non-recursively).

// Reload Wasm's configuration file
func reloadWasm(fullPath string) { _ = "STUB: not implemented"; return }

// get WasmPluginWrapper

// register plugin

// Check if the file exists
func fileExist(file string) bool { _ = "STUB: not implemented"; return false }

// Check the file suffix of wasm
func pathIsWasmFile(fullPath string) bool { _ = "STUB: not implemented"; return false }
