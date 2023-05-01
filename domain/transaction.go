package domain

import "time"

type (
	Transaction struct {
		ID         string    `json:"id"`
		OwnedBy    string    `json:"owned_by"`
		Status     TrxStatus `json:"status"`
		WalletID   string    `json:"by"`
		CreatedAt  time.Time `json:"create_at"`
		UpdateDate time.Time `json:"update_at"`
	}
)
