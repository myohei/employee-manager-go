package app

import (
	"context"
	"github.com/myohei/employee-manager-go/internal/employee/domain"
	"github.com/myohei/employee-manager-go/internal/util"
)

type (
	DeleteEmployee interface {
		DeleteByID(ctx context.Context) error
	}

	deleteEmployee struct {
		repo domain.EmployeeRepo
	}
)

func NewDeleteEmployee(repo domain.EmployeeRepo) DeleteEmployee {
	return &deleteEmployee{repo: repo}
}

func (d *deleteEmployee) DeleteByID(ctx context.Context) error {
	id := util.GetInput()
	employeeID, err := domain.EmployeeIDString(id)
	if err != nil {
		return err
	}
	if err := d.repo.DeleteByID(ctx, employeeID); err != nil {
		return err
	}
	return nil
}
