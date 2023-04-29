package domain

type (
	Transaction struct {
		ID          string `json:"id"`
		OwnedBy     string `json:"owned_by"`
		Status      Status `json:"status"`
		By          string `json:"by"`
		DateCreated string `json:"date_created"`
	}
)
