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

package oss

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"

	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/jinzhu/copier"
)

var (
	Int64ToTime = copier.TypeConverter{
		SrcType: int64(0),
		DstType: &time.Time{},
		Fn: func(src interface{}) (interface{}, error) {
			s, _ := src.(int64)
			t := time.Unix(s, 0)
			return &t, nil
		},
	}
	TimeToInt64 = copier.TypeConverter{
		SrcType: &time.Time{},
		DstType: int64(0),
		Fn: func(src interface{}) (interface{}, error) {
			s, _ := src.(*time.Time)
			return s.Unix(), nil
		},
	}
	TimeValueToInt64 = copier.TypeConverter{
		SrcType: time.Time{},
		DstType: int64(0),
		Fn: func(src interface{}) (interface{}, error) {
			s, _ := src.(time.Time)
			return s.Unix(), nil
		},
	}
)

func GetGetObjectOutput(ob *s3.GetObjectOutput) (*GetObjectOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetPutObjectOutput(resp *manager.UploadOutput) (*PutObjectOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetDeleteObjectOutput(resp *s3.DeleteObjectOutput) (*DeleteObjectOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetDeleteObjectTaggingOutput(resp *s3.DeleteObjectTaggingOutput) (*DeleteObjectTaggingOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetGetObjectTaggingOutput(resp *s3.GetObjectTaggingOutput) (*GetObjectTaggingOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetCopyObjectOutput(resp *s3.CopyObjectOutput) (*CopyObjectOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetListObjectsOutput(resp *s3.ListObjectsOutput) (*ListObjectsOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if not return NextMarker, use the value of the last Key in the response as the marker

func GetGetObjectCannedAclOutput(resp *s3.GetObjectAclOutput) (*GetObjectCannedAclOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetUploadPartOutput(resp *s3.UploadPartOutput) (*UploadPartOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetUploadPartCopyOutput(resp *s3.UploadPartCopyOutput) (*UploadPartCopyOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetListPartsOutput(resp *s3.ListPartsOutput) (*ListPartsOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetListMultipartUploadsOutput(resp *s3.ListMultipartUploadsOutput) (*ListMultipartUploadsOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetListObjectVersionsOutput(resp *s3.ListObjectVersionsOutput) (*ListObjectVersionsOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
