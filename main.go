package main

import (
	"gin_vue_admin_learn/core"
	"gin_vue_admin_learn/global"
)

func main() {
	global.GVA_VP = core.Viper()
	global.GVA_LOG = core.Zap()

}
