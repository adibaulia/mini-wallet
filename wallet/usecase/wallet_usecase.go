package usecase

import (
	"log"
	"mini-wallet/domain"
	"time"
)

type (
	walletUseCase struct {
		walletRepo     domain.WalletRepository
		accountUseCase domain.AccountUseCase
	}
)

func NewWalletUseCase(walletRepo domain.WalletRepository, accountUsecase domain.AccountUseCase) domain.WalletUseCase {
	return &walletUseCase{walletRepo, accountUsecase}
}

func (u *walletUseCase) EnableWallet(token string) (*domain.Wallet, error) {
	xid, err := u.accountUseCase.ValidateAccountToken(token)
	if err != nil {
		log.Printf("error validating account token: %v", err)
		return nil, err
	}

	wallet, err := u.walletRepo.GetWallet(xid)
	if err != nil {
		log.Printf("error getting wallet: %v", err)
		return nil, err
	}

	if wallet == nil {
		wallet = &domain.Wallet{
			OwnedBy:   xid,
			Status:    domain.ENABLED,
			EnabledAt: time.Now(),
			Balance:   0,
		}

		err := u.walletRepo.CreateWallet(wallet)
		if err != nil {
			log.Printf("error creating wallet: %v", err)
			return nil, err
		}

		return wallet, nil
	}

	if wallet.Status == domain.ENABLED {
		return nil, domain.ErrWalletAlreadyEnabled
	}

	wallet.EnabledAt = time.Now()
	wallet.Status = domain.ENABLED
	err = u.walletRepo.UpdateWallet(wallet)
	if err != nil {
		log.Printf("error when updating wallet: %v", err)
		return nil, err
	}

	return wallet, nil

}

func (u *walletUseCase) DisableWallet(token string) (*domain.Wallet, error) {
	xid, err := u.accountUseCase.ValidateAccountToken(token)
	if err != nil {
		log.Printf("error validating account token: %v", err)
		return nil, err
	}

	wallet, err := u.walletRepo.GetWallet(xid)
	if err != nil {
		log.Printf("error getting wallet: %v", err)
		return nil, err
	}

	if wallet == nil {
		return nil, domain.ErrWalletNotFound
	}

	if wallet.Status == domain.DISABLE {
		return nil, domain.ErrWalletAlreadyDisabled
	}

	wallet.Status = domain.DISABLE
	err = u.walletRepo.UpdateWallet(wallet)
	if err != nil {
		log.Printf("error when updating wallet: %v", err)
		return nil, err
	}

	return wallet, nil
}

func (u *walletUseCase) GetWalletBalance(token string) (*domain.Wallet, error) {
	xid, err := u.accountUseCase.ValidateAccountToken(token)
	if err != nil {
		log.Printf("error validating account token: %v", err)
		return nil, err
	}

	wallet, err := u.walletRepo.GetWallet(xid)
	if err != nil {
		log.Printf("error getting wallet: %v", err)
		return nil, err
	}

	if wallet == nil {
		return nil, domain.ErrWalletNotFound
	}

	if wallet.Status != domain.ENABLED {
		return nil, domain.ErrWalletMustEnabled
	}

	err = u.walletRepo.UpdateWallet(wallet)
	if err != nil {
		log.Printf("error when updating wallet: %v", err)
		return nil, err
	}

	return wallet, nil
}
func (u *walletUseCase) GetWalletTransactions(token string) ([]domain.Transaction, error) {
	xid, err := u.accountUseCase.ValidateAccountToken(token)
	if err != nil {
		log.Printf("error validating account token: %v", err)
		return nil, err
	}

	transactions, err := u.walletRepo.GetWalletTransactions(xid)
	if err != nil {
		log.Printf("error when getting wallet transactions: %v", err)
		return nil, err
	}
	return transactions, nil
}

func (u *walletUseCase) DepositMoneyWallet(token string, refID string, amount int64) (*domain.Wallet, error) {
	xid, err := u.accountUseCase.ValidateAccountToken(token)
	if err != nil {
		log.Printf("error validating account token: %v", err)
		return nil, err
	}

	wallet, err := u.walletRepo.GetWallet(xid)
	if err != nil {
		log.Printf("error getting wallet: %v", err)
		return nil, err
	}

	if wallet == nil {
		return nil, domain.ErrWalletNotFound
	}

	if wallet.Status != domain.ENABLED {
		return nil, domain.ErrWalletMustEnabled
	}

	wallet.Balance = wallet.Balance + amount

	err = u.walletRepo.UpdateBalance(wallet, refID, domain.DEPOSIT, amount)
	if err != nil {
		log.Printf("error when updating wallet: %v", err)
		return nil, err
	}

	return wallet, nil
}

func (u *walletUseCase) WithdrawMoneyWallet(token string, refID string, amount int64) (*domain.Wallet, error) {
	xid, err := u.accountUseCase.ValidateAccountToken(token)
	if err != nil {
		log.Printf("error validating account token: %v", err)
		return nil, err
	}

	wallet, err := u.walletRepo.GetWallet(xid)
	if err != nil {
		log.Printf("error getting wallet: %v", err)
		return nil, err
	}

	if wallet == nil {
		return nil, domain.ErrWalletNotFound
	}

	if wallet.Status != domain.ENABLED {
		return nil, domain.ErrWalletMustEnabled
	}

	if wallet.Balance-amount < 0 {
		return nil, domain.ErrWalletInsufficantBalance
	}
	wallet.Balance = wallet.Balance - amount

	err = u.walletRepo.UpdateBalance(wallet, refID, domain.WITHDRAW, amount)
	if err != nil {
		log.Printf("error when updating wallet: %v", err)
		return nil, err
	}
	return wallet, nil
}
