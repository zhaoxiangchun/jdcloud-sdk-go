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


type WatermarkTemplate struct {

    /* x轴偏移量
- 单位: 像素
 (Optional) */
    OffSetX int `json:"offSetX"`

    /* y轴偏移量
- 单位: 像素
 (Optional) */
    OffSetY int `json:"offSetY"`

    /* 水印宽度
- 单位: 像素
 (Optional) */
    Width int `json:"width"`

    /* 水印高度
- 单位: 像素
 (Optional) */
    Height int `json:"height"`

    /* 水印地址
 (Optional) */
    Url string `json:"url"`

    /* 水印模板自定义名称
 (Optional) */
    Template string `json:"template"`
}
