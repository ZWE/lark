// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// DeleteReserve 删除一个预约
//
// 只能删除归属于自己的预约；删除后数据不可恢复
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/reserve/delete
func (r *VCAPI) DeleteReserve(ctx context.Context, request *DeleteReserveReq, options ...MethodOptionFunc) (*DeleteReserveResp, *Response, error) {
	r.cli.logInfo(ctx, "[lark] VC#DeleteReserve call api")
	r.cli.logDebug(ctx, "[lark] VC#DeleteReserve request: %s", jsonString(request))

	if r.cli.mock.mockVCDeleteReserve != nil {
		r.cli.logDebug(ctx, "[lark] VC#DeleteReserve mock enable")
		return r.cli.mock.mockVCDeleteReserve(ctx, request, options...)
	}

	req := &RawRequestReq{
		Method:              "DELETE",
		URL:                 "https://open.feishu.cn/open-apis/vc/v1/reserves/:reserve_id",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(deleteReserveResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] VC#DeleteReserve DELETE https://open.feishu.cn/open-apis/vc/v1/reserves/:reserve_id failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] VC#DeleteReserve DELETE https://open.feishu.cn/open-apis/vc/v1/reserves/:reserve_id failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("VC", "DeleteReserve", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] VC#DeleteReserve request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockVCDeleteReserve(f func(ctx context.Context, request *DeleteReserveReq, options ...MethodOptionFunc) (*DeleteReserveResp, *Response, error)) {
	r.mockVCDeleteReserve = f
}

func (r *Mock) UnMockVCDeleteReserve() {
	r.mockVCDeleteReserve = nil
}

type DeleteReserveReq struct {
	ReserveID string `path:"reserve_id" json:"-"` // 预约ID, 示例值："6911188411932033028"
}

type deleteReserveResp struct {
	Code int                `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string             `json:"msg,omitempty"`  // 错误描述
	Data *DeleteReserveResp `json:"data,omitempty"`
}

type DeleteReserveResp struct{}