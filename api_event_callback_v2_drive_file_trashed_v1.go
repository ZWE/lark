// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// EventV2DriveFileTrashedV1
//
// 了解事件订阅的使用场景和配置流程，请点击查看 [事件订阅概述](https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM)
// 文件被删除到回收站将触发此事件。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/event/delete-file-to-trash-can
func (r *EventCallbackService) HandlerEventV2DriveFileTrashedV1(f eventV2DriveFileTrashedV1Handler) {
	r.cli.eventHandler.eventV2DriveFileTrashedV1Handler = f
}

type eventV2DriveFileTrashedV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2DriveFileTrashedV1) (string, error)

type EventV2DriveFileTrashedV1 struct {
	FileToken  string                               `json:"file_token,omitempty"`  // 文件token. 如: doccnxxxxxx
	FileType   FileType                             `json:"file_type,omitempty"`   // 文件类型，目前有doc、sheet. 如: doc
	OperatorID *EventV2DriveFileTrashedV1OperatorID `json:"operator_id,omitempty"` // 操作者id
}

type EventV2DriveFileTrashedV1OperatorID struct {
	OpenID  string `json:"open_id,omitempty"`  // 如: ou_xxxxxx
	UnionID string `json:"union_id,omitempty"` // 如: on_xxxxxx
	UserID  string `json:"user_id,omitempty"`  // 如: xxxxxx
}