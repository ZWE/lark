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

// EventV2IMMessageReadV1 用户阅读机器人发送的单聊消息后触发此事件。{使用示例}(url=/api/tools/api_explore/api_explore_config?project=im&version=v1&resource=message&event=message_read)
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/home/develop-a-bot-in-5-minutes/create-an-app)
// - 需要订阅 [消息与群组] 分类下的 [消息已读] 事件
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/events/message_read
func (r *EventCallbackService) HandlerEventV2IMMessageReadV1(f EventV2IMMessageReadV1Handler) {
	r.cli.eventHandler.eventV2IMMessageReadV1Handler = f
}

// EventV2IMMessageReadV1Handler event EventV2IMMessageReadV1 handler
type EventV2IMMessageReadV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2IMMessageReadV1) (string, error)

// EventV2IMMessageReadV1 ...
type EventV2IMMessageReadV1 struct {
	Reader        *EventV2IMMessageReadV1Reader `json:"reader,omitempty"`
	MessageIDList []string                      `json:"message_id_list,omitempty"` // 消息列表
}

// EventV2IMMessageReadV1Reader ...
type EventV2IMMessageReadV1Reader struct {
	ReaderID  *EventV2IMMessageReadV1ReaderReaderID `json:"reader_id,omitempty"`  // 用户 ID
	ReadTime  string                                `json:"read_time,omitempty"`  // 阅读时间
	TenantKey string                                `json:"tenant_key,omitempty"` // tenant key
}

// EventV2IMMessageReadV1ReaderReaderID ...
type EventV2IMMessageReadV1ReaderReaderID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 user ID
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}