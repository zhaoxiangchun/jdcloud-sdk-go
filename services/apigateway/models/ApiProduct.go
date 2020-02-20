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


type ApiProduct struct {

    /* 分组ID (Optional) */
    ApiGroupId string `json:"apiGroupId"`

    /* 名称 (Optional) */
    GroupName string `json:"groupName"`

    /* 描述 (Optional) */
    Description string `json:"description"`

    /* 版本号 (Optional) */
    Version string `json:"version"`

    /* 区域 (Optional) */
    RegionId string `json:"regionId"`

    /* 计价方式 (Optional) */
    Price string `json:"price"`

    /* 发布日期，格式为毫秒级时间戳 (Optional) */
    DeploymentDate int64 `json:"deploymentDate"`

    /* 购买状态，1已购买，0未购买 (Optional) */
    IsBuyed int `json:"isBuyed"`

    /* 发布日期，格式为毫秒级时间戳 (Optional) */
    BuyDate int64 `json:"buyDate"`
}