package usecases

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/tanyudii/balance-api/internal/domain/entities"
	"github.com/tanyudii/balance-api/internal/domain/repositories"
	"github.com/tanyudii/balance-api/internal/pkg/errutil"
	"github.com/tanyudii/balance-api/internal/pkg/logger"
	"math/rand"
	"strconv"
)

type Account interface {
	Register(ctx context.Context, r *entities.AccountDaftarRequest) (string, error)
	Tabung(ctx context.Context, r *entities.AccountMutationRequest) (float64, error)
	Tarik(ctx context.Context, r *entities.AccountMutationRequest) (float64, error)
	Saldo(ctx context.Context, noRekening string) (float64, error)
}

type AccountUsecase struct {
	accountRepo repositories.Account
}

var _ Account = (*AccountUsecase)(nil)

func NewAccountUsecase(
	accountRepo repositories.Account,
) *AccountUsecase {
	return &AccountUsecase{accountRepo}
}

func (u *AccountUsecase) Register(ctx context.Context, r *entities.AccountDaftarRequest) (string, error) {
	if err := r.Validate(); err != nil {
		return "", err
	}

	account := &entities.Account{
		Nama:  r.Nama,
		Nik:   r.Nik,
		NoHp:  r.NoHp,
		Saldo: 0,
	}

	err := u.validateUniq(ctx, account)
	if err != nil {
		return "", err
	}

	account.NoRekening, err = u.genNoRekening(ctx)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"nama":  account.Nama,
			"nik":   account.Nik,
			"no_hp": account.NoHp,
		}).Errorf("error when generating no rekening: %v", err)

		return "", fmt.Errorf("error when generating no rekening: %w", err)
	}

	if err = u.accountRepo.CreateAccount(ctx, account); err != nil {
		logger.WithFields(logrus.Fields{
			"nama":  account.Nama,
			"nik":   account.Nik,
			"no_hp": account.NoHp,
		}).Errorf("error when creating account: %v", err)

		return "", err
	}

	logger.Infof("account created: %v", account.NoRekening)

	return account.NoRekening, nil
}

func (u *AccountUsecase) Tabung(ctx context.Context, r *entities.AccountMutationRequest) (float64, error) {
	if err := r.Validate(); err != nil {
		return 0, err
	}

	account, err := u.getAccountTrx(ctx, r.NoRekening)
	if err != nil {
		return 0, err
	}

	if err = u.accountRepo.LogBalance(ctx, account.ID, r.Nominal); err != nil {
		logger.WithFields(logrus.Fields{
			"field":       "no_rekening",
			"no_rekening": r.NoRekening,
		}).Errorf("error when tabung balance: %v", err)

		return 0, err
	}

	account.Saldo += r.Nominal

	logger.Infof("account balance logged [tabung]: %v", account)

	return account.Saldo, nil
}

func (u *AccountUsecase) Tarik(ctx context.Context, r *entities.AccountMutationRequest) (float64, error) {
	if err := r.Validate(); err != nil {
		return 0, err
	}

	account, err := u.getAccountTrx(ctx, r.NoRekening)
	if err != nil {
		return 0, err
	}

	if account.Saldo < r.Nominal {
		return 0, errutil.NewBadRequestError("saldo tidak cukup")
	}

	realValue := -r.Nominal
	if err = u.accountRepo.LogBalance(ctx, account.ID, realValue); err != nil {
		logger.WithFields(logrus.Fields{
			"field":       "no_rekening",
			"no_rekening": r.NoRekening,
		}).Errorf("error when tarik balance: %v", err)

		return 0, err
	}

	logger.Infof("account balance logged [tarik]: %v", account)

	return account.Saldo + realValue, nil
}

func (u *AccountUsecase) Saldo(ctx context.Context, noRekening string) (float64, error) {
	account, err := u.getAccountTrx(ctx, noRekening)
	if err != nil {
		return 0, err
	}

	return account.Saldo, nil
}

func (u *AccountUsecase) getAccountTrx(ctx context.Context, noRekening string) (*entities.Account, error) {
	account, err := u.accountRepo.GetAccountByNoRekening(ctx, noRekening)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"field":       "no_rekening",
			"no_rekening": noRekening,
		}).Warnf("error when getting account for trx by no rekening: %v", err)

		return nil, err
	}

	return account, nil
}

func (u *AccountUsecase) genNoRekening(ctx context.Context) (string, error) {
	low := 1000000000
	hi := 9999999999
	rand := strconv.Itoa(low + rand.Intn(hi-low))

	_, err := u.accountRepo.GetAccountByNoRekening(ctx, rand)
	if err == nil {
		return u.genNoRekening(ctx)
	}

	if !errors.Is(err, repositories.ErrorAccountNotFound) {
		logger.WithFields(logrus.Fields{
			"field": "no_rekening",
			"rand":  rand,
		}).Errorf("error when generating no rekening: %v", err)

		return "", err
	}

	return rand, nil
}

func (u *AccountUsecase) validateUniq(ctx context.Context, account *entities.Account) error {
	_, err := u.accountRepo.GetAccountByNIK(ctx, account.Nik)
	if err == nil {
		return errutil.NewBadRequestError("nik sudah terdaftar")
	} else if !errors.Is(err, repositories.ErrorAccountNotFound) {
		logger.WithFields(logrus.Fields{
			"field": "nik",
			"nik":   account.Nik,
		}).Errorf("error when validating unique: %v", err)

		return fmt.Errorf("error validating unique: %w", err)
	}

	_, err = u.accountRepo.GetAccountByNoHp(ctx, account.NoHp)
	if err == nil {
		return errutil.NewBadRequestError("no hp sudah terdaftar")
	} else if !errors.Is(err, repositories.ErrorAccountNotFound) {
		logger.WithFields(logrus.Fields{
			"field": "no_hp",
			"no_hp": account.NoHp,
		}).Errorf("error when validating unique: %v", err)

		return fmt.Errorf("error validating unique: %w", err)
	}

	return nil
}
