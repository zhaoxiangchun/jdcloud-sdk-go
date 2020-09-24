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

type AddBandwidthPackageIpRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* 共享带宽ID  */
    BandwidthPackageId string `json:"bandwidthPackageId"`

    /* 由客户端生成，用于保证请求的幂等性，长度不能超过36个字符；<br/>
如果多个请求使用了相同的clientToken，只会执行第一个请求，之后的请求直接返回第一个请求的结果<br/>
 (Optional) */
    ClientToken *string `json:"clientToken"`

    /* 弹性公网IP ID  */
    ElasticIpIds []string `json:"elasticIpIds"`
}

/*
 * param regionId: 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域 (Required)
 * param bandwidthPackageId: 共享带宽ID (Required)
 * param elasticIpIds: 弹性公网IP ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddBandwidthPackageIpRequest(
    regionId string,
    bandwidthPackageId string,
    elasticIpIds []string,
) *AddBandwidthPackageIpRequest {

	return &AddBandwidthPackageIpRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/bandwidthPackages/{bandwidthPackageId}:addBandwidthPackageIp",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        BandwidthPackageId: bandwidthPackageId,
        ElasticIpIds: elasticIpIds,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域 (Required)
 * param bandwidthPackageId: 共享带宽ID (Required)
 * param clientToken: 由客户端生成，用于保证请求的幂等性，长度不能超过36个字符；<br/>
如果多个请求使用了相同的clientToken，只会执行第一个请求，之后的请求直接返回第一个请求的结果<br/>
 (Optional)
 * param elasticIpIds: 弹性公网IP ID (Required)
 */
func NewAddBandwidthPackageIpRequestWithAllParams(
    regionId string,
    bandwidthPackageId string,
    clientToken *string,
    elasticIpIds []string,
) *AddBandwidthPackageIpRequest {

    return &AddBandwidthPackageIpRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/bandwidthPackages/{bandwidthPackageId}:addBandwidthPackageIp",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        BandwidthPackageId: bandwidthPackageId,
        ClientToken: clientToken,
        ElasticIpIds: elasticIpIds,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddBandwidthPackageIpRequestWithoutParam() *AddBandwidthPackageIpRequest {

    return &AddBandwidthPackageIpRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/bandwidthPackages/{bandwidthPackageId}:addBandwidthPackageIp",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeEdCPSRegions）获取分布式云物理服务器支持的地域(Required) */
func (r *AddBandwidthPackageIpRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param bandwidthPackageId: 共享带宽ID(Required) */
func (r *AddBandwidthPackageIpRequest) SetBandwidthPackageId(bandwidthPackageId string) {
    r.BandwidthPackageId = bandwidthPackageId
}

/* param clientToken: 由客户端生成，用于保证请求的幂等性，长度不能超过36个字符；<br/>
如果多个请求使用了相同的clientToken，只会执行第一个请求，之后的请求直接返回第一个请求的结果<br/>
(Optional) */
func (r *AddBandwidthPackageIpRequest) SetClientToken(clientToken string) {
    r.ClientToken = &clientToken
}

/* param elasticIpIds: 弹性公网IP ID(Required) */
func (r *AddBandwidthPackageIpRequest) SetElasticIpIds(elasticIpIds []string) {
    r.ElasticIpIds = elasticIpIds
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddBandwidthPackageIpRequest) GetRegionId() string {
    return r.RegionId
}

type AddBandwidthPackageIpResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddBandwidthPackageIpResult `json:"result"`
}

type AddBandwidthPackageIpResult struct {
    Success bool `json:"success"`
}