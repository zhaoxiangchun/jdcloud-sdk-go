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

package client

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    resourcetag "github.com/jdcloud-api/jdcloud-sdk-go/services/resourcetag/apis"
    "encoding/json"
    "errors"
)

type ResourcetagClient struct {
    core.JDCloudClient
}

func NewResourcetagClient(credential *core.Credential) *ResourcetagClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("resourcetag.jdcloud-api.com")

    return &ResourcetagClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "resourcetag",
            Revision:    "0.1.7",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *ResourcetagClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *ResourcetagClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

/* 获得资源与对应标签列表详情，不含资源名称和可用区。<br/>
注意查询cdn的资源时url中regionId必须指定为cn-all。<br/>
该接口目前不支持分页功能。
 */
func (c *ResourcetagClient) DescribeResources(request *resourcetag.DescribeResourcesRequest) (*resourcetag.DescribeResourcesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &resourcetag.DescribeResourcesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取资源标签。<br/>
注意查询cdn资源的标签时url中regionId必须指定为cn-all。<br/>
注意查询不限制地域时url中regionId必须指定为all-region。
 */
func (c *ResourcetagClient) DescribeTags(request *resourcetag.DescribeTagsRequest) (*resourcetag.DescribeTagsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &resourcetag.DescribeTagsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 资源标签绑定。<br/>
注意cdn资源绑定标签时url中regionId必须指定为cn-all。
 */
func (c *ResourcetagClient) TagResources(request *resourcetag.TagResourcesRequest) (*resourcetag.TagResourcesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &resourcetag.TagResourcesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 资源标签解绑。<br/>
注意cdn资源解绑标签时url中regionId必须指定为cn-all。
 */
func (c *ResourcetagClient) UnTagResources(request *resourcetag.UnTagResourcesRequest) (*resourcetag.UnTagResourcesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &resourcetag.UnTagResourcesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据标签查找资源。 <br/>
若要查找cdn产品线的资源则url中的regionId必须指定为cn-all。
 */
func (c *ResourcetagClient) QueryResource(request *resourcetag.QueryResourceRequest) (*resourcetag.QueryResourceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &resourcetag.QueryResourceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

