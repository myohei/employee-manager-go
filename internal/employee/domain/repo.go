package domain

import "context"

type EmployeeRepo interface {
	NextIdentity(ctx context.Context) (EmployeeID,error)
	FindAll(ctx context.Context) (Employees, error)
	FindByID(ctx context.Context, id EmployeeID) (*Employee, error)
	Save(ctx context.Context, employee *Employee) error
	DeleteByID(ctx context.Context, id EmployeeID) error
}
