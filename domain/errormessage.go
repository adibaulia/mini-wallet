package domain

import "fmt"

var (
	ErrWalletAlreadyEnabled     = fmt.Errorf("Wallet already enabled")
	ErrWalletAlreadyDisabled    = fmt.Errorf("Wallet already disabled")
	ErrWalletNotFound           = fmt.Errorf("Wallet not found")
	ErrWalletMustEnabled        = fmt.Errorf("Wallet must be enabled")
	ErrWalletInsufficantBalance = fmt.Errorf("Insufficant balance")
)
