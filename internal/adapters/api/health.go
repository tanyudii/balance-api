package api

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func NewRegisterHealthAPI(
	r *echo.Group,
) HealthAPI {
	h := &healthAPI{}
	h.Register(r)
	return h
}

type HealthAPI interface {
	Register(r *echo.Group)
}

type healthAPI struct {
	db *gorm.DB
}

func (h *healthAPI) Register(r *echo.Group) {
	r.GET("_health", h.HealthCheck)
}

func (h *healthAPI) HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}
