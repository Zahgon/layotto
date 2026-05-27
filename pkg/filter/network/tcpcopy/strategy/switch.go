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

package strategy

import (
	"sync"

	"mosn.io/layotto/pkg/filter/network/tcpcopy/model"
)

const (
	minInterval = 30
	maxInterval = 60 * 60

	defaultCpuMaxRate = 80
	defaultMemMaxRate = 70

	defaultDuration = 1

	kindOn       = "ON"        // ON
	kindOff      = "OFF"       // OFF
	kindForceOff = "FORCE_OFF" // Forced shutdown
)

var (
	appDumpConfig = &model.DumpConfig{
		Switch:     kindOff,
		Interval:   minInterval,
		Duration:   defaultDuration,
		CpuMaxRate: defaultCpuMaxRate,
		MemMaxRate: defaultMemMaxRate,
	}

	globalDumpConfig = &model.DumpConfig{
		Switch:     kindOff,
		Interval:   minInterval,
		Duration:   defaultDuration,
		CpuMaxRate: defaultCpuMaxRate,
		MemMaxRate: defaultMemMaxRate,
	}

	// switch status
	DumpSwitch = true

	// Sampling Flag, 0 means no sampling, 1 means sampling
	DumpSampleFlag int32

	// cpu fuse threshold
	DumpCpuMaxRate float64 = defaultCpuMaxRate

	// mem fuse threshold
	DumpMemMaxRate float64 = defaultMemMaxRate

	// Dump Interval
	DumpInterval = minInterval

	// Single sampling duration
	DumpDuration = defaultDuration

	// Dump uuid
	DumpSampleUuid = "inituuid"

	// Sampling status of different Business
	DumpBusinessCache = new(sync.Map)

	initOnce = new(sync.Once)
)

// For hot reloading app-level dumpConfig
func UpdateAppDumpConfig(value string) bool { _ = "STUB: not implemented"; return false }

// unmarshal

// validate

// publish config

// For hot reloading global dumpConfig
func UpdateGlobalDumpConfig(value string) bool { _ = "STUB: not implemented"; return false }

// unmarshal

// validate

// publish config

func updateDumpConfig() { _ = "STUB: not implemented"; return }

func isDumpSwitchOpen() bool { _ = "STUB: not implemented"; return false }

func getDumpInterval() int { _ = "STUB: not implemented"; return 0 }

func getDumpDuration() int { _ = "STUB: not implemented"; return 0 }

func getDumpCpuMaxRate() float64 { _ = "STUB: not implemented"; return 0 }

func getDumpMemMaxRate() float64 { _ = "STUB: not implemented"; return 0 }

func updateSampleFlag() {
	_ = "STUB: not implemented"

	// Default sampling interval is 30s
	return
}

// Update the sampling flag

// Continuous sampling

// Send a sampling token (reset the counter) to each business. 0 means the number of samplings in this sampling period is 0, i.e., available for sampling
