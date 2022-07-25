// Code generated by lark_sdk_gen. DO NOT EDIT.
/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lark

import (
	"context"
)

// GetContactCustomAttrList 获取企业自定义的用户字段配置信息
//
// 调用该接口前, 需要先确认[企业管理员](https://www.feishu.cn/hc/zh-CN/articles/360049067822)在[企业管理后台 - 组织架构 - 成员字段管理](http://www.feishu.cn/admin/contacts/employee-field-new/custom) 自定义字段管理栏开启了“允许开放平台API调用“。
// ![通讯录.gif](//sf3-cn.feishucdn.com/obj/open-platform-opendoc/544738c94f13ef0b9ebaff53a5133cc7_E9EGMkXyzX.gif)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/custom_attr/list
func (r *ContactService) GetContactCustomAttrList(ctx context.Context, request *GetContactCustomAttrListReq, options ...MethodOptionFunc) (*GetContactCustomAttrListResp, *Response, error) {
	if r.cli.mock.mockContactGetContactCustomAttrList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Contact#GetContactCustomAttrList mock enable")
		return r.cli.mock.mockContactGetContactCustomAttrList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "GetContactCustomAttrList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/custom_attrs",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getContactCustomAttrListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockContactGetContactCustomAttrList mock ContactGetContactCustomAttrList method
func (r *Mock) MockContactGetContactCustomAttrList(f func(ctx context.Context, request *GetContactCustomAttrListReq, options ...MethodOptionFunc) (*GetContactCustomAttrListResp, *Response, error)) {
	r.mockContactGetContactCustomAttrList = f
}

// UnMockContactGetContactCustomAttrList un-mock ContactGetContactCustomAttrList method
func (r *Mock) UnMockContactGetContactCustomAttrList() {
	r.mockContactGetContactCustomAttrList = nil
}

// GetContactCustomAttrListReq ...
type GetContactCustomAttrListReq struct {
	PageSize  *int64  `query:"page_size" json:"-"`  // 分页大小, 示例值: 10, 最大值: `100`
	PageToken *string `query:"page_token" json:"-"` // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "AQD9/Rn9eij9Pm39ED40/RYU5lvOM4s6zgbeeNNaWd%2BVKwAsoreeRWk0J2noGvJy"
}

// GetContactCustomAttrListResp ...
type GetContactCustomAttrListResp struct {
	Items     []*GetContactCustomAttrListRespItem `json:"items,omitempty"`      // 自定义字段定义
	PageToken string                              `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore   bool                                `json:"has_more,omitempty"`   // 是否还有更多项
}

// GetContactCustomAttrListRespItem ...
type GetContactCustomAttrListRespItem struct {
	ID       string                                      `json:"id,omitempty"`        // 自定义字段id
	Type     string                                      `json:"type,omitempty"`      // 自定义字段类型, 可选值有: `TEXT`: 纯文本, 用于纯文本描述人员, 如备注, `HREF`: 静态 URL, 用于人员 Profile 跳转链接, `ENUMERATION`: 枚举, 用于结构化描述人员, 如民族, `GENERIC_USER`: 用户, 用于描述人和人关系, 如 HRBP, `PICTURE_ENUM`: 枚举图片, 以结构化的图片描述人员, 如在人员 Profile 展示荣誉徽章
	Options  *GetContactCustomAttrListRespItemOptions    `json:"options,omitempty"`   // 选项定义, 当type为`ENUMERATION`或者`PICTURE_ENUM`时此项有值, 列举所有可选项
	I18nName []*GetContactCustomAttrListRespItemI18nName `json:"i18n_name,omitempty"` // 自定义字段的字段名称
}

// GetContactCustomAttrListRespItemI18nName ...
type GetContactCustomAttrListRespItemI18nName struct {
	Locale string `json:"locale,omitempty"` // 语言版本
	Value  string `json:"value,omitempty"`  // 字段名
}

// GetContactCustomAttrListRespItemOptions ...
type GetContactCustomAttrListRespItemOptions struct {
	DefaultOptionID string                                           `json:"default_option_id,omitempty"` // 默认选项id
	OptionType      string                                           `json:"option_type,omitempty"`       // 选项类型, 可选值有: TEXT: 文本选项, PICTURE: 图片选项
	Options         []*GetContactCustomAttrListRespItemOptionsOption `json:"options,omitempty"`           // 选项列表
}

// GetContactCustomAttrListRespItemOptionsOption ...
type GetContactCustomAttrListRespItemOptionsOption struct {
	ID    string `json:"id,omitempty"`    // 枚举类型选项id
	Value string `json:"value,omitempty"` // 枚举选项值, 当option_type为`TEXT`为文本值, 当option_type为`PICTURE`时为图片链接
	Name  string `json:"name,omitempty"`  // 名称, 仅option_type为PICTURE时有效
}

// getContactCustomAttrListResp ...
type getContactCustomAttrListResp struct {
	Code int64                         `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                        `json:"msg,omitempty"`  // 错误描述
	Data *GetContactCustomAttrListResp `json:"data,omitempty"`
}