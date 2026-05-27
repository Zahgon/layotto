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

package aliyun

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// Prefix is an option to set prefix parameter
func Prefix(value string) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// KeyMarker is an option to set key-marker parameter
func KeyMarker(value string) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// MaxUploads is an option to set max-uploads parameter
func MaxUploads(value int) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// Delimiter is an option to set delimiler parameter
func Delimiter(value string) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// UploadIDMarker is an option to set upload-id-marker parameter
func UploadIDMarker(value string) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// VersionId is an option to set versionId parameter
func VersionId(value string) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// ObjectACL is an option to set X-Oss-Object-Acl header
func ObjectACL(value string) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// CacheControl is an option to set Cache-Control header
func CacheControl(value string) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// ContentEncoding is an option to set Content-Encoding header
func ContentEncoding(value string) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// ACL is an option to set X-Oss-Acl header
func ACL(acl string) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// ContentType is an option to set Content-Type header
func ContentType(value string) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// ContentLength is an option to set Content-Length header
func ContentLength(length int64) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// ContentDisposition is an option to set Content-Disposition header
func ContentDisposition(value string) oss.Option {
	_ = "STUB: not implemented"
	return *new(oss.Option)
}

// SetTagging is an option to set object tagging
func SetTagging(value map[string]string) oss.Option {
	_ = "STUB: not implemented"
	return *new(oss.Option)
}

// ContentLanguage is an option to set Content-Language header
func ContentLanguage(value string) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// ContentMD5 is an option to set Content-MD5 header
func ContentMD5(value string) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// Expires is an option to set Expires header
func Expires(t int64) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// AcceptEncoding is an option to set Accept-Encoding header
func AcceptEncoding(value string) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// IfModifiedSince is an option to set If-Modified-Since header
func IfModifiedSince(t int64) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// IfUnmodifiedSince is an option to set If-Unmodified-Since header
func IfUnmodifiedSince(t int64) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// IfMatch is an option to set If-Match header
func IfMatch(value string) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// IfNoneMatch is an option to set IfNoneMatch header
func IfNoneMatch(value string) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// Range is an option to set Range header, [start, end]
func Range(start, end int64) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// CopySourceIfMatch is an option to set X-Oss-Copy-Source-If-Match header
func CopySourceIfMatch(value string) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// CopySourceIfNoneMatch is an option to set X-Oss-Copy-Source-If-None-Match header
func CopySourceIfNoneMatch(value string) oss.Option {
	_ = "STUB: not implemented"
	return *new(oss.Option)
}

// CopySourceIfModifiedSince is an option to set X-Oss-CopySource-If-Modified-Since header
func CopySourceIfModifiedSince(t int64) oss.Option {
	_ = "STUB: not implemented"
	return *new(oss.Option)
}

// CopySourceIfUnmodifiedSince is an option to set X-Oss-Copy-Source-If-Unmodified-Since header
func CopySourceIfUnmodifiedSince(t int64) oss.Option {
	_ = "STUB: not implemented"
	return *new(oss.Option)
}

// MetadataDirective is an option to set X-Oss-Metadata-Directive header
func MetadataDirective(value string) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// Meta is an option to set Meta header
func Meta(key, value string) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// ServerSideEncryption is an option to set X-Oss-Server-Side-Encryption header
func ServerSideEncryption(value string) oss.Option {
	_ = "STUB: not implemented"
	return *new(oss.Option)
}

// ServerSideEncryptionKeyID is an option to set X-Oss-Server-Side-Encryption-Key-Id header
func ServerSideEncryptionKeyID(value string) oss.Option {
	_ = "STUB: not implemented"
	return *new(oss.Option)
}

// ServerSideDataEncryption is an option to set X-Oss-Server-Side-Data-Encryption header
func ServerSideDataEncryption(value string) oss.Option {
	_ = "STUB: not implemented"
	return *new(oss.Option)
}

// SSECAlgorithm is an option to set X-Oss-Server-Side-Encryption-Customer-Algorithm header
func SSECAlgorithm(value string) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// SSECKey is an option to set X-Oss-Server-Side-Encryption-Customer-Key header
func SSECKey(value string) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// SSECKeyMd5 is an option to set X-Oss-Server-Side-Encryption-Customer-Key-Md5 header
func SSECKeyMd5(value string) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// Origin is an option to set Origin header
func Origin(value string) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// RangeBehavior  is an option to set Range value, such as "standard"
func RangeBehavior(value string) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

func PartHashCtxHeader(value string) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

func PartMd5CtxHeader(value string) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

func PartHashCtxParam(value string) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

func PartMd5CtxParam(value string) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// Marker is an option to set marker parameter
func Marker(value string) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// MaxKeys is an option to set maxkeys parameter
func MaxKeys(value int) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// EncodingType is an option to set encoding-type parameter
func EncodingType(value string) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// VersionIdMarker is an option to set version-id-marker parameter
func VersionIdMarker(value string) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// MaxParts is an option to set max-parts parameter
func MaxParts(value int) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// PartNumberMarker is an option to set part-number-marker parameter
func PartNumberMarker(value int) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// StorageClass bucket storage class
func StorageClass(value string) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// ResponseContentType is an option to set response-content-type param
func ResponseContentType(value string) oss.Option {
	_ = "STUB: not implemented"
	return *new(oss.Option)
}

// ResponseContentLanguage is an option to set response-content-language param
func ResponseContentLanguage(value string) oss.Option {
	_ = "STUB: not implemented"
	return *new(oss.Option)
}

// ResponseExpires is an option to set response-expires param
func ResponseExpires(value string) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// ResponseCacheControl is an option to set response-cache-control param
func ResponseCacheControl(value string) oss.Option {
	_ = "STUB: not implemented"
	return *new(oss.Option)
}

// ResponseContentDisposition is an option to set response-content-disposition param
func ResponseContentDisposition(value string) oss.Option {
	_ = "STUB: not implemented"
	return *new(oss.Option)
}

// ResponseContentEncoding is an option to set response-content-encoding param
func ResponseContentEncoding(value string) oss.Option {
	_ = "STUB: not implemented"
	return *new(oss.Option)
}

// Process is an option to set x-oss-process param
func Process(value string) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// TrafficLimitParam is a option to set x-oss-traffic-limit
func TrafficLimitParam(value int64) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }

// SetHeader Allow users to set personalized http headers
func SetHeader(key string, value interface{}) oss.Option {
	_ = "STUB: not implemented"
	return *new(oss.Option)
}

// AddParam Allow users to set personalized http params
func AddParam(key string, value interface{}) oss.Option {
	_ = "STUB: not implemented"
	return *new(oss.Option)
}

// RequestPayer is an option to set payer who pay for the request
func RequestPayer(value string) oss.Option { _ = "STUB: not implemented"; return *new(oss.Option) }
