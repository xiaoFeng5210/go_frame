package main

import (
	_ "my_goframe/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"my_goframe/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
