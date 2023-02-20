package products

import (
	"bmstuInformaticsTechnologies/internal/api"
	"bmstuInformaticsTechnologies/internal/product_service"
	"bmstuInformaticsTechnologies/pkg/logging"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"html/template"
	"net/http"
)

const (
	BaseApiPath = "/api/product"

	ListAllProductRURL   = "/all"
	ListAllProductMethod = "GET"

	ListProductURL    = "/product/{product_name}"
	ListProductMethod = "GET"

	DeleteProductURL    = BaseApiPath + "/delete/{product_id}"
	DeleteProductMethod = "DELETE"

	CreateProductURL    = BaseApiPath + "/create"
	CreateProductMethod = "POST"

	UpdateProductURL    = BaseApiPath + "/update"
	UpdateProductMethod = "PATCH"

	GetProductsURL    = "/api/products"
	GetProductsMethod = "GET"

	GetProductsFromCategoryURL    = "/api/{category_api_name}/products"
	GetProductsFromCategoryMethod = "GET"
)

type Handler struct {
	logger         logging.LoggerInterface
	productService product_service.ProductServiceInterface
}

func NewProductHandler(logger logging.LoggerInterface, productService product_service.ProductServiceInterface) *Handler {
	return &Handler{
		logger:         logger,
		productService: productService,
	}
}

func (h *Handler) Register(router *mux.Router) {
	router.HandleFunc(ListAllProductRURL, api.Middleware(h.ListAllProducts)).Methods(ListAllProductMethod)
	router.HandleFunc(ListProductURL, api.Middleware(h.ListProduct)).Methods(ListProductMethod)
	router.HandleFunc(DeleteProductURL, api.Middleware(h.DeleteProduct)).Methods(DeleteProductMethod)
	router.HandleFunc(CreateProductURL, api.Middleware(h.CreateProduct)).Methods(CreateProductMethod)
	router.HandleFunc(UpdateProductURL, api.Middleware(h.UpdateProduct)).Methods(UpdateProductMethod)
	router.HandleFunc(GetProductsURL, api.Middleware(h.GetProducts)).Methods(GetProductsMethod)
	router.HandleFunc(GetProductsFromCategoryURL, api.Middleware(h.GetProductsByCategory)).Methods(GetProductsFromCategoryMethod)
}

func (h *Handler) ListAllProducts(w http.ResponseWriter, req *http.Request) error {
	products, err := h.productService.GetProducts()
	if err != nil {
		h.logger.Error("unable to get products", zap.Error(err))
	}

	ts, err := template.ParseFiles("./templates/all.page.tmpl")
	if err != nil {
		h.logger.Error("template parse error", zap.Any("error", err))
		return api.NewAppError("tmpl errors", http.StatusInternalServerError)
	}

	err = ts.Execute(w, products)
	if err != nil {
		h.logger.Error("template execute error", zap.Any("error", err))
		return api.NewAppError("tmpl error", http.StatusInternalServerError)
	}

	return nil
}

func (h *Handler) ListProduct(w http.ResponseWriter, req *http.Request) error {
	RequestedProduct := mux.Vars(req)["product_name"]
	product, err := h.productService.GetProduct(RequestedProduct)
	if err != nil {
		h.logger.Error("unable to get product form database", zap.Error(err))
		return api.ErrNotFound
	}

	ts, err := template.ParseFiles("./templates/product.page.tmpl")
	if err != nil {
		h.logger.Error("template parse error", zap.Any("error", err))
		return api.NewAppError("tmpl errors", http.StatusInternalServerError)
	}

	err = ts.Execute(w, *product)
	if err != nil {
		h.logger.Error("template execute error", zap.Any("error", err))
		return api.NewAppError("tmpl errors", http.StatusInternalServerError)
	}

	return nil
}

// DeleteProduct ... Delete product by id
// @Summary Deletion of product
// @Description  Delete product by UUID
// @Tags Product
// @Param id path string true "UUID of product"
// @Success 200 {object} ProductDeleteResponse
// @Failure 500 {object} api.AppError
// @Router /api/product/delete [delete]
func (h *Handler) DeleteProduct(w http.ResponseWriter, req *http.Request) error {
	RequestedProduct := mux.Vars(req)["product_id"]
	err := h.productService.DeleteProductByID(RequestedProduct)
	if err != nil {
		h.logger.Error("unable to delete product form database", zap.Error(err))
		return api.APIError("unable to delete", http.StatusInternalServerError)
	}

	var resp ProductDeleteResponse
	resp.Status = "OK"
	err = api.RenderJSONResponse(w, resp, http.StatusOK)
	if err != nil {
		h.logger.Error("unable to write response", zap.Error(err))
		return api.APIError("unable to write response", http.StatusInternalServerError)
	}

	return nil
}

// CreateProduct ... CreateProduct from json
// @Summary Creating New Product
// @Description  CreateNewProduct with json struct
// @Tags Product
// @Param product body product_service.NewProductDTO true "Product"
// @Success 200 {object} ProductCreateResponse
// @Failure 400 {object} api.AppError
// @Failure 500 {object} api.AppError
// @Router /api/product/create [post]
func (h *Handler) CreateProduct(w http.ResponseWriter, req *http.Request) error {
	var product product_service.NewProductDTO
	err := json.NewDecoder(req.Body).Decode(&product)
	if err != nil {
		h.logger.Error("unable to unmarshal json", zap.Any("body", req.Body), zap.Error(err))
		return api.BadRequestError("bad post data")
	}

	err = h.productService.NewProduct(&product)
	if err != nil {
		h.logger.Error("unable to create product form post data", zap.Error(err))
		return api.APIError("unable to create product, db error", http.StatusInternalServerError)
	}

	var resp ProductCreateResponse
	resp.Status = "OK"
	err = api.RenderJSONResponse(w, resp, http.StatusOK)
	if err != nil {
		h.logger.Error("unable to write response", zap.Error(err))
		return api.APIError("unable to write response", http.StatusInternalServerError)
	}

	return nil
}

// UpdateProduct ... Update fields of a product
// @Summary Updating product information
// @Description Update any field / fields of product
// @Description Only ID is required
// @Tags Product
// @Param product body product_service.ProductUpdateDTO false "Product"
// @Success 200 {object} ProductCreateResponse
// @Failure 400 {object} api.AppError
// @Failure 500 {object} api.AppError
// @Router /api/product/update [patch]
func (h *Handler) UpdateProduct(w http.ResponseWriter, req *http.Request) error {
	var product product_service.ProductUpdateDTO
	err := json.NewDecoder(req.Body).Decode(&product)
	if err != nil {
		h.logger.Error("unable to unmarshal json", zap.Any("body", req.Body), zap.Error(err))
		return api.BadRequestError("bad post data")
	}

	err = h.productService.UpdateProduct(&product)
	if err != nil {
		e, ok := err.(*api.AppError)
		if ok {
			h.logger.Error("unable to update product form post data", zap.Error(e))
			return err
		}
		h.logger.Error("unable to update product form post data", zap.Error(err))
		return api.APIError("unable to update product, db error", http.StatusInternalServerError)
	}

	var resp ProductUpdateResponse
	resp.Status = "OK"
	err = api.RenderJSONResponse(w, resp, http.StatusOK)
	if err != nil {
		h.logger.Error("unable to write response", zap.Error(err))
		return api.APIError("unable to write response", http.StatusInternalServerError)
	}

	return nil
}

// GetProducts ... Get a list of all products, user-based
// @Summary Getting a list of user products
// @Description Getting a list of user products
// @Tags Product
// @Success 200 {object} ProductList
// @Failure 400 {object} api.AppError
// @Failure 500 {object} api.AppError
// @Router /api/products [get]
func (h *Handler) GetProducts(w http.ResponseWriter, req *http.Request) error {
	products, err := h.productService.GetProducts()
	if err != nil {
		h.logger.Error("unable to get products", zap.Error(err))
	}

	var resp ProductList
	resp.Status = "OK"
	resp.Count = len(products)
	for _, product := range products {
		resp.Products = append(resp.Products, product.ToUserProduct())
	}
	err = api.RenderJSONResponse(w, resp, http.StatusOK)
	if err != nil {
		h.logger.Error("unable to write response", zap.Error(err))
		return api.APIError("unable to write response", http.StatusInternalServerError)
	}

	return nil
}

// GetProductsByCategory ... Get a list of all products based on category
// @Summary Getting a list of user products based on category
// @Description Getting a list of user products based on category
// @Tags Product
// @Param category_api_name path string true "Category"
// @Success 200 {object} ProductList
// @Failure 400 {object} api.AppError
// @Failure 500 {object} api.AppError
// @Router /api/{category_api_name}/products [get]
func (h *Handler) GetProductsByCategory(w http.ResponseWriter, req *http.Request) error {
	products, err := h.productService.GetProductsByCategory(mux.Vars(req)["category_api_name"])
	if err != nil {
		h.logger.Error("unable to get products", zap.Error(err))
	}

	var resp ProductList
	resp.Status = "OK"
	for _, product := range products {
		resp.Products = append(resp.Products, product.ToUserProduct())
	}
	err = api.RenderJSONResponse(w, resp, http.StatusOK)
	if err != nil {
		h.logger.Error("unable to write response", zap.Error(err))
		return api.APIError("unable to write response", http.StatusInternalServerError)
	}

	return nil
}
