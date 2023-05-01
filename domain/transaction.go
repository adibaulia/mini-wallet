package domain

type (
	Transaction struct {
		ID          string       `json:"id"`
		OwnedBy     string       `json:"owned_by"`
		Status      WalletStatus `json:"status"`
		By          string       `json:"by"`
		DateCreated string       `json:"date_created"`
	}
)
