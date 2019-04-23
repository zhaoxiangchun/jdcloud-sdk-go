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


type InstanceCloudDisk struct {

    /* 云硬盘ID (Optional) */
    DiskId string `json:"diskId"`

    /* 所属AZ (Optional) */
    Az string `json:"az"`

    /* 硬盘名称 (Optional) */
    Name string `json:"name"`

    /* 硬盘描述 (Optional) */
    Description string `json:"description"`

    /* 磁盘类型，取值为 ssd, premium-hdd 之一 (Optional) */
    DiskType string `json:"diskType"`

    /* 磁盘大小（GiB） (Optional) */
    DiskSize int `json:"diskSize"`

    /* 用户指定购买的iops值，目前只支持 ssd.io1 类型云盘 (Optional) */
    Iops int `json:"iops"`

    /* 云硬盘状态，取值为 creating、available、in-use、extending、restoring、deleting、deleted、error_creating、error_deleting、error_restoring、error_extending 之一 (Optional) */
    Status string `json:"status"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`
}
