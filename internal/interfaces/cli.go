package interfaces

import (
	"context"
	"errors"
	"fmt"
	"github.com/myohei/employee-manager-go/internal/employee/app"
	"github.com/myohei/employee-manager-go/internal/util"
	"strconv"
)

type CLI struct {
	create app.CreateEmployee
	show   app.ShowEmployee
	delete app.DeleteEmployee
}

func NewCLI(create app.CreateEmployee, show app.ShowEmployee, delete app.DeleteEmployee) *CLI {
	return &CLI{create: create, show: show, delete: delete}
}

func (c *CLI) Start(ctx context.Context) error {
	for {
		c.showMenu()
		menu, err := c.selectMenu()
		if err != nil {
			fmt.Println("不正な入力です")
			continue
		}
		if err := c.do(ctx, menu); err != nil {
			fmt.Printf("err:%v\n", err)
		}
	}
}

func (c *CLI) showMenu() {
	fmt.Println("========================")
	fmt.Println("1. 登録")
	fmt.Println("2. 一覧")
	fmt.Println("3. 更新")
	fmt.Println("4. 削除")
	fmt.Println("5. 終了")
	fmt.Println("========================")
}

func (c *CLI) selectMenu() (int, error) {
	fmt.Println("選択> (1-5)")
	in := util.GetInput()
	s, err := strconv.Atoi(in)
	if err != nil {
		return 0, err
	}
	if s >= 1 && s <= 5 {
		return s, nil
	}
	return 0, errors.New("aaa")
}

func (c *CLI) do(ctx context.Context, menu int) error {
	var err error
	switch menu {
	case 1:
		_, err = c.create.Create(ctx)
	case 2:
		err = c.show.ShowAll(ctx)
	case 3:
		fmt.Println("unsupport")
		return nil
	case 4:
		err = c.delete.DeleteByID(ctx)
	}
	return err
}
