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
    monitor "github.com/jdcloud-api/jdcloud-sdk-go/services/monitor/apis"
    "encoding/json"
    "errors"
)

type MonitorClient struct {
    core.JDCloudClient
}

func NewMonitorClient(credential *core.Credential) *MonitorClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("monitor.jdcloud-api.com")

    return &MonitorClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "monitor",
            Revision:    "2.0.0",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *MonitorClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *MonitorClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

/* 创建报警规则 */
func (c *MonitorClient) CreateAlarm(request *monitor.CreateAlarmRequest) (*monitor.CreateAlarmResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &monitor.CreateAlarmResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查看某资源单个监控项数据，metric介绍：<a href="https://docs.jdcloud.com/cn/monitoring/metrics">Metrics</a>，可以使用接口<a href="https://docs.jdcloud.com/cn/monitoring/metrics">describeMetrics</a>：查询产品线可用的metric列表。 */
func (c *MonitorClient) DescribeMetricData(request *monitor.DescribeMetricDataRequest) (*monitor.DescribeMetricDataResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &monitor.DescribeMetricDataResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询监控图可用的产品线列表 */
func (c *MonitorClient) DescribeServices(request *monitor.DescribeServicesRequest) (*monitor.DescribeServicesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &monitor.DescribeServicesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询可用于创建监控规则的指标列表,metric介绍：<a href="https://docs.jdcloud.com/cn/monitoring/metrics">Metrics</a> */
func (c *MonitorClient) DescribeMetricsForAlarm(request *monitor.DescribeMetricsForAlarmRequest) (*monitor.DescribeMetricsForAlarmResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &monitor.DescribeMetricsForAlarmResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询规则的报警联系人 */
func (c *MonitorClient) DescribeAlarmContacts(request *monitor.DescribeAlarmContactsRequest) (*monitor.DescribeAlarmContactsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &monitor.DescribeAlarmContactsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询报警历史 */
func (c *MonitorClient) DescribeAlarmHistory(request *monitor.DescribeAlarmHistoryRequest) (*monitor.DescribeAlarmHistoryResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &monitor.DescribeAlarmHistoryResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询规则详情 */
func (c *MonitorClient) DescribeAlarm(request *monitor.DescribeAlarmRequest) (*monitor.DescribeAlarmResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &monitor.DescribeAlarmResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除规则 */
func (c *MonitorClient) DeleteAlarms(request *monitor.DeleteAlarmsRequest) (*monitor.DeleteAlarmsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &monitor.DeleteAlarmsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据不同的聚合方式将metric的数据聚合为一个点。downAggrType：last(最后一个点)、max(最大值)、min(最小值)、avg(平均值)。该接口返回值为上报metric的原始值，没有做单位转换。metric介绍：<a href="https://docs.jdcloud.com/cn/monitoring/metrics">Metrics</a> */
func (c *MonitorClient) LastDownsample(request *monitor.LastDownsampleRequest) (*monitor.LastDownsampleResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &monitor.LastDownsampleResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询可用于创建监控规则的产品列表 */
func (c *MonitorClient) DescribeProductsForAlarm(request *monitor.DescribeProductsForAlarmRequest) (*monitor.DescribeProductsForAlarmResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &monitor.DescribeProductsForAlarmResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改已创建的报警规则 */
func (c *MonitorClient) UpdateAlarm(request *monitor.UpdateAlarmRequest) (*monitor.UpdateAlarmResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &monitor.UpdateAlarmResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 启用、禁用规则 */
func (c *MonitorClient) EnableAlarms(request *monitor.EnableAlarmsRequest) (*monitor.EnableAlarmsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &monitor.EnableAlarmsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 该接口为自定义监控数据上报的接口，方便您将自己采集的时序数据上报到云监控。不同region域名上报不同region的数据，参考：<a href="https://docs.jdcloud.com/cn/monitoring/reporting-monitoring-data">调用说明</a>可上报原始数据和已聚合的统计数据。支持批量上报方式。单次请求最多包含 50 个数据点；数据大小不超过 256k。 */
func (c *MonitorClient) PutMetricData(request *monitor.PutMetricDataRequest) (*monitor.PutMetricDataResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &monitor.PutMetricDataResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询规则列表 */
func (c *MonitorClient) DescribeAlarms(request *monitor.DescribeAlarmsRequest) (*monitor.DescribeAlarmsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &monitor.DescribeAlarmsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 根据产品线查询可用监控项列表,metric介绍：<a href="https://docs.jdcloud.com/cn/monitoring/metrics">Metrics</a> */
func (c *MonitorClient) DescribeMetrics(request *monitor.DescribeMetricsRequest) (*monitor.DescribeMetricsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &monitor.DescribeMetricsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

