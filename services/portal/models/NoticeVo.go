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


type NoticeVo struct {

    /* uuid (Optional) */
    Uuid int `json:"uuid"`

    /* 主键id (Optional) */
    Id int `json:"id"`

    /* 标题 (Optional) */
    Title string `json:"title"`

    /* 公告类型; 1:产品公告; 2:域名公告; 3:活动公告; 4:其他公告 (Optional) */
    Type string `json:"type"`

    /* 置顶; 100:不置顶; 1;2;3;4;5:置顶位置(数字不能重复) (Optional) */
    GoTop string `json:"goTop"`

    /* 位置; 0:不显示; 1:左边; 2:左中; 3:中; 4:右中; 5:右 (Optional) */
    Inlet string `json:"inlet"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 修改时间 (Optional) */
    UpdateTime string `json:"updateTime"`

    /* 创建人 (Optional) */
    CreatePin string `json:"createPin"`

    /* 是否失效; 0:生效; 1:失效 (Optional) */
    Yn string `json:"yn"`

    /* 公告内容 (Optional) */
    Content string `json:"content"`

    /* 发送开始时间 (Optional) */
    StartTime string `json:"startTime"`

    /* 发送结束时间 (Optional) */
    EndTime string `json:"endTime"`

    /* 位置; 1:置顶; 2:入口 (Optional) */
    Site int `json:"site"`

    /* 页码数 (Optional) */
    PageNum int `json:"pageNum"`

    /* 页显示数量 (Optional) */
    PageSize int `json:"pageSize"`

    /* 语言 (Optional) */
    Lang string `json:"lang"`

    /* 中英文关联id (Optional) */
    LangId int `json:"langId"`

    /* 查询时间 (Optional) */
    Ts string `json:"ts"`
}
