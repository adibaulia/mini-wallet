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
func (repo *postgreAccountRepo) GetAccount(customerXID string) (*domain.Account, error) {
	// Query the database for the account with the given CustomerXID
	var account domain.Account
	err := repo.DB.QueryRow("SELECT id, customer_xid, token FROM accounts WHERE customer_xid = $1", customerXID).
		Scan(&account.ID, &account.CustomerXID, &account.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to query database: %w", err)
	}

	return &account, nil
}
