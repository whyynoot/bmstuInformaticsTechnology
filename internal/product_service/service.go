package product_service

import (
	"bmstuInformaticsTechnologies/pkg/client/postrgresql"
	"bmstuInformaticsTechnologies/pkg/logging"
	"context"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
	"go.uber.org/zap"
)

const (
	ProductTable = "public.product"
)

type ProductService struct {
	DataBaseClient postrgresql.Client
	logger         logging.LoggerInterface
}

func NewProductService(logger logging.LoggerInterface, client postrgresql.Client) *ProductService {
	return &ProductService{
		DataBaseClient: client,
		logger:         logger,
	}
}

func (p *ProductService) GetProducts() ([]ProductStorage, error) {
	query := sqlbuilder.Select(
		"id",
		"name_ru",
		"api_name",
		"description",
		"image_path",
		"price",
		"category_id").From(ProductTable).String()
	productRows, err := p.DataBaseClient.Query(context.Background(), query)
	if err != nil {
		p.logger.Error("Failed to query rows", zap.Error(err))
		return nil, err
	}

	var Products []ProductStorage

	for productRows.Next() {
		var product ProductStorage
		if err := productRows.Scan(&product.ID, &product.NameRu, &product.ApiName, &product.Description,
			&product.ImagePath, &product.Price, &product.CategoryID); err != nil {
			return nil, err
		}
		Products = append(Products, product)
	}

	return Products, nil
}

func (p *ProductService) GetProduct(apiName string) (*ProductStorage, error) {
	sb := sqlbuilder.NewSelectBuilder()
	query := sb.Select(
		"id",
		"name_ru",
		"api_name",
		"description",
		"image_path",
		"price",
		"category_id").From(ProductTable).Where(fmt.Sprintf("api_name = '%s'", apiName)).String()
	productRows := p.DataBaseClient.QueryRow(context.Background(), query)

	var product ProductStorage
	if err := productRows.Scan(&product.ID, &product.NameRu, &product.ApiName, &product.Description,
		&product.ImagePath, &product.Price, &product.CategoryID); err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductService) DeleteProductByID(id string) error {
	sb := sqlbuilder.NewDeleteBuilder()
	query := sb.DeleteFrom(ProductTable).Where(fmt.Sprintf("id = '%v'", id)).String()

	_, err := p.DataBaseClient.Query(context.Background(), query)
	if err != nil {
		p.logger.Error("unable to delete product with id", zap.String("id", id),
			zap.Error(err))
		return err
	}

	return nil
}
