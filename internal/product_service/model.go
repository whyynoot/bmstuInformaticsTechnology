package product_service

// ProductStorage represents product in system with all information about it
type ProductStorage struct {
	ID          string
	NameRu      string
	ApiName     string
	Description string
	ImagePath   string
	Price       int64
	CategoryID  int
}

// UserProduct represents only a product information only for user purpose
type UserProduct struct {
	NameRu      string `json:"name"`
	ApiName     string `json:"api_name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	ImagePath   string `json:"image_path"`
}
