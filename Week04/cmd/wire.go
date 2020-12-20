// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"Week04/internal/dao"
	"Week04/internal/service"

	"github.com/google/wire"
)

func InitUserService(userDAO *dao.UserDAO) service.UserService {
	wire.Build(service.MakeUserServiceImpl)
	return service.UserServiceImpl{}
}

// ctx := context.Background()

// userService := service.MakeUserServiceImpl(&dao.UserDAOImpl{})

// userEndpoints := &endpoint.UserEndpoints{
// 	QueryUserEndpoint: endpoint.MakeQueryUserEndpoint(userService),
// }

// r := transport.MakeHTTPHandler(ctx, userEndpoints)
