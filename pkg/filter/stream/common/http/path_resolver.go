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

package http

// PathResolver is a util to extract elements in http path.
type PathResolver struct {
	// e.g. /a/b/c/d
	rawPath string
	//unsolved path,which starts by "/". e.g.  /b/c/d
	unresolved string
}

func NewPathResolver(path string) *PathResolver { _ = "STUB: not implemented"; return nil }

func (p *PathResolver) HasNext() bool { _ = "STUB: not implemented"; return false }

func (p *PathResolver) Next() string { _ = "STUB: not implemented"; return "" }

// /a/b/c
// remove first /

// a/b/c
// find first /

// a

// /b/c

func (p *PathResolver) UnresolvedPath() string { _ = "STUB: not implemented"; return "" }
