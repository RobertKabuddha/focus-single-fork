// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WfMainFlowDao is the data access object for table wf_main_flow.
type WfMainFlowDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns WfMainFlowColumns // columns contains all the column names of Table for convenient usage.
}

// WfMainFlowColumns defines and stores column names for table wf_main_flow.
type WfMainFlowColumns struct {
	Id                string //
	FlowOrder         string // TDP系统流程Id
	ApplyType         string // 流程对应类型(客户禁用、客户合并详见代码枚举)
	FlowStatus        string // 流程状态(1:pending,2:complete-reject,3:complete-approved)
	SubmitUser        string // 提交申请用户itcode
	RevokedDate       string // 流程撤回时间
	RevokedReason     string // 撤回理由
	RequestJson       string // 请求数据参数json字符串
	CreateName        string // 创建人
	CreateDate        string // 创建时间
	UpdateName        string // 更新人
	UpdateDate        string // 更新时间
	Deleted           string //
	SubmitDate        string // 提交时间
	EucId             string // cdbid
	CusFullName       string // 客户全称
	CompanyName       string // 发起申请所在bu
	Ucpid             string // 用户cdbid
	SubmitUserName    string // 提交申请用户姓名
	ApplicationReason string // 发起申请原因
}

// wfMainFlowColumns holds the columns for table wf_main_flow.
var wfMainFlowColumns = WfMainFlowColumns{
	Id:                "id",
	FlowOrder:         "flow_order",
	ApplyType:         "apply_type",
	FlowStatus:        "flow_status",
	SubmitUser:        "submit_user",
	RevokedDate:       "revoked_date",
	RevokedReason:     "revoked_reason",
	RequestJson:       "request_json",
	CreateName:        "create_name",
	CreateDate:        "create_date",
	UpdateName:        "update_name",
	UpdateDate:        "update_date",
	Deleted:           "deleted",
	SubmitDate:        "submit_date",
	EucId:             "euc_id",
	CusFullName:       "cus_full_name",
	CompanyName:       "company_name",
	Ucpid:             "ucpid",
	SubmitUserName:    "submit_user_name",
	ApplicationReason: "application_reason",
}

// NewWfMainFlowDao creates and returns a new DAO object for table data access.
func NewWfMainFlowDao() *WfMainFlowDao {
	return &WfMainFlowDao{
		group:   "default",
		table:   "wf_main_flow",
		columns: wfMainFlowColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *WfMainFlowDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *WfMainFlowDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *WfMainFlowDao) Columns() WfMainFlowColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *WfMainFlowDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *WfMainFlowDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *WfMainFlowDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
