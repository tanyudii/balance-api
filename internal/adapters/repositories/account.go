package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/tanyudii/balance-api/internal/domain/entities"
	"github.com/tanyudii/balance-api/internal/domain/repositories"
	"gorm.io/gorm"
)

type Account struct {
	db *gorm.DB
}

var _ repositories.Account = (*Account)(nil)

func NewAccountRepository(db *gorm.DB) *Account {
	return &Account{db}
}

func (r *Account) GetAccountByNIK(ctx context.Context, nik string) (*entities.Account, error) {
	var data entities.Account
	err := r.db.WithContext(ctx).Where("nik = ?", nik).First(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repositories.ErrorAccountNotFound
		}
		return nil, fmt.Errorf("error when get account by nik: %w", err)
	}
	return &data, nil
}

func (r *Account) GetAccountByNoHp(ctx context.Context, noHp string) (*entities.Account, error) {
	var data entities.Account
	err := r.db.WithContext(ctx).Where("no_hp = ?", noHp).First(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repositories.ErrorAccountNotFound
		}
		return nil, fmt.Errorf("error when get account by no hp: %w", err)
	}
	return &data, nil
}

func (r *Account) GetAccountByNoRekening(ctx context.Context, noRekening string) (*entities.Account, error) {
	var data entities.Account
	err := r.db.WithContext(ctx).Where("no_rekening = ?", noRekening).First(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repositories.ErrorAccountNotFound
		}
		return nil, fmt.Errorf("error when get account by no rekening: %w", err)
	}
	return &data, nil
}

func (r *Account) CreateAccount(ctx context.Context, data *entities.Account) error {
	if err := r.db.WithContext(ctx).Create(data).Error; err != nil {
		return fmt.Errorf("error when create account: %w", err)
	}
	return nil
}

func (r *Account) LogBalance(ctx context.Context, id uint, amount float64) error {
	trx := r.db.WithContext(ctx).Begin()

	var err error
	defer func() {
		if err != nil {
			trx.Rollback()
		}
	}()

	if err = trx.Model(Account{}).
		Where("id = ?", id).
		Update("saldo", gorm.Expr("saldo + ?", amount)).Error; err != nil {
		return fmt.Errorf("error when log balance: %w", err)
	}

	if err = trx.Create(&entities.AccountMutation{
		AccountID: id,
		Amount:    amount,
	}).Error; err != nil {
		return fmt.Errorf("error when log balance: %w", err)
	}

	if err = trx.Commit().Error; err != nil {
		trx.Rollback()
		return fmt.Errorf("error when log balance: %w", err)
	}

	return err
}
