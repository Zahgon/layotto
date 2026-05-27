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
package common

const (
	TimeoutCode int = iota
	UnavailebleCode
	InternalCode
	InvalidArgsCode
)

type CommonError interface {
	Code() int
	Msg() string
	Error() string
}

type commonError struct {
	code int
	msg  string
}

func (le *commonError) Code() int { _ = "STUB: not implemented"; return 0 }

func (le *commonError) Msg() string { _ = "STUB: not implemented"; return "" }

func (le *commonError) Error() string { _ = "STUB: not implemented"; return "" }

func Error(code int, msg string) CommonError { _ = "STUB: not implemented"; return *new(CommonError) }

func Errorf(code int, format string, a ...interface{}) CommonError {
	_ = "STUB: not implemented"
	return *new(CommonError)
}

func ToGrpcError(err error) error { _ = "STUB: not implemented"; return nil }
