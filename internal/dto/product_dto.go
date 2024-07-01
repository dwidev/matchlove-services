package dto

type (
	ProductDTO struct {
		ProductID   string `json:"productId"`
		ProductName string `json:"productName" validate:"required"`
		ProductCode string
		Description string              `json:"description" validate:"required"`
		Stock       int                 `json:"stock" validate:"required"`
		Price       float64             `json:"price" validate:"required"`
		ImageUrl    []string            `json:"imageUrl" validate:"required"`
		MerchantId  string              `json:"-"`
		Merchant    *MerchantProfileDTO `json:"merchant"`
	}
)
