// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// EventV2DriveFilePermissionMemberAddedV1
//
// 了解事件订阅的使用场景和配置流程，请点击查看 [事件订阅概述](https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM)
// 文件协作者添加用户/群时将触发此事件。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/event/file-collaborator-add
func (r *EventCallbackService) HandlerEventV2DriveFilePermissionMemberAddedV1(f eventV2DriveFilePermissionMemberAddedV1Handler) {
	r.cli.eventHandler.eventV2DriveFilePermissionMemberAddedV1Handler = f
}

type eventV2DriveFilePermissionMemberAddedV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2DriveFilePermissionMemberAddedV1) (string, error)

type EventV2DriveFilePermissionMemberAddedV1 struct{}