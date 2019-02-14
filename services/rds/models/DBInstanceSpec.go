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

type DBInstanceSpec struct {

    /* 实例名，具体规则可参见帮助中心文档:[名称及密码限制](../../../documentation/Cloud-Database-and-Cache/RDS/Introduction/Restrictions/SQLServer-Restrictions.md) (Optional) */
    InstanceName *string `json:"instanceName"`

    /* 实例引擎类型，参见[枚举参数定义](../Enum-Definitions/Enum-Definitions.md)  */
    Engine string `json:"engine"`

    /* 实例引擎版本，参见[枚举参数定义](../Enum-Definitions/Enum-Definitions.md)  */
    EngineVersion string `json:"engineVersion"`

    /* 实例规格代码，可以查看文档[MySQL 实例规格](../Instance-Specifications/Instance-Specifications-MySQL.md)、[SQL Server实例规格](../Instance-Specifications/Instance-Specifications-SQLServer.md)  */
    InstanceClass string `json:"instanceClass"`

    /* 磁盘大小，单位GB  */
    InstanceStorageGB int `json:"instanceStorageGB"`

    /* 可用区ID， 第一个ID必须为主实例所在的可用区。如两个可用区一样，也需输入两个azId  */
    AzId []string `json:"azId"`

    /* VPC的ID  */
    VpcId string `json:"vpcId"`

    /* 子网ID  */
    SubnetId string `json:"subnetId"`

    /* 参数组ID, 缺省系统会创建一个默认参数组<br>- 仅支持MySQL (Optional) */
    ParameterGroup *string `json:"parameterGroup"`

    /* 计费规格，包括计费类型，计费周期等  */
    ChargeSpec *charge.ChargeSpec `json:"chargeSpec"`
}
