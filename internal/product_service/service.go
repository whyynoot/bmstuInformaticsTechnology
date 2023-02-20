package product_service

import (
	"bmstuInformaticsTechnologies/internal/api"
	"bmstuInformaticsTechnologies/pkg/client/postrgresql"
	"bmstuInformaticsTechnologies/pkg/logging"
	"context"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
	"go.uber.org/zap"
	"strconv"
)

const (
	ProductTable    = "public.product"
	CategoriesTable = "public.category"
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

func (p *ProductService) GetProductsByCategory(apiName string) ([]ProductStorage, error) {
	builder := sqlbuilder.PostgreSQL.NewSelectBuilder()
	query, args := builder.Select(
		"p.id",
		"p.name_ru",
		"p.api_name",
		"p.description",
		"p.image_path",
		"p.price",
		"p.category_id").From(ProductTable+" p").Join(CategoriesTable+" cat", "cat.id = p.category_id").Where(builder.Equal("cat.api_name", apiName)).Build()
	productRows, err := p.DataBaseClient.Query(context.Background(), query, args...)
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
	sb := sqlbuilder.PostgreSQL.NewDeleteBuilder()
	query, args := sb.DeleteFrom(ProductTable).Where(sb.Equal("id", id)).Build()

	_, err := p.DataBaseClient.Query(context.Background(), query, args...)
	if err != nil {
		p.logger.Error("unable to delete product with id", zap.String("id", id),
			zap.Error(err))
		return err
	}

	return nil
}

func (p *ProductService) NewProduct(product *NewProductDTO) error {
	builder := sqlbuilder.PostgreSQL.NewInsertBuilder()
	query, args := builder.InsertInto(ProductTable).Cols(
		"name_ru",
		"api_name",
		"description",
		"image_path",
		"price",
		"category_id").Values(product.NameRu, product.ApiName, product.Description,
		product.ImagePath, product.Price, product.CategoryID).Build()

	_, err := p.DataBaseClient.Query(context.Background(), query, args...)
	if err != nil {
		p.logger.Error("unable to create product", zap.String("name", product.NameRu),
			zap.String("api_name", product.ApiName), zap.String("desc", product.Description),
			zap.String("image_path", product.ImagePath), zap.Int("price", product.Price),
			zap.Int("cat_id", product.CategoryID), zap.Error(err))
		return err
	}

	return nil
}

func (p *ProductService) UpdateProduct(product *ProductUpdateDTO) error {
	updatesCount := 0
	updatesKeys := make([]string, 0, 8)
	updatesValues := make([]string, 0, 8)
	switch {
	case product.NameRu != nil:
		updatesKeys = append(updatesKeys, "name_ru")
		updatesValues = append(updatesValues, *product.NameRu)
		updatesCount += 1
	case product.CategoryID != nil:
		updatesKeys = append(updatesKeys, "category_id")
		updatesValues = append(updatesValues, strconv.Itoa(*product.CategoryID))
		updatesCount += 1
	case product.Price != nil:
		updatesKeys = append(updatesKeys, "price")
		updatesValues = append(updatesValues, strconv.Itoa(*product.Price))
		updatesCount += 1
	case product.ImagePath != nil:
		updatesKeys = append(updatesKeys, "image_path")
		updatesValues = append(updatesValues, *product.ImagePath)
		updatesCount += 1
	case product.ApiName != nil:
		updatesKeys = append(updatesKeys, "api_name")
		updatesValues = append(updatesValues, *product.ApiName)
		updatesCount += 1
	case product.Description != nil:
		updatesKeys = append(updatesKeys, "description")
		updatesValues = append(updatesValues, *product.Description)
		updatesCount += 1
	}
	// fmt.Println(updatesValues, updatesKeys, updatesCount)

	if updatesCount == 0 {
		return api.BadRequestError("zero fields")
	}

	SqlSetString := ""
	for i := 0; i < updatesCount; i++ {
		SqlSetString += fmt.Sprintf(" %v = '%v'", updatesKeys[i], updatesValues[i])
	}

	builder := sqlbuilder.PostgreSQL.NewUpdateBuilder()
	query, args := builder.Update(ProductTable).Set(SqlSetString).Where(builder.Equal("id", product.Id)).Build()

	_, err := p.DataBaseClient.Query(context.Background(), query, args...)
	if err != nil {
		p.logger.Error("unable to create product", zap.Stringp("name", product.NameRu),
			zap.Stringp("api_name", product.ApiName), zap.Stringp("desc", product.Description),
			zap.Stringp("image_path", product.ImagePath), zap.Intp("price", product.Price),
			zap.Intp("cat_id", product.CategoryID), zap.Error(err))
		return err
	}

	return nil
}
