// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetContactGroupMember 通过该接口可查询某个用户组的成员(目前成员仅支持用户，未来会支持部门)列表，如果应用的通讯录权限范围是“全部员工”，则可查询企业内任何用户组的成员列表。如果应用的通讯录权限范围不是“全部员工”，则仅可查询通讯录权限范围中的用户组的成员列表，[点击了解通讯录权限范围](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/scope/overview)。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group-member/simplelist
func (r *ContactService) GetContactGroupMember(ctx context.Context, request *GetContactGroupMemberReq, options ...MethodOptionFunc) (*GetContactGroupMemberResp, *Response, error) {
	if r.cli.mock.mockContactGetContactGroupMember != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Contact#GetContactGroupMember mock enable")
		return r.cli.mock.mockContactGetContactGroupMember(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "GetContactGroupMember",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/group/:group_id/member/simplelist",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getContactGroupMemberResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockContactGetContactGroupMember(f func(ctx context.Context, request *GetContactGroupMemberReq, options ...MethodOptionFunc) (*GetContactGroupMemberResp, *Response, error)) {
	r.mockContactGetContactGroupMember = f
}

func (r *Mock) UnMockContactGetContactGroupMember() {
	r.mockContactGetContactGroupMember = nil
}

type GetContactGroupMemberReq struct {
	PageSize     *int64  `query:"page_size" json:"-"`      // 分页大小, 示例值：50, 默认值: `50`
	PageToken    *string `query:"page_token" json:"-"`     // 分页标记，第一次请求不填，表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token，下次遍历可采用该 page_token 获取查询结果, 示例值："AQD9/Rn9eij9Pm39ED40/dk53s4Ebp882DYfFaPFbz00L4CMZJrqGdzNyc8BcZtDbwVUvRmQTvyMYicnGWrde9X56TgdBuS+JKiSIkdexPw="
	MemberIDType *IDType `query:"member_id_type" json:"-"` // 欲获取成员ID类型。当member_type =user时候，member_id_type表示user_id_type，枚举值为open_id, union_id, user_id, 示例值："open_id", 可选值有: `open_id`：member_type =user时候，表示用户的open_id, `union_id`：member_type =user时候，表示用户的union_id, `user_id`：member_type =user时候，表示用户的user_id, 默认值: `open_id`
	MemberType   *string `query:"member_type" json:"-"`    // 期待获取的用户组成员的类型，取值为 user, 示例值："user", 可选值有: `user`：user, 默认值: `user`
	GroupID      string  `path:"group_id" json:"-"`        // 用户组ID, 示例值："g128187"
}

type getContactGroupMemberResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *GetContactGroupMemberResp `json:"data,omitempty"`
}

type GetContactGroupMemberResp struct {
	Memberlist []*GetContactGroupMemberRespMember `json:"memberlist,omitempty"` // 成员列表
	PageToken  string                             `json:"page_token,omitempty"` // 下次分页获取的page_token
	HasMore    bool                               `json:"has_more,omitempty"`   // 是否还需要分页获取
}

type GetContactGroupMemberRespMember struct {
	MemberID   string `json:"member_id,omitempty"`   // 成员ID
	MemberType string `json:"member_type,omitempty"` // 用户组成员的类型，目前取值为 user。未来将支持department
}