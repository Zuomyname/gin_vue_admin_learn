package core

import (
	"flag"
	"fmt"
	"gin_vue_admin_learn/global"
	"gin_vue_admin_learn/utils"
	"github.com/fsnotify/fsnotify"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file")
		flag.Parse()
		if config == "" {
			if configEnv := os.Getenv(utils.ConfigEnv); configEnv == "" {
				config = utils.ConfigFile
				fmt.Printf("你正在使用的是默认的配置文件：%v\n", utils.ConfigFile)
			} else {
				config = configEnv
				fmt.Printf("你正在使用 GVA_CONFIG 环境变量 ：%v\n", configEnv)
			}
		} else {
			fmt.Printf("你正在使用命令行-c参数传递的值，config的路径是%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("你在正在使用 func Viper() 传递的值， config的路径是%v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed: ", e.Name)
		if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&global.GVA_CONFIG); err != nil {
		fmt.Println(err)
	}

	global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	global.BlackCache = local_cache.NewCache(
		//local_cache.SetDefaultExpire(time.Second * time.Duration(global.GVA_CONFIG.JWT.ExpiresTime)),
	)
	return v

}
