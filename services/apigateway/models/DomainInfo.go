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


type DomainInfo struct {

    /* api分组id (Optional) */
    ApiGroupId string `json:"apiGroupId"`

    /* 域名id (Optional) */
    DomainId string `json:"domainId"`

    /* 域名 (Optional) */
    Domain string `json:"domain"`

    /* 解析的cname (Optional) */
    Cname string `json:"cname"`

    /* 域名使用的协议 (Optional) */
    Protocol string `json:"protocol"`

    /* 域名创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 域名状态 (Optional) */
    Status string `json:"status"`
}