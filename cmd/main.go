package main

import (
	"context"
	"github.com/Alexanderbr1/posts-log/internal/config"
	"github.com/Alexanderbr1/posts-log/internal/repository"
	"github.com/Alexanderbr1/posts-log/internal/server"
	"github.com/Alexanderbr1/posts-log/internal/service"
	"github.com/Alexanderbr1/posts-log/pkg/database"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	CONFIG_DIR  = "configs"
	CONFIG_FILE = "main"
)

func main() {
	cfg, err := config.New(CONFIG_DIR, CONFIG_FILE)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Ctx.Ttl)
	defer cancel()

	dbClient, err := database.NewMongoClient(ctx, database.ConnectionInfo{
		URI:      cfg.DB.URI,
		Username: cfg.DB.Username,
		Password: cfg.DB.Password,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := dbClient.Disconnect(ctx); err != nil {
			return
		}
	}()
	db := dbClient.Database(cfg.DB.Database)

	logRepo := repository.NewRepository(cfg, db)
	logService := service.NewService(logRepo)

	srv := server.NewServer(logService.Logs)
	go func() {
		if err := srv.Run(cfg.Server.Host, cfg.Server.Port); err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("Posts-log started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("Posts-log stopped")

	srv.Stop()
}
