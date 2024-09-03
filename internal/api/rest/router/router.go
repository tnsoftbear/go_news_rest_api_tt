package router

import (
	"fiber-reform-rest/internal/api/rest/controller"
	"fiber-reform-rest/internal/api/rest/middleware"
	"fiber-reform-rest/internal/infra/config"
	"fiber-reform-rest/internal/infra/security/jwt"
	"fiber-reform-rest/internal/infra/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"gopkg.in/reform.v1"
)

func Setup(app *fiber.App, reformDB *reform.DB, config *config.Config) {
	newsRepo := storage.NewNewsRepositoryMysql(reformDB)
	jm := jwt.NewJWTManager(&config.Auth.Jwt)
	app.Use(logger.New())
	
	limitCfg := limiter.Config{
		Max: 60,
	}

	pub := app.Group("", limiter.New(limitCfg))
	pub.Get("/ping", controller.GetPing)
	pub.Get("/dashboard", monitor.New())
	pub.Post("/login", func(ctx *fiber.Ctx) error { return controller.PostLogin(ctx, jm) })

	api := app.Group("", limiter.New(limitCfg))
	api.Use(middleware.Auth(&config.Auth))
	api.Get("/list", func(ctx *fiber.Ctx) error { return controller.GetNewsList(ctx, newsRepo) })
	api.Post("/add", func(ctx *fiber.Ctx) error { return controller.PostNewsAdd(ctx, newsRepo) })
	api.Post("/add-category/:NewsId/:CatId", func(ctx *fiber.Ctx) error { return controller.PostNewsAddCategory(ctx, newsRepo) })
	api.Post("/edit/:Id", func(ctx *fiber.Ctx) error { return controller.PostNewsEditById(ctx, newsRepo) })
	api.Delete("/:NewsId", func(ctx *fiber.Ctx) error { return controller.DeleteNewsById(ctx, newsRepo) })
}
