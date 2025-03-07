package main

import (
	"log"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"sagepulse.ai/uhdy/user-service/handler"
	"sagepulse.ai/uhdy/user-service/repository"
	"sagepulse.ai/uhdy/utils/config"
	utils_logger "sagepulse.ai/uhdy/utils/logger"

	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
)

type UserConfig struct {
	JWT struct {
		Secret string `mapstructure:"secret"`
	} `mapstructure:"jwt"`
}

var ServiceCfg UserConfig

func main() {
	// Fiber app with logger middleware
	app := fiber.New()
	app.Use(logger.New())
	app.Use(utils_logger.InitLoggerMiddleware())

	// API with Huma
	api := humafiber.New(app, huma.DefaultConfig("Uhdy User API", "1.0.0"))

	dbCfg := config.DefaultCfg.Database
	repo := repository.NewUserPostgresRepository(dbCfg.User, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.Name)
	var serviceCfg UserConfig
	config.ReadConfig("user", "yaml", "/etc/uhdy", &serviceCfg)
	userHandler, err := handler.NewUserHandler(repo, serviceCfg.JWT.Secret)
	if err != nil {
		log.Fatalf("Fail to make a handler: %v", err)
	}

	// Register SignUp/SignIn with OpenAPI documentation
	huma.Register(api, huma.Operation{
		OperationID:   "signup",
		Method:        http.MethodPost,
		Path:          "/signup",
		Summary:       "Register new user",
		DefaultStatus: 201,
		Description:   "Register new user with ID/Password",
		Tags:          []string{"Auth"},
	}, userHandler.SignUp)

	huma.Register(api, huma.Operation{
		OperationID: "signin",
		Method:      http.MethodPost,
		Path:        "/signin",
		Summary:     "SignIn the service",
		Description: "SignIn with ID/Password",
		Tags:        []string{"Auth"},
	}, userHandler.SignIn)

	log.Fatal(app.Listen(config.DefaultCfg.Server.Port))
}
