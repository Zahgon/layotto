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
	"os"
)

const (
	storeName = "oss_demo"
)

// TestGetObjectInput retrieves an object from an S3-compatible object storage service.
// Parameters:
// - bucket: the name of the bucket to which the object belongs.
// - fileName: the name of the object to retrieve
func TestGetObjectInput(bucket, fileName string) {
	_ = "STUB: not implemented"
	// Connect to the object store.
	return
}

// Create a client for the object store.

// Create a request to get an object from the specified bucket.

// Retrieve the object using the client and the request.

// Read the object data into a byte array.

// Convert and print the byte array as a string.

// TestPutObject puts an object into an S3 bucket with the given filename and value.
func TestPutObject(bucket, fileName string, value string) { _ = "STUB: not implemented"; return }

// Check if there's an error in connecting to the server

// Create an S3 object storage service client

// Create a PutObjectInput object with the given parameters

// Send a stream of PutObject requests to the server

// Set the body of the request to the value provided

// Send the request to the server

// Close the stream and wait for the response

// TestListObjects connects to a gRPC service and lists objects in a bucket
// under a specified store, by iterating through the objects in the bucket
// using markers.
func TestListObjects(bucket string) {
	_ = "STUB: not implemented"
	// Connect to the gRPC service.
	return
}

// Create a new ObjectStorageServiceClient.

// Initialize the marker to start at the beginning of the list.

// Iterate through the objects in the bucket using markers.

// Create a new ListObjectsInput request with the store name, bucket name, max keys, and marker.

// Send the ListObjects request to the service and get the response.

// Set the marker to the value of NextMarker in the response.

// If the response is not truncated, print the objects in the bucket and return.

// If the response is truncated, print the objects in the bucket and continue iterating.

func TestDeleteObject(bucket, fileName string) { _ = "STUB: not implemented"; return }

func TestDeleteObjects(bucket, fileName1, fileName2 string) { _ = "STUB: not implemented"; return }

func TestTagging(bucket, name string) { _ = "STUB: not implemented"; return }

func TestAcl(bucket, name string) { _ = "STUB: not implemented"; return }

func TestCopyObject(bucket, name string) { _ = "STUB: not implemented"; return }

func TestPart(bucket, name string) { _ = "STUB: not implemented"; return }

//req4 := &s3.AbortMultipartUploadInput{StoreName: storeName, Bucket: bucket, Key: "海贼王.jpeg", UploadId: "EEE5317D0EB841AC9B80D0B6A2F811AA"}
//resp4, err := c.AbortMultipartUpload(context.Background(), req4)
//if err != nil {
//	fmt.Printf("AbortMultipartUpload fail, err: %+v \n", err)
//	return
//}
//fmt.Printf("AbortMultipartUpload success, resp: %+v \n", resp4)

func TestListVersion(bucket string) { _ = "STUB: not implemented"; return }

func TestRestore(bucket, name string) { _ = "STUB: not implemented"; return }

func TestObjectExist(bucket, name string) { _ = "STUB: not implemented"; return }

func TestAbortMultipartUpload(bucket string) { _ = "STUB: not implemented"; return }

func TestSign(bucket, name, method string) { _ = "STUB: not implemented"; return }

func TestAppend(bucket, fileName, data, position string) { _ = "STUB: not implemented"; return }

func TestHeadObject(bucket, fileName string) { _ = "STUB: not implemented"; return }

func main() {

	if os.Args[1] == "put" {
		TestPutObject(os.Args[2], os.Args[3], os.Args[4])
	}
	if os.Args[1] == "get" {
		TestGetObjectInput(os.Args[2], os.Args[3])
	}
	if os.Args[1] == "del" {
		TestDeleteObject(os.Args[2], os.Args[3])
	}
	if os.Args[1] == "dels" {
		TestDeleteObjects(os.Args[2], os.Args[3], os.Args[4])
	}

	if os.Args[1] == "list" {
		TestListObjects(os.Args[2])
	}

	if os.Args[1] == "tag" {
		TestTagging(os.Args[2], os.Args[3])
	}

	if os.Args[1] == "acl" {
		TestAcl(os.Args[2], os.Args[3])
	}

	if os.Args[1] == "copy" {
		TestCopyObject(os.Args[2], os.Args[3])
	}

	if os.Args[1] == "part" {
		TestPart(os.Args[2], os.Args[3])
	}

	if os.Args[1] == "version" {
		TestListVersion(os.Args[2])
	}

	if os.Args[1] == "restore" {
		TestRestore(os.Args[2], os.Args[3])
	}
	if os.Args[1] == "exist" {
		TestObjectExist(os.Args[2], os.Args[3])
	}
	if os.Args[1] == "abort" {
		TestAbortMultipartUpload(os.Args[2])
	}

	if os.Args[1] == "sign" {
		TestSign(os.Args[2], os.Args[3], os.Args[4])
	}

	if os.Args[1] == "append" {
		TestAppend(os.Args[2], os.Args[3], os.Args[4], os.Args[5])
	}

	if os.Args[1] == "head" {
		TestHeadObject(os.Args[2], os.Args[3])
	}
}
