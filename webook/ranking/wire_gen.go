// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"gitee.com/geekbang/basic-go/webook/ranking/grpc"
	"gitee.com/geekbang/basic-go/webook/ranking/ioc"
	"gitee.com/geekbang/basic-go/webook/ranking/repository"
	"gitee.com/geekbang/basic-go/webook/ranking/repository/cache"
	"gitee.com/geekbang/basic-go/webook/ranking/service"
	"github.com/google/wire"
)

// Injectors from wire.go:

func Init() *App {
	interactiveServiceClient := ioc.InitInterActiveRpcClient()
	articleServiceClient := ioc.InitArticleRpcClient()
	cmdable := ioc.InitRedis()
	redisRankingCache := cache.NewRedisRankingCache(cmdable)
	rankingLocalCache := cache.NewRankingLocalCache()
	rankingRepository := repository.NewCachedRankingRepository(redisRankingCache, rankingLocalCache)
	rankingService := service.NewBatchRankingService(interactiveServiceClient, articleServiceClient, rankingRepository)
	rankingServiceServer := grpc.NewRankingServiceServer(rankingService)
	server := ioc.InitGRPCxServer(rankingServiceServer)
	app := &App{
		server: server,
	}
	return app
}

// wire.go:

var serviceProviderSet = wire.NewSet(cache.NewRankingLocalCache, cache.NewRedisRankingCache, repository.NewCachedRankingRepository, service.NewBatchRankingService)

var thirdProvider = wire.NewSet(ioc.InitRedis, ioc.InitInterActiveRpcClient, ioc.InitArticleRpcClient)
