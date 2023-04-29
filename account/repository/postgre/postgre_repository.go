package postgre

import (
	"database/sql"
	"fmt"
	"mini-wallet/domain"

	"github.com/google/uuid"
)

type (
	postgreAccountRepo struct {
		DB *sql.DB
	}
)

func NewAccountRepository(db *sql.DB) domain.AccountRepository {
	return &postgreAccountRepo{db}
}

func (repo *postgreAccountRepo) CreateAccount(account domain.Account) error {
	id := uuid.New().String()
	query := `
        INSERT INTO accounts (id, customer_xid, token)
        VALUES ($1, $2, $3)
    `
	_, err := repo.DB.Exec(query, id, account.CustomerXID, account.Token)
	if err != nil {
		return fmt.Errorf("error inserting account: %w", err)
	}
	return nil
}
func (repo *postgreAccountRepo) GetAccount(string) (domain.Account, error) {
	return domain.Account{}, nil
}
