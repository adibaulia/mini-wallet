package domain

type Status string

var (
	DISABLE  Status = "disabled"
	ENABLED  Status = "enabled"
	WITHDRAW Status = "withdraw"
	DEPOSIT  Status = "deposit"
)
