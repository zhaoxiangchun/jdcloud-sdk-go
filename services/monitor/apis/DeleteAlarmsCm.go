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

type DeleteAlarmsCmRequest struct {

    core.JDCloudRequest

    /* region  */
    RegionId string `json:"regionId"`

    /* ids,多个id用|分隔  */
    Ids string `json:"ids"`
}

/*
 * param regionId: region (Required)
 * param ids: ids,多个id用|分隔 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteAlarmsCmRequest(
    regionId string,
    ids string,
) *DeleteAlarmsCmRequest {

	return &DeleteAlarmsCmRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/cm/alarms",
			Method:  "DELETE",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Ids: ids,
	}
}

/*
 * param regionId: region (Required)
 * param ids: ids,多个id用|分隔 (Required)
 */
func NewDeleteAlarmsCmRequestWithAllParams(
    regionId string,
    ids string,
) *DeleteAlarmsCmRequest {

    return &DeleteAlarmsCmRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cm/alarms",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Ids: ids,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteAlarmsCmRequestWithoutParam() *DeleteAlarmsCmRequest {

    return &DeleteAlarmsCmRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cm/alarms",
            Method:  "DELETE",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: region(Required) */
func (r *DeleteAlarmsCmRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param ids: ids,多个id用|分隔(Required) */
func (r *DeleteAlarmsCmRequest) SetIds(ids string) {
    r.Ids = ids
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteAlarmsCmRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteAlarmsCmResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteAlarmsCmResult `json:"result"`
}

type DeleteAlarmsCmResult struct {
}