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
    jdro "github.com/jdcloud-api/jdcloud-sdk-go/services/jdro/models"
)

type DescribeResourceTypeListRequest struct {

    core.JDCloudRequest

    /* 地域 ID  */
    RegionId string `json:"regionId"`

    /* 当前所在页，默认为1 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 页面大小，默认为20；取值范围[1, 100] (Optional) */
    PageSize *int `json:"pageSize"`

    /* 产品线类型，比如 VM (Optional) */
    Product *string `json:"product"`

    /* 搜索的内容 (Optional) */
    Search *string `json:"search"`
}

/*
 * param regionId: 地域 ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeResourceTypeListRequest(
    regionId string,
) *DescribeResourceTypeListRequest {

	return &DescribeResourceTypeListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/resourcetypes",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域 ID (Required)
 * param pageNumber: 当前所在页，默认为1 (Optional)
 * param pageSize: 页面大小，默认为20；取值范围[1, 100] (Optional)
 * param product: 产品线类型，比如 VM (Optional)
 * param search: 搜索的内容 (Optional)
 */
func NewDescribeResourceTypeListRequestWithAllParams(
    regionId string,
    pageNumber *int,
    pageSize *int,
    product *string,
    search *string,
) *DescribeResourceTypeListRequest {

    return &DescribeResourceTypeListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/resourcetypes",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        PageNumber: pageNumber,
        PageSize: pageSize,
        Product: product,
        Search: search,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeResourceTypeListRequestWithoutParam() *DescribeResourceTypeListRequest {

    return &DescribeResourceTypeListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/resourcetypes",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域 ID(Required) */
func (r *DescribeResourceTypeListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param pageNumber: 当前所在页，默认为1(Optional) */
func (r *DescribeResourceTypeListRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}

/* param pageSize: 页面大小，默认为20；取值范围[1, 100](Optional) */
func (r *DescribeResourceTypeListRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param product: 产品线类型，比如 VM(Optional) */
func (r *DescribeResourceTypeListRequest) SetProduct(product string) {
    r.Product = &product
}

/* param search: 搜索的内容(Optional) */
func (r *DescribeResourceTypeListRequest) SetSearch(search string) {
    r.Search = &search
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeResourceTypeListRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeResourceTypeListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeResourceTypeListResult `json:"result"`
}

type DescribeResourceTypeListResult struct {
    ResourceTypeList []jdro.DescribeResourceTypeListItem `json:"resourceTypeList"`
    TotalCount int64 `json:"totalCount"`
}