package postgre

import (
	"database/sql"
	"mini-wallet/domain"
)

type (
	postgreAccountRepo struct {
		DB *sql.DB
	}
)

func NewAccountRepository(db *sql.DB) domain.AccountRepository {
	return &postgreAccountRepo{db}
}

func (repo *postgreAccountRepo) CreateAccount(a domain.Account) error {
	return nil
}
func (repo *postgreAccountRepo) GetAccount(string) (domain.Account, error) {
	return domain.Account{}, nil
}
