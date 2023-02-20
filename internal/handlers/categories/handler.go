package categories

import (
	"bmstuInformaticsTechnologies/internal/api"
	"bmstuInformaticsTechnologies/internal/category_service"
	"bmstuInformaticsTechnologies/pkg/logging"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

const (
	BaseApiPath = "/api/categories"

	GetCategoriesURL    = BaseApiPath
	GetCategoriesMethod = "GET"

	CreateCategoryURL    = BaseApiPath + "/create"
	CreateCategoryMethod = "POST"

	UpdateCategoryURL    = BaseApiPath + "/update"
	UpdateCategoryMethod = "PATCH"
)

type Handler struct {
	logger          logging.LoggerInterface
	categoryService category_service.CategoryServiceInterface
}

func NewCategoriesHandler(logger logging.LoggerInterface, categoryService category_service.CategoryServiceInterface) *Handler {
	return &Handler{
		logger:          logger,
		categoryService: categoryService,
	}
}

func (h *Handler) Register(router *mux.Router) {
	router.HandleFunc(GetCategoriesURL, api.Middleware(h.GetCategories)).Methods(GetCategoriesMethod)
	router.HandleFunc(UpdateCategoryURL, api.Middleware(h.UpdateCategory)).Methods(UpdateCategoryMethod)
	router.HandleFunc(CreateCategoryURL, api.Middleware(h.CreateCategory)).Methods(CreateCategoryMethod)
}

// GetCategories ... Get a list of all categories
// @Summary Get a list of all categories
// @Description Get a list of all categories
// @Tags Category
// @Success 200 {object} CategoryList
// @Failure 400 {object} api.AppError
// @Failure 500 {object} api.AppError
// @Router /api/categories [get]
func (h *Handler) GetCategories(w http.ResponseWriter, req *http.Request) error {
	categories, err := h.categoryService.GetCategories()
	if err != nil {
		h.logger.Error("unable to get categories", zap.Error(err))
	}

	var resp CategoryList
	resp.Status = "OK"
	for _, category := range categories {
		resp.Categories = append(resp.Categories, category.ToUserCategory())
	}
	err = api.RenderJSONResponse(w, resp, http.StatusOK)
	if err != nil {
		h.logger.Error("unable to write response", zap.Error(err))
		return api.APIError("unable to write response", http.StatusInternalServerError)
	}

	return nil
}

// CreateCategory ... CreateCategory from json
// @Summary Creating New Category
// @Description  CreateNewProduct with json struct
// @Tags Category
// @Param product body category_service.CategoryCreateDTO true "Category"
// @Success 200 {object} CategoryCreateResponse
// @Failure 400 {object} api.AppError
// @Failure 500 {object} api.AppError
// @Router /api/categories/create [post]
func (h *Handler) CreateCategory(w http.ResponseWriter, req *http.Request) error {

	var category category_service.CategoryCreateDTO
	err := json.NewDecoder(req.Body).Decode(&category)
	if err != nil {
		h.logger.Error("unable to unmarshal json", zap.Any("body", req.Body), zap.Error(err))
	}

	err = h.categoryService.NewCategory(&category)
	if err != nil {
		h.logger.Error("unable to create category form post data", zap.Error(err))
		return api.APIError("unable to create category, db error", http.StatusInternalServerError)
	}

	var resp CategoryCreateResponse
	resp.Status = "OK"
	err = api.RenderJSONResponse(w, resp, http.StatusOK)
	if err != nil {
		h.logger.Error("unable to write response", zap.Error(err))
		return api.APIError("unable to write response", http.StatusInternalServerError)
	}

	return nil
}

// UpdateCategory ... Update fields of a category
// @Summary Updating category information
// @Description Update any field / fields of category
// @Description Only ID is required
// @Tags Category
// @Param product body category_service.CategoryUpdateDTO false "Category"
// @Success 200 {object} CategoryUpdateResponse
// @Failure 400 {object} api.AppError
// @Failure 500 {object} api.AppError
// @Router /api/categories/update [patch]
func (h *Handler) UpdateCategory(w http.ResponseWriter, req *http.Request) error {

	var category category_service.CategoryUpdateDTO
	err := json.NewDecoder(req.Body).Decode(&category)
	if err != nil {
		h.logger.Error("unable to unmarshal json", zap.Any("body", req.Body), zap.Error(err))
	}

	err = h.categoryService.UpdateCategory(&category)
	if err != nil {
		e, ok := err.(*api.AppError)
		if ok {
			h.logger.Error("unable to update category form post data", zap.Error(e))
			return err
		}
		h.logger.Error("unable to update category form post data", zap.Error(err))
		return api.APIError("unable to update category, db error", http.StatusInternalServerError)
	}

	var resp CategoryUpdateResponse
	resp.Status = "OK"
	err = api.RenderJSONResponse(w, resp, http.StatusOK)
	if err != nil {
		h.logger.Error("unable to write response", zap.Error(err))
		return api.APIError("unable to write response", http.StatusInternalServerError)
	}

	return nil
}
