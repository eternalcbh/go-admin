package users

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/global/consts"
	"go-admin/app/http/controller/web"
	"go-admin/app/http/validator/core/data_transfer"
	"go-admin/app/utils/response"
)

type Destroy struct {
	// 表单参数验证结构体支持匿名结构体嵌套、以及匿名结构体与普通字段组合
	Id
}

func (d Destroy) CheckParams(context *gin.Context) {
	if err := context.ShouldBind(&d); err != nil {
		// 将表单参数验证器出现的错误直接交给错误翻译器统一处理即可
		response.ValidatorError(context, err)
		return
	}

	//  该函数主要是将本结构体的字段（成员）按照 consts.ValidatorPrefix+ json标签对应的 键 => 值 形式绑定在上下文，便于下一步（控制器）可以直接通过 context.Get(键) 获取相关值
	extraAddBindDataContext := data_transfer.DataAddContext(d, consts.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, "UserShow表单参数验证器json化失败", "")
		return
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&web.Users{}).Destroy(extraAddBindDataContext)
	}
}
