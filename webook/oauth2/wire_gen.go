// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"gitee.com/geekbang/basic-go/webook/oauth2/grpc"
	"gitee.com/geekbang/basic-go/webook/oauth2/ioc"
	"github.com/google/wire"
)

// Injectors from wire.go:

func Init() *App {
	loggerV1 := ioc.InitLogger()
	service := ioc.InitPrometheus(loggerV1)
	oauth2ServiceServer := grpc.NewOauth2ServiceServer(service)
	server := ioc.InitGRPCxServer(oauth2ServiceServer)
	app := &App{
		server: server,
	}
	return app
}

// wire.go:

var thirdProvider = wire.NewSet(ioc.InitLogger)