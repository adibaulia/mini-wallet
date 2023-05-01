package postgre

import (
	"database/sql"
	"fmt"
	"mini-wallet/domain"

	"github.com/google/uuid"
	"github.com/lib/pq"
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

func (r *postgreWalletRepo) updateWalletTx(tx *sql.Tx, wallet *domain.Wallet) error {
	_, err := tx.Exec("UPDATE wallets SET owned_by = $1, status = $3, enabled_at = $4, balance = $5, updated_at = NOW() WHERE owned_by = $1",
		wallet.OwnedBy, wallet.Status, wallet.EnabledAt, wallet.Balance)
	if err != nil {
		return fmt.Errorf("failed to update wallet: %w", err)
	}

	return nil
}

func (r *postgreWalletRepo) UpdateBalance(wallet *domain.Wallet, refID string, trxStatus domain.TrxStatus, amount int64) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	err = r.updateWalletTx(tx, wallet)
	if err != nil {
		return err
	}

	err = r.InsertTrx(tx, wallet, refID, trxStatus, amount)
	if err != nil {
		return err
	}

	return nil

}

func (r *postgreWalletRepo) InsertTrx(tx *sql.Tx, wallet *domain.Wallet, refID string, trxStatus domain.TrxStatus, amount int64) error {
	err := r.createTransaction(tx, &domain.Transaction{
		ID:       refID,
		OwnedBy:  wallet.OwnedBy,
		Status:   trxStatus,
		WalletID: wallet.ID,
		Amount:   amount,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *postgreWalletRepo) createTransaction(tx *sql.Tx, trx *domain.Transaction) error {
	stmt := `INSERT INTO transactions (id, owned_by, status, wallet_id, amount, created_at)
             VALUES ($1, $2, $3, $4, $5, NOW())`
	_, err := tx.Exec(stmt, trx.ID, trx.OwnedBy, trx.Status, trx.WalletID, trx.Amount)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok && pgErr.Code == "23505" {
			return domain.ErrRefIDTransactionAlreadyExists
		}
		return err
	}
	return nil
}
