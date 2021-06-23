package app

import "github.com/google/wire"

var SuperSet = wire.NewSet(
	NewCreateEmployee,
	NewShowEmployee,
	NewDeleteEmployee,
)
