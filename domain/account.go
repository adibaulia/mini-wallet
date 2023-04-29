package domain

type (
	Account struct {
		ID          string `json:"id"`
		CustomerXID string `json:"customer_xid"`
		Token       string `json:"token"`
	}
)
