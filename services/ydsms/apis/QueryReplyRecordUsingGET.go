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

type QueryReplyRecordUsingGETRequest struct {

    core.JDCloudRequest

    /* 应用id  */
    AppId string `json:"appId"`

    /* 手机号码 (Optional) */
    SendNumber *string `json:"sendNumber"`

    /* 页码 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小 (Optional) */
    PageSize *int `json:"pageSize"`

    /* 结束时间,pattern ="yyyy-MM-dd HH:mm:ss",timezone="GMT+8"  */
    EndTime string `json:"endTime"`

    /* 开始时间,pattern ="yyyy-MM-dd HH:mm:ss",timezone="GMT+8"  */
    StartTime string `json:"startTime"`
}

/*
 * param appId: 应用id (Required)
 * param endTime: 结束时间,pattern ="yyyy-MM-dd HH:mm:ss",timezone="GMT+8" (Required)
 * param startTime: 开始时间,pattern ="yyyy-MM-dd HH:mm:ss",timezone="GMT+8" (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryReplyRecordUsingGETRequest(
    appId string,
    endTime string,
    startTime string,
) *QueryReplyRecordUsingGETRequest {

	return &QueryReplyRecordUsingGETRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/smsApps/{appId}:queryReplyRecord",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        AppId: appId,
        EndTime: endTime,
        StartTime: startTime,
	}
}

/*
 * param appId: 应用id (Required)
 * param sendNumber: 手机号码 (Optional)
 * param pageNumber: 页码 (Optional)
 * param pageSize: 分页大小 (Optional)
 * param endTime: 结束时间,pattern ="yyyy-MM-dd HH:mm:ss",timezone="GMT+8" (Required)
 * param startTime: 开始时间,pattern ="yyyy-MM-dd HH:mm:ss",timezone="GMT+8" (Required)
 */
func NewQueryReplyRecordUsingGETRequestWithAllParams(
    appId string,
    sendNumber *string,
    pageNumber *int,
    pageSize *int,
    endTime string,
    startTime string,
) *QueryReplyRecordUsingGETRequest {

    return &QueryReplyRecordUsingGETRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/smsApps/{appId}:queryReplyRecord",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        AppId: appId,
        SendNumber: sendNumber,
        PageNumber: pageNumber,
        PageSize: pageSize,
        EndTime: endTime,
        StartTime: startTime,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryReplyRecordUsingGETRequestWithoutParam() *QueryReplyRecordUsingGETRequest {

    return &QueryReplyRecordUsingGETRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/smsApps/{appId}:queryReplyRecord",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param appId: 应用id(Required) */
func (r *QueryReplyRecordUsingGETRequest) SetAppId(appId string) {
    r.AppId = appId
}

/* param sendNumber: 手机号码(Optional) */
func (r *QueryReplyRecordUsingGETRequest) SetSendNumber(sendNumber string) {
    r.SendNumber = &sendNumber
}

/* param pageNumber: 页码(Optional) */
func (r *QueryReplyRecordUsingGETRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小(Optional) */
func (r *QueryReplyRecordUsingGETRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param endTime: 结束时间,pattern ="yyyy-MM-dd HH:mm:ss",timezone="GMT+8"(Required) */
func (r *QueryReplyRecordUsingGETRequest) SetEndTime(endTime string) {
    r.EndTime = endTime
}

/* param startTime: 开始时间,pattern ="yyyy-MM-dd HH:mm:ss",timezone="GMT+8"(Required) */
func (r *QueryReplyRecordUsingGETRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryReplyRecordUsingGETRequest) GetRegionId() string {
    return ""
}

type QueryReplyRecordUsingGETResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryReplyRecordUsingGETResult `json:"result"`
}

type QueryReplyRecordUsingGETResult struct {
    SmsApps []ydsms.ReplyRecord `json:"smsApps"`
    TotalCount int64 `json:"totalCount"`
}