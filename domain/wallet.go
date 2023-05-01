package domain

import "time"

type (
	Wallet struct {
		ID         string       `json:"id"`
		OwnedBy    string       `json:"owned_by"`
		Status     WalletStatus `json:"status"`
		EnabledAt  time.Time    `json:"enabled_at"`
		Balance    int64        `json:"balance"`
		CreatedAt  time.Time    `json:"create_at"`
		UpdateDate time.Time    `json:"update_at"`
	}

	WalletUseCase interface {
		EnableWallet(string) (*Wallet, error)
		DisableWallet(string) (*Wallet, error)
		GetWalletBalance(string) (*Wallet, error)
		DepositMoneyWallet(string, int64) (*Wallet, error)
		WithdrawMoneyWallet(string, int64) (*Wallet, error)
	}

	WalletRepository interface {
		CreateWallet(*Wallet) error
		GetWallet(string) (*Wallet, error)
		UpdateWallet(*Wallet) error
	}
)
