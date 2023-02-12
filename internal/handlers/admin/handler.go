package admin

import (
	"bmstuInformaticsTechnologies/internal/api"
	"bmstuInformaticsTechnologies/internal/product_service"
	"bmstuInformaticsTechnologies/pkg/logging"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"html/template"
	"net/http"
)

// TODO: Add support of login and auth

const (
	AdminPageURL    = "/admin"
	AdminPageMethod = "GET"
)

type Handler struct {
	logger         logging.LoggerInterface
	productService product_service.ProductServiceInterface
}

func NewAdminHandler(logger logging.LoggerInterface, productService product_service.ProductServiceInterface) *Handler {
	return &Handler{
		logger:         logger,
		productService: productService,
	}
}

func (h *Handler) Register(router *mux.Router) {
	router.HandleFunc(AdminPageURL, api.Middleware(h.AdminPage)).Methods(AdminPageMethod)
}

func (h *Handler) AdminPage(w http.ResponseWriter, req *http.Request) error {
	products, err := h.productService.GetProducts()
	if err != nil {
		h.logger.Error("unable to get products", zap.Error(err))
	}
	ts, err := template.ParseFiles("./templates/admin.page.tmpl")
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
