// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
)

type CreateSmsTaskUsingPOSTRequest struct {

    core.JDCloudRequest

    /* 应用id  */
    AppId string `json:"appId"`

    /* 短信发送号码集合文件 (Optional) */
    SendNumberUrl *string `json:"sendNumberUrl"`

    /* 短信发送时间，不传表示立即发送 (Optional) */
    SendTime *string `json:"sendTime"`

    /* 短信签名id  */
    SignId string `json:"signId"`

    /* 任务名称  */
    TaskName string `json:"taskName"`

    /* 短信模板id  */
    TemplateId string `json:"templateId"`
}

/*
 * param appId: 应用id (Required)
 * param signId: 短信签名id (Required)
 * param taskName: 任务名称 (Required)
 * param templateId: 短信模板id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateSmsTaskUsingPOSTRequest(
    appId string,
    signId string,
    taskName string,
    templateId string,
) *CreateSmsTaskUsingPOSTRequest {

	return &CreateSmsTaskUsingPOSTRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/smsTasks",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        AppId: appId,
        SignId: signId,
        TaskName: taskName,
        TemplateId: templateId,
	}
}

/*
 * param appId: 应用id (Required)
 * param sendNumberUrl: 短信发送号码集合文件 (Optional)
 * param sendTime: 短信发送时间，不传表示立即发送 (Optional)
 * param signId: 短信签名id (Required)
 * param taskName: 任务名称 (Required)
 * param templateId: 短信模板id (Required)
 */
func NewCreateSmsTaskUsingPOSTRequestWithAllParams(
    appId string,
    sendNumberUrl *string,
    sendTime *string,
    signId string,
    taskName string,
    templateId string,
) *CreateSmsTaskUsingPOSTRequest {

    return &CreateSmsTaskUsingPOSTRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/smsTasks",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        AppId: appId,
        SendNumberUrl: sendNumberUrl,
        SendTime: sendTime,
        SignId: signId,
        TaskName: taskName,
        TemplateId: templateId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateSmsTaskUsingPOSTRequestWithoutParam() *CreateSmsTaskUsingPOSTRequest {

    return &CreateSmsTaskUsingPOSTRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/smsTasks",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param appId: 应用id(Required) */
func (r *CreateSmsTaskUsingPOSTRequest) SetAppId(appId string) {
    r.AppId = appId
}

/* param sendNumberUrl: 短信发送号码集合文件(Optional) */
func (r *CreateSmsTaskUsingPOSTRequest) SetSendNumberUrl(sendNumberUrl string) {
    r.SendNumberUrl = &sendNumberUrl
}

/* param sendTime: 短信发送时间，不传表示立即发送(Optional) */
func (r *CreateSmsTaskUsingPOSTRequest) SetSendTime(sendTime string) {
    r.SendTime = &sendTime
}

/* param signId: 短信签名id(Required) */
func (r *CreateSmsTaskUsingPOSTRequest) SetSignId(signId string) {
    r.SignId = signId
}

/* param taskName: 任务名称(Required) */
func (r *CreateSmsTaskUsingPOSTRequest) SetTaskName(taskName string) {
    r.TaskName = taskName
}

/* param templateId: 短信模板id(Required) */
func (r *CreateSmsTaskUsingPOSTRequest) SetTemplateId(templateId string) {
    r.TemplateId = templateId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateSmsTaskUsingPOSTRequest) GetRegionId() string {
    return ""
}

type CreateSmsTaskUsingPOSTResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateSmsTaskUsingPOSTResult `json:"result"`
}

type CreateSmsTaskUsingPOSTResult struct {
    TaskId string `json:"taskId"`
}