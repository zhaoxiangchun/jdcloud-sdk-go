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

type DescribeVpcRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（queryEdCPSRegions）获取分布式云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 私有网络ID  */
    VpcId string `json:"vpcId"`
}

/*
 * param regionId: 地域ID，可调用接口（queryEdCPSRegions）获取分布式云物理服务器支持的地域 (Required)
 * param vpcId: 私有网络ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeVpcRequest(
    regionId string,
    vpcId string,
) *DescribeVpcRequest {

	return &DescribeVpcRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/vpcs/{vpcId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        VpcId: vpcId,
	}
}

/*
 * param regionId: 地域ID，可调用接口（queryEdCPSRegions）获取分布式云物理服务器支持的地域 (Required)
 * param vpcId: 私有网络ID (Required)
 */
func NewDescribeVpcRequestWithAllParams(
    regionId string,
    vpcId string,
) *DescribeVpcRequest {

    return &DescribeVpcRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vpcs/{vpcId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        VpcId: vpcId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeVpcRequestWithoutParam() *DescribeVpcRequest {

    return &DescribeVpcRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/vpcs/{vpcId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（queryEdCPSRegions）获取分布式云物理服务器支持的地域(Required) */
func (r *DescribeVpcRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param vpcId: 私有网络ID(Required) */
func (r *DescribeVpcRequest) SetVpcId(vpcId string) {
    r.VpcId = vpcId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeVpcRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeVpcResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeVpcResult `json:"result"`
}

type DescribeVpcResult struct {
    Vpc edcps.Vpc `json:"vpc"`
}