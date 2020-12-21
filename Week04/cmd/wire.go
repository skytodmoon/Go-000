//+build wireinject

package cmd

import (
	"Week04/app"
	"Week04/config"
	"Week04/internal/biz"
	"Week04/internal/dao"
	"Week04/internal/router"
	"Week04/internal/service"
	"Week04/library/logger"

	"github.com/google/wire"
)

func BuildApp(path string) (*app.App, error) {
	panic(wire.Build(
		config.WireSet,
		logger.New,
		dao.WireSet,
		service.WireSet,
		biz.WireSet,
		router.WireSet,
		app.WireSet,
	))
}
