package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/tanyudii/balance-api/config"
	"github.com/tanyudii/balance-api/internal/adapters/api"
	"github.com/tanyudii/balance-api/internal/adapters/repositories"
	"github.com/tanyudii/balance-api/internal/domain/usecases"
	"github.com/tanyudii/balance-api/internal/pkg/gracefully"
	"gorm.io/gorm"
)

func main() {
	cfg := config.GetConfig()

	e := setupEcho()

	db := config.GetDatabase()
	defer config.CloseDatabase(db)

	registerRouter(e, db)

	gracefully.RunEchoGracefully(e, cfg.AppPort)
}

func setupEcho() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	cfg := config.GetConfig()
	e.Logger.SetLevel(echoLogLevel(cfg.AppLogLevel))
	return e
}

func echoLogLevel(level string) log.Lvl {
	switch level {
	case "debug":
		return log.DEBUG
	case "warn":
		return log.WARN
	case "error":
		return log.ERROR
	default:
		return log.INFO
	}
}

func registerRouter(e *echo.Echo, db *gorm.DB) {
	baseGroup := e.Group("/")

	accountRepo := repositories.NewAccountRepository(db)

	userUc := usecases.NewAccountUsecase(accountRepo)

	api.NewRegisterHealthAPI(baseGroup)
	api.NewRegisterAccountAPI(baseGroup, userUc)
}
