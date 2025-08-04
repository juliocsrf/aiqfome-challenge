package favorite

type CreateFavoriteRequest struct {
	CustomerID string `json:"customer_id" validate:"required,uuid"`
	ProductID  int64  `json:"product_id" validate:"required"`
}

type DeleteFavoriteRequest struct {
	CustomerID string `json:"customer_id" validate:"required,uuid"`
	ProductID  int64  `json:"product_id" validate:"required"`
}
