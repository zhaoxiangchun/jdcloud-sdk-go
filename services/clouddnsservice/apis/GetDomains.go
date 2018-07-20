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
    clouddnsservice "github.com/jdcloud-api/jdcloud-sdk-go/services/clouddnsservice/models"
)

type GetDomainsRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 当前页数，起始值为1，默认为1  */
    PageNumber int `json:"pageNumber"`

    /* 分页查询时设置的每页行数  */
    PageSize int `json:"pageSize"`

    /* 关键字，按照”%domainName%”模式搜索主域名 (Optional) */
    DomainName *string `json:"domainName"`
}

/*
 * param regionId: Region ID (Required)
 * param pageNumber: 当前页数，起始值为1，默认为1 (Required)
 * param pageSize: 分页查询时设置的每页行数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetDomainsRequest(
    regionId string,
    pageNumber int,
    pageSize int,
) *GetDomainsRequest {

	return &GetDomainsRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/domain",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
	}
}

/*
 * param regionId: Region ID (Required)
 * param pageNumber: 当前页数，起始值为1，默认为1 (Required)
 * param pageSize: 分页查询时设置的每页行数 (Required)
 * param domainName: 关键字，按照”%domainName%”模式搜索主域名 (Optional)
 */
func NewGetDomainsRequestWithAllParams(
    regionId string,
    pageNumber int,
    pageSize int,
    domainName *string,
) *GetDomainsRequest {

    return &GetDomainsRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        DomainName: domainName,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetDomainsRequestWithoutParam() *GetDomainsRequest {

    return &GetDomainsRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/domain",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *GetDomainsRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 当前页数，起始值为1，默认为1(Required) */
func (r *GetDomainsRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = pageNumber
}

/* param pageSize: 分页查询时设置的每页行数(Required) */
func (r *GetDomainsRequest) SetPageSize(pageSize int) {
    r.PageSize = pageSize
}

/* param domainName: 关键字，按照”%domainName%”模式搜索主域名(Optional) */
func (r *GetDomainsRequest) SetDomainName(domainName string) {
    r.DomainName = &domainName
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetDomainsRequest) GetRegionId() string {
    return r.RegionId
}

type GetDomainsResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetDomainsResult `json:"result"`
}

type GetDomainsResult struct {
    DataList []clouddnsservice.Domain `json:"dataList"`
    CurrentCount int `json:"currentCount"`
    TotalCount int `json:"totalCount"`
}