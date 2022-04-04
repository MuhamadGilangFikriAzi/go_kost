package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gokost.com/m/delivery/api"
	"gokost.com/m/manager"
)

type serverConfig struct {
	gin            *gin.Engine
	Name           string
	Port           string
	InfraManager   manager.InfraManager
	RepoManager    manager.RepoManager
	UseCaseManager manager.UseCaseManager
}

func (s *serverConfig) initHeader() {
	s.gin.GET("/test", api.Test())
}

func (s *serverConfig) Run() {
	s.initHeader()
	s.gin.Run(fmt.Sprintf("%s:%s", s.Name, s.Port))
}

func Server() *serverConfig {
	ginStart := gin.Default()
	infra := manager.NewInfraManager()
	repo := manager.NewRepoManager(infra.PostgreConn())
	usecase := manager.NewUseCaseManager(repo)
	return &serverConfig{
		ginStart, "localhost", "8000", infra, repo, usecase,
	}
}
