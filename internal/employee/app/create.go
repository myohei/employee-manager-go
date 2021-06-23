package app

import (
	"context"
	"fmt"
	"github.com/myohei/employee-manager-go/internal/common"
	"github.com/myohei/employee-manager-go/internal/employee/app/dto"
	"github.com/myohei/employee-manager-go/internal/employee/domain"
	"github.com/myohei/employee-manager-go/internal/util"
)

type (
	CreateEmployee interface {
		Create(ctx context.Context) (*dto.EmployeeData, error)
	}
	createEmployee struct {
		repo domain.EmployeeRepo
	}
)

func NewCreateEmployee(repo domain.EmployeeRepo) CreateEmployee {
	return &createEmployee{repo: repo}
}

func (c *createEmployee) Create(ctx context.Context) (*dto.EmployeeData, error) {
	op := "createEmployee.Create"
	name := c.getName()
	gender := c.getGender()
	birthday := c.getBirthday()
	salary := c.getSalary()
	id, err := c.repo.NextIdentity(ctx)
	if err != nil {
		return nil, &common.Error{
			Op:  op,
			Err: err,
		}
	}
	employee := domain.NewEmployee(id, name, gender, birthday, salary)
	if err := c.repo.Save(ctx, employee); err != nil {
		return nil, &common.Error{
			Op:  op,
			Err: err,
		}
	}
	return &dto.EmployeeData{
		ID:       employee.ID().String(),
		Name:     employee.Name().String(),
		Birthday: employee.Birthday().Time(),
		Salary:   employee.Salary().Uint(),
	}, nil
}

func (c *createEmployee) getName() domain.Name {
	for {
		fmt.Println("名前を入力>")
		in := util.GetInput()
		name, err := domain.NameString(in)
		if err != nil {
			fmt.Println(common.ErrorMessage(err))
			continue
		}
		return name
	}
}

func (c *createEmployee) getGender() domain.Gender {
	for {
		fmt.Println("性別を入力> (男性 or 女性)")
		in := util.GetInput()
		gender, err := domain.GenderString(in)
		if err != nil {
			fmt.Println("男性 or 女性で入力してください")
			continue
		}
		return gender
	}
}

func (c *createEmployee) getBirthday() domain.Birthday {
	for {
		fmt.Println("誕生日を入力> (例: 1987-09-10)")
		in := util.GetInput()
		birthday, err := domain.BirthdayString(in)
		if err != nil {
			fmt.Println(common.ErrorMessage(err))
			continue
		}
		return birthday
	}
}

func (c *createEmployee) getSalary() domain.Salary {
	for {
		fmt.Println("給与を入力>")
		in := util.GetInput()
		salary, err := domain.SalaryString(in)
		if err != nil {
			fmt.Println(common.ErrorMessage(err))
			continue
		}
		return salary
	}
}
