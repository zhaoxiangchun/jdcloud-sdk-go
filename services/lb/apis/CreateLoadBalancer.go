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

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    charge "github.com/jdcloud-api/jdcloud-sdk-go/services/charge/models"
    vpc "github.com/jdcloud-api/jdcloud-sdk-go/services/vpc/models"
    lb "github.com/jdcloud-api/jdcloud-sdk-go/services/lb/models"
)

type CreateLoadBalancerRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* LoadBalancer的名称,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符  */
    LoadBalancerName string `json:"loadBalancerName"`

    /* LoadBalancer所属子网的Id  */
    SubnetId string `json:"subnetId"`

    /* LoadBalancer的类型，取值：alb、nlb、dnlb，默认为alb (Optional) */
    Type *string `json:"type"`

    /* 【alb，nlb】LoadBalancer所属availability Zone列表,对于alb,nlb是必选参数 <br>【dnlb】全可用区可用，不必传该参数  */
    Azs []string `json:"azs"`

    /* 【alb】支持按用量计费，默认为按用量。【nlb】支持按用量计费。【dnlb】支持按配置计费 (Optional) */
    ChargeSpec *charge.ChargeSpec `json:"chargeSpec"`

    /* 负载均衡关联的弹性IP规格 (Optional) */
    ElasticIp *vpc.ElasticIpSpec `json:"elasticIp"`

    /* 【alb】 安全组 ID列表 (Optional) */
    SecurityGroupIds []string `json:"securityGroupIds"`

    /* LoadBalancer的描述信息,允许输入UTF-8编码下的全部字符，不超过256字符 (Optional) */
    Description *string `json:"description"`

    /* 删除保护，取值为True(开启)或False(关闭)，默认为False (Optional) */
    DeleteProtection *bool `json:"deleteProtection"`

    /* 用户tag 信息 (Optional) */
    UserTags []lb.Tag `json:"userTags"`
}

/*
 * param regionId: Region ID (Required)
 * param loadBalancerName: LoadBalancer的名称,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符 (Required)
 * param subnetId: LoadBalancer所属子网的Id (Required)
 * param azs: 【alb，nlb】LoadBalancer所属availability Zone列表,对于alb,nlb是必选参数 <br>【dnlb】全可用区可用，不必传该参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateLoadBalancerRequest(
    regionId string,
    loadBalancerName string,
    subnetId string,
    azs []string,
) *CreateLoadBalancerRequest {

	return &CreateLoadBalancerRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/loadBalancers/",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        LoadBalancerName: loadBalancerName,
        SubnetId: subnetId,
        Azs: azs,
	}
}

/*
 * param regionId: Region ID (Required)
 * param loadBalancerName: LoadBalancer的名称,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符 (Required)
 * param subnetId: LoadBalancer所属子网的Id (Required)
 * param type_: LoadBalancer的类型，取值：alb、nlb、dnlb，默认为alb (Optional)
 * param azs: 【alb，nlb】LoadBalancer所属availability Zone列表,对于alb,nlb是必选参数 <br>【dnlb】全可用区可用，不必传该参数 (Required)
 * param chargeSpec: 【alb】支持按用量计费，默认为按用量。【nlb】支持按用量计费。【dnlb】支持按配置计费 (Optional)
 * param elasticIp: 负载均衡关联的弹性IP规格 (Optional)
 * param securityGroupIds: 【alb】 安全组 ID列表 (Optional)
 * param description: LoadBalancer的描述信息,允许输入UTF-8编码下的全部字符，不超过256字符 (Optional)
 * param deleteProtection: 删除保护，取值为True(开启)或False(关闭)，默认为False (Optional)
 * param userTags: 用户tag 信息 (Optional)
 */
func NewCreateLoadBalancerRequestWithAllParams(
    regionId string,
    loadBalancerName string,
    subnetId string,
    type_ *string,
    azs []string,
    chargeSpec *charge.ChargeSpec,
    elasticIp *vpc.ElasticIpSpec,
    securityGroupIds []string,
    description *string,
    deleteProtection *bool,
    userTags []lb.Tag,
) *CreateLoadBalancerRequest {

    return &CreateLoadBalancerRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/loadBalancers/",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        LoadBalancerName: loadBalancerName,
        SubnetId: subnetId,
        Type: type_,
        Azs: azs,
        ChargeSpec: chargeSpec,
        ElasticIp: elasticIp,
        SecurityGroupIds: securityGroupIds,
        Description: description,
        DeleteProtection: deleteProtection,
        UserTags: userTags,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateLoadBalancerRequestWithoutParam() *CreateLoadBalancerRequest {

    return &CreateLoadBalancerRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/loadBalancers/",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CreateLoadBalancerRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param loadBalancerName: LoadBalancer的名称,只允许输入中文、数字、大小写字母、英文下划线“_”及中划线“-”，不允许为空且不超过32字符(Required) */
func (r *CreateLoadBalancerRequest) SetLoadBalancerName(loadBalancerName string) {
    r.LoadBalancerName = loadBalancerName
}

/* param subnetId: LoadBalancer所属子网的Id(Required) */
func (r *CreateLoadBalancerRequest) SetSubnetId(subnetId string) {
    r.SubnetId = subnetId
}

/* param type_: LoadBalancer的类型，取值：alb、nlb、dnlb，默认为alb(Optional) */
func (r *CreateLoadBalancerRequest) SetType(type_ string) {
    r.Type = &type_
}

/* param azs: 【alb，nlb】LoadBalancer所属availability Zone列表,对于alb,nlb是必选参数 <br>【dnlb】全可用区可用，不必传该参数(Required) */
func (r *CreateLoadBalancerRequest) SetAzs(azs []string) {
    r.Azs = azs
}

/* param chargeSpec: 【alb】支持按用量计费，默认为按用量。【nlb】支持按用量计费。【dnlb】支持按配置计费(Optional) */
func (r *CreateLoadBalancerRequest) SetChargeSpec(chargeSpec *charge.ChargeSpec) {
    r.ChargeSpec = chargeSpec
}

/* param elasticIp: 负载均衡关联的弹性IP规格(Optional) */
func (r *CreateLoadBalancerRequest) SetElasticIp(elasticIp *vpc.ElasticIpSpec) {
    r.ElasticIp = elasticIp
}

/* param securityGroupIds: 【alb】 安全组 ID列表(Optional) */
func (r *CreateLoadBalancerRequest) SetSecurityGroupIds(securityGroupIds []string) {
    r.SecurityGroupIds = securityGroupIds
}

/* param description: LoadBalancer的描述信息,允许输入UTF-8编码下的全部字符，不超过256字符(Optional) */
func (r *CreateLoadBalancerRequest) SetDescription(description string) {
    r.Description = &description
}

/* param deleteProtection: 删除保护，取值为True(开启)或False(关闭)，默认为False(Optional) */
func (r *CreateLoadBalancerRequest) SetDeleteProtection(deleteProtection bool) {
    r.DeleteProtection = &deleteProtection
}

/* param userTags: 用户tag 信息(Optional) */
func (r *CreateLoadBalancerRequest) SetUserTags(userTags []lb.Tag) {
    r.UserTags = userTags
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateLoadBalancerRequest) GetRegionId() string {
    return r.RegionId
}

type CreateLoadBalancerResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateLoadBalancerResult `json:"result"`
}

type CreateLoadBalancerResult struct {
    LoadBalancerId string `json:"loadBalancerId"`
}