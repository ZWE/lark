// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// EventV1TripApproval
//
// 了解事件订阅的使用场景和配置流程，请点击查看 [事件订阅概述](https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM)
// 「审批」应用的表单里如果包含 [出差控件组]，则在此表单审批通过后触发此事件。
// * 依赖权限：[访问审批应用]
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uIDO24iM4YjLygjN/event/business-trip
func (r *EventCallbackService) HandlerEventV1TripApproval(f eventV1TripApprovalHandler) {
	r.cli.eventHandler.eventV1TripApprovalHandler = f
}

type eventV1TripApprovalHandler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV1, event *EventV1TripApproval) (string, error)

type EventV1TripApproval struct {
	AppID        string                              `json:"app_id,omitempty"`        // 如: cli_xxx
	TenantKey    string                              `json:"tenant_key,omitempty"`    // 如: xxx
	Type         string                              `json:"type,omitempty"`          // 如: trip_approval
	InstanceCode string                              `json:"instance_code,omitempty"` // 审批实例Code. 如: xxx
	EmployeeID   string                              `json:"employee_id,omitempty"`   // 用户id. 如: xxx
	StartTime    int64                               `json:"start_time,omitempty"`    // 审批发起时间. 如: 1502199207
	EndTime      int64                               `json:"end_time,omitempty"`      // 审批结束时间. 如: 1502199307
	Schedules    []*EventV1TripApprovalEventSchedule `json:"schedules,omitempty"`
	TripInterval int64                               `json:"trip_interval,omitempty"` // 行程总时长（秒）. 如: 3600
	TripReason   string                              `json:"trip_reason,omitempty"`   // 出差事由. 如: xxx
	TripPeers    []*EventV1TripApprovalEventTripPeer `json:"trip_peers,omitempty"`    // 同行人
}

type EventV1TripApprovalEventSchedule struct {
	TripStartTime  string `json:"trip_start_time,omitempty"` // 行程开始时间. 如: 2018-12-01 12:00:00
	TripEndTime    string `json:"trip_end_time,omitempty"`   // 行程结束时间. 如: 2018-12-01 12:00:00
	TripInterval   int64  `json:"trip_interval,omitempty"`   // 行程时长（秒）. 如: 3600
	Departure      string `json:"departure,omitempty"`       // 出发地. 如: xxx
	Destination    string `json:"destination,omitempty"`     // 目的地. 如: xxx
	Transportation string `json:"transportation,omitempty"`  // 交通工具. 如: xxx
	TripType       string `json:"trip_type,omitempty"`       // 单程/往返. 如: 单程
	Remark         string `json:"remark,omitempty"`          // 备注. 如: 备注
}

type EventV1TripApprovalEventTripPeer struct {
	string `json:",omitempty"` // 如: xxx
}