package domain

type (
	Wallet struct {
		ID        string `json:"id"`
		OwnedBy   string `json:"owned_by"`
		Status    Status `json:"status"`
		EnabledAt string `json:"enabled_at"`
		Balance   int64  `json:"balance"`
	}
)
