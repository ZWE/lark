// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// CreateTaskCollaborator 该接口用于新增任务协作者
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-collaborator/create
func (r *TaskService) CreateTaskCollaborator(ctx context.Context, request *CreateTaskCollaboratorReq, options ...MethodOptionFunc) (*CreateTaskCollaboratorResp, *Response, error) {
	if r.cli.mock.mockTaskCreateTaskCollaborator != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Task#CreateTaskCollaborator mock enable")
		return r.cli.mock.mockTaskCreateTaskCollaborator(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Task",
		API:                   "CreateTaskCollaborator",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/task/v1/tasks/:task_id/collaborators",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createTaskCollaboratorResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockTaskCreateTaskCollaborator(f func(ctx context.Context, request *CreateTaskCollaboratorReq, options ...MethodOptionFunc) (*CreateTaskCollaboratorResp, *Response, error)) {
	r.mockTaskCreateTaskCollaborator = f
}

func (r *Mock) UnMockTaskCreateTaskCollaborator() {
	r.mockTaskCreateTaskCollaborator = nil
}

type CreateTaskCollaboratorReq struct {
	UserIDType *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`,, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	TaskID     string  `path:"task_id" json:"-"`       // 任务 ID, 示例值："83912691-2e43-47fc-94a4-d512e03984fa"
	ID         string  `json:"id,omitempty"`           // 任务协作者的 ID, 示例值："ou_99e1a581b36ecc4862cbfbce473f1234"
}

type createTaskCollaboratorResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *CreateTaskCollaboratorResp `json:"data,omitempty"`
}

type CreateTaskCollaboratorResp struct {
	Collaborator *CreateTaskCollaboratorRespCollaborator `json:"collaborator,omitempty"` // 返回创建成功后的任务协作者
}

type CreateTaskCollaboratorRespCollaborator struct {
	ID string `json:"id,omitempty"` // 任务协作者的 ID
}