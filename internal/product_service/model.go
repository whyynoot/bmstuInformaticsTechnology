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

func (s *ProductStorage) ToUserProduct() UserProduct {
	return UserProduct{
		NameRu:      s.NameRu,
		ApiName:     s.ApiName,
		Description: s.Description,
		Price:       int(s.Price),
		ImagePath:   s.ImagePath,
	}
}

// UserProduct represents only a product information only for user purpose
type UserProduct struct {
	NameRu      string `json:"name"`
	ApiName     string `json:"api_name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	ImagePath   string `json:"image_path"`
}

// NewProductDTO represents only a product information only for user purpose
type NewProductDTO struct {
	NameRu      string `json:"name"`
	ApiName     string `json:"api_name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	ImagePath   string `json:"image_path"`
	CategoryID  int    `json:"category_id"`
}

// UserProduct represents only a product information only for user purpose
type ProductUpdateDTO struct {
	Id          string  `json:"id"`
	NameRu      *string `json:"name,omitempty"`
	ApiName     *string `json:"api_name,omitempty"`
	Description *string `json:"description,omitempty"`
	Price       *int    `json:"price,omitempty"`
	ImagePath   *string `json:"image_path,omitempty"`
	CategoryID  *int    `json:"category_id,omitempty"`
}
