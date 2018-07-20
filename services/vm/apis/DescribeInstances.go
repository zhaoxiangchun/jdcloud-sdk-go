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
    vm "github.com/jdcloud-api/jdcloud-sdk-go/services/vm/models"
    common "github.com/jdcloud-api/jdcloud-sdk-go/services/common/models"
)

type DescribeInstancesRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 页码；默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认为20；取值范围[10, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* Tag筛选条件 (Optional) */
    Tags []vm.TagFilter `json:"tags"`

    /* instanceId - 实例ID，精确匹配，支持多个
privateIpAddress - 主网卡IP地址，模糊匹配，支持单个
az - 可用区，精确匹配，支持多个
vpcId - 私有网络ID，精确匹配，支持多个
status - 云主机状态，精确匹配，支持多个
name - 实例名称，模糊匹配，支持单个
imageId - 镜像ID，模糊匹配，支持单个
networkInterfaceId - 弹性网卡ID，精确匹配，支持多个
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: Region ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeInstancesRequest(
    regionId string,
) *DescribeInstancesRequest {

	return &DescribeInstancesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param pageNumber: 页码；默认为1 (Optional)
 * param pageSize: 分页大小；默认为20；取值范围[10, 100] (Optional)
 * param tags: Tag筛选条件 (Optional)
 * param filters: instanceId - 实例ID，精确匹配，支持多个
privateIpAddress - 主网卡IP地址，模糊匹配，支持单个
az - 可用区，精确匹配，支持多个
vpcId - 私有网络ID，精确匹配，支持多个
status - 云主机状态，精确匹配，支持多个
name - 实例名称，模糊匹配，支持单个
imageId - 镜像ID，模糊匹配，支持单个
networkInterfaceId - 弹性网卡ID，精确匹配，支持多个
 (Optional)
 */
func NewDescribeInstancesRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    tags []vm.TagFilter,
    filters []common.Filter,
) *DescribeInstancesRequest {

    return &DescribeInstancesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Tags: tags,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeInstancesRequestWithoutParam() *DescribeInstancesRequest {

    return &DescribeInstancesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DescribeInstancesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 页码；默认为1(Optional) */
func (r *DescribeInstancesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小；默认为20；取值范围[10, 100](Optional) */
func (r *DescribeInstancesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param tags: Tag筛选条件(Optional) */
func (r *DescribeInstancesRequest) SetTags(tags []vm.TagFilter) {
    r.Tags = tags
}

/* param filters: instanceId - 实例ID，精确匹配，支持多个
privateIpAddress - 主网卡IP地址，模糊匹配，支持单个
az - 可用区，精确匹配，支持多个
vpcId - 私有网络ID，精确匹配，支持多个
status - 云主机状态，精确匹配，支持多个
name - 实例名称，模糊匹配，支持单个
imageId - 镜像ID，模糊匹配，支持单个
networkInterfaceId - 弹性网卡ID，精确匹配，支持多个
(Optional) */
func (r *DescribeInstancesRequest) SetFilters(filters []common.Filter) {
    r.Filters = filters
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeInstancesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeInstancesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeInstancesResult `json:"result"`
}

type DescribeInstancesResult struct {
    Instances []vm.Instance `json:"instances"`
    TotalCount int `json:"totalCount"`
}