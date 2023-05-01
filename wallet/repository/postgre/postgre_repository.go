package postgre

import (
	"database/sql"
	"fmt"
	"mini-wallet/domain"

	"github.com/google/uuid"
)

type (
	postgreWalletRepo struct {
		DB *sql.DB
	}
)

func NewWalletRepository(db *sql.DB) domain.WalletRepository {
	return &postgreWalletRepo{db}
}

func (r *postgreWalletRepo) CreateWallet(wallet *domain.Wallet) error {
	wallet.ID = uuid.New().String()
	_, err := r.DB.Exec("INSERT INTO wallets (id, owned_by, status, enabled_at, balance) VALUES ($1, $2, $3, $4, $5)",
		wallet.ID, wallet.OwnedBy, wallet.Status, wallet.EnabledAt, wallet.Balance)
	if err != nil {
		return fmt.Errorf("failed to insert wallet: %w", err)
	}

	return nil
}

func (r *postgreWalletRepo) GetWallet(xid string) (*domain.Wallet, error) {
	var wallet domain.Wallet
	err := r.DB.QueryRow("SELECT id, owned_by, status, enabled_at, balance FROM wallets WHERE owned_by = $1", xid).
		Scan(&wallet.ID, &wallet.OwnedBy, &wallet.Status, &wallet.EnabledAt, &wallet.Balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get wallet: %w", err)
	}

	return &wallet, nil
}

func (r *postgreWalletRepo) UpdateWallet(wallet *domain.Wallet) error {
	_, err := r.DB.Exec("UPDATE wallets SET owned_by = $1, status = $3, enabled_at = $4, balance = $5, updated_at = NOW() WHERE owned_by = $1",
		wallet.OwnedBy, wallet.Status, wallet.EnabledAt, wallet.Balance)
	if err != nil {
		return fmt.Errorf("failed to update wallet: %w", err)
	}

	return nil
}
