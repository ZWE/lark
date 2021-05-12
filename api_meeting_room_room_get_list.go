// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetRoomList 该接口用于获取指定建筑下的会议室。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uADOyUjLwgjM14CM4ITN
func (r *MeetingRoomAPI) GetRoomList(ctx context.Context, request *GetRoomListReq, options ...MethodOptionFunc) (*GetRoomListResp, *Response, error) {
	r.cli.logInfo(ctx, "[lark] MeetingRoom#GetRoomList call api")
	r.cli.logDebug(ctx, "[lark] MeetingRoom#GetRoomList request: %s", jsonString(request))

	if r.cli.mock.mockMeetingRoomGetRoomList != nil {
		r.cli.logDebug(ctx, "[lark] MeetingRoom#GetRoomList mock enable")
		return r.cli.mock.mockMeetingRoomGetRoomList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/meeting_room/room/list?building_id=omb_8ec170b937536a5d87c23b418b83f9bb&page_size=1&page_token=0&order_by=name-asc&fields=*",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getRoomListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] MeetingRoom#GetRoomList GET https://open.feishu.cn/open-apis/meeting_room/room/list?building_id=omb_8ec170b937536a5d87c23b418b83f9bb&page_size=1&page_token=0&order_by=name-asc&fields=* failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] MeetingRoom#GetRoomList GET https://open.feishu.cn/open-apis/meeting_room/room/list?building_id=omb_8ec170b937536a5d87c23b418b83f9bb&page_size=1&page_token=0&order_by=name-asc&fields=* failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("MeetingRoom", "GetRoomList", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] MeetingRoom#GetRoomList request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockMeetingRoomGetRoomList(f func(ctx context.Context, request *GetRoomListReq, options ...MethodOptionFunc) (*GetRoomListResp, *Response, error)) {
	r.mockMeetingRoomGetRoomList = f
}

func (r *Mock) UnMockMeetingRoomGetRoomList() {
	r.mockMeetingRoomGetRoomList = nil
}

type GetRoomListReq struct {
	BuildingID string  `query:"building_id" json:"-"` // 被查询的建筑物 ID
	PageSize   *int    `query:"page_size" json:"-"`   // 请求期望返回的会议室数量，不足则返回全部，该值默认为 100，最大为 1000
	PageToken  *string `query:"page_token" json:"-"`  // 用于标记当前请求的分页标记，将返回以当前分页标记开始，往后 page_size 个元素
	OrderBy    *string `query:"order_by" json:"-"`    // 提供用于对名称/楼层进行升序/降序排序的方式查询，可选项有："name-asc,name-desc,floor_name-asc,floor_name-desc"，传入其他字符串不做处理，默认无序
	Fields     *string `query:"fields" json:"-"`      // 用于指定返回的字段名，每个字段名之间用逗号 "," 分隔，如：“id,name”，"*" 表示返回全部字段，可选字段有："id,name,description,capacity,building_id,building_name,floor_name,is_disabled,display_id"，默认返回所有字段
}

type getRoomListResp struct {
	Code int              `json:"code,omitempty"` // 返回码，非 0 表示失败
	Msg  string           `json:"msg,omitempty"`  // 返回码的描述，"success" 表示成功，其他为错误提示信息
	Data *GetRoomListResp `json:"data,omitempty"` // 返回业务信息
}

type GetRoomListResp struct {
	PageToken string                `json:"page_token,omitempty"` // 分页标记，存在下一页时返回
	HasMore   bool                  `json:"has_more,omitempty"`   // 存在下一页时，该值为 true，否则为 false
	Rooms     *GetRoomListRespRooms `json:"rooms,omitempty"`      // 会议室列表
}

type GetRoomListRespRooms struct {
	RoomID       string `json:"room_id,omitempty"`       // 会议室 ID
	BuildingID   string `json:"building_id,omitempty"`   // 会议室所属建筑物 ID
	BuildingName string `json:"building_name,omitempty"` // 会议室所属建筑物名称
	Capacity     int    `json:"capacity,omitempty"`      // 会议室能容纳的人数
	Description  string `json:"description,omitempty"`   // 会议室的相关描述
	DisplayID    string `json:"display_id,omitempty"`    // 会议室的展示 ID
	FloorName    string `json:"floor_name,omitempty"`    // 会议室所在楼层名称
	IsDisabled   bool   `json:"is_disabled,omitempty"`   // 会议室是否不可用，若会议室不可用，则该值为 true，否则为 false
	Name         string `json:"name,omitempty"`          // 会议室名称
}