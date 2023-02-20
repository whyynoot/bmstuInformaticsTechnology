package products

import "bmstuInformaticsTechnologies/internal/product_service"

// ProductList api response with status and product list.
type ProductList struct {
	Products []product_service.UserProduct `json:"products"`
	Status   string                        `json:"status"`
	Count    int                           `json:"count"`
}

// Product api response with status and single product.
type Product struct {
	Product product_service.UserProduct `json:"product"`
	Status  string                      `json:"status"`
}

type ProductDeleteResponse struct {
	Status string `json:"status"`
}

type ProductCreateResponse struct {
	Status string `json:"status"`
}

type ProductUpdateResponse struct {
	Status string `json:"status"`
}
