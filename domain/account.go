package domain

import "time"

type (
	Account struct {
		ID          string    `json:"id"`
		CustomerXID string    `json:"customer_xid"`
		Token       string    `json:"token"`
		CreatedAt   time.Time `json:"created_at"`
		UpdateDate  time.Time `json:"updated_at"`
	}

	AccountUseCase interface {
		CreateAccount(string) (string, error)
		ValidateAccountToken(string) (string, error)
	}

	AccountRepository interface {
		CreateAccount(Account) error
		GetAccount(string) (*Account, error)
	}
)
