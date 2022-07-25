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

// GetHireEmployee 通过员工 ID 获取入职信息
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/employee/get
func (r *HireService) GetHireEmployee(ctx context.Context, request *GetHireEmployeeReq, options ...MethodOptionFunc) (*GetHireEmployeeResp, *Response, error) {
	if r.cli.mock.mockHireGetHireEmployee != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Hire#GetHireEmployee mock enable")
		return r.cli.mock.mockHireGetHireEmployee(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "GetHireEmployee",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/hire/v1/employees/:employee_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getHireEmployeeResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHireGetHireEmployee mock HireGetHireEmployee method
func (r *Mock) MockHireGetHireEmployee(f func(ctx context.Context, request *GetHireEmployeeReq, options ...MethodOptionFunc) (*GetHireEmployeeResp, *Response, error)) {
	r.mockHireGetHireEmployee = f
}

// UnMockHireGetHireEmployee un-mock HireGetHireEmployee method
func (r *Mock) UnMockHireGetHireEmployee() {
	r.mockHireGetHireEmployee = nil
}

// GetHireEmployeeReq ...
type GetHireEmployeeReq struct {
	EmployeeID string  `path:"employee_id" json:"-"`   // 员工ID, 示例值: "123"
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 用户的 open id, union_id: 用户的 union id, user_id: 用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetHireEmployeeResp ...
type GetHireEmployeeResp struct {
	Employee *GetHireEmployeeRespEmployee `json:"employee,omitempty"` // 员工信息
}

// GetHireEmployeeRespEmployee ...
type GetHireEmployeeRespEmployee struct {
	ID                     string       `json:"id,omitempty"`                       // 员工ID
	ApplicationID          string       `json:"application_id,omitempty"`           // 投递ID
	OnboardStatus          int64        `json:"onboard_status,omitempty"`           // 入职状态, 可选值有: 1: 已入职, 2: 已离职
	ConversionStatus       int64        `json:"conversion_status,omitempty"`        // 转正状态, 可选值有: 1: 未转正, 2: 已转正
	OnboardTime            int64        `json:"onboard_time,omitempty"`             // 实际入职时间
	ExpectedConversionTime int64        `json:"expected_conversion_time,omitempty"` // 预期转正时间
	ActualConversionTime   int64        `json:"actual_conversion_time,omitempty"`   // 实际转正时间
	OverboardTime          int64        `json:"overboard_time,omitempty"`           // 离职时间
	OverboardNote          string       `json:"overboard_note,omitempty"`           // 离职原因
	OnboardCityCode        string       `json:"onboard_city_code,omitempty"`        // 办公地点
	Department             string       `json:"department,omitempty"`               // 入职部门
	Leader                 string       `json:"leader,omitempty"`                   // 直属上级
	Sequence               string       `json:"sequence,omitempty"`                 // 序列
	Level                  string       `json:"level,omitempty"`                    // 职级
	EmployeeType           EmployeeType `json:"employee_type,omitempty"`            // 员工类型
}

// getHireEmployeeResp ...
type getHireEmployeeResp struct {
	Code int64                `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string               `json:"msg,omitempty"`  // 错误描述
	Data *GetHireEmployeeResp `json:"data,omitempty"`
}