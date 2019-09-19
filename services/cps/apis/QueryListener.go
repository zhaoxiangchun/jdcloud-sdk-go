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
)

type QueryListenerRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（queryCPSLBRegions）获取云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 监听器ID  */
    ListenerId string `json:"listenerId"`
}

/*
 * param regionId: 地域ID，可调用接口（queryCPSLBRegions）获取云物理服务器支持的地域 (Required)
 * param listenerId: 监听器ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryListenerRequest(
    regionId string,
    listenerId string,
) *QueryListenerRequest {

	return &QueryListenerRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/listeners/{listenerId}",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ListenerId: listenerId,
	}
}

/*
 * param regionId: 地域ID，可调用接口（queryCPSLBRegions）获取云物理服务器支持的地域 (Required)
 * param listenerId: 监听器ID (Required)
 */
func NewQueryListenerRequestWithAllParams(
    regionId string,
    listenerId string,
) *QueryListenerRequest {

    return &QueryListenerRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/listeners/{listenerId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ListenerId: listenerId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryListenerRequestWithoutParam() *QueryListenerRequest {

    return &QueryListenerRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/listeners/{listenerId}",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（queryCPSLBRegions）获取云物理服务器支持的地域(Required) */
func (r *QueryListenerRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param listenerId: 监听器ID(Required) */
func (r *QueryListenerRequest) SetListenerId(listenerId string) {
    r.ListenerId = listenerId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryListenerRequest) GetRegionId() string {
    return r.RegionId
}

type QueryListenerResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryListenerResult `json:"result"`
}

type QueryListenerResult struct {
    Listener cps.Listener `json:"listener"`
}