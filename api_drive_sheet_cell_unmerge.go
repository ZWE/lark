// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// UnmergeSheetCell 该接口用于根据 spreadsheetToken 和维度信息拆分单元格；单次操作不超过5000行，100列。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uATNzUjLwUzM14CM1MTN
func (r *DriveService) UnmergeSheetCell(ctx context.Context, request *UnmergeSheetCellReq, options ...MethodOptionFunc) (*UnmergeSheetCellResp, *Response, error) {
	if r.cli.mock.mockDriveUnmergeSheetCell != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#UnmergeSheetCell mock enable")
		return r.cli.mock.mockDriveUnmergeSheetCell(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "UnmergeSheetCell",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/unmerge_cells",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(unmergeSheetCellResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockDriveUnmergeSheetCell(f func(ctx context.Context, request *UnmergeSheetCellReq, options ...MethodOptionFunc) (*UnmergeSheetCellResp, *Response, error)) {
	r.mockDriveUnmergeSheetCell = f
}

func (r *Mock) UnMockDriveUnmergeSheetCell() {
	r.mockDriveUnmergeSheetCell = nil
}

type UnmergeSheetCellReq struct {
	SpreadSheetToken string `path:"spreadsheetToken" json:"-"` // spreadsheet 的 token，获取方式见[在线表格开发指南](/ssl:ttdoc/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/overview)
	Range            string `json:"range,omitempty"`           // 查询范围，包含 sheetId 与单元格范围两部分，目前支持四种索引方式，详见[在线表格开发指南](/ssl:ttdoc/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/overview)
}

type unmergeSheetCellResp struct {
	Code int64                 `json:"code,omitempty"`
	Msg  string                `json:"msg,omitempty"`
	Data *UnmergeSheetCellResp `json:"data,omitempty"`
}

type UnmergeSheetCellResp struct {
	SpreadSheetToken string `json:"spreadsheetToken,omitempty"` // spreadsheet 的 token
}
