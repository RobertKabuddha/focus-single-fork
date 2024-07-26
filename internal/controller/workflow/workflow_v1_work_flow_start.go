package workflow

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"focus-single/api/workflow/v1"
)

func (c *ControllerV1) WorkFlowStart(ctx context.Context, req *v1.WorkFlowStartReq) (res *v1.WorkFlowStartRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
