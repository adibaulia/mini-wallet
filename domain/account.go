package domain

type (
	Account struct {
		ID          string `json:"id"`
		CustomerXID string `json:"customer_xid"`
		Token       string `json:"token"`
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
