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

type StopPodRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Pod ID  */
    PodId string `json:"podId"`
}

/*
 * param regionId: Region ID (Required)
 * param podId: Pod ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewStopPodRequest(
    regionId string,
    podId string,
) *StopPodRequest {

	return &StopPodRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/pods/{podId}:stopPod",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        PodId: podId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param podId: Pod ID (Required)
 */
func NewStopPodRequestWithAllParams(
    regionId string,
    podId string,
) *StopPodRequest {

    return &StopPodRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/pods/{podId}:stopPod",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PodId: podId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewStopPodRequestWithoutParam() *StopPodRequest {

    return &StopPodRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/pods/{podId}:stopPod",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *StopPodRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param podId: Pod ID(Required) */
func (r *StopPodRequest) SetPodId(podId string) {
    r.PodId = podId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r StopPodRequest) GetRegionId() string {
    return r.RegionId
}

type StopPodResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result StopPodResult `json:"result"`
}

type StopPodResult struct {
}