package model

type Merchant struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	APIKey    string `json:"api_key"`
	CreatedAt string `json:"created_at"`
}
