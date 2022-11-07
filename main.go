package main

import (
	_ "github.com/go-mogu/mogu-picture/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"

	"github.com/go-mogu/mogu-picture/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
