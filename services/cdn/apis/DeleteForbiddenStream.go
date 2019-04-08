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
    cdn "github.com/jdcloud-api/jdcloud-sdk-go/services/cdn/models"
)

type DeleteForbiddenStreamRequest struct {

    core.JDCloudRequest

    /* 用户域名  */
    Domain string `json:"domain"`

    /* 要删除的禁播流 (Optional) */
    DeleteStreams []cdn.DeleteStream `json:"deleteStreams"`
}

/*
 * param domain: 用户域名 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDeleteForbiddenStreamRequest(
    domain string,
) *DeleteForbiddenStreamRequest {

	return &DeleteForbiddenStreamRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/liveDomain/{domain}/stream:unForbidden",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Domain: domain,
	}
}

/*
 * param domain: 用户域名 (Required)
 * param deleteStreams: 要删除的禁播流 (Optional)
 */
func NewDeleteForbiddenStreamRequestWithAllParams(
    domain string,
    deleteStreams []cdn.DeleteStream,
) *DeleteForbiddenStreamRequest {

    return &DeleteForbiddenStreamRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveDomain/{domain}/stream:unForbidden",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Domain: domain,
        DeleteStreams: deleteStreams,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDeleteForbiddenStreamRequestWithoutParam() *DeleteForbiddenStreamRequest {

    return &DeleteForbiddenStreamRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveDomain/{domain}/stream:unForbidden",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param domain: 用户域名(Required) */
func (r *DeleteForbiddenStreamRequest) SetDomain(domain string) {
    r.Domain = domain
}

/* param deleteStreams: 要删除的禁播流(Optional) */
func (r *DeleteForbiddenStreamRequest) SetDeleteStreams(deleteStreams []cdn.DeleteStream) {
    r.DeleteStreams = deleteStreams
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DeleteForbiddenStreamRequest) GetRegionId() string {
    return ""
}

type DeleteForbiddenStreamResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DeleteForbiddenStreamResult `json:"result"`
}

type DeleteForbiddenStreamResult struct {
}