// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// WfMainFlow is the golang structure of table wf_main_flow for DAO operations like Where/Data.
type WfMainFlow struct {
	g.Meta            `orm:"table:wf_main_flow, do:true"`
	Id                interface{} //
	FlowOrder         interface{} // TDP系统流程Id
	ApplyType         interface{} // 流程对应类型(客户禁用、客户合并详见代码枚举)
	FlowStatus        interface{} // 流程状态(1:pending,2:complete-reject,3:complete-approved)
	SubmitUser        interface{} // 提交申请用户itcode
	RevokedDate       *gtime.Time // 流程撤回时间
	RevokedReason     interface{} // 撤回理由
	RequestJson       interface{} // 请求数据参数json字符串
	CreateName        interface{} // 创建人
	CreateDate        *gtime.Time // 创建时间
	UpdateName        interface{} // 更新人
	UpdateDate        *gtime.Time // 更新时间
	Deleted           interface{} //
	SubmitDate        *gtime.Time // 提交时间
	EucId             interface{} // cdbid
	CusFullName       interface{} // 客户全称
	CompanyName       interface{} // 发起申请所在bu
	Ucpid             interface{} // 用户cdbid
	SubmitUserName    interface{} // 提交申请用户姓名
	ApplicationReason interface{} // 发起申请原因
}
