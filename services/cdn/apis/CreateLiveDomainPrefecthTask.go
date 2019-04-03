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

type CreateLiveDomainPrefecthTaskRequest struct {

    core.JDCloudRequest

    /* 预热的URL (Optional) */
    UrlList []string `json:"urlList"`

    /* 预热时长 (Optional) */
    PrefetchTime *int `json:"prefetchTime"`

    /* 操作类型只能是[start,stop]中的一种 (Optional) */
    Action *string `json:"action"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateLiveDomainPrefecthTaskRequest(
) *CreateLiveDomainPrefecthTaskRequest {

	return &CreateLiveDomainPrefecthTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/task:createLivePrefetchTask",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param urlList: 预热的URL (Optional)
 * param prefetchTime: 预热时长 (Optional)
 * param action: 操作类型只能是[start,stop]中的一种 (Optional)
 */
func NewCreateLiveDomainPrefecthTaskRequestWithAllParams(
    urlList []string,
    prefetchTime *int,
    action *string,
) *CreateLiveDomainPrefecthTaskRequest {

    return &CreateLiveDomainPrefecthTaskRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/task:createLivePrefetchTask",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        UrlList: urlList,
        PrefetchTime: prefetchTime,
        Action: action,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateLiveDomainPrefecthTaskRequestWithoutParam() *CreateLiveDomainPrefecthTaskRequest {

    return &CreateLiveDomainPrefecthTaskRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/task:createLivePrefetchTask",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param urlList: 预热的URL(Optional) */
func (r *CreateLiveDomainPrefecthTaskRequest) SetUrlList(urlList []string) {
    r.UrlList = urlList
}

/* param prefetchTime: 预热时长(Optional) */
func (r *CreateLiveDomainPrefecthTaskRequest) SetPrefetchTime(prefetchTime int) {
    r.PrefetchTime = &prefetchTime
}

/* param action: 操作类型只能是[start,stop]中的一种(Optional) */
func (r *CreateLiveDomainPrefecthTaskRequest) SetAction(action string) {
    r.Action = &action
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateLiveDomainPrefecthTaskRequest) GetRegionId() string {
    return ""
}

type CreateLiveDomainPrefecthTaskResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateLiveDomainPrefecthTaskResult `json:"result"`
}

type CreateLiveDomainPrefecthTaskResult struct {
}