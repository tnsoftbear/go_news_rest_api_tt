package bootstrap

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"frr-news/internal/api/rest/router"
	"frr-news/internal/infra/config"
	"frr-news/internal/infra/env"
	"frr-news/internal/infra/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// Run initializes and starts web service with REST API.
// Graceful shutdown considered.
func Run() {
	cfg := readConfig()
	app := setupApp(cfg)

	go func() {
		listenAddr := fmt.Sprintf("%s:%d", cfg.App.Host, cfg.App.Port)
		if err := app.Listen(listenAddr); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	fmt.Println("Gracefully shutting down...")
	if err := app.Shutdown(); err != nil {
		log.Fatal("Application shutdown failed with error:", err)
	}
	fmt.Println("Application was successful shutdown.")
}

func readConfig() *config.Config {
	configPath := flag.String("config", "./config/core.yaml", "load configurations from a file")
	flag.Parse()

	cfg, err := config.NewConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}

func setupApp(cfg *config.Config) *fiber.App {
	env.Setup()
	reformDB := storage.Setup(&cfg.MysqlStorage)
	app := fiber.New(fiber.Config{
		AppName:      cfg.App.Name,
		ServerHeader: cfg.App.ServerHeader,
	})
	app.Use(recover.New())
	app.Use(logger.New())
	router.Setup(app, reformDB, cfg)
	return app
}
