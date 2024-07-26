// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"focus-single/internal/model"
)

type (
	IReply interface {
		// Create 创建回复
		Create(ctx context.Context, in model.ReplyCreateInput) error
		// Delete 删除回复(硬删除)
		Delete(ctx context.Context, id uint) error
		// 删除回复(硬删除)
		DeleteByUserContentId(ctx context.Context, userId uint, contentId uint) error
		// 获取回复列表
		GetList(ctx context.Context, in model.ReplyGetListInput) (out *model.ReplyGetListOutput, err error)
	}
)

var (
	localReply IReply
)

func Reply() IReply {
	if localReply == nil {
		panic("implement not found for interface IReply, forgot register?")
	}
	return localReply
}

func RegisterReply(i IReply) {
	localReply = i
}
