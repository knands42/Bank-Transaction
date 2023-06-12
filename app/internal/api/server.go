package api

import (
	"fmt"

	db "github.com/caiofernandes00/Database-Transactions-Simulation.git/app/internal/db/sqlc"
	"github.com/caiofernandes00/Database-Transactions-Simulation.git/app/internal/token"
	"github.com/caiofernandes00/Database-Transactions-Simulation.git/app/internal/token/factory"
	"github.com/caiofernandes00/Database-Transactions-Simulation.git/app/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config        util.Config
	store         db.Store
	tokenMaker    token.Maker
	router        *gin.Engine
	hashingConfig *util.HashingConfig
}

func NewServer(config util.Config, store db.Store, hashingConfig *util.HashingConfig) (*Server, error) {
	tokenMaker, err := factory.TokenFactory(&config)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:        config,
		store:         store,
		tokenMaker:    tokenMaker,
		hashingConfig: hashingConfig,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.POST("/tokens/renew_access", server.renewAccessToken)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts/", server.listAccount)

	authRoutes.POST("/transfers/", server.createTransfer)

	server.router = router
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
