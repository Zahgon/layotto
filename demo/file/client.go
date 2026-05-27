package main

import (
	"fmt"
	"os"
)

const (
	storeName   = "file_demo"
	storageType = "Standard"
)

func TestGet(fileName string) {
	_ = "STUB: not implemented"
	// Dial to the gRPC server
	return
}

// Create a new client

// Make a request to get a file

// Receive data from the server and store in pic

// Print the result of the GetFile request

func TestPut(fileName string, value string) {
	_ = "STUB: not implemented"
	// Dial to the gRPC server
	return
}

// Create metadata and a new client

// Make a request to put a file

func TestList(bucketName string) {
	_ = "STUB: not implemented"
	// Dial to the gRPC server
	return
}

// Create metadata and a new client

// Make a request to list files

// TestDel deletes a file with the given fileName from the server
func TestDel(fileName string) {
	_ = "STUB: not implemented"
	// Dial a connection to the server
	return
}

// Define metadata to be sent with the request

// Create a new runtime client using the connection

// Create a file request with the given storeName, fileName and metadata

// Create a delete file request using the file request created above

// Send the delete file request to the server and check for errors

// If successful, print a message indicating success

// TestStat retrieves metadata for the file with the given fileName from the server
func TestStat(fileName string) {
	_ = "STUB: not implemented"
	// Dial a connection to the server
	return
}

// Define metadata to be sent with the request

// Create a new runtime client using the connection

// Create a file request with the given storeName, fileName and metadata

// Create a get file metadata request using the file request created above

// Send the get file metadata request to the server and check for errors

// Check if the error returned is a "not found" error and print a message if so

// If it's not a "not found" error and not nil, print an error message

// If successful, print metadata for the file

// CreateBucket function creates a new bucket in the specified S3-compatible object storage service.
// Parameters:
//   - bn (string): the name of the new bucket to be created
func CreateBucket(bn string) {
	_ = "STUB: not implemented"
	// Set the connection parameters for the S3-compatible object storage service.
	return
}

// Initialize minio client object with the connection parameters.

// Make a new bucket with the specified name and location.

// Check to see if we already own this bucket (which happens if you run this twice)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("you can use client like: client put/get/del/list fileName/directryName")
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
	if os.Args[1] == "bucket" {
		CreateBucket(os.Args[2])
	}
}
