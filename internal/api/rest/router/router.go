package router

import (
	"frr-news/internal/api/rest/auth"
	"frr-news/internal/api/rest/controller"
	"frr-news/internal/infra/config"
	"frr-news/internal/infra/security/jwt"
	"frr-news/internal/infra/storage"

	_ "frr-news/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/swagger"
	"gopkg.in/reform.v1"
)

// @title           News service
// @version         0.0.1
// @description     This is a testing task for implementing JSON REST API with fiber and reform.
// @termsOfService  http://swagger.io/terms/

// @contact.name   	Igor
// @contact.url    	http://github.com/tnsoftbear
// @contact.email  	myg0t@inbox.lv

// @license.name  	MIT
// @license.url   	https://rem.mit-license.org/

// @host      		localhost:4000
// @BasePath  		/

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @externalDocs.description  REST API details
// @externalDocs.url          https://github.com/tnsoftbear/go_news_rest_api_tt
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
	pub.Get("/swagger/*", swagger.HandlerDefault)

	api := app.Group("", limiter.New(limitCfg))
	api.Use(auth.Handler(&config.Auth))
	api.Get("/list", func(ctx *fiber.Ctx) error { return controller.GetNewsList(ctx, newsRepo) })
	api.Post("/add", func(ctx *fiber.Ctx) error { return controller.PostNewsAdd(ctx, newsRepo) })
	api.Post("/add-category/:NewsId/:CatId", func(ctx *fiber.Ctx) error { return controller.PostNewsAddCategory(ctx, newsRepo) })
	api.Post("/edit/:Id", func(ctx *fiber.Ctx) error { return controller.PostNewsEditById(ctx, newsRepo) })
	api.Delete("/:NewsId", func(ctx *fiber.Ctx) error { return controller.DeleteNewsById(ctx, newsRepo) })
}
