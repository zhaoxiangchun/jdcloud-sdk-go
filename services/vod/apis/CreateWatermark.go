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
)

type CreateWatermarkRequest struct {

    core.JDCloudRequest

    /* 水印名称 (Optional) */
    Name *string `json:"name"`

    /* 图片地址 (Optional) */
    ImgUrl *string `json:"imgUrl"`

    /* 宽度 (Optional) */
    Width *int `json:"width"`

    /* 高度 (Optional) */
    Height *int `json:"height"`

    /* 水印位置 (Optional) */
    Position *string `json:"position"`

    /* 偏移单位 (Optional) */
    Unit *string `json:"unit"`

    /* 水平偏移 (Optional) */
    OffsetX *int `json:"offsetX"`

    /* 竖直偏移 (Optional) */
    OffsetY *int `json:"offsetY"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateWatermarkRequest(
) *CreateWatermarkRequest {

	return &CreateWatermarkRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/watermarks",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param name: 水印名称 (Optional)
 * param imgUrl: 图片地址 (Optional)
 * param width: 宽度 (Optional)
 * param height: 高度 (Optional)
 * param position: 水印位置 (Optional)
 * param unit: 偏移单位 (Optional)
 * param offsetX: 水平偏移 (Optional)
 * param offsetY: 竖直偏移 (Optional)
 */
func NewCreateWatermarkRequestWithAllParams(
    name *string,
    imgUrl *string,
    width *int,
    height *int,
    position *string,
    unit *string,
    offsetX *int,
    offsetY *int,
) *CreateWatermarkRequest {

    return &CreateWatermarkRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/watermarks",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Name: name,
        ImgUrl: imgUrl,
        Width: width,
        Height: height,
        Position: position,
        Unit: unit,
        OffsetX: offsetX,
        OffsetY: offsetY,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateWatermarkRequestWithoutParam() *CreateWatermarkRequest {

    return &CreateWatermarkRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/watermarks",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param name: 水印名称(Optional) */
func (r *CreateWatermarkRequest) SetName(name string) {
    r.Name = &name
}

/* param imgUrl: 图片地址(Optional) */
func (r *CreateWatermarkRequest) SetImgUrl(imgUrl string) {
    r.ImgUrl = &imgUrl
}

/* param width: 宽度(Optional) */
func (r *CreateWatermarkRequest) SetWidth(width int) {
    r.Width = &width
}

/* param height: 高度(Optional) */
func (r *CreateWatermarkRequest) SetHeight(height int) {
    r.Height = &height
}

/* param position: 水印位置(Optional) */
func (r *CreateWatermarkRequest) SetPosition(position string) {
    r.Position = &position
}

/* param unit: 偏移单位(Optional) */
func (r *CreateWatermarkRequest) SetUnit(unit string) {
    r.Unit = &unit
}

/* param offsetX: 水平偏移(Optional) */
func (r *CreateWatermarkRequest) SetOffsetX(offsetX int) {
    r.OffsetX = &offsetX
}

/* param offsetY: 竖直偏移(Optional) */
func (r *CreateWatermarkRequest) SetOffsetY(offsetY int) {
    r.OffsetY = &offsetY
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateWatermarkRequest) GetRegionId() string {
    return ""
}

type CreateWatermarkResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateWatermarkResult `json:"result"`
}

type CreateWatermarkResult struct {
    Id int64 `json:"id"`
    Name string `json:"name"`
    ImgUrl string `json:"imgUrl"`
    Width int `json:"width"`
    Height int `json:"height"`
    Position string `json:"position"`
    Unit string `json:"unit"`
    OffsetX int `json:"offsetX"`
    OffsetY int `json:"offsetY"`
    CreateTime string `json:"createTime"`
    UpdateTime string `json:"updateTime"`
}