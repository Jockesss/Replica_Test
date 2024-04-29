package app

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"replica/internal/config"
	"replica/internal/http/v1"
	"replica/internal/server"
	"syscall"
	"time"
)

func Run() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Printf(err.Error())
	}
	log.Printf("URL_SERVER ---> %s:%s", cfg.HTTP.Host, cfg.HTTP.Port)

	handler := v1.NewHandler()
	srv := server.NewServer(cfg, handler.Routes())

	log.Println("Starting server...")

	go func() {
		if err := srv.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Failed to start server: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Stop(ctx); err != nil {
		log.Fatalf("Failed to gracefully stop server: %s", err)
	} else {
		log.Println("Server stopped gracefully")
	}
}
