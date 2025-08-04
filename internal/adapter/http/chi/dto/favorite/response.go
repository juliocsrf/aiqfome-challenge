package favorite

// FavoriteResponse representa a resposta após operação com favoritos
type FavoriteResponse struct {
	CustomerID string `json:"customer_id"`
	ProductID  int64  `json:"product_id"`
	Message    string `json:"message"`
}

// ErrorResponse representa uma resposta de erro
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

// SuccessResponse representa uma resposta de sucesso genérica
type SuccessResponse struct {
	Message string `json:"message"`
}
