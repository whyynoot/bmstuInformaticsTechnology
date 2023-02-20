package category_service

type CategoryStorage struct {
	Id      int
	NameRu  string
	ApiName string
}

func (c *CategoryStorage) ToUserCategory() UserCategory {
	return UserCategory{
		NameRu:  c.NameRu,
		ApiName: c.ApiName,
	}
}

type UserCategory struct {
	NameRu  string `json:"name_ru"`
	ApiName string `json:"api_name"`
}

type CategoryUpdateDTO struct {
	Id      int     `json:"id"`
	NameRu  *string `json:"name_ru"`
	ApiName *string `json:"api_name"`
}

type CategoryCreateDTO struct {
	NameRu  string `json:"name_ru"`
	ApiName string `json:"api_name"`
}

type CategoryDeleteDTO struct {
	Id int `json:"id"`
}
