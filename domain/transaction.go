package domain

import "time"

type (
	Transaction struct {
		ID        string    `json:"id"`
		OwnedBy   string    `json:"owned_by"`
		Status    TrxStatus `json:"status"`
		WalletID  string    `json:"wallet_id"`
		Amount    int64     `json:"amount"`
		CreatedAt time.Time `json:"create_at"`
	}
)
