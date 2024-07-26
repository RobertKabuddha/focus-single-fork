package content

import (
	"context"

	"focus-single/internal/service"
	"github.com/gogf/gf/v2/util/gutil"

	"focus-single/internal/consts"
	"focus-single/internal/dao"
	"focus-single/internal/model"
	"focus-single/internal/model/entity"
)

//自动生成Service
//   1.在internal/logic/模块/xx.go内写逻辑
//   2.开发/更新完成logic业务模块后，您需要手动执行一下 gf gen service 命令

type sWorkflow struct{}

func init() {
	service.RegisterWorkflow(New())
}

func New() *sWorkflow {
	return &sWorkflow{}
}

// GetList 查询内容列表
func (s *sWorkflow) GetList(ctx context.Context, in model.ContentGetListInput) (out *model.ContentGetListOutput, err error) {
	var (
		m = dao.Content.Ctx(ctx)
	)
	out = &model.ContentGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}
	// 默认查询topic
	if in.Type != "" {
		m = m.Where(dao.Content.Columns().Type, in.Type)
	} else {
		m = m.Where(dao.Content.Columns().Type, consts.ContentTypeTopic)
	}
	// 栏目检索
	if in.CategoryId > 0 {
		idArray, err := service.Category().GetSubIdList(ctx, in.CategoryId)
		if err != nil {
			return out, err
		}
		m = m.Where(dao.Content.Columns().CategoryId, idArray)
	}
	// 管理员可以查看所有文章
	if in.UserId > 0 && !service.User().IsAdmin(ctx, in.UserId) {
		m = m.Where(dao.Content.Columns().UserId, in.UserId)
	}
	// 分配查询
	listModel := m.Page(in.Page, in.Size)
	// 排序方式
	switch in.Sort {
	case consts.ContentSortHot:
		listModel = listModel.OrderDesc(dao.Content.Columns().ViewCount)

	case consts.ContentSortActive:
		listModel = listModel.OrderDesc(dao.Content.Columns().UpdatedAt)

	default:
		listModel = listModel.OrderDesc(dao.Content.Columns().Id)
	}
	// 执行查询
	var list []*entity.Content
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}
	// Content
	if err := listModel.ScanList(&out.List, "Content"); err != nil {
		return out, err
	}
	// Category
	err = dao.Category.Ctx(ctx).
		Fields(model.ContentListCategoryItem{}).
		Where(dao.Category.Columns().Id, gutil.ListItemValuesUnique(out.List, "Content", "CategoryId")).
		ScanList(&out.List, "Category", "Content", "id:CategoryId")
	if err != nil {
		return out, err
	}
	// User
	err = dao.User.Ctx(ctx).
		Fields(model.ContentListUserItem{}).
		Where(dao.User.Columns().Id, gutil.ListItemValuesUnique(out.List, "Content", "UserId")).
		ScanList(&out.List, "User", "Content", "id:UserId")
	if err != nil {
		return out, err
	}
	return
}
