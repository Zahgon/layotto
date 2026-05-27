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

package tencentcloud

import (
	"context"
	"errors"

	tcsms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"

	"mosn.io/layotto/components/sms"
)

var (
	ErrTemplateParams = errors.New("error: template parameters should be key-value pairs of the form [idx, param], with idx starting at 0 and ending at the length of parameters - 1")
)

// SmsClient defines the methods of the tencentcloud sms client.
type SmsClient interface {
	SendSmsWithContext(ctx context.Context, request *tcsms.SendSmsRequest) (response *tcsms.SendSmsResponse, err error)
}

// InitConfig is the information required to initialize the sms client
type InitConfig struct {
	SecretId  string
	SecretKey string
	Region    string
}

// NewInitConfig create init config, which will be used in NewSmsClient.
// It checks metadata from sms.Config
// The `accessKeyID` should not be empty.
// The `accessKeySecret` should not be empty.
// The `region` should not be empty.
// Please make sure `region` value is available,
// you can refer https://cloud.tencent.com/document/api/382/52071
func NewInitConfig(config *sms.Config) (*InitConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewSmsClient create tencentcloud sms client instance by InitConfig.
func NewSmsClient(conf *InitConfig) (SmsClient, error) {
	_ = "STUB: not implemented"
	return *new(SmsClient), nil
}

// NewSendSmsRequest create tencentcloud sms request struct by sms.SendSmsWithTemplateRequest.
// The PhoneNumbers should not be empty.
// The Template should not be empty.
// In Template, the TemplateParams should be key-value pairs of the form [idx, param],
// with idx starting at 0 and ending at the length of parameters - 1,
// each [idx, param] key-value pair indicates that the value of the variable at the idx position is param.
// Each parameter and its index must match the variable position of the template corresponding to TemplateId,
// you can refer https://cloud.tencent.com/document/api/382/55981
func NewSendSmsRequest(req *sms.SendSmsWithTemplateRequest) (*tcsms.SendSmsRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// required fields

// optional fields

// ConvertSmsResponse convert tcsms.SendSmsResponse to sms.SendSmsWithTemplateResponse
func ConvertSmsResponse(resp *tcsms.SendSmsResponse) *sms.SendSmsWithTemplateResponse {
	_ = "STUB: not implemented"
	return nil
}

// Sms implemented sms.Sms, is used to send request to tencentcloud sms
type Sms struct {
	client SmsClient
}

// NewSms create empty sms client for tencentcloud
func NewSms() sms.SmsService {
	_ = "STUB: not implemented"

	// Init used to init tencentcloud sms client
	// It checks metadata from sms.Config
	// The `accessKeyID` should not be empty.
	// The `accessKeySecret` should not be empty.
	// The `region` should not be empty.
	// Please make sure `region` value is available,
	// you can refer https://cloud.tencent.com/document/api/382/52071
	return *new(sms.SmsService)
}

func (s *Sms) Init(ctx context.Context, config *sms.Config) error {
	_ = "STUB: not implemented"
	return nil
}

// SendSmsWithTemplate used to send sms with template to tencentcloud sms
// Before calling this method, you should make sure that the Init method is called.
// The PhoneNumbers should not be empty.
// The Template should not be empty.
// In Template, the TemplateParams should be key-value pairs of the form [idx, param],
// with idx starting at 0 and ending at the length of parameters - 1,
// each [idx, param] key-value pair indicates that the value of the variable at the idx position is param.
// Each parameter and its index must match the variable position of the template corresponding to TemplateId,
// you can refer https://cloud.tencent.com/document/api/382/55981
func (s *Sms) SendSmsWithTemplate(ctx context.Context, request *sms.SendSmsWithTemplateRequest) (*sms.SendSmsWithTemplateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Sms) getClient() (SmsClient, error) {
	_ = "STUB: not implemented"
	return *new(SmsClient), nil
}
