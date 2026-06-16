package main

import (
	"context"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"cl_system/internal/cmd"
)

func main() {
	cmd.Main.Run(context.Background())
}
