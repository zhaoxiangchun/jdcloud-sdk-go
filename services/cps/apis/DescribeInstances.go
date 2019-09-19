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

type DescribeInstancesRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 页码；默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 分页大小；默认为20；取值范围[20, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 可用区，精确匹配 (Optional) */
    Az *string `json:"az"`

    /* 云物理服务器名称，支持模糊匹配 (Optional) */
    Name *string `json:"name"`

    /* 网络类型，精确匹配，支持basic，vpc (Optional) */
    NetworkType *string `json:"networkType"`

    /* 实例类型，精确匹配，调用接口（describeDeviceTypes）获取实例类型 (Optional) */
    DeviceType *string `json:"deviceType"`

    /* 子网ID (Optional) */
    SubnetId *string `json:"subnetId"`

    /* 密钥对ID (Optional) */
    KeypairId *string `json:"keypairId"`

    /* 是否启用外网, yes/no (Optional) */
    EnableInternet *string `json:"enableInternet"`

    /* instanceId - 云物理服务器ID，精确匹配，支持多个<br/>
privateIp - 云物理服务器内网IP，精确匹配，支持多个<br/>
status - 云物理服务器状态，参考云物理服务器状态，精确匹配，支持多个
 (Optional) */
    Filters []common.Filter `json:"filters"`
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
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
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 * param pageNumber: 页码；默认为1 (Optional)
 * param pageSize: 分页大小；默认为20；取值范围[20, 100] (Optional)
 * param az: 可用区，精确匹配 (Optional)
 * param name: 云物理服务器名称，支持模糊匹配 (Optional)
 * param networkType: 网络类型，精确匹配，支持basic，vpc (Optional)
 * param deviceType: 实例类型，精确匹配，调用接口（describeDeviceTypes）获取实例类型 (Optional)
 * param subnetId: 子网ID (Optional)
 * param keypairId: 密钥对ID (Optional)
 * param enableInternet: 是否启用外网, yes/no (Optional)
 * param filters: instanceId - 云物理服务器ID，精确匹配，支持多个<br/>
privateIp - 云物理服务器内网IP，精确匹配，支持多个<br/>
status - 云物理服务器状态，参考云物理服务器状态，精确匹配，支持多个
 (Optional)
 */
func NewDescribeInstancesRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    az *string,
    name *string,
    networkType *string,
    deviceType *string,
    subnetId *string,
    keypairId *string,
    enableInternet *string,
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
        Az: az,
        Name: name,
        NetworkType: networkType,
        DeviceType: deviceType,
        SubnetId: subnetId,
        KeypairId: keypairId,
        EnableInternet: enableInternet,
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

/* param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域(Required) */
func (r *DescribeInstancesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 页码；默认为1(Optional) */
func (r *DescribeInstancesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 分页大小；默认为20；取值范围[20, 100](Optional) */
func (r *DescribeInstancesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param az: 可用区，精确匹配(Optional) */
func (r *DescribeInstancesRequest) SetAz(az string) {
    r.Az = &az
}

/* param name: 云物理服务器名称，支持模糊匹配(Optional) */
func (r *DescribeInstancesRequest) SetName(name string) {
    r.Name = &name
}

/* param networkType: 网络类型，精确匹配，支持basic，vpc(Optional) */
func (r *DescribeInstancesRequest) SetNetworkType(networkType string) {
    r.NetworkType = &networkType
}

/* param deviceType: 实例类型，精确匹配，调用接口（describeDeviceTypes）获取实例类型(Optional) */
func (r *DescribeInstancesRequest) SetDeviceType(deviceType string) {
    r.DeviceType = &deviceType
}

/* param subnetId: 子网ID(Optional) */
func (r *DescribeInstancesRequest) SetSubnetId(subnetId string) {
    r.SubnetId = &subnetId
}

/* param keypairId: 密钥对ID(Optional) */
func (r *DescribeInstancesRequest) SetKeypairId(keypairId string) {
    r.KeypairId = &keypairId
}

/* param enableInternet: 是否启用外网, yes/no(Optional) */
func (r *DescribeInstancesRequest) SetEnableInternet(enableInternet string) {
    r.EnableInternet = &enableInternet
}

/* param filters: instanceId - 云物理服务器ID，精确匹配，支持多个<br/>
privateIp - 云物理服务器内网IP，精确匹配，支持多个<br/>
status - 云物理服务器状态，参考云物理服务器状态，精确匹配，支持多个
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
    Instances []cps.Instance `json:"instances"`
    PageNumber int `json:"pageNumber"`
    PageSize int `json:"pageSize"`
    TotalCount int `json:"totalCount"`
}