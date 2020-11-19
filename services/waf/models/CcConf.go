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


type CcConf struct {

    /* 是否使能 0表示否 (Optional) */
    Enable int `json:"enable"`

    /* 0表示正常，1表示攻击紧急 (Optional) */
    CcMode int `json:"ccMode"`

    /* qps配置 (Optional) */
    Qps int `json:"qps"`

    /* 是否支持自定义cc，0表示否 (Optional) */
    EnableUserDefine int `json:"enableUserDefine"`

    /* cc自定义规则个数 (Optional) */
    RulesCount int `json:"rulesCount"`
}