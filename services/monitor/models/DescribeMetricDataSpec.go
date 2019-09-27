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


type DescribeMetricDataSpec struct {

    /* 聚合方式，可选值参考:sum、avg、min、max (Optional) */
    AggrType string `json:"aggrType"`

    /* 资源的维度。当serviceCode下存在多个维度时，查询数据必须指定相应的维度 (Optional) */
    Dimension string `json:"dimension"`

    /* 采样方式，可选值参考：sum、avg、last、min、max (Optional) */
    DownSampleType string `json:"downSampleType"`

    /* 查询时间范围的结束时间， UTC时间，格式：2016-12-11T00:00:00+0800（为空时，将由startTime与timeInterval计算得出）（注意在url中+要转译为%2B故url中为2016-12-11T00:00:00%2B0800）
in: query (Optional) */
    EndTime string `json:"endTime"`

    /* 是否对查询的tags分组
in: query (Optional) */
    GroupBy bool `json:"groupBy"`

    /* 是否求速率
in: query (Optional) */
    Rate bool `json:"rate"`

    /* 资源的uuid  */
    ResourceId string `json:"resourceId"`

    /* 资源的类型，取值vm, lb, ip, database 等,<a href="https://docs.jdcloud.com/cn/monitoring/api/describeservices?content=API&SOP=JDCloud">describeServices</a>：查询己接入云监控的产品线列表，当产品线下有多个分组时，查询分组对应的监控项，serviceCode请传对应分组的groupCode字段值 (Optional) */
    ServiceCode string `json:"serviceCode"`

    /* 查询时间范围的开始时间， UTC时间，格式：2016-12-11T00:00:00+0800（注意在url中+要转译为%2B故url中为2016-12-11T00:00:00%2B0800）
in: query (Optional) */
    StartTime string `json:"startTime"`

    /* 监控指标数据的维度信息,根据tags来筛选指标数据不同的维度
in: query (Optional) */
    Tags []TagFilter `json:"tags"`

    /* 时间间隔：1h，6h，12h，1d，3d，7d，14d，固定时间间隔，timeInterval默认为1h，当前时间往 前1h
in: query (Optional) */
    TimeInterval string `json:"timeInterval"`
}
