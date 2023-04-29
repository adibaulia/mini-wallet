package usecase

import "mini-wallet/domain"

type (
	accountUseCase struct {
		accountRepo domain.AccountRepository
	}
)

func NewAccountUseCase(accountRepo domain.AccountRepository) domain.AccountUseCase {
	return &accountUseCase{accountRepo}
}

func (c *accountUseCase) CreateAccount(xid string) (string, error) {
	return "", nil
}
func (c *accountUseCase) ValidateAccount(token string) bool {
	return false
}
