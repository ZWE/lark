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

// CreateBitableField 该接口用于在数据表中新增一个字段
//
// 该接口支持调用频率上限为 10 QPS
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-field/create
func (r *BitableService) CreateBitableField(ctx context.Context, request *CreateBitableFieldReq, options ...MethodOptionFunc) (*CreateBitableFieldResp, *Response, error) {
	if r.cli.mock.mockBitableCreateBitableField != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#CreateBitableField mock enable")
		return r.cli.mock.mockBitableCreateBitableField(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bitable",
		API:                   "CreateBitableField",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/bitable/v1/apps/:app_token/tables/:table_id/fields",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createBitableFieldResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBitableCreateBitableField mock BitableCreateBitableField method
func (r *Mock) MockBitableCreateBitableField(f func(ctx context.Context, request *CreateBitableFieldReq, options ...MethodOptionFunc) (*CreateBitableFieldResp, *Response, error)) {
	r.mockBitableCreateBitableField = f
}

// UnMockBitableCreateBitableField un-mock BitableCreateBitableField method
func (r *Mock) UnMockBitableCreateBitableField() {
	r.mockBitableCreateBitableField = nil
}

// CreateBitableFieldReq ...
type CreateBitableFieldReq struct {
	AppToken    string                            `path:"app_token" json:"-"`    // bitable app token, 示例值: "appbcbWCzen6D8dezhoCH2RpMAh"
	TableID     string                            `path:"table_id" json:"-"`     // table id, 示例值: "tblsRc9GRRXKqhvW"
	FieldName   string                            `json:"field_name,omitempty"`  // 多维表格字段名, 示例值: "多行文本"
	Type        int64                             `json:"type,omitempty"`        // 多维表格字段类型, 示例值: 1, 可选值有: 1: 多行文本, 2: 数字, 3: 单选, 4: 多选, 5: 日期, 7: 复选框, 11: 人员, 15: 超链接, 17: 附件, 18: 关联, 20: 公式, 21: 双向关联, 1001: 创建时间, 1002: 最后更新时间, 1003: 创建人, 1004: 修改人, 1005: 自动编号, 13: 电话号码, 22: 地理位置
	Property    *CreateBitableFieldReqProperty    `json:"property,omitempty"`    // 字段属性, 具体参考: [字段编辑指南](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-field/guide)
	Description *CreateBitableFieldReqDescription `json:"description,omitempty"` // 字段的描述
}

// CreateBitableFieldReqDescription ...
type CreateBitableFieldReqDescription struct {
	DisableSync *bool   `json:"disable_sync,omitempty"` // 是否禁止同步, 如果为true, 表示禁止同步该描述内容到表单的问题描述（只在新增、修改字段时生效）, 示例值: ture, 默认值: `ture`
	Text        *string `json:"text,omitempty"`         // 字段描述内容, 示例值: "这是一个字段描述"
}

// CreateBitableFieldReqProperty ...
type CreateBitableFieldReqProperty struct {
	Options           []*CreateBitableFieldReqPropertyOption   `json:"options,omitempty"`            // 单选、多选字段的选项信息
	Formatter         *string                                  `json:"formatter,omitempty"`          // 数字、公式字段的显示格式, 示例值: "0"
	DateFormatter     *string                                  `json:"date_formatter,omitempty"`     // 日期、创建时间、最后更新时间字段的显示格式, 示例值: "日期格式"
	AutoFill          *bool                                    `json:"auto_fill,omitempty"`          // 日期字段中新纪录自动填写创建时间, 示例值: false
	Multiple          *bool                                    `json:"multiple,omitempty"`           // 人员字段中允许添加多个成员, 单向关联、双向关联中允许添加多个记录, 示例值: false
	TableID           *string                                  `json:"table_id,omitempty"`           // 单向关联、双向关联字段中关联的数据表的id, 示例值: "tblsRc9GRRXKqhvW"
	TableName         *string                                  `json:"table_name,omitempty"`         // 单向关联、双向关联字段中关联的数据表的名字, 示例值: ""table2""
	BackFieldName     *string                                  `json:"back_field_name,omitempty"`    // 双向关联字段中关联的数据表中对应的双向关联字段的名字, 示例值: ""table1-双向关联""
	AutoSerial        *CreateBitableFieldReqPropertyAutoSerial `json:"auto_serial,omitempty"`        // 自动编号类型
	Location          *CreateBitableFieldReqPropertyLocation   `json:"location,omitempty"`           // 地理位置输入方式
	FormulaExpression *string                                  `json:"formula_expression,omitempty"` // 公式字段的表达式, 示例值: "bitable::$table[tblNj92WQBAasdEf].$field[fldMV60rYs]*2"
}

// CreateBitableFieldReqPropertyAutoSerial ...
type CreateBitableFieldReqPropertyAutoSerial struct {
	Type    string                                           `json:"type,omitempty"`    // 自动编号类型, 示例值: "auto_increment_number", 可选值有: custom: 自定义编号, auto_increment_number: 自增数字
	Options []*CreateBitableFieldReqPropertyAutoSerialOption `json:"options,omitempty"` // 自动编号规则列表
}

// CreateBitableFieldReqPropertyAutoSerialOption ...
type CreateBitableFieldReqPropertyAutoSerialOption struct {
	Type  string `json:"type,omitempty"`  // 自动编号的可选规则项类型, 示例值: "created_time", 可选值有: system_number: 自增数字位, value范围1-9, fixed_text: 固定字符, 最大长度: 20, created_time: 创建时间, 支持格式 "yyyyMMdd"、"yyyyMM"、"yyyy"、"MMdd"、"MM"、"dd"
	Value string `json:"value,omitempty"` // 与自动编号的可选规则项类型相对应的取值, 示例值: "yyyyMMdd"
}

// CreateBitableFieldReqPropertyLocation ...
type CreateBitableFieldReqPropertyLocation struct {
	InputType string `json:"input_type,omitempty"` // 地理位置输入限制, 示例值: "not_limit", 可选值有: only_mobile: 只允许移动端上传, not_limit: 无限制
}

// CreateBitableFieldReqPropertyOption ...
type CreateBitableFieldReqPropertyOption struct {
	Name  *string `json:"name,omitempty"`  // 选项名, 示例值: "红色"
	ID    *string `json:"id,omitempty"`    // 选项 ID, 创建时不允许指定 ID, 示例值: "optKl35lnG"
	Color *int64  `json:"color,omitempty"` // 选项颜色, 示例值: 0, 取值范围: `0` ～ `54`
}

// CreateBitableFieldResp ...
type CreateBitableFieldResp struct {
	Field *CreateBitableFieldRespField `json:"field,omitempty"` // 字段
}

// CreateBitableFieldRespField ...
type CreateBitableFieldRespField struct {
	FieldID     string                                  `json:"field_id,omitempty"`    // 多维表格字段 id
	FieldName   string                                  `json:"field_name,omitempty"`  // 多维表格字段名
	Type        int64                                   `json:"type,omitempty"`        // 多维表格字段类型, 可选值有: 1: 多行文本, 2: 数字, 3: 单选, 4: 多选, 5: 日期, 7: 复选框, 11: 人员, 15: 超链接, 17: 附件, 18: 关联, 20: 公式, 21: 双向关联, 1001: 创建时间, 1002: 最后更新时间, 1003: 创建人, 1004: 修改人, 1005: 自动编号, 13: 电话号码, 22: 地理位置
	Property    *CreateBitableFieldRespFieldProperty    `json:"property,omitempty"`    // 字段属性, 具体参考: [字段编辑指南](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-field/guide)
	Description *CreateBitableFieldRespFieldDescription `json:"description,omitempty"` // 字段的描述
}

// CreateBitableFieldRespFieldDescription ...
type CreateBitableFieldRespFieldDescription struct {
	DisableSync bool   `json:"disable_sync,omitempty"` // 是否禁止同步, 如果为true, 表示禁止同步该描述内容到表单的问题描述（只在新增、修改字段时生效）
	Text        string `json:"text,omitempty"`         // 字段描述内容
}

// CreateBitableFieldRespFieldProperty ...
type CreateBitableFieldRespFieldProperty struct {
	Options           []*CreateBitableFieldRespFieldPropertyOption   `json:"options,omitempty"`            // 单选、多选字段的选项信息
	Formatter         string                                         `json:"formatter,omitempty"`          // 数字、公式字段的显示格式
	DateFormatter     string                                         `json:"date_formatter,omitempty"`     // 日期、创建时间、最后更新时间字段的显示格式
	AutoFill          bool                                           `json:"auto_fill,omitempty"`          // 日期字段中新纪录自动填写创建时间
	Multiple          bool                                           `json:"multiple,omitempty"`           // 人员字段中允许添加多个成员, 单向关联、双向关联中允许添加多个记录
	TableID           string                                         `json:"table_id,omitempty"`           // 单向关联、双向关联字段中关联的数据表的id
	TableName         string                                         `json:"table_name,omitempty"`         // 单向关联、双向关联字段中关联的数据表的名字
	BackFieldName     string                                         `json:"back_field_name,omitempty"`    // 双向关联字段中关联的数据表中对应的双向关联字段的名字
	AutoSerial        *CreateBitableFieldRespFieldPropertyAutoSerial `json:"auto_serial,omitempty"`        // 自动编号类型
	Location          *CreateBitableFieldRespFieldPropertyLocation   `json:"location,omitempty"`           // 地理位置输入方式
	FormulaExpression string                                         `json:"formula_expression,omitempty"` // 公式字段的表达式
}

// CreateBitableFieldRespFieldPropertyAutoSerial ...
type CreateBitableFieldRespFieldPropertyAutoSerial struct {
	Type    string                                                 `json:"type,omitempty"`    // 自动编号类型, 可选值有: custom: 自定义编号, auto_increment_number: 自增数字
	Options []*CreateBitableFieldRespFieldPropertyAutoSerialOption `json:"options,omitempty"` // 自动编号规则列表
}

// CreateBitableFieldRespFieldPropertyAutoSerialOption ...
type CreateBitableFieldRespFieldPropertyAutoSerialOption struct {
	Type  string `json:"type,omitempty"`  // 自动编号的可选规则项类型, 可选值有: system_number: 自增数字位, value范围1-9, fixed_text: 固定字符, 最大长度: 20, created_time: 创建时间, 支持格式 "yyyyMMdd"、"yyyyMM"、"yyyy"、"MMdd"、"MM"、"dd"
	Value string `json:"value,omitempty"` // 与自动编号的可选规则项类型相对应的取值
}

// CreateBitableFieldRespFieldPropertyLocation ...
type CreateBitableFieldRespFieldPropertyLocation struct {
	InputType string `json:"input_type,omitempty"` // 地理位置输入限制, 可选值有: only_mobile: 只允许移动端上传, not_limit: 无限制
}

// CreateBitableFieldRespFieldPropertyOption ...
type CreateBitableFieldRespFieldPropertyOption struct {
	Name  string `json:"name,omitempty"`  // 选项名
	ID    string `json:"id,omitempty"`    // 选项 ID, 创建时不允许指定 ID
	Color int64  `json:"color,omitempty"` // 选项颜色
}

// createBitableFieldResp ...
type createBitableFieldResp struct {
	Code int64                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *CreateBitableFieldResp `json:"data,omitempty"`
}
