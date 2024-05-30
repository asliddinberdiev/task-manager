package main

import (
	"github.com/asliddinberdiev/task-manager/conf"
	"github.com/asliddinberdiev/task-manager/pkgs/handlers"
	"github.com/asliddinberdiev/task-manager/pkgs/repositories"
	"github.com/asliddinberdiev/task-manager/pkgs/server"
	"github.com/asliddinberdiev/task-manager/pkgs/services"
	"github.com/asliddinberdiev/task-manager/pkgs/utils"
	"go.uber.org/zap"
	_ "github.com/lib/pq"
)

func main() {
	// initialize envs
	cfg := conf.Load()
	if cfg.AppPort == "" {
		cfg.AppPort = "9000"
	}

	// initialize logger
	utils.InitLogger(cfg.AppMode)
	defer utils.Logger.Sync()

	// initialize db
	db, postgres_err := repositories.NewPostgresDB(repositories.PostgresConfig{
		User:     cfg.Postgres.User,
		Password: cfg.Postgres.Password,
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		Name:     cfg.Postgres.Name,
		SSLMode:  cfg.Postgres.SSLMode,
	})
	if postgres_err != nil {
		utils.Logger.Error("failed to initialize db", zap.Error(postgres_err))
	}

	// new repository
	repos := repositories.NewRepository(db)
	// new services
	services := services.NewService(repos)
	// new handlers
	handlers := handlers.NewHandler(services)

	// new server
	utils.Logger.Info("app run", zap.String("port", cfg.AppPort))
	srv := new(server.Server)
	if err := srv.Run(cfg.AppPort, handlers.InitRoutes()); err != nil {
		utils.Logger.Fatal("error running http server", zap.Error(err))
	}
}
