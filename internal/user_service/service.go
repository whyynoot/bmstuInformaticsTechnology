package user_service

import (
	"bmstuInformaticsTechnologies/internal/api"
	"bmstuInformaticsTechnologies/pkg/client/postrgresql"
	"bmstuInformaticsTechnologies/pkg/logging"
	"context"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
	"go.uber.org/zap"
	"net/http"
)

const (
	UserTable = "public.user"
)

type UserService struct {
	DataBaseClient postrgresql.Client
	logger         logging.LoggerInterface
}

func NewUserService(logger logging.LoggerInterface, client postrgresql.Client) *UserService {
	return &UserService{
		DataBaseClient: client,
		logger:         logger,
	}
}

func (u *UserService) GetUserByUsernameAndPassword(username, password string) (*User, error) {
	sb := sqlbuilder.PostgreSQL.NewSelectBuilder()
	query, args := sb.Select(
		"id", "admin").From(UserTable).Where(sb.Equal("username", username), fmt.Sprintf("password = crypt('%v', password)", password)).Build()
	userRow := u.DataBaseClient.QueryRow(context.Background(), query, args...)

	var userId *string
	var adminBool *bool
	if err := userRow.Scan(&userId, &adminBool); err != nil {
		return nil, api.UnauthorizedError("wrong pass or wrong user")
	}

	if adminBool == nil {
		return &User{
			Username: username,
			Password: password,
			Admin:    false,
		}, nil
	} else {
		return &User{
			Username: username,
			Password: password,
			Admin:    *adminBool,
		}, nil
	}
}

func (u *UserService) Create(dto UserDTO) (*User, error) {
	query := fmt.Sprintf("INSERT INTO public.user (username, password, admin) VALUES ('%v', crypt('%v', gen_salt('bf')), FALSE)",
		dto.Username, dto.Password)
	_, err := u.DataBaseClient.Exec(context.Background(), query)
	if err != nil {
		u.logger.Error("cannot inster user", zap.Error(err))
		return nil, api.APIError("duplicate user", http.StatusBadRequest)
	}
	//if rows.Err() != nil {
	//	u.logger.Error("cannot inster user", zap.Error(err))
	//	return nil, api.APIError("unable to create user", http.StatusBadRequest)
	//}
	//fmt.Println(rows, rows.Err(), err)

	return &User{
		Username: dto.Username,
		Password: dto.Password,
		Admin:    false,
	}, nil
}
