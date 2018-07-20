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


type Domain struct {

    /* 域名的唯一ID (Optional) */
    Id int `json:"id"`

    /* 域名字符串 (Optional) */
    DomainName string `json:"domainName"`

    /* 创建时间，格式Unix timestamp (Optional) */
    CreateTime int `json:"createTime"`

    /* 过期时间，格式Unix timestamp (Optional) */
    ExpirationDate int `json:"expirationDate"`

    /* 套餐类型，0-免费 1-企业版 2-高级版 (Optional) */
    PackId int `json:"packId"`
}
