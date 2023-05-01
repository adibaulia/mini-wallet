package domain

type (
	Wallet struct {
		ID        string       `json:"id"`
		OwnedBy   string       `json:"owned_by"`
		Status    WalletStatus `json:"status"`
		EnabledAt string       `json:"enabled_at"`
		Balance   int64        `json:"balance"`
	}

	WalletUseCase interface {
		ChangeWalletStatus(string, WalletStatus) (*Wallet, error)
		GetWalletBalance(string) (*Wallet, error)
		UpdateMoneyWallet(string, TrxStatus, int64) (*Wallet, error)
	}

	WalletRepository interface {
		CreateWallet(Wallet) error
		GetWallet(string) (*Wallet, error)
		UpdateWallet(Wallet) error
	}
)
