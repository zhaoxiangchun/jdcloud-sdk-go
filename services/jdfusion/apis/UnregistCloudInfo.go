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
)

type UnregistCloudInfoRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 云信息ID  */
    CloudId string `json:"cloudId"`
}

/*
 * param regionId: 地域ID (Required)
 * param cloudId: 云信息ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUnregistCloudInfoRequest(
    regionId string,
    cloudId string,
) *UnregistCloudInfoRequest {

	return &UnregistCloudInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/cloud_info/{cloudId}",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        CloudId: cloudId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param cloudId: 云信息ID (Required)
 */
func NewUnregistCloudInfoRequestWithAllParams(
    regionId string,
    cloudId string,
) *UnregistCloudInfoRequest {

    return &UnregistCloudInfoRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cloud_info/{cloudId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        CloudId: cloudId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUnregistCloudInfoRequestWithoutParam() *UnregistCloudInfoRequest {

    return &UnregistCloudInfoRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cloud_info/{cloudId}",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *UnregistCloudInfoRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param cloudId: 云信息ID(Required) */
func (r *UnregistCloudInfoRequest) SetCloudId(cloudId string) {
    r.CloudId = cloudId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UnregistCloudInfoRequest) GetRegionId() string {
    return r.RegionId
}

type UnregistCloudInfoResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UnregistCloudInfoResult `json:"result"`
}

type UnregistCloudInfoResult struct {
}