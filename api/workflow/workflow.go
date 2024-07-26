// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package workflow

import (
	"context"

	"focus-single/api/workflow/v1"
)

type IWorkflowV1 interface {
	WorkFlowList(ctx context.Context, req *v1.WorkFlowListReq) (res *v1.WorkFlowListRes, err error)
	WorkFlowStart(ctx context.Context, req *v1.WorkFlowStartReq) (res *v1.WorkFlowStartRes, err error)
}
