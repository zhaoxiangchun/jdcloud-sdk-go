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


type SmsSignVO struct {

    /* 应用id (Optional) */
    AppId string `json:"appId"`

    /* 申请状态，1申请中2拒绝3通过 (Optional) */
    ApplyStatus int `json:"applyStatus"`

    /* 创建时间 (Optional) */
    CreateTime string `json:"createTime"`

    /* 用户pin (Optional) */
    Pin string `json:"pin"`

    /* 签名内容 (Optional) */
    SignContent string `json:"signContent"`

    /* 短信签名id (Optional) */
    SignId string `json:"signId"`

    /* 短信签名类型id，调用 listSmsSignTypesUsingGET 接口查询 (Optional) */
    SignTypeId int `json:"signTypeId"`

    /* 短信签名状态，0未启用 1启用 (Optional) */
    Status int `json:"status"`

    /* 短信签名审核说明 (Optional) */
    AuditorExplanation string `json:"auditorExplanation"`

    /* 短信签名用途 0自用1他用 (Optional) */
    SignPurpose int `json:"signPurpose"`

    /* 授权委托书下载地址 (Optional) */
    SignAttorneyUrl string `json:"signAttorneyUrl"`

    /* 证明材料下载地址 (Optional) */
    SignCertificateUrl string `json:"signCertificateUrl"`

    /* 其他证明材料下载地址 (Optional) */
    SignOtherCertificateUrl string `json:"signOtherCertificateUrl"`

    /*  (Optional) */
    UpdateTime string `json:"updateTime"`
}
