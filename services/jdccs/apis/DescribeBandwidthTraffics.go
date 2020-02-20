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
    jdccs "github.com/jdcloud-api/jdcloud-sdk-go/services/jdccs/models"
)

type DescribeBandwidthTrafficsRequest struct {

    core.JDCloudRequest

    /* IDC机房ID  */
    Idc string `json:"idc"`

    /* 页码, 默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小，默认为20 (Optional) */
    PageSize *int `json:"pageSize"`
}

/*
 * param idc: IDC机房ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeBandwidthTrafficsRequest(
    idc string,
) *DescribeBandwidthTrafficsRequest {

	return &DescribeBandwidthTrafficsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/idcs/{idc}/bandwidthTraffics",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        Idc: idc,
	}
}

/*
 * param idc: IDC机房ID (Required)
 * param pageNumber: 页码, 默认为1 (Optional)
 * param pageSize: 分页大小，默认为20 (Optional)
 */
func NewDescribeBandwidthTrafficsRequestWithAllParams(
    idc string,
    pageNumber *int,
    pageSize *int,
) *DescribeBandwidthTrafficsRequest {

    return &DescribeBandwidthTrafficsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/idcs/{idc}/bandwidthTraffics",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        Idc: idc,
        PageNumber: pageNumber,
        PageSize: pageSize,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeBandwidthTrafficsRequestWithoutParam() *DescribeBandwidthTrafficsRequest {

    return &DescribeBandwidthTrafficsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/idcs/{idc}/bandwidthTraffics",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param idc: IDC机房ID(Required) */
func (r *DescribeBandwidthTrafficsRequest) SetIdc(idc string) {
    r.Idc = idc
}

/* param pageNumber: 页码, 默认为1(Optional) */
func (r *DescribeBandwidthTrafficsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小，默认为20(Optional) */
func (r *DescribeBandwidthTrafficsRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeBandwidthTrafficsRequest) GetRegionId() string {
    return ""
}

type DescribeBandwidthTrafficsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeBandwidthTrafficsResult `json:"result"`
}

type DescribeBandwidthTrafficsResult struct {
    BandwidthTraffics []jdccs.DescribeBandwidthTraffic `json:"bandwidthTraffics"`
    PageNumber int `json:"pageNumber"`
    PageSize int `json:"pageSize"`
    TotalCount int `json:"totalCount"`
}