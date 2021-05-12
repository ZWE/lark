// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// BatchGetBuildingID 该接口用于删除建筑物（办公大楼）。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzMxYjLzMTM24yMzEjN
func (r *MeetingRoomAPI) BatchGetBuildingID(ctx context.Context, request *BatchGetBuildingIDReq, options ...MethodOptionFunc) (*BatchGetBuildingIDResp, *Response, error) {
	r.cli.logInfo(ctx, "[lark] MeetingRoom#BatchGetBuildingID call api")
	r.cli.logDebug(ctx, "[lark] MeetingRoom#BatchGetBuildingID request: %s", jsonString(request))

	if r.cli.mock.mockMeetingRoomBatchGetBuildingID != nil {
		r.cli.logDebug(ctx, "[lark] MeetingRoom#BatchGetBuildingID mock enable")
		return r.cli.mock.mockMeetingRoomBatchGetBuildingID(ctx, request, options...)
	}

	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/meeting_room/building/delete",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(batchGetBuildingIDResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] MeetingRoom#BatchGetBuildingID POST https://open.feishu.cn/open-apis/meeting_room/building/delete failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] MeetingRoom#BatchGetBuildingID POST https://open.feishu.cn/open-apis/meeting_room/building/delete failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("MeetingRoom", "BatchGetBuildingID", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] MeetingRoom#BatchGetBuildingID request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockMeetingRoomBatchGetBuildingID(f func(ctx context.Context, request *BatchGetBuildingIDReq, options ...MethodOptionFunc) (*BatchGetBuildingIDResp, *Response, error)) {
	r.mockMeetingRoomBatchGetBuildingID = f
}

func (r *Mock) UnMockMeetingRoomBatchGetBuildingID() {
	r.mockMeetingRoomBatchGetBuildingID = nil
}

type BatchGetBuildingIDReq struct {
	BuildingID string `json:"building_id,omitempty"` // 要删除的建筑ID
}

type batchGetBuildingIDResp struct {
	Code int                     `json:"code,omitempty"` // 返回码，非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 返回码的描述，"success" 表示成功，其他为错误提示信息
	Data *BatchGetBuildingIDResp `json:"data,omitempty"`
}

type BatchGetBuildingIDResp struct{}