package main

import (
	"context"
	"github.com/myohei/employee-manager-go/internal/registry"
	"log"
)

func main() {
	ctx := context.Background()
	cli := registry.InitCLI()
	err := cli.Start(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
