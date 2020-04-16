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
    baseanti "github.com/jdcloud-api/jdcloud-sdk-go/services/baseanti/models"
)

type DescribeIpResourcesRequest struct {

    core.JDCloudRequest

    /* 地域编码. 基础防护已支持华北-北京, 华东-宿迁, 华东-上海, 华南-广州  */
    RegionId string `json:"regionId"`

    /* IP 模糊匹配 (Optional) */
    Ip *string `json:"ip"`
}

/*
 * param regionId: 地域编码. 基础防护已支持华北-北京, 华东-宿迁, 华东-上海, 华南-广州 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeIpResourcesRequest(
    regionId string,
) *DescribeIpResourcesRequest {

	return &DescribeIpResourcesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/ipResources",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域编码. 基础防护已支持华北-北京, 华东-宿迁, 华东-上海, 华南-广州 (Required)
 * param ip: IP 模糊匹配 (Optional)
 */
func NewDescribeIpResourcesRequestWithAllParams(
    regionId string,
    ip *string,
) *DescribeIpResourcesRequest {

    return &DescribeIpResourcesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/ipResources",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Ip: ip,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeIpResourcesRequestWithoutParam() *DescribeIpResourcesRequest {

    return &DescribeIpResourcesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/ipResources",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域编码. 基础防护已支持华北-北京, 华东-宿迁, 华东-上海, 华南-广州(Required) */
func (r *DescribeIpResourcesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param ip: IP 模糊匹配(Optional) */
func (r *DescribeIpResourcesRequest) SetIp(ip string) {
    r.Ip = &ip
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeIpResourcesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeIpResourcesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeIpResourcesResult `json:"result"`
}

type DescribeIpResourcesResult struct {
    DataList []baseanti.IpResource `json:"dataList"`
    TotalCount int `json:"totalCount"`
}