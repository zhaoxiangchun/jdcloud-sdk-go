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
    iotcore "github.com/jdcloud-api/jdcloud-sdk-go/services/iotcore/models"
)

type GetDeviceGroupListRequest struct {

    core.JDCloudRequest

    /* 设备归属的实例ID  */
    InstanceId string `json:"instanceId"`

    /* 设备归属的实例所在区域  */
    RegionId string `json:"regionId"`

    /* 获取用户 NULL为当前用户 (Optional) */
    QueryUserPin *string `json:"queryUserPin"`
}

/*
 * param instanceId: 设备归属的实例ID (Required)
 * param regionId: 设备归属的实例所在区域 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetDeviceGroupListRequest(
    instanceId string,
    regionId string,
) *GetDeviceGroupListRequest {

	return &GetDeviceGroupListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/devicegrouplist",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        InstanceId: instanceId,
        RegionId: regionId,
	}
}

/*
 * param instanceId: 设备归属的实例ID (Required)
 * param regionId: 设备归属的实例所在区域 (Required)
 * param queryUserPin: 获取用户 NULL为当前用户 (Optional)
 */
func NewGetDeviceGroupListRequestWithAllParams(
    instanceId string,
    regionId string,
    queryUserPin *string,
) *GetDeviceGroupListRequest {

    return &GetDeviceGroupListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/devicegrouplist",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        InstanceId: instanceId,
        RegionId: regionId,
        QueryUserPin: queryUserPin,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetDeviceGroupListRequestWithoutParam() *GetDeviceGroupListRequest {

    return &GetDeviceGroupListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/devicegrouplist",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param instanceId: 设备归属的实例ID(Required) */
func (r *GetDeviceGroupListRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param regionId: 设备归属的实例所在区域(Required) */
func (r *GetDeviceGroupListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param queryUserPin: 获取用户 NULL为当前用户(Optional) */
func (r *GetDeviceGroupListRequest) SetQueryUserPin(queryUserPin string) {
    r.QueryUserPin = &queryUserPin
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetDeviceGroupListRequest) GetRegionId() string {
    return r.RegionId
}

type GetDeviceGroupListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetDeviceGroupListResult `json:"result"`
}

type GetDeviceGroupListResult struct {
    Data []iotcore.DeviceGroupResp `json:"data"`
}