package main

import (
	"go-url-shortener/internal/config"
	"go-url-shortener/internal/library/logger/sl"
	"go-url-shortener/internal/storage/sqlite"
	"log/slog"
	"os"
)
const (
	envLocal = "local"
	envDev = "dev"
	envProd = "prod"
)

func main() {
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)		
	
	log.Info("Starting URL Shortener", slog.String("env", cfg.Env))
	log.Debug("Debug mode enabled")

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("Failed to initialize storage", sl.Error(err))
		os.Exit(1)
	}

	_ = storage 
	// TODO: init router: chi,
	// TODO: run server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),)
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),)
		
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),)
	}
	return log
}