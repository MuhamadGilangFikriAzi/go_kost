package main

import "gokost.com/m/manager"

type Server interface {
}

type serverConfig struct {
	InfraManager   manager.InfraManager
	RepoManager    manager.RepoManager
	UseCaseManager manager.UseCaseManager
}

func NewServer() Server {
	infra := manager.NewInfraManager()
	repo := manager.NewRepoManager(infra.PostgreConn())
	usecase := manager.NewUseCaseManager(repo)
	return &serverConfig{
		infra, repo, usecase,
	}
}
