package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gokost.com/m/config"
	"gokost.com/m/delivery/api"
	"gokost.com/m/delivery/middleware"
	"gokost.com/m/manager"
)

type AppServer interface {
	Run()
}

type serverConfig struct {
	gin            *gin.Engine
	Name           string
	Port           string
	InfraManager   manager.InfraManager
	RepoManager    manager.RepoManager
	UseCaseManager manager.UseCaseManager
	Config         *config.Config
	Middleware     *middleware.AuthTokenMiddleware
}

func (s *serverConfig) initHeader() {
	s.gin.Use(s.Middleware.TokenAuthMiddleware())
	s.routeGroupApi()
}

func (s *serverConfig) routeGroupApi() {
	apiTesting := s.gin.Group("/testing")
	api.NewTestingApi(apiTesting)

	apiProduct := s.gin.Group("/product")
	api.NewProductApo(apiProduct, s.UseCaseManager.AllProductUseCase())

	apiGroupCustomer := s.gin.Group("/customer")
	api.NewCustomerApi(apiGroupCustomer, s.UseCaseManager.ListCustomerUseCase())

	apiGroupBoardingRoom := s.gin.Group("boarding_room")
	api.NewBoardingApi(apiGroupBoardingRoom, s.UseCaseManager.AvailableRoomUseCase())

	apiGroupTransaction := s.gin.Group("transaction")
	api.NewTransactionApi(apiGroupTransaction, s.UseCaseManager.InsertTransactionUseCase(), s.UseCaseManager.UpdateCustomerUseCase(), s.UseCaseManager.CustomerTransactionUseCase())

	apiLogin := s.gin.Group("login")
	api.NewLoginApi(apiLogin, s.UseCaseManager.LoginAdminUseCase(), s.Config.ConfigToken)
}

func (s *serverConfig) Run() {
	s.initHeader()
	s.gin.Run(fmt.Sprintf("%s:%s", s.Name, s.Port))
}

func Server() AppServer {
	ginStart := gin.Default()
	config := config.NewConfig()
	infra := manager.NewInfraManager(config.ConfigDatabase)
	repo := manager.NewRepoManager(infra.PostgreConn(), infra.MysqlConn())
	usecase := manager.NewUseCaseManager(repo)
	middleware := middleware.NewAuthTokenMiddleware(config.ConfigToken)
	return &serverConfig{
		ginStart,
		config.ConfigServer.Url,
		config.ConfigServer.Port,
		infra,
		repo,
		usecase,
		config,
		middleware,
	}
}
