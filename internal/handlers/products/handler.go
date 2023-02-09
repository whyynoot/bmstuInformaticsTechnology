package products

import (
	"bmstuInformaticsTechnologies/internal/product_service"
	"bmstuInformaticsTechnologies/pkg/logging"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"html/template"
	"net/http"
	"strings"
)

const (
	ListAllProductRURL   = "/all"
	ListAllProductMethod = "GET"

	ListProductURL    = "/product/{product_name}"
	ListProductMethod = "GET"
)

type Handler struct {
	// + PRODUCT SERVICE
	logger   logging.LoggerInterface
	products map[string]product_service.Product
}

// TODO: Error handling

func NewProductHandler(logger logging.LoggerInterface) *Handler {
	h := Handler{
		logger: logger,
	}

	// FOR FIRST LAB BREAKER
	products := make(map[string]product_service.Product)
	products["milk"] = product_service.Product{
		Name:        "Milk",
		Description: "2%",
		Price:       100,
		Image:       "/static/milk.jpg",
		Stock:       10,
	}
	products["egg"] = product_service.Product{
		Name:        "Egg",
		Description: "10 in a package",
		Price:       110,
		Image:       "/static/egg.jpg",
		Stock:       10,
	}

	h.products = products
	// TODO: add service
	return &h
}

func (h *Handler) Register(router *mux.Router) {
	router.HandleFunc(ListAllProductRURL, h.ListAllProducts).Methods(ListAllProductMethod)
	router.HandleFunc(ListProductURL, h.ListProduct).Methods(ListProductMethod)
}

func (h *Handler) ListAllProducts(w http.ResponseWriter, req *http.Request) {
	ts, err := template.ParseFiles("./templates/all.page.tmpl")
	if err != nil {
		return
	}

	err = ts.Execute(w, h.products)
	if err != nil {
		w.WriteHeader(501)
	}
}

func (h *Handler) ListProduct(w http.ResponseWriter, req *http.Request) {
	RequestedProduct := mux.Vars(req)["product_name"]
	product, ok := h.products[strings.ToLower(RequestedProduct)]
	if !ok {
		ts, err := template.ParseFiles("./templates/notfound.page.tmpl")
		if err != nil {
			h.logger.Error("error", zap.Any("error", err))
			w.WriteHeader(501)
			return
		}
		w.WriteHeader(404)
		err = ts.Execute(w, nil)
		if err != nil {
			h.logger.Error("error", zap.Any("error", err))
			w.WriteHeader(501)
			return
		}
	}

	ts, err := template.ParseFiles("./templates/product.page.tmpl")
	if err != nil {
		h.logger.Error("error", zap.Any("error", err))
	}

	err = ts.Execute(w, product)
	if err != nil {
		w.WriteHeader(501)
		h.logger.Error("error", zap.Any("error", err))
	}
}
