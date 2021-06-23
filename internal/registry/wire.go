//+build wireinject

package registry

import (
	"github.com/google/wire"
	"github.com/myohei/employee-manager-go/internal/employee/app"
	"github.com/myohei/employee-manager-go/internal/employee/infra/inmemory"
	"github.com/myohei/employee-manager-go/internal/interfaces"
)

func InitCLI() (c *interfaces.CLI) {
	wire.Build(
		inmemory.SuperSet,
		app.SuperSet,
		interfaces.NewCLI,
	)
	return
}
