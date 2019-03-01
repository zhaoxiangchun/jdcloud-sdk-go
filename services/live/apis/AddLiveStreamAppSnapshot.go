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

type AddLiveStreamAppSnapshotRequest struct {

    core.JDCloudRequest

    /* 直播流所属应用名称  */
    AppName string `json:"appName"`

    /* 您的推流加速域名  */
    PublishDomain string `json:"publishDomain"`

    /* 截图模版  */
    Template string `json:"template"`
}

/*
 * param appName: 直播流所属应用名称 (Required)
 * param publishDomain: 您的推流加速域名 (Required)
 * param template: 截图模版 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddLiveStreamAppSnapshotRequest(
    appName string,
    publishDomain string,
    template string,
) *AddLiveStreamAppSnapshotRequest {

	return &AddLiveStreamAppSnapshotRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/snapshotApps:template",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        AppName: appName,
        PublishDomain: publishDomain,
        Template: template,
	}
}

/*
 * param appName: 直播流所属应用名称 (Required)
 * param publishDomain: 您的推流加速域名 (Required)
 * param template: 截图模版 (Required)
 */
func NewAddLiveStreamAppSnapshotRequestWithAllParams(
    appName string,
    publishDomain string,
    template string,
) *AddLiveStreamAppSnapshotRequest {

    return &AddLiveStreamAppSnapshotRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/snapshotApps:template",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        AppName: appName,
        PublishDomain: publishDomain,
        Template: template,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddLiveStreamAppSnapshotRequestWithoutParam() *AddLiveStreamAppSnapshotRequest {

    return &AddLiveStreamAppSnapshotRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/snapshotApps:template",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param appName: 直播流所属应用名称(Required) */
func (r *AddLiveStreamAppSnapshotRequest) SetAppName(appName string) {
    r.AppName = appName
}

/* param publishDomain: 您的推流加速域名(Required) */
func (r *AddLiveStreamAppSnapshotRequest) SetPublishDomain(publishDomain string) {
    r.PublishDomain = publishDomain
}

/* param template: 截图模版(Required) */
func (r *AddLiveStreamAppSnapshotRequest) SetTemplate(template string) {
    r.Template = template
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddLiveStreamAppSnapshotRequest) GetRegionId() string {
    return ""
}

type AddLiveStreamAppSnapshotResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddLiveStreamAppSnapshotResult `json:"result"`
}

type AddLiveStreamAppSnapshotResult struct {
}