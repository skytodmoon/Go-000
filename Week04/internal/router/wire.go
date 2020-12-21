package router

import "github.com/google/wire"

var WireSet = wire.NewSet(
	NewRouter,
)
