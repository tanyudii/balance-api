package repositories

import (
	"context"
	"github.com/tanyudii/balance-api/internal/domain/entities"
	"github.com/tanyudii/balance-api/internal/pkg/errutil"
)

var (
	ErrorAccountNotFound = errutil.NewNotFoundError("akun tidak ditemukan")
)

type Account interface {
	GetAccountByNIK(ctx context.Context, nik string) (*entities.Account, error)
	GetAccountByNoHp(ctx context.Context, noHp string) (*entities.Account, error)
	GetAccountByNoRekening(ctx context.Context, noRekening string) (*entities.Account, error)
	CreateAccount(ctx context.Context, data *entities.Account) error
	LogBalance(ctx context.Context, id uint, amount float64) error
}
