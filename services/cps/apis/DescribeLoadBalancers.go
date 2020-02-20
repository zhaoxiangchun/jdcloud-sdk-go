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
    cps "github.com/jdcloud-api/jdcloud-sdk-go/services/cps/models"
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribeLoadBalancersRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeCPSLBRegions）获取云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 页码；默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认为20；取值范围[20, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 状态 (Optional) */
    Status *string `json:"status"`

    /* 名称 (Optional) */
    Name *string `json:"name"`

    /* 私有网络ID，精确匹配 (Optional) */
    VpcId *string `json:"vpcId"`

    /* 是否绑定eip (Optional) */
    BindEip *bool `json:"bindEip"`

    /* loadBalancerId - 负载均衡实例ID，精确匹配，支持多个
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: 地域ID，可调用接口（describeCPSLBRegions）获取云物理服务器支持的地域 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeLoadBalancersRequest(
    regionId string,
) *DescribeLoadBalancersRequest {

	return &DescribeLoadBalancersRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/slbs",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeCPSLBRegions）获取云物理服务器支持的地域 (Required)
 * param pageNumber: 页码；默认为1 (Optional)
 * param pageSize: 分页大小；默认为20；取值范围[20, 100] (Optional)
 * param status: 状态 (Optional)
 * param name: 名称 (Optional)
 * param vpcId: 私有网络ID，精确匹配 (Optional)
 * param bindEip: 是否绑定eip (Optional)
 * param filters: loadBalancerId - 负载均衡实例ID，精确匹配，支持多个
 (Optional)
 */
func NewDescribeLoadBalancersRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    status *string,
    name *string,
    vpcId *string,
    bindEip *bool,
    filters []common.Filter,
) *DescribeLoadBalancersRequest {

    return &DescribeLoadBalancersRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/slbs",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Status: status,
        Name: name,
        VpcId: vpcId,
        BindEip: bindEip,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeLoadBalancersRequestWithoutParam() *DescribeLoadBalancersRequest {

    return &DescribeLoadBalancersRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/slbs",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeCPSLBRegions）获取云物理服务器支持的地域(Required) */
func (r *DescribeLoadBalancersRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 页码；默认为1(Optional) */
func (r *DescribeLoadBalancersRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小；默认为20；取值范围[20, 100](Optional) */
func (r *DescribeLoadBalancersRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param status: 状态(Optional) */
func (r *DescribeLoadBalancersRequest) SetStatus(status string) {
    r.Status = &status
}

/* param name: 名称(Optional) */
func (r *DescribeLoadBalancersRequest) SetName(name string) {
    r.Name = &name
}

/* param vpcId: 私有网络ID，精确匹配(Optional) */
func (r *DescribeLoadBalancersRequest) SetVpcId(vpcId string) {
    r.VpcId = &vpcId
}

/* param bindEip: 是否绑定eip(Optional) */
func (r *DescribeLoadBalancersRequest) SetBindEip(bindEip bool) {
    r.BindEip = &bindEip
}

/* param filters: loadBalancerId - 负载均衡实例ID，精确匹配，支持多个
(Optional) */
func (r *DescribeLoadBalancersRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeLoadBalancersRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeLoadBalancersResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeLoadBalancersResult `json:"result"`
}

type DescribeLoadBalancersResult struct {
    LoadBalancers []cps.LoadBalancer `json:"loadBalancers"`
    PageNumber int `json:"pageNumber"`
    PageSize int `json:"pageSize"`
    TotalCount int `json:"totalCount"`
}