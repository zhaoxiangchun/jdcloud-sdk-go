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

type AddEdgeWithCoreRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* IoTCore实例编号  */
    InstanceId string `json:"instanceId"`

    /* Edge名称  */
    EdgeName string `json:"edgeName"`

    /* 硬件平台编号，具体参数参见硬件平台参数文档  */
    Architecture string `json:"architecture"`

    /* 操作系统编号，具体参数参见操作系统参数文档  */
    Os string `json:"os"`

    /* Edge厂家信息  */
    Manufacturer string `json:"manufacturer"`

    /* Edge型号信息  */
    EdgeModel string `json:"edgeModel"`

    /* Edge描述 (Optional) */
    EdgeDesc *string `json:"edgeDesc"`
}

/*
 * param regionId: 地域ID (Required)
 * param instanceId: IoTCore实例编号 (Required)
 * param edgeName: Edge名称 (Required)
 * param architecture: 硬件平台编号，具体参数参见硬件平台参数文档 (Required)
 * param os: 操作系统编号，具体参数参见操作系统参数文档 (Required)
 * param manufacturer: Edge厂家信息 (Required)
 * param edgeModel: Edge型号信息 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewAddEdgeWithCoreRequest(
    regionId string,
    instanceId string,
    edgeName string,
    architecture string,
    os string,
    manufacturer string,
    edgeModel string,
) *AddEdgeWithCoreRequest {

	return &AddEdgeWithCoreRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/instances/{instanceId}/edges/{edgeName}:addEdge",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        InstanceId: instanceId,
        EdgeName: edgeName,
        Architecture: architecture,
        Os: os,
        Manufacturer: manufacturer,
        EdgeModel: edgeModel,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param instanceId: IoTCore实例编号 (Required)
 * param edgeName: Edge名称 (Required)
 * param architecture: 硬件平台编号，具体参数参见硬件平台参数文档 (Required)
 * param os: 操作系统编号，具体参数参见操作系统参数文档 (Required)
 * param manufacturer: Edge厂家信息 (Required)
 * param edgeModel: Edge型号信息 (Required)
 * param edgeDesc: Edge描述 (Optional)
 */
func NewAddEdgeWithCoreRequestWithAllParams(
    regionId string,
    instanceId string,
    edgeName string,
    architecture string,
    os string,
    manufacturer string,
    edgeModel string,
    edgeDesc *string,
) *AddEdgeWithCoreRequest {

    return &AddEdgeWithCoreRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/edges/{edgeName}:addEdge",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        InstanceId: instanceId,
        EdgeName: edgeName,
        Architecture: architecture,
        Os: os,
        Manufacturer: manufacturer,
        EdgeModel: edgeModel,
        EdgeDesc: edgeDesc,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewAddEdgeWithCoreRequestWithoutParam() *AddEdgeWithCoreRequest {

    return &AddEdgeWithCoreRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/instances/{instanceId}/edges/{edgeName}:addEdge",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *AddEdgeWithCoreRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param instanceId: IoTCore实例编号(Required) */
func (r *AddEdgeWithCoreRequest) SetInstanceId(instanceId string) {
    r.InstanceId = instanceId
}

/* param edgeName: Edge名称(Required) */
func (r *AddEdgeWithCoreRequest) SetEdgeName(edgeName string) {
    r.EdgeName = edgeName
}

/* param architecture: 硬件平台编号，具体参数参见硬件平台参数文档(Required) */
func (r *AddEdgeWithCoreRequest) SetArchitecture(architecture string) {
    r.Architecture = architecture
}

/* param os: 操作系统编号，具体参数参见操作系统参数文档(Required) */
func (r *AddEdgeWithCoreRequest) SetOs(os string) {
    r.Os = os
}

/* param manufacturer: Edge厂家信息(Required) */
func (r *AddEdgeWithCoreRequest) SetManufacturer(manufacturer string) {
    r.Manufacturer = manufacturer
}

/* param edgeModel: Edge型号信息(Required) */
func (r *AddEdgeWithCoreRequest) SetEdgeModel(edgeModel string) {
    r.EdgeModel = edgeModel
}

/* param edgeDesc: Edge描述(Optional) */
func (r *AddEdgeWithCoreRequest) SetEdgeDesc(edgeDesc string) {
    r.EdgeDesc = &edgeDesc
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r AddEdgeWithCoreRequest) GetRegionId() string {
    return r.RegionId
}

type AddEdgeWithCoreResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result AddEdgeWithCoreResult `json:"result"`
}

type AddEdgeWithCoreResult struct {
}