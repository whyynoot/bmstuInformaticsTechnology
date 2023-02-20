package category_service

// CategoryServiceInterface Interface to interact with categories
type CategoryServiceInterface interface {
	GetCategories() ([]CategoryStorage, error)
	DeleteCategoryByID(id string) error
	NewCategory(category *CategoryCreateDTO) error
	UpdateCategory(category *CategoryUpdateDTO) error
}
