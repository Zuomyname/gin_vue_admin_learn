package system

import (
	"gin_vue_admin_learn/service"
)

type ApiGroup struct {
	SystemApiApi
}

var (
	apiService = service.ServiceGroupApp.SystemServiceGroup.ApiService
)
