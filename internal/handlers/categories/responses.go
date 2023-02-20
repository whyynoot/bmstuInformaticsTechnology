package categories

import (
	"bmstuInformaticsTechnologies/internal/category_service"
)

// CategoryList api response with status and product list.
type CategoryList struct {
	Categories []category_service.UserCategory `json:"categories"`
	Status     string                          `json:"status"`
}

// Category api response with status and single product.
type Category struct {
	Category category_service.UserCategory `json:"category"`
	Status   string                        `json:"status"`
}

type CategoryCreateResponse struct {
	Status string `json:"status"`
}

type CategoryUpdateResponse struct {
	Status string `json:"status"`
}
