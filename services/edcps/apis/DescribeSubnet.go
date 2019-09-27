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
    edcps "github.com/jdcloud-api/jdcloud-sdk-go/services/edcps/models"
)

type DescribeSubnetRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（queryEdCPSRegions）获取分布式云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 子网ID  */
    SubnetId string `json:"subnetId"`
}

/*
 * param regionId: 地域ID，可调用接口（queryEdCPSRegions）获取分布式云物理服务器支持的地域 (Required)
 * param subnetId: 子网ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeSubnetRequest(
    regionId string,
    subnetId string,
) *DescribeSubnetRequest {

	return &DescribeSubnetRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/subnets/{subnetId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        SubnetId: subnetId,
	}
}

/*
 * param regionId: 地域ID，可调用接口（queryEdCPSRegions）获取分布式云物理服务器支持的地域 (Required)
 * param subnetId: 子网ID (Required)
 */
func NewDescribeSubnetRequestWithAllParams(
    regionId string,
    subnetId string,
) *DescribeSubnetRequest {

    return &DescribeSubnetRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/subnets/{subnetId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        SubnetId: subnetId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeSubnetRequestWithoutParam() *DescribeSubnetRequest {

    return &DescribeSubnetRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/subnets/{subnetId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（queryEdCPSRegions）获取分布式云物理服务器支持的地域(Required) */
func (r *DescribeSubnetRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param subnetId: 子网ID(Required) */
func (r *DescribeSubnetRequest) SetSubnetId(subnetId string) {
    r.SubnetId = subnetId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeSubnetRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeSubnetResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeSubnetResult `json:"result"`
}

type DescribeSubnetResult struct {
    Subnet edcps.Subnet `json:"subnet"`
}