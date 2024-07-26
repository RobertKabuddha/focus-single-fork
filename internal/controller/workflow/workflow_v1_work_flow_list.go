package workflow

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"focus-single/api/workflow/v1"
)

func (c *ControllerV1) WorkFlowList(ctx context.Context, req *v1.WorkFlowListReq) (res *v1.WorkFlowListRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
