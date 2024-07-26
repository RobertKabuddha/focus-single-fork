// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// WfMainFlow is the golang structure for table wf_main_flow.
type WfMainFlow struct {
	Id                int         `json:"id"                orm:"id"                 description:""`
	FlowOrder         string      `json:"flowOrder"         orm:"flow_order"         description:"TDP系统流程Id"`
	ApplyType         string      `json:"applyType"         orm:"apply_type"         description:"流程对应类型(客户禁用、客户合并详见代码枚举)"`
	FlowStatus        string      `json:"flowStatus"        orm:"flow_status"        description:"流程状态(1:pending,2:complete-reject,3:complete-approved)"`
	SubmitUser        string      `json:"submitUser"        orm:"submit_user"        description:"提交申请用户itcode"`
	RevokedDate       *gtime.Time `json:"revokedDate"       orm:"revoked_date"       description:"流程撤回时间"`
	RevokedReason     string      `json:"revokedReason"     orm:"revoked_reason"     description:"撤回理由"`
	RequestJson       string      `json:"requestJson"       orm:"request_json"       description:"请求数据参数json字符串"`
	CreateName        string      `json:"createName"        orm:"create_name"        description:"创建人"`
	CreateDate        *gtime.Time `json:"createDate"        orm:"create_date"        description:"创建时间"`
	UpdateName        string      `json:"updateName"        orm:"update_name"        description:"更新人"`
	UpdateDate        *gtime.Time `json:"updateDate"        orm:"update_date"        description:"更新时间"`
	Deleted           int         `json:"deleted"           orm:"deleted"            description:""`
	SubmitDate        *gtime.Time `json:"submitDate"        orm:"submit_date"        description:"提交时间"`
	EucId             string      `json:"eucId"             orm:"euc_id"             description:"cdbid"`
	CusFullName       string      `json:"cusFullName"       orm:"cus_full_name"      description:"客户全称"`
	CompanyName       string      `json:"companyName"       orm:"company_name"       description:"发起申请所在bu"`
	Ucpid             string      `json:"ucpid"             orm:"ucpid"              description:"用户cdbid"`
	SubmitUserName    string      `json:"submitUserName"    orm:"submit_user_name"   description:"提交申请用户姓名"`
	ApplicationReason string      `json:"applicationReason" orm:"application_reason" description:"发起申请原因"`
}
