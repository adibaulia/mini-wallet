package domain

type WalletStatus string
type TrxStatus string

var (
	DISABLE  WalletStatus = "disabled"
	ENABLED  WalletStatus = "enabled"
	WITHDRAW TrxStatus    = "withdraw"
	DEPOSIT  TrxStatus    = "deposit"
)
