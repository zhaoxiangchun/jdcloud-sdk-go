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

package models


type EventOut struct {

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 资源ID (Optional) */
    PhysicalId string `json:"physicalId"`

    /* 资源运行操作 (Optional) */
    ResourceAction string `json:"resourceAction"`

    /* 资源名称 (Optional) */
    ResourceName string `json:"resourceName"`

    /* 资源运行状态 (Optional) */
    ResourceStatus string `json:"resourceStatus"`

    /* 资源运行状态原因 (Optional) */
    ResourceStatusReason string `json:"resourceStatusReason"`

    /* 资源类型 (Optional) */
    ResourceType string `json:"resourceType"`

    /* 资源栈ID (Optional) */
    StackId string `json:"stackId"`

    /* 唯一标识 (Optional) */
    Uuid string `json:"uuid"`
}