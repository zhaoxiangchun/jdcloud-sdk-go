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


type ClusterNetworkSpec struct {

    /* kube-apiserver是否可公网访问，false则kube-apiserver不绑定公网地址，true绑定公网地址  */
    PublicApiServer bool `json:"publicApiServer"`

    /* master网络的cidr  */
    MasterCidr string `json:"masterCidr"`

    /* service网络的cidr  */
    ServiceCidr string `json:"serviceCidr"`

    /* 用户侧承载node和pod的vpc id  */
    VpcId string `json:"vpcId"`

    /* 初始pod的子网id  */
    PodSubnetId string `json:"podSubnetId"`

    /* 初始loadbalancer类型的service所创建的lb所在的subnet  */
    LbSubnetId string `json:"lbSubnetId"`

    /* 初始的node子网ID  */
    NodeSubnetId string `json:"nodeSubnetId"`

    /* nat网关配置  */
    NatGateway NatGatewaySpec `json:"natGateway"`
}