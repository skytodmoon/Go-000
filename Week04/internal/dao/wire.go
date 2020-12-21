package dao

import "github.com/google/wire"

var WireSet = wire.NewSet(
	NewDao,
	NewDatabaseMYSQL,
)
