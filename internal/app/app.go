package app

import (
	"l0/api"
	"l0/api/middleware"
	"l0/internal/db"
	"l0/internal/handler"
	"l0/internal/service"
	"context"
	"fmt"
	"l0/internal/cfg"
	"net/http"
	"time"
	"github.com/go-pg/pg"
	"github.com/rs/zerolog"
	// "github.com/sirupsen/logrus"
)

type appServer struct {
	config cfg.Cfg
	srv    *http.Server
	db     *pg.DB
	logger *zerolog.Logger
}

func NewServer(config cfg.Cfg, logger *zerolog.Logger) *appServer { 
	return &appServer{
		config: config,
		logger: logger,
	}
}

func (server *appServer) Serve(ctx context.Context) error {
	server.logger.Info().Msg("Starting server")
 
	a := server.config.GetDBString()

	dbPool := pg.Connect(&a)
	defer dbPool.Close()
	server.db=dbPool
	orderStorage := db.NewOrderStorage(dbPool)

	err, orders := orderStorage.MigrateDb()
	if err != nil{
		return err
	}
 
	orderService := service.NewOrderService(orderStorage,orders)
	orderHandler := handler.NewOrderHandler(orderService)
	routes := api.CreateRoutes(orderHandler) 
	routes.Use(middleware.RequestLog)                                                       

	server.srv = &http.Server{ 
		Addr:    "0.0.0.0:" + server.config.Port,
		Handler: routes,
	}

	server.logger.Info().Msg("Server started.")

	err = server.srv.ListenAndServe() 

	if err != nil && err != http.ErrServerClosed {
		server.logger.Err(err).Msg("Failure while serving")
		return err
	}

	return nil
}

func (server *appServer) Shutdown() error {
	server.logger.Info().Msg("server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	server.db.Close() 

	defer func() {
		cancel()
	}()

	var err error

	if err = server.srv.Shutdown(ctxShutDown); err != nil { 
		server.logger.Err(err)

		err = fmt.Errorf("server shutdown failed %w. ", err)

		return err
	}
	server.logger.Info().Msg("Shutdown!")

	return nil
}