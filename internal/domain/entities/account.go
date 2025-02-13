package entities

import "time"

type Account struct {
	ID         uint      `gorm:"primaryKey;autoIncrement"`
	Nama       string    `gorm:"type:varchar(255);not null"`
	Nik        string    `gorm:"type:varchar(20);unique;not null"`
	NoHp       string    `gorm:"type:varchar(15);unique;not null"`
	NoRekening string    `gorm:"type:varchar(30);unique;not null"`
	Saldo      float64   `gorm:"type:numeric(15,2);default:0;not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

type AccountMutation struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	AccountID uint      `gorm:"not null;index"`
	Amount    float64   `gorm:"type:numeric(15,2);not null;check:amount > 0"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
