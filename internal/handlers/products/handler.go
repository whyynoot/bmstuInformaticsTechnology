package products

import (
	"bmstuInformaticsTechnologies/internal/api"
	"bmstuInformaticsTechnologies/internal/product_service"
	"bmstuInformaticsTechnologies/pkg/logging"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"html/template"
	"net/http"
)

const (
	BaseApiPath = "/api"

	ListAllProductRURL   = "/all"
	ListAllProductMethod = "GET"

	ListProductURL    = "/product/{product_name}"
	ListProductMethod = "GET"

	DeleteProductByID   = BaseApiPath + "/delete/{product_id}"
	DeleteProductMethod = "DELETE"
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
	router.HandleFunc(DeleteProductByID, api.Middleware(h.DeleteProduct)).Methods(DeleteProductMethod)
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
