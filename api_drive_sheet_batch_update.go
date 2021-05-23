// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// BatchUpdateSheet 该接口用于根据 spreadsheetToken 操作表格，如增加工作表，复制工作表、删除工作表。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uYTMzUjL2EzM14iNxMTN
func (r *DriveService) BatchUpdateSheet(ctx context.Context, request *BatchUpdateSheetReq, options ...MethodOptionFunc) (*BatchUpdateSheetResp, *Response, error) {
	if r.cli.mock.mockDriveBatchUpdateSheet != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#BatchUpdateSheet mock enable")
		return r.cli.mock.mockDriveBatchUpdateSheet(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "BatchUpdateSheet",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/sheets_batch_update",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(batchUpdateSheetResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockDriveBatchUpdateSheet(f func(ctx context.Context, request *BatchUpdateSheetReq, options ...MethodOptionFunc) (*BatchUpdateSheetResp, *Response, error)) {
	r.mockDriveBatchUpdateSheet = f
}

func (r *Mock) UnMockDriveBatchUpdateSheet() {
	r.mockDriveBatchUpdateSheet = nil
}

type BatchUpdateSheetReq struct {
	SpreadSheetToken string                       `path:"spreadsheetToken" json:"-"` // spreadsheet 的 token，获取方式见[在线表格开发指南](/ssl:ttdoc/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/overview)
	Requests         *BatchUpdateSheetReqRequests `json:"requests,omitempty"`        // 请求操作，支持增、删、复制工作表 ，三个操作选一个
}

type BatchUpdateSheetReqRequests struct {
	AddSheet    *BatchUpdateSheetReqRequestsAddSheet    `json:"addSheet,omitempty"`    // 增加工作表
	CopySheet   *BatchUpdateSheetReqRequestsCopySheet   `json:"copySheet,omitempty"`   // 复制工作表
	DeleteSheet *BatchUpdateSheetReqRequestsDeleteSheet `json:"deleteSheet,omitempty"` // 删除 sheet
}

type BatchUpdateSheetReqRequestsAddSheet struct {
	Properties *BatchUpdateSheetReqRequestsAddSheetProperties `json:"properties,omitempty"` // 工作表属性
}

type BatchUpdateSheetReqRequestsAddSheetProperties struct {
	Title string `json:"title,omitempty"` // 工作表标题
	Index *int64 `json:"index,omitempty"` // 新增工作表的位置，不填默认往前增加工作表
}

type BatchUpdateSheetReqRequestsCopySheet struct {
	Source      *BatchUpdateSheetReqRequestsCopySheetSource      `json:"source,omitempty"`      // 需要复制的工作表资源
	Destination *BatchUpdateSheetReqRequestsCopySheetDestination `json:"destination,omitempty"` // 工作表 的属性
}

type BatchUpdateSheetReqRequestsCopySheetSource struct {
	SheetID string `json:"sheetId,omitempty"` // 源 sheetId
}

type BatchUpdateSheetReqRequestsCopySheetDestination struct {
	Title *string `json:"title,omitempty"` // 目标工作表名称。不填为 old_title(副本_0)
}

type BatchUpdateSheetReqRequestsDeleteSheet struct {
	SheetID string `json:"sheetId,omitempty"` // sheetId
}

type batchUpdateSheetResp struct {
	Code int64                 `json:"code,omitempty"`
	Msg  string                `json:"msg,omitempty"`
	Data *BatchUpdateSheetResp `json:"data,omitempty"`
}

type BatchUpdateSheetResp struct {
	Replies *BatchUpdateSheetRespReplies `json:"replies,omitempty"` // 返回本次相关操作工作表的结果
}

type BatchUpdateSheetRespReplies struct {
	AddSheet    *BatchUpdateSheetRespRepliesAddSheet    `json:"addSheet,omitempty"`    // 增加/复制工作表的属性
	CopySheet   *BatchUpdateSheetRespRepliesCopySheet   `json:"copySheet,omitempty"`   // 增加/复制工作表的属性
	DeleteSheet *BatchUpdateSheetRespRepliesDeleteSheet `json:"deleteSheet,omitempty"` // 删除工作表
}

type BatchUpdateSheetRespRepliesAddSheet struct {
	SheetID string `json:"sheetId,omitempty"` // sheetId
	Title   string `json:"title,omitempty"`   // 工作表标题
	Index   int64  `json:"index,omitempty"`   // 工作表位置
}

type BatchUpdateSheetRespRepliesCopySheet struct {
	SheetID string `json:"sheetId,omitempty"` // sheetId
	Title   string `json:"title,omitempty"`   // 工作表标题
	Index   int64  `json:"index,omitempty"`   // 工作表位置
}

type BatchUpdateSheetRespRepliesDeleteSheet struct {
	Result  bool   `json:"result,omitempty"`  // 删除工作表是否成功
	SheetID string `json:"sheetId,omitempty"` // sheetId
}
