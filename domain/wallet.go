package domain

import "time"

type (
	Wallet struct {
		ID         string       `json:"id"`
		OwnedBy    string       `json:"owned_by"`
		Status     WalletStatus `json:"status"`
		EnabledAt  time.Time    `json:"enabled_at"`
		Balance    int64        `json:"balance"`
		CreatedAt  time.Time    `json:"created_at,omitempty"`
		UpdateDate time.Time    `json:"updated_at,omitempty"`
	}

	WalletUseCase interface {
		EnableWallet(string) (*Wallet, error)
		DisableWallet(string) (*Wallet, error)
		GetWalletBalance(string) (*Wallet, error)
		DepositMoneyWallet(string, string, int64) (*Wallet, error)
		WithdrawMoneyWallet(string, string, int64) (*Wallet, error)
		GetWalletTransactions(string) ([]Transaction, error)
	}

	WalletRepository interface {
		CreateWallet(*Wallet) error
		GetWallet(string) (*Wallet, error)
		UpdateWallet(*Wallet) error
		UpdateBalance(*Wallet, string, TrxStatus, int64) error
		GetWalletTransactions(string) ([]Transaction, error)
	}
)
