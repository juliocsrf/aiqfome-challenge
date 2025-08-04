package customer

import "github.com/juliocsrf/aiqfome-challenge/internal/domain/entity"

type CustomerResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Favorites []Product `json:"favorites"`
}

type Product struct {
	ID    int64   `json:"id"`
	Title string  `json:"title"`
	Image string  `json:"image"`
	Price float64 `json:"price"`
}

func FromEntity(customer *entity.Customer) *CustomerResponse {
	favorites := make([]Product, len(customer.Favorites))
	for i, product := range customer.Favorites {
		favorites[i] = Product{
			ID:    product.Id,
			Title: product.Title,
			Image: product.Image,
			Price: product.Price,
		}
	}

	return &CustomerResponse{
		ID:        customer.Id,
		Name:      customer.Name,
		Email:     customer.Email,
		Favorites: favorites,
	}
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}
