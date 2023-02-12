package product_service

type ProductServiceInterface interface {
	GetProducts() ([]ProductStorage, error)
	GetProduct(apiName string) (*ProductStorage, error)
	DeleteProductByID(id string) error
}
