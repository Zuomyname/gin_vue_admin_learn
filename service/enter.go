package service

import "gin_vue_admin_learn/service/system"

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
