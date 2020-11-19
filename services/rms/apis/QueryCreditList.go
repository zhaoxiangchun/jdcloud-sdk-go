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
    rms "github.com/jdcloud-api/jdcloud-sdk-go/services/rms/models"
)

type QueryCreditListRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 应用ID  */
    AppId string `json:"appId"`

    /* 第几页  */
    PageNum int `json:"pageNum"`

    /* 每页多少条记录  */
    PageSize int `json:"pageSize"`
}

/*
 * param regionId: Region ID (Required)
 * param appId: 应用ID (Required)
 * param pageNum: 第几页 (Required)
 * param pageSize: 每页多少条记录 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryCreditListRequest(
    regionId string,
    appId string,
    pageNum int,
    pageSize int,
) *QueryCreditListRequest {

	return &QueryCreditListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/queryCreditList",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        AppId: appId,
        PageNum: pageNum,
        PageSize: pageSize,
	}
}

/*
 * param regionId: Region ID (Required)
 * param appId: 应用ID (Required)
 * param pageNum: 第几页 (Required)
 * param pageSize: 每页多少条记录 (Required)
 */
func NewQueryCreditListRequestWithAllParams(
    regionId string,
    appId string,
    pageNum int,
    pageSize int,
) *QueryCreditListRequest {

    return &QueryCreditListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/queryCreditList",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppId: appId,
        PageNum: pageNum,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryCreditListRequestWithoutParam() *QueryCreditListRequest {

    return &QueryCreditListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/queryCreditList",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *QueryCreditListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param appId: 应用ID(Required) */
func (r *QueryCreditListRequest) SetAppId(appId string) {
    r.AppId = appId
}

/* param pageNum: 第几页(Required) */
func (r *QueryCreditListRequest) SetPageNum(pageNum int) {
    r.PageNum = pageNum
}

/* param pageSize: 每页多少条记录(Required) */
func (r *QueryCreditListRequest) SetPageSize(pageSize int) {
    r.PageSize = pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryCreditListRequest) GetRegionId() string {
    return r.RegionId
}

type QueryCreditListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryCreditListResult `json:"result"`
}

type QueryCreditListResult struct {
    Data rms.RespCreditPageResult `json:"data"`
    Status bool `json:"status"`
    Code string `json:"code"`
    Message string `json:"message"`
}