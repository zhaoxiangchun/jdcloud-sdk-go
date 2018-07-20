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

type AssociateElasticIpRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* Instance ID  */
    InstanceId string `json:"instanceId"`

    /* 弹性IP ID  */
    ElasticIpId string `json:"elasticIpId"`
}

/*
 * param regionId: Region ID (Required)
 * param instanceId: Instance ID (Required)
 * param elasticIpId: 弹性IP ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAssociateElasticIpRequest(
    regionId string,
    instanceId string,
    elasticIpId string,
) *AssociateElasticIpRequest {

	return &AssociateElasticIpRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}:associateElasticIp",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        ElasticIpId: elasticIpId,
	}
}

/*
 * param regionId: Region ID (Required)
 * param instanceId: Instance ID (Required)
 * param elasticIpId: 弹性IP ID (Required)
 */
func NewAssociateElasticIpRequestWithAllParams(
    regionId string,
    instanceId string,
    elasticIpId string,
) *AssociateElasticIpRequest {

    return &AssociateElasticIpRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:associateElasticIp",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        ElasticIpId: elasticIpId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAssociateElasticIpRequestWithoutParam() *AssociateElasticIpRequest {

    return &AssociateElasticIpRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}:associateElasticIp",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *AssociateElasticIpRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: Instance ID(Required) */
func (r *AssociateElasticIpRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param elasticIpId: 弹性IP ID(Required) */
func (r *AssociateElasticIpRequest) SetElasticIpId(elasticIpId string) {
    r.ElasticIpId = elasticIpId
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AssociateElasticIpRequest) GetRegionId() string {
    return r.RegionId
}

type AssociateElasticIpResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AssociateElasticIpResult `json:"result"`
}

type AssociateElasticIpResult struct {
}