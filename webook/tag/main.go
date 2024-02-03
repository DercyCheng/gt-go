package main

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	initViperV2Watch()
	// 我要在初始化的过程中，把缓存加载好
	// 谁来加载？
	// dao, cache, repository, svc 谁来加载？
	app := Init()
	err := app.GRPCServer.Serve()
	if err != nil {
		panic(err)
	}
}

func initViperV2Watch() {
	cfile := pflag.String("config",
		"config/dev.yaml", "配置文件路径")
	pflag.Parse()
	// 直接指定文件路径
	viper.SetConfigFile(*cfile)
	viper.WatchConfig()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}