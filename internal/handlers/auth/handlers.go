package auth

import (
	"bmstuInformaticsTechnologies/internal/api"
	"bmstuInformaticsTechnologies/internal/user_service"
	"bmstuInformaticsTechnologies/pkg/jwt"
	"bmstuInformaticsTechnologies/pkg/logging"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

const (
	BaseApiPath = "/api/auth"

	SignUpURL    = BaseApiPath + "/signup"
	SignUpMethod = "POST"

	LogInURL    = BaseApiPath + "/login"
	LogInMethod = "POST"
)

type Handler struct {
	logger      logging.LoggerInterface
	userService user_service.UserServiceInterface
	jwtHelper   jwt.HelperInterface
}

func NewAuthHandler(logger logging.LoggerInterface, userService user_service.UserServiceInterface) *Handler {
	return &Handler{
		logger:      logger,
		userService: userService,
		jwtHelper:   jwt.NewHelper(logger),
	}
}

func (h *Handler) Register(router *mux.Router) {
	router.HandleFunc(SignUpURL, api.Middleware(h.CreateUser)).Methods(SignUpMethod)
	router.HandleFunc(LogInURL, api.Middleware(h.LogInUser)).Methods(LogInMethod)
}

// CreateUser ... CreateUser with username and password
// @Summary Signing up new user
// @Description  Sign up a new user with username and password
// @Tags User
// @Param product body user_service.UserDTO true "User"
// @Success 200 {object} AuthorizationResponse
// @Failure 400 {object} api.AppError
// @Failure 500 {object} api.AppError
// @Router /api/auth/signup [post]
func (h *Handler) CreateUser(w http.ResponseWriter, req *http.Request) error {

	var dto user_service.UserDTO
	err := json.NewDecoder(req.Body).Decode(&dto)
	if err != nil {
		h.logger.Error("unable to unmarshal json", zap.Any("body", req.Body), zap.Error(err))
		return api.BadRequestError("bad json")
	}

	user, err := h.userService.Create(dto)
	if err != nil {
		e, ok := err.(*api.AppError)
		if ok {
			h.logger.Error("unable to create user", zap.Error(e))
			return err
		}
		h.logger.Error("unable to create user form post data", zap.Error(err))
		return api.APIError("unable to create user, db error", http.StatusInternalServerError)
	}

	auth, err := h.jwtHelper.GenerateToken(user)
	if err != nil {
		h.logger.Error("unable to generate token", zap.Error(err))
		return api.APIError("unable to generate token", http.StatusInternalServerError)
	}

	resp := AuthorizationResponse{auth}
	err = api.RenderJSONResponse(w, resp, http.StatusOK)
	if err != nil {
		h.logger.Error("unable to write response", zap.Error(err))
		return api.APIError("unable to write response", http.StatusInternalServerError)
	}

	return nil
}

// LogInUser ... Log in with username and password
// @Summary Log in
// @Description  Log in with username and password
// @Tags User
// @Param product body UserLoginRequest true "User"
// @Success 200 {object} AuthorizationResponse
// @Failure 400 {object} api.AppError
// @Failure 500 {object} api.AppError
// @Router /api/auth/login [post]
func (h *Handler) LogInUser(w http.ResponseWriter, req *http.Request) error {

	var creds UserLoginRequest
	err := json.NewDecoder(req.Body).Decode(&creds)
	if err != nil {
		h.logger.Error("unable to unmarshal json", zap.Any("body", req.Body), zap.Error(err))
		return api.BadRequestError("bad json")
	}

	user, err := h.userService.GetUserByUsernameAndPassword(creds.Username, creds.Password)
	if err != nil {
		e, ok := err.(*api.AppError)
		if ok {
			h.logger.Error("unable to auth user", zap.Error(e))
			return err
		}
		h.logger.Error("unable to auth use", zap.Error(err))
		return api.APIError("unable to auth user", http.StatusInternalServerError)
	}

	auth, err := h.jwtHelper.GenerateToken(user)
	if err != nil {
		h.logger.Error("unable to generate token", zap.Error(err))
		return api.APIError("unable to generate token", http.StatusInternalServerError)
	}

	resp := AuthorizationResponse{auth}
	err = api.RenderJSONResponse(w, resp, http.StatusOK)
	if err != nil {
		h.logger.Error("unable to write response", zap.Error(err))
		return api.APIError("unable to write response", http.StatusInternalServerError)
	}

	return nil
}
