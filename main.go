package main

import (
	"gin_vue_admin_learn/core"
	"gin_vue_admin_learn/global"
	"gin_vue_admin_learn/initialize"
)

func main() {
	global.GVA_VP = core.Viper()
	global.GVA_LOG = core.Zap()
	global.GVA_DB = initialize.Gorm()
	if global.GVA_DB != nil {
		initialize.RegisterTables(global.GVA_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
