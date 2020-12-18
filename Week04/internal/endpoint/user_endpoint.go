package endpoint

import (
	"context"

	"Week04/internal/service"

	"github.com/go-kit/kit/endpoint"
)

type UserEndpoints struct {
	QueryUserEndpoint endpoint.Endpoint
}

type QueryUserRequest struct {
	Email string
}

type QueryUserResponse struct {
	UserInfo *service.UserInfoDTO `json:"user_info"`
}

func MakeQueryUserEndpoint(userService service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*QueryUserRequest)
		userInfo, err := userService.QueryUser(ctx, req.Email)
		return &QueryUserResponse{UserInfo: userInfo}, err

	}
}
