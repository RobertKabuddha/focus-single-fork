package v1

import "github.com/gogf/gf/v2/frame/g"

type WorkFlowListReq struct {
	g.Meta `path:"/workflow/{UserId}" method:"get" summary:"查看用户自己的工作流" tags:"工作流"`
	UserId uint `json:"userId" in:"path" dc:"用户ID"`
}

type WorkFlowListRes struct {
	Id                string `json:"id"`
	FlowOrder         string `json:"flowOrder"         dc:"TDP系统流程Id"`
	ApplyType         string `json:"applyType"         dc:"流程对应类型(客户禁用、客户合并详见代码枚举)"`
	FlowStatus        string `json:"flowStatus"        dc:"流程状态(1:pending,2:complete-reject,3:complete-approved)"`
	SubmitUser        string `json:"submitUser"        dc:"提交申请用户itcode"`
	RevokedDate       string `json:"revokedDate"       dc:"流程撤回时间"`
	RevokedReason     string `json:"revokedReason"     dc:"撤回理由"`
	RequestJson       string `json:"requestJson"       dc:"请求数据参数json字符串"`
	CreateName        string `json:"createName"        dc:"创建人"`
	CreateDate        string `json:"createDate"        dc:"创建时间"`
	UpdateName        string `json:"updateName"        dc:"更新人"`
	UpdateDate        string `json:"updateDate"        dc:"更新时间"`
	Deleted           string `json:"deleted"           dc:""`
	SubmitDate        string `json:"submitDate"        dc:"提交时间"`
	EucId             string `json:"eucId"             dc:"cdbid"`
	CusFullName       string `json:"cusFullName"       dc:"客户全称"`
	CompanyName       string `json:"companyName"       dc:"发起申请所在bu"`
	Ucpid             string `json:"ucpid"             dc:"用户cdbid"`
	SubmitUserName    string `json:"submitUserName"    dc:"提交申请用户姓名"`
	ApplicationReason string `json:"applicationReason" dc:"发起申请原因"`
}

type WorkFlowStartReq struct {
	g.Meta            `path:"/wf/callback" method:"post" summary:"开启工作流" tags:"工作流"`
	ApplyItCode       string   `json:"applyItCode" v:"required#请输入开启工作流关键字" dc:"description"`
	DxpUserId         string   `json:"dxpUserId" v:"required#请输入开启工作流关键字"`
	ApplyUserName     string   `json:"applyUserName" v:"required#请输入开启工作流关键字"`
	ApplyIdentityType string   `json:"applyIdentityType" v:"required#请输入开启工作流关键字"`
	CompanyName       string   `json:"companyName" v:"required#请输入开启工作流关键字"`
	SalesUserList     []string `json:"salesUserList,omitempty"`
	GtcTeamUserList   []string `json:"gtcTeamUserList,omitempty"`
	WorkFlowKey       string   `json:"workFlowKey" v:"required#请输入开启工作流关键字"`
	EucId             string   `json:"eucId" v:"required#请输入开启工作流关键字"`
	EnvelopeId        string   `json:"envelopeId,omitempty"`
	ApplicationReason string   `json:"applicationReason,omitempty"`
}

type WorkFlowStartRes struct {
	FlowOrder string `json:"flowOrder"  dc:"TDP系统流程Id"`
}
