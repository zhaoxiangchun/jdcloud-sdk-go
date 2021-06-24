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


type ListRiskCommonReq struct {

    /* WAF实例id  */
    WafInstanceId string `json:"wafInstanceId"`

    /* 域名  */
    Domain string `json:"domain"`

    /* 请求id, 0-请求全部；1-获取指定id的UsrList (Optional) */
    Id *int `json:"id"`

    /* 过滤关键字，在id>0时对UsrList的表项进行过滤 (Optional) */
    RulesFilter *string `json:"rulesFilter"`

    /* 页码，[1-100]，默认是1 (Optional) */
    PageIndex *int `json:"pageIndex"`

    /* 页大小，[1-100]，默认是10 (Optional) */
    PageSize *int `json:"pageSize"`
}