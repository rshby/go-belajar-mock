package dto

type GetProductResponse struct {
	ID             int    `json:"id"`
	IdentityNumber string `json:"identity_number"`
	FullName       string `json:"full_name"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}
