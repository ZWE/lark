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

// MoveDriveFile 将文件或者文件夹移动到用户云空间的其他位置。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/move
func (r *DriveService) MoveDriveFile(ctx context.Context, request *MoveDriveFileReq, options ...MethodOptionFunc) (*MoveDriveFileResp, *Response, error) {
	if r.cli.mock.mockDriveMoveDriveFile != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#MoveDriveFile mock enable")
		return r.cli.mock.mockDriveMoveDriveFile(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "MoveDriveFile",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/files/:file_token/move",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(moveDriveFileResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveMoveDriveFile mock DriveMoveDriveFile method
func (r *Mock) MockDriveMoveDriveFile(f func(ctx context.Context, request *MoveDriveFileReq, options ...MethodOptionFunc) (*MoveDriveFileResp, *Response, error)) {
	r.mockDriveMoveDriveFile = f
}

// UnMockDriveMoveDriveFile un-mock DriveMoveDriveFile method
func (r *Mock) UnMockDriveMoveDriveFile() {
	r.mockDriveMoveDriveFile = nil
}

// MoveDriveFileReq ...
type MoveDriveFileReq struct {
	FileToken   string  `path:"file_token" json:"-"`    // 需要移动的文件token, 示例值: "boxcnrHpsg1QDqXAAAyachabcef"
	Type        *string `json:"type,omitempty"`         // 文件类型, 如果该值为空或者与文件实际类型不匹配, 接口会返回失败, 示例值: "file", 可选值有: `file`: 普通文件类型, `docx`: 新版文档类型, `bitable`: 多维表格类型, `doc`: 文档类型, `sheet`: 电子表格类型, `mindnote`: 思维笔记类型, `folder`: 文件夹类型
	FolderToken *string `json:"folder_token,omitempty"` // 目标文件夹token, 示例值: "fldbcO1UuPz8VwnpPx5a92abcef"
}

// MoveDriveFileResp ...
type MoveDriveFileResp struct {
	TaskID string `json:"task_id,omitempty"` // 异步任务id, 移动文件夹时返回
}

// moveDriveFileResp ...
type moveDriveFileResp struct {
	Code int64              `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string             `json:"msg,omitempty"`  // 错误描述
	Data *MoveDriveFileResp `json:"data,omitempty"`
}