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
    ag "github.com/jdcloud-api/jdcloud-sdk-go/services/ag/apis"
    "encoding/json"
    "errors"
)

type AgClient struct {
    core.JDCloudClient
}

func NewAgClient(credential *core.Credential) *AgClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("ag.jdcloud-api.com")

    return &AgClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "ag",
            Revision:    "0.4.0",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *AgClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *AgClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

/* 修改一个高可用组的信息 */
func (c *AgClient) UpdateAg(request *ag.UpdateAgRequest) (*ag.UpdateAgResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ag.UpdateAgResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据 id 删除高可用组，需确保 AG 中云主机实例已全部删除 */
func (c *AgClient) DeleteAg(request *ag.DeleteAgRequest) (*ag.DeleteAgResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ag.DeleteAgResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改高可用组的实例模板 */
func (c *AgClient) SetInstanceTemplate(request *ag.SetInstanceTemplateRequest) (*ag.SetInstanceTemplateResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ag.SetInstanceTemplateResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询(ag)配额 */
func (c *AgClient) DescribeQuotas(request *ag.DescribeQuotasRequest) (*ag.DescribeQuotasResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ag.DescribeQuotasResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 从高可用组中剔除实例 */
func (c *AgClient) AbandonInstances(request *ag.AbandonInstancesRequest) (*ag.AbandonInstancesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ag.AbandonInstancesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建一个高可用组 */
func (c *AgClient) CreateAg(request *ag.CreateAgRequest) (*ag.CreateAgResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ag.CreateAgResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据 id 查询高可用组详情 */
func (c *AgClient) DescribeAg(request *ag.DescribeAgRequest) (*ag.DescribeAgResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ag.DescribeAgResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 使用过滤条件查询一个或多个高可用组 */
func (c *AgClient) DescribeAgs(request *ag.DescribeAgsRequest) (*ag.DescribeAgsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &ag.DescribeAgsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}
