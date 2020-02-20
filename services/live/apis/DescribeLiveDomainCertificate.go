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

type DescribeLiveDomainCertificateRequest struct {

    core.JDCloudRequest

    /* (直播or时移)播放域名
- 仅支持精确匹配
  */
    PlayDomain string `json:"playDomain"`
}

/*
 * param playDomain: (直播or时移)播放域名
- 仅支持精确匹配
 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeLiveDomainCertificateRequest(
    playDomain string,
) *DescribeLiveDomainCertificateRequest {

	return &DescribeLiveDomainCertificateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/liveDomainCertificate",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        PlayDomain: playDomain,
	}
}

/*
 * param playDomain: (直播or时移)播放域名
- 仅支持精确匹配
 (Required)
 */
func NewDescribeLiveDomainCertificateRequestWithAllParams(
    playDomain string,
) *DescribeLiveDomainCertificateRequest {

    return &DescribeLiveDomainCertificateRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveDomainCertificate",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        PlayDomain: playDomain,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeLiveDomainCertificateRequestWithoutParam() *DescribeLiveDomainCertificateRequest {

    return &DescribeLiveDomainCertificateRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/liveDomainCertificate",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param playDomain: (直播or时移)播放域名
- 仅支持精确匹配
(Required) */
func (r *DescribeLiveDomainCertificateRequest) SetPlayDomain(playDomain string) {
    r.PlayDomain = playDomain
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeLiveDomainCertificateRequest) GetRegionId() string {
    return ""
}

type DescribeLiveDomainCertificateResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeLiveDomainCertificateResult `json:"result"`
}

type DescribeLiveDomainCertificateResult struct {
    PlayDomain string `json:"playDomain"`
    CertStatus string `json:"certStatus"`
    Cert string `json:"cert"`
    Title string `json:"title"`
}