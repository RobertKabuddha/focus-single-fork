// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IInteract interface {
		// 赞
		Zan(ctx context.Context, targetType string, targetId uint) error
		// 取消赞
		CancelZan(ctx context.Context, targetType string, targetId uint) error
		// 我是否有对指定内容赞
		DidIZan(ctx context.Context, targetType string, targetId uint) (bool, error)
		// 踩
		Cai(ctx context.Context, targetType string, targetId uint) error
		// 取消踩
		CancelCai(ctx context.Context, targetType string, targetId uint) error
		// 我是否有对指定内容踩
		DidICai(ctx context.Context, targetType string, targetId uint) (bool, error)
	}
)

var (
	localInteract IInteract
)

func Interact() IInteract {
	if localInteract == nil {
		panic("implement not found for interface IInteract, forgot register?")
	}
	return localInteract
}

func RegisterInteract(i IInteract) {
	localInteract = i
}
