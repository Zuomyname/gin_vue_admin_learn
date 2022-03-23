package system

import (
	"gin_vue_admin_learn/model/common/response"
	"gin_vue_admin_learn/model/system"
	"gin_vue_admin_learn/utils"
	"github.com/gin-gonic/gin"
)

type SystemApiApi struct {
}

func (s *SystemApiApi) CreateApi(c *gin.Context) {
	var api system.SysApi
	_ = c.ShouldBindJSON(&api)

	if err := utils.Verify(api, utils.ApiVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := apiService.CreateApi(api); err != nil {
		response.FailWithMessage("创建失败", c)
		return
	}

	response.OkWithMessage("创建成功", c)
}
