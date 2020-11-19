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
    ydsms "github.com/jdcloud-api/jdcloud-sdk-go/services/ydsms/models"
)

type GetSmsTaskIdUsingGETRequest struct {

    core.JDCloudRequest

    /* 任务id  */
    TaskId string `json:"taskId"`
}

/*
 * param taskId: 任务id (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetSmsTaskIdUsingGETRequest(
    taskId string,
) *GetSmsTaskIdUsingGETRequest {

	return &GetSmsTaskIdUsingGETRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/smsTasks/{taskId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        TaskId: taskId,
	}
}

/*
 * param taskId: 任务id (Required)
 */
func NewGetSmsTaskIdUsingGETRequestWithAllParams(
    taskId string,
) *GetSmsTaskIdUsingGETRequest {

    return &GetSmsTaskIdUsingGETRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/smsTasks/{taskId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        TaskId: taskId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetSmsTaskIdUsingGETRequestWithoutParam() *GetSmsTaskIdUsingGETRequest {

    return &GetSmsTaskIdUsingGETRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/smsTasks/{taskId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param taskId: 任务id(Required) */
func (r *GetSmsTaskIdUsingGETRequest) SetTaskId(taskId string) {
    r.TaskId = taskId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetSmsTaskIdUsingGETRequest) GetRegionId() string {
    return ""
}

type GetSmsTaskIdUsingGETResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetSmsTaskIdUsingGETResult `json:"result"`
}

type GetSmsTaskIdUsingGETResult struct {
    SmsTaskVO ydsms.SmsTaskVO `json:"smsTaskVO"`
}