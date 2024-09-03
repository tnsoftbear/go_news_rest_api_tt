package bootstrap

import (
	"fiber-reform-rest/internal/api/rest/router"
	"fiber-reform-rest/internal/infra/config"
	"fiber-reform-rest/internal/infra/env"
	"fiber-reform-rest/internal/infra/storage"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// Run initializes and starts web service with REST API
func Run() {
	config := readConfig()
	app := setupApp(config)

	go func() {
		listenAddr := fmt.Sprintf("%s:%d", config.App.Host,	config.App.Port)
		if err := app.Listen(listenAddr); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	fmt.Println("Gracefully shutting down...")
	app.Shutdown()
	fmt.Println("Application was successful shutdown.")
}

func readConfig() *config.Config {
	configPath := flag.String("config", "./config/core.yaml", "load configurations from a file")
	flag.Parse()
	
	config, err := config.NewConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	return config
}

func setupApp(config *config.Config) *fiber.App {
	env.Setup()
	reformDB := storage.Setup(&config.MysqlStorage)
	app := fiber.New(fiber.Config{
		AppName:       config.App.Name,
		ServerHeader:  config.App.ServerHeader,
	})
	app.Use(recover.New())
	app.Use(logger.New())
	router.Setup(app, reformDB, config)
	return app
}