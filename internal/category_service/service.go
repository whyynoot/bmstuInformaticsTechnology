package category_service

import (
	"bmstuInformaticsTechnologies/internal/api"
	"bmstuInformaticsTechnologies/pkg/client/postrgresql"
	"bmstuInformaticsTechnologies/pkg/logging"
	"context"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
	"go.uber.org/zap"
)

const (
	CategoriesTable = "public.category"
)

type CategoryService struct {
	DataBaseClient postrgresql.Client
	logger         logging.LoggerInterface
}

func NewCategoryService(logger logging.LoggerInterface, client postrgresql.Client) *CategoryService {
	return &CategoryService{
		DataBaseClient: client,
		logger:         logger,
	}
}

func (c *CategoryService) GetCategories() ([]CategoryStorage, error) {
	query := sqlbuilder.Select(
		"id",
		"name_ru",
		"api_name").From(CategoriesTable).String()
	categoryRows, err := c.DataBaseClient.Query(context.Background(), query)
	if err != nil {
		c.logger.Error("Failed to query rows", zap.Error(err))
		return nil, err
	}

	var Categories []CategoryStorage

	for categoryRows.Next() {
		var category CategoryStorage
		if err := categoryRows.Scan(&category.Id, &category.NameRu, &category.ApiName); err != nil {
			return nil, err
		}
		Categories = append(Categories, category)
	}

	return Categories, nil
}

func (c *CategoryService) DeleteCategoryByID(id string) error {
	sb := sqlbuilder.PostgreSQL.NewDeleteBuilder()
	query, args := sb.DeleteFrom(CategoriesTable).Where(sb.Equal("id", id)).Build()

	_, err := c.DataBaseClient.Query(context.Background(), query, args...)
	if err != nil {
		c.logger.Error("unable to delete category with id", zap.String("id", id),
			zap.Error(err))
		return err
	}

	return nil
}

func (c *CategoryService) NewCategory(category *CategoryCreateDTO) error {
	builder := sqlbuilder.PostgreSQL.NewInsertBuilder()
	query, args := builder.InsertInto(CategoriesTable).Cols(
		"name_ru",
		"api_name").Values(category.NameRu, category.ApiName).Build()

	_, err := c.DataBaseClient.Query(context.Background(), query, args...)
	if err != nil {
		c.logger.Error("unable to create category", zap.String("name", category.NameRu),
			zap.String("api_name", category.ApiName), zap.Error(err))
		return err
	}

	return nil
}

func (c *CategoryService) UpdateCategory(category *CategoryUpdateDTO) error {
	updatesCount := 0
	updatesKeys := make([]string, 0, 8)
	updatesValues := make([]string, 0, 8)
	switch {
	case category.NameRu != nil:
		updatesKeys = append(updatesKeys, "name_ru")
		updatesValues = append(updatesValues, *category.NameRu)
		updatesCount += 1
	case category.ApiName != nil:
		updatesKeys = append(updatesKeys, "api_name")
		updatesValues = append(updatesValues, *category.ApiName)
		updatesCount += 1
	}

	if updatesCount == 0 {
		return api.BadRequestError("zero fields")
	}

	SqlSetString := ""
	for i := 0; i < updatesCount; i++ {
		SqlSetString += fmt.Sprintf(" %v = '%v'", updatesKeys[i], updatesValues[i])
	}

	builder := sqlbuilder.PostgreSQL.NewUpdateBuilder()
	query, args := builder.Update(CategoriesTable).Set(SqlSetString).Where(builder.Equal("id", category.Id)).Build()

	_, err := c.DataBaseClient.Query(context.Background(), query, args...)
	if err != nil {
		c.logger.Error("unable to create category", zap.Stringp("name", category.NameRu),
			zap.Stringp("api_name", category.ApiName), zap.Error(err))
		return err
	}

	return nil
}
