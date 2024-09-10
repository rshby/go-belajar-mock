package dto

type CreateProductRequest struct {
	IdentityNumber string `json:"identity_number"`
	FullName       string `json:"full_name"`
}
