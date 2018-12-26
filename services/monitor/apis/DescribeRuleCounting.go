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
    monitor "github.com/jdcloud-api/jdcloud-sdk-go/services/monitor/models"
)

type DescribeRuleCountingRequest struct {

    core.JDCloudRequest

    /* 服务码列表
filter name 为serviceCodes表示查询多个产品线的规则 (Optional) */
    Filters []monitor.Filter `json:"filters"`

    /* pin (Optional) */
    Pin *string `json:"pin"`

    /* 要查询的地域，为空则查询所有的 (Optional) */
    Datacenter *string `json:"datacenter"`

    /* 主帐号 (Optional) */
    AdminPin *string `json:"adminPin"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeRuleCountingRequest(
) *DescribeRuleCountingRequest {

	return &DescribeRuleCountingRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/overview/queryRuleCountings",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param filters: 服务码列表
filter name 为serviceCodes表示查询多个产品线的规则 (Optional)
 * param pin: pin (Optional)
 * param datacenter: 要查询的地域，为空则查询所有的 (Optional)
 * param adminPin: 主帐号 (Optional)
 */
func NewDescribeRuleCountingRequestWithAllParams(
    filters []monitor.Filter,
    pin *string,
    datacenter *string,
    adminPin *string,
) *DescribeRuleCountingRequest {

    return &DescribeRuleCountingRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/overview/queryRuleCountings",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        Filters: filters,
        Pin: pin,
        Datacenter: datacenter,
        AdminPin: adminPin,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeRuleCountingRequestWithoutParam() *DescribeRuleCountingRequest {

    return &DescribeRuleCountingRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/overview/queryRuleCountings",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param filters: 服务码列表
filter name 为serviceCodes表示查询多个产品线的规则(Optional) */
func (r *DescribeRuleCountingRequest) SetFilters(filters []monitor.Filter) {
    r.Filters = filters
}

/* param pin: pin(Optional) */
func (r *DescribeRuleCountingRequest) SetPin(pin string) {
    r.Pin = &pin
}

/* param datacenter: 要查询的地域，为空则查询所有的(Optional) */
func (r *DescribeRuleCountingRequest) SetDatacenter(datacenter string) {
    r.Datacenter = &datacenter
}

/* param adminPin: 主帐号(Optional) */
func (r *DescribeRuleCountingRequest) SetAdminPin(adminPin string) {
    r.AdminPin = &adminPin
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeRuleCountingRequest) GetRegionId() string {
    return ""
}

type DescribeRuleCountingResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeRuleCountingResult `json:"result"`
}

type DescribeRuleCountingResult struct {
    AlarmRuleCount int64 `json:"alarmRuleCount"`
    DisableRuleCount int64 `json:"disableRuleCount"`
    NormalRuleCount int64 `json:"normalRuleCount"`
    ServiceCodes []string `json:"serviceCodes"`
    SubUserPermission bool `json:"subUserPermission"`
    TotalRuleCount int64 `json:"totalRuleCount"`
    UnknownRuleCount int64 `json:"unknownRuleCount"`
}