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


type ClusterSpec struct {

    /* 集群名称(不能少于6字符不能超过20字符，除下划线外不能包含特殊符号)  */
    Name string `json:"name"`

    /* 集群root用户密码(须包含大小写字母、数字及特殊字符其中三类，且不能少于8字符不能超过30字符)  */
    Password string `json:"password"`

    /* 集群版本，默认版本为JMR2.0.0  */
    Version string `json:"version"`

    /* 集群计费类型，支持按配置和包年包月计费  */
    PayType string `json:"payType"`

    /* 主节点数量  */
    MasterNodeCount int `json:"masterNodeCount"`

    /* Master节点CPU (Optional) */
    MasterCore *int `json:"masterCore"`

    /* Master节点内存(推荐至少8G内存，否则服务可能会部署失败) (Optional) */
    MasterMemory *int `json:"masterMemory"`

    /* Master系统硬盘类型：ssd.gp1,ssd.io1和hdd.std1  */
    MasterSystemDiskType string `json:"masterSystemDiskType"`

    /* Master系统硬盘大小，单位GB  */
    MasterSystemDiskVolume int `json:"masterSystemDiskVolume"`

    /* Master系统硬盘iops，只有在硬盘类型是ssd.gp1,ssd.io1时，才需要有iops，200起步，步长为10  */
    MasterSystemDiskIops int `json:"masterSystemDiskIops"`

    /* Master数据盘类型：ssd.gp1,ssd.io1和hdd.std1  */
    MasterDiskType string `json:"masterDiskType"`

    /* Master数据盘大小，单位GB  */
    MasterDiskVolume int `json:"masterDiskVolume"`

    /* Master数据盘ipos，只有在硬盘类型是ssd.gp1,ssd.io1时，才需要有iops，200起步，步长为10  */
    MasterDiskIops int `json:"masterDiskIops"`

    /* master节点规格  */
    MasterInstanceType string `json:"masterInstanceType"`

    /* Slave节点数量  */
    SlaveNodeCount int `json:"slaveNodeCount"`

    /* Slave节点CPU (Optional) */
    SlaveCore *int `json:"slaveCore"`

    /* Slave节点内存(推荐至少4G内存，否则服务可能会部署失败) (Optional) */
    SlaveMemory *int `json:"slaveMemory"`

    /* Slave系统硬盘类型：ssd.gp1,ssd.io1和hdd.std1  */
    SlaveSystemDiskType string `json:"slaveSystemDiskType"`

    /* Slave系统硬盘大小，单位GB  */
    SlaveSystemDiskVolume int `json:"slaveSystemDiskVolume"`

    /* Slave系统硬盘iops，只有在硬盘类型是ssd.gp1,ssd.io1时，才需要有iops，200起步，步长为10  */
    SlaveSystemDiskIops int `json:"slaveSystemDiskIops"`

    /* Slave数据盘类型：ssd.gp1,ssd.io1和hdd.std1  */
    SlaveDiskType string `json:"slaveDiskType"`

    /* Slave数据盘大小，单位GB  */
    SlaveDiskVolume int `json:"slaveDiskVolume"`

    /* Slave数据盘ipos，只有在硬盘类型是ssd.gp1,ssd.io1时，才需要有iops，200起步，步长为10  */
    SlaveDiskIops int `json:"slaveDiskIops"`

    /* slave节点规格  */
    CoreInstanceType string `json:"coreInstanceType"`

    /* 关联JSS  */
    JssFlag bool `json:"jssFlag"`

    /* 数据中心，即regionId  */
    DataCenter string `json:"dataCenter"`

    /* 软件列表  */
    SoftwareList string `json:"softwareList"`

    /* 集群是否为高可用，默认为高可用集群  */
    HaFlag bool `json:"haFlag"`

    /* Vpc网络ID  */
    VpcId string `json:"vpcId"`

    /* Vpc子网ID  */
    VpcSubnetId string `json:"vpcSubnetId"`

    /* 数据中心的可用区  */
    Az string `json:"az"`
}