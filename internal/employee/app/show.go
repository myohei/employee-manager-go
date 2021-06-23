package app

import (
	"context"
	"fmt"
	"github.com/myohei/employee-manager-go/internal/employee/domain"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type (
	ShowEmployee interface {
		ShowAll(ctx context.Context) error
	}
	showEmployee struct {
		repo domain.EmployeeRepo
	}
)

func NewShowEmployee(repo domain.EmployeeRepo) ShowEmployee {
	return &showEmployee{repo: repo}
}

func (s *showEmployee) ShowAll(ctx context.Context) error {
	employees, err := s.repo.FindAll(ctx)
	if err != nil {
		return err
	}
	if len(employees) == 0 {
		fmt.Println("登録された社員はいません")
		return nil
	}
	fmt.Println("id\tname\tgender\tage\tsalary")
	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	for _, e := range employees {
		s.printEmployee(*e)
	}
	return nil
}

func (s *showEmployee) printEmployee(e domain.Employee) {
	p := message.NewPrinter(language.English)
	fmt.Printf("%s\t%s\t%s\t%d\t%s\n", e.ID(), e.Name(), e.Gender(), e.Birthday().Age(), p.Sprintf("%d", e.Salary()))
}
