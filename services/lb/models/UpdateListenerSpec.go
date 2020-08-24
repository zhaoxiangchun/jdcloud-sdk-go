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


type UpdateListenerSpec struct {

    /* 监听器名称,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符 (Optional) */
    ListenerName string `json:"listenerName"`

    /* Listener状态, 取值为On或者为Off (Optional) */
    Status string `json:"status"`

    /* 【alb Https和Tls协议】Listener绑定的默认证书，只支持一个证书 (Optional) */
    CertificateSpecs []CertificateSpec `json:"certificateSpecs"`

    /* 【alb、nlb】空闲连接超时时间, 范围为[1,86400]。 <br>（Tcp和Tls协议）默认为：1800s <br>（Http和Https协议）默认为：60s <br>【dnlb】不支持该功能 (Optional) */
    ConnectionIdleTimeSeconds int `json:"connectionIdleTimeSeconds"`

    /* 默认后端服务Id (Optional) */
    BackendId string `json:"backendId"`

    /* 【alb Https和Http协议】转发规则组Id (Optional) */
    UrlMapId string `json:"urlMapId"`

    /* 监听器描述,允许输入UTF-8编码下的全部字符，不超过256字符 (Optional) */
    Description string `json:"description"`
}
