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
    user "github.com/jdcloud-api/jdcloud-sdk-go/services/user/apis"
    "encoding/json"
    "errors"
)

type UserClient struct {
    core.JDCloudClient
}

func NewUserClient(credential *core.Credential) *UserClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("user.jdcloud-api.com")

    return &UserClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "user",
            Revision:    "0.1.5",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *UserClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *UserClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

func (c *UserClient) DisableLogger() {
    c.Logger = core.NewDummyLogger()
}

/* 查询用户信息 */
func (c *UserClient) DescribeUser(request *user.DescribeUserRequest) (*user.DescribeUserResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &user.DescribeUserResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

