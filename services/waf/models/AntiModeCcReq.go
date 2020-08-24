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


type AntiModeCcReq struct {

    /* WAF实例id  */
    WafInstanceId string `json:"wafInstanceId"`

    /* 域名  */
    Domain string `json:"domain"`

    /* 0表示正常，1表示攻击紧急，2全局模式 3单ip模式 (Optional) */
    CcMode int `json:"ccMode"`

    /* cc qps配置，ccMode 为0/1时， 该字段传0， 表示不可配；  ccMode为2/3时，qps限制[1-20000] (Optional) */
    Qps int `json:"qps"`
}
