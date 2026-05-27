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

package nacos

import (
	log "mosn.io/layotto/kit/logger"
)

const (
	DEBUG = "debug"
	INFO  = "info"
	WARN  = "warn"
	ERROR = "error"
)

// An adapter to implement log.LoggerInterface in agollo package.
type DefaultLogger struct {
	logger log.Logger
}

func NewDefaultLogger(logger log.Logger) *DefaultLogger { _ = "STUB: not implemented"; return nil }

func (d *DefaultLogger) Debugf(format string, params ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (d *DefaultLogger) Infof(format string, params ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (d *DefaultLogger) Warnf(format string, params ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (d *DefaultLogger) Errorf(format string, params ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (d *DefaultLogger) Debug(v ...interface{}) { _ = "STUB: not implemented"; return }

func (d *DefaultLogger) Info(v ...interface{}) { _ = "STUB: not implemented"; return }

func (d *DefaultLogger) Warn(v ...interface{}) { _ = "STUB: not implemented"; return }

func (d *DefaultLogger) Error(v ...interface{}) { _ = "STUB: not implemented"; return }
