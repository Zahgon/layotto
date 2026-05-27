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
	"context"

	dm20151123 "github.com/alibabacloud-go/dm-20151123/v2/client"

	"mosn.io/layotto/components/email"
)

type AliyunEmail struct {
	client *dm20151123.Client
}

func NewAliyunEmail() email.EmailService {
	_ = "STUB: not implemented"
	return *new(email.EmailService)
}

var _ email.EmailService = (*AliyunEmail)(nil)

func (a *AliyunEmail) Init(ctx context.Context, conf *email.Config) error {
	_ = "STUB: not implemented"
	return nil
}

// accessKey ID

// accessKey Secret

// endpoint, ref https://api.aliyun.com/product/Dm

// SendEmail .
func (a *AliyunEmail) SendEmail(ctx context.Context, req *email.SendEmailRequest) (*email.SendEmailResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AccountName is the email send from

// ToAddress is target addresses the email send to

// AddressType = 1: use the email send from
// ref https://help.aliyun.com/document_detail/29444.html

// ReplyToAddress = false: the email no need to reply
// ref https://help.aliyun.com/document_detail/29444.html

func (a *AliyunEmail) checkSendRequest(r *email.SendEmailRequest) bool {
	_ = "STUB: not implemented"
	// make sure content not empty
	return false
}

// SendEmailWithTemplate .
// template must have been applied in aliyun console, and there need the template name
// receivers must have been filled in aliyun console, and there need the receivers list name
func (a *AliyunEmail) SendEmailWithTemplate(ctx context.Context, req *email.SendEmailWithTemplateRequest) (*email.SendEmailWithTemplateResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AccountName is the email send from

// ReceiversName is the name of the recipient list that is created in advance and uploaded with recipients.
// Only take the element with index zero

// AddressType = 1: use the email send from
// ref https://help.aliyun.com/document_detail/29444.html

func (a *AliyunEmail) checkSendWithTemplateRequest(r *email.SendEmailWithTemplateRequest) bool {
	_ = "STUB: not implemented"
	// make sure template exist
	return false
}
