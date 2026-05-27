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
	"fmt"
	"os"
)

const (
	storeName = "file_demo"
)

func TestGet(fileName string) { _ = "STUB: not implemented"; return }

func TestPut(fileName string, value string) { _ = "STUB: not implemented"; return }

func TestList(bucketName string) { _ = "STUB: not implemented"; return }

func TestDel(fileName string) { _ = "STUB: not implemented"; return }

func TestStat(fileName string) { _ = "STUB: not implemented"; return }

//here use grpc error code check file exist or not.

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("you can use client like: client put/get/del/list fileName/directryName\n")
		fmt.Println("eg:")
		fmt.Println(" ./main put dir/a.txt aaa")
		fmt.Println(" ./main get dir/a.txt")
		fmt.Println(" ./main list dir/")
		fmt.Println(" ./main del dir/a.txt")
		return
	}
	if os.Args[1] == "put" {
		TestPut(os.Args[2], os.Args[3])
	}
	if os.Args[1] == "get" {
		TestGet(os.Args[2])
	}
	if os.Args[1] == "del" {
		TestDel(os.Args[2])
	}
	if os.Args[1] == "list" {
		TestList(os.Args[2])
	}
	if os.Args[1] == "stat" {
		TestStat(os.Args[2])
	}
}
