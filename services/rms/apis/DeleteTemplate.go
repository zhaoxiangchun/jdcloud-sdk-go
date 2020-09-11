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
    rms "github.com/jdcloud-api/jdcloud-sdk-go/services/rms/models"
)

type DeleteTemplateRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 模板ID  */
    TemplateId string `json:"templateId"`

    /* 应用ID  */
    AppId string `json:"appId"`
}

/*
 * param regionId: Region ID (Required)
 * param templateId: 模板ID (Required)
 * param appId: 应用ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteTemplateRequest(
    regionId string,
    templateId string,
    appId string,
) *DeleteTemplateRequest {

	return &DeleteTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/deleteTemplate",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        TemplateId: templateId,
        AppId: appId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param templateId: 模板ID (Required)
 * param appId: 应用ID (Required)
 */
func NewDeleteTemplateRequestWithAllParams(
    regionId string,
    templateId string,
    appId string,
) *DeleteTemplateRequest {

    return &DeleteTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/deleteTemplate",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        TemplateId: templateId,
        AppId: appId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteTemplateRequestWithoutParam() *DeleteTemplateRequest {

    return &DeleteTemplateRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/deleteTemplate",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *DeleteTemplateRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param templateId: 模板ID(Required) */
func (r *DeleteTemplateRequest) SetTemplateId(templateId string) {
    r.TemplateId = templateId
}

/* param appId: 应用ID(Required) */
func (r *DeleteTemplateRequest) SetAppId(appId string) {
    r.AppId = appId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteTemplateRequest) GetRegionId() string {
    return r.RegionId
}

type DeleteTemplateResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteTemplateResult `json:"result"`
}

type DeleteTemplateResult struct {
    Data rms.RespTemplateData `json:"data"`
    Status bool `json:"status"`
    Code string `json:"code"`
    Message string `json:"message"`
}