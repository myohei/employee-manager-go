package inmemory

import "github.com/google/wire"

var SuperSet = wire.NewSet(NewEmployeeRepo)
