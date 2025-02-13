package api

import (
	"github.com/labstack/echo/v4"
	"github.com/tanyudii/balance-api/internal/domain/entities"
	"github.com/tanyudii/balance-api/internal/domain/usecases"
	"github.com/tanyudii/balance-api/internal/pkg/api"
)

type Account interface {
	Register(r *echo.Group)
}

type AccountAPI struct {
	accountUc usecases.Account
}

var _ Account = (*AccountAPI)(nil)

func NewRegisterAccountAPI(
	r *echo.Group,
	accountUc usecases.Account,
) Account {
	h := &AccountAPI{accountUc: accountUc}
	h.Register(r)
	return h
}

func (h *AccountAPI) Register(r *echo.Group) {
	r.POST("daftar", h.Daftar)
	r.POST("tabung", h.Tabung)
	r.POST("tarik", h.Tarik)
	r.GET("saldo/:no_rekening", h.Saldo)
}

func (h *AccountAPI) Daftar(c echo.Context) error {
	req := new(entities.AccountDaftarRequest)
	if err := c.Bind(req); err != nil {
		return api.EchoErrorResponse(c, err)
	}

	noRekening, err := h.accountUc.Register(c.Request().Context(), req)
	if err != nil {
		return api.EchoErrorResponse(c, err)
	}

	return api.EchoResponse(c, 201, map[string]string{
		"no_rekening": noRekening,
	})
}

func (h *AccountAPI) Tabung(c echo.Context) error {
	req := new(entities.AccountMutationRequest)
	if err := c.Bind(req); err != nil {
		return api.EchoErrorResponse(c, err)
	}

	balance, err := h.accountUc.Tabung(c.Request().Context(), req)
	if err != nil {
		return api.EchoErrorResponse(c, err)
	}

	return api.EchoResponse(c, 200, map[string]float64{
		"balance": balance,
	})
}

func (h *AccountAPI) Tarik(c echo.Context) error {
	req := new(entities.AccountMutationRequest)
	if err := c.Bind(req); err != nil {
		return api.EchoErrorResponse(c, err)
	}

	balance, err := h.accountUc.Tarik(c.Request().Context(), req)
	if err != nil {
		return api.EchoErrorResponse(c, err)
	}

	return api.EchoResponse(c, 200, map[string]float64{
		"balance": balance,
	})
}

func (h *AccountAPI) Saldo(c echo.Context) error {
	noRekening := c.Param("no_rekening")

	balance, err := h.accountUc.Saldo(c.Request().Context(), noRekening)
	if err != nil {
		return api.EchoErrorResponse(c, err)
	}

	return api.EchoResponse(c, 200, map[string]float64{
		"balance": balance,
	})
}
