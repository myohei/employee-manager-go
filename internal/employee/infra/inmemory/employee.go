package inmemory

import (
	"context"
	"github.com/google/uuid"
	"github.com/myohei/employee-manager-go/internal/common"
	"github.com/myohei/employee-manager-go/internal/employee/domain"
)

type employeeRepo struct {
	employees domain.Employees
}

func NewEmployeeRepo() domain.EmployeeRepo {
	return &employeeRepo{employees: make(domain.Employees, 0)}
}

func (e *employeeRepo) FindByID(ctx context.Context, id domain.EmployeeID) (*domain.Employee, error) {
	panic("implement me")
}

func (e *employeeRepo) Save(ctx context.Context, employee *domain.Employee) error {
	e.employees = append(e.employees, employee)
	return nil
}

func (e *employeeRepo) NextIdentity(ctx context.Context) (domain.EmployeeID, error) {
	return domain.EmployeeID(uuid.New()), nil
}

func (e *employeeRepo) FindAll(ctx context.Context) (domain.Employees, error) {
	return e.employees, nil
}

func (e *employeeRepo) DeleteByID(ctx context.Context, id domain.EmployeeID) error {
	var idx = -1
	for i, e := range e.employees {
		if e.ID() == id {
			idx = i
			break
		}
	}
	if idx < 0 {
		return &common.Error{
			Code:    common.ErrCodeNotFound,
			Message: "employee was not found",
			Op:      "employeeRepo.DeleteByID",
		}
	}
	s := e.employees
	e.employees = append(s[:idx], s[idx+1:]...)
	return nil
}
