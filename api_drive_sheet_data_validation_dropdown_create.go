// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// CreateSheetDataValidationDropdown 该接口根据 spreadsheetToken 、range 和下拉列表属性给单元格设置下拉列表规则；单次设置范围不超过5000行，100列。当一个数据区域中已有数据，支持将有效数据直接转为选项。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/datavalidation/set-dropdown
func (r *DriveService) CreateSheetDataValidationDropdown(ctx context.Context, request *CreateSheetDataValidationDropdownReq, options ...MethodOptionFunc) (*CreateSheetDataValidationDropdownResp, *Response, error) {
	if r.cli.mock.mockDriveCreateSheetDataValidationDropdown != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#CreateSheetDataValidationDropdown mock enable")
		return r.cli.mock.mockDriveCreateSheetDataValidationDropdown(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "CreateSheetDataValidationDropdown",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/dataValidation",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createSheetDataValidationDropdownResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockDriveCreateSheetDataValidationDropdown(f func(ctx context.Context, request *CreateSheetDataValidationDropdownReq, options ...MethodOptionFunc) (*CreateSheetDataValidationDropdownResp, *Response, error)) {
	r.mockDriveCreateSheetDataValidationDropdown = f
}

func (r *Mock) UnMockDriveCreateSheetDataValidationDropdown() {
	r.mockDriveCreateSheetDataValidationDropdown = nil
}

type CreateSheetDataValidationDropdownReq struct {
	SpreadSheetToken   string                                              `path:"spreadsheetToken" json:"-"`    // spreadsheet 的 token，获取方式见[在线表格开发指南](/ssl:ttdoc/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/overview)
	Range              string                                              `json:"range,omitempty"`              // 查询范围，包含 sheetId 与单元格范围两部分，目前支持四种索引方式，详见 [在线表格开发指南](/ssl:ttdoc/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/overview)
	DataValidationType string                                              `json:"dataValidationType,omitempty"` // 下拉列表填"list"
	DataValidation     *CreateSheetDataValidationDropdownReqDataValidation `json:"dataValidation,omitempty"`     // 下拉列表规则属性
}

type CreateSheetDataValidationDropdownReqDataValidation struct {
	ConditionValues []string                                                   `json:"conditionValues,omitempty"` // 下拉列表选项值, 需为字符串,不能包含","，选项值最长100字符,选项个数最多500个
	Options         *CreateSheetDataValidationDropdownReqDataValidationOptions `json:"options,omitempty"`         // 可选属性
}

type CreateSheetDataValidationDropdownReqDataValidationOptions struct {
	MultipleValues     *bool    `json:"multipleValues,omitempty"`     // 单选填false, 多选填true，不填默认为false
	HighlightValidData *bool    `json:"highlightValidData,omitempty"` // 是否设置颜色和胶囊样式, 不填默认为false
	Colors             []string `json:"colors,omitempty"`             // 当highlightValidData为true时，color需填颜色,与conditionValues中的值一一对应。需是RGB16进制格式,如"#fffd00"
}

type createSheetDataValidationDropdownResp struct {
	Code int64                                  `json:"code,omitempty"` // 状态码，0代表成功
	Msg  *string                                `json:"msg,omitempty"`  // 状态信息
	Data *CreateSheetDataValidationDropdownResp `json:"data,omitempty"`
}

type CreateSheetDataValidationDropdownResp struct{}
