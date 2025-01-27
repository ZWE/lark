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

// UnsubscribeCalendarEvent 该接口用于以用户身份取消订阅指定日历下的日程变更事件。
//
// 当前身份必须对日历有reader、writer或owner权限（调用[获取日历](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/get)接口, role字段可查看权限）。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/unsubscription
func (r *CalendarService) UnsubscribeCalendarEvent(ctx context.Context, request *UnsubscribeCalendarEventReq, options ...MethodOptionFunc) (*UnsubscribeCalendarEventResp, *Response, error) {
	if r.cli.mock.mockCalendarUnsubscribeCalendarEvent != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Calendar#UnsubscribeCalendarEvent mock enable")
		return r.cli.mock.mockCalendarUnsubscribeCalendarEvent(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Calendar",
		API:                 "UnsubscribeCalendarEvent",
		Method:              "POST",
		URL:                 r.cli.openBaseURL + "/open-apis/calendar/v4/calendars/:calendar_id/events/unsubscription",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(unsubscribeCalendarEventResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCalendarUnsubscribeCalendarEvent mock CalendarUnsubscribeCalendarEvent method
func (r *Mock) MockCalendarUnsubscribeCalendarEvent(f func(ctx context.Context, request *UnsubscribeCalendarEventReq, options ...MethodOptionFunc) (*UnsubscribeCalendarEventResp, *Response, error)) {
	r.mockCalendarUnsubscribeCalendarEvent = f
}

// UnMockCalendarUnsubscribeCalendarEvent un-mock CalendarUnsubscribeCalendarEvent method
func (r *Mock) UnMockCalendarUnsubscribeCalendarEvent() {
	r.mockCalendarUnsubscribeCalendarEvent = nil
}

// UnsubscribeCalendarEventReq ...
type UnsubscribeCalendarEventReq struct {
	CalendarID string `path:"calendar_id" json:"-"` // 日历ID。参见[日历ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/introduction), 示例值: "feishu.cn_xxxxxxxxxx@group.calendar.feishu.cn"
}

// UnsubscribeCalendarEventResp ...
type UnsubscribeCalendarEventResp struct {
}

// unsubscribeCalendarEventResp ...
type unsubscribeCalendarEventResp struct {
	Code int64                         `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                        `json:"msg,omitempty"`  // 错误描述
	Data *UnsubscribeCalendarEventResp `json:"data,omitempty"`
}
