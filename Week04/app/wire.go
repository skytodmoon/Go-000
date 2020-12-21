package app

import (
	"Week04/app/web"

	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	NewGinEngine,
	NewApp,
	web.NewServer,
)
