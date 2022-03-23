package system

import (
	"gin_vue_admin_learn/global"
	"gin_vue_admin_learn/model/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type ApiService struct {
}

var ApiServiceApp = new(ApiService)

// 创建 api
func (s *ApiService) CreateApi(api system.SysApi) (err error) {
	if errors.Is(global.GVA_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同的api")
	}
	return global.GVA_DB.Create(&api).Error
}
