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

import charge "github.com/jdcloud-api/jdcloud-sdk-go/services/charge/models"

type Instance struct {

    /* 云物理服务器实例ID (Optional) */
    InstanceId string `json:"instanceId"`

    /* 区域代码, 如 cn-east-1 (Optional) */
    Region string `json:"region"`

    /* 可用区, 如 cn-east-1a (Optional) */
    Az string `json:"az"`

    /* 实例类型, 如 cps.c.normal (Optional) */
    DeviceType string `json:"deviceType"`

    /* 云物理服务器名称 (Optional) */
    Name string `json:"name"`

    /* 云物理服务器描述 (Optional) */
    Description string `json:"description"`

    /* 云物理服务器生命周期状态 (Optional) */
    Status string `json:"status"`

    /* 是否启用外网, 如 yes/no (Optional) */
    EnableInternet string `json:"enableInternet"`

    /* 是否启用IPv6, 如 yes/no (Optional) */
    EnableIpv6 string `json:"enableIpv6"`

    /* 带宽, 单位Mbps (Optional) */
    Bandwidth int `json:"bandwidth"`

    /* 镜像类型, 如 standard (Optional) */
    ImageType string `json:"imageType"`

    /* 操作系统类型ID (Optional) */
    OsTypeId string `json:"osTypeId"`

    /* 操作系统名称 (Optional) */
    OsName string `json:"osName"`

    /* 操作系统类型, 如 ubuntu/centos (Optional) */
    OsType string `json:"osType"`

    /* 操作系统版本, 如 16.04 (Optional) */
    OsVersion string `json:"osVersion"`

    /* 系统盘RAID类型ID (Optional) */
    SysRaidTypeId string `json:"sysRaidTypeId"`

    /* 系统盘RAID类型, 如 NORAID, RAID0, RAID1 (Optional) */
    SysRaidType string `json:"sysRaidType"`

    /* 数据盘RAID类型ID (Optional) */
    DataRaidTypeId string `json:"dataRaidTypeId"`

    /* 数据盘RAID类型, 如 NORAID, RAID0, RAID1 (Optional) */
    DataRaidType string `json:"dataRaidType"`

    /* 网络类型, 如 basic, vpc (Optional) */
    NetworkType string `json:"networkType"`

    /* 私有网络ID (Optional) */
    VpcId string `json:"vpcId"`

    /* 私有网络名称 (Optional) */
    VpcName string `json:"vpcName"`

    /* 子网编号 (Optional) */
    SubnetId string `json:"subnetId"`

    /* 子网名称 (Optional) */
    SubnetName string `json:"subnetName"`

    /* 内网IP (Optional) */
    PrivateIp string `json:"privateIp"`

    /* 外网链路类型, 如 bgp (Optional) */
    LineType string `json:"lineType"`

    /* 弹性公网IPID (Optional) */
    ElasticIpId string `json:"elasticIpId"`

    /* 公网IP (Optional) */
    PublicIp string `json:"publicIp"`

    /* 公网IPv6 (Optional) */
    PublicIpv6 string `json:"publicIpv6"`

    /* 密钥对id (Optional) */
    KeypairId string `json:"keypairId"`

    /* agent状态 (Optional) */
    AgentStatus string `json:"agentStatus"`

    /* 计费信息 (Optional) */
    Charge charge.Charge `json:"charge"`
}
