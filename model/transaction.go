package models

type Transaction struct {
	ID         string `json:"id"`
	MerchantID string `json:"merchant_id"`
	Status     string `json:"status"`
	Amount     string `json:"amount"`
	CreatedAt  string `json:"created_at"`
}
