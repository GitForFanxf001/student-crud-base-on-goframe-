package main

import (
	_ "student/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "student/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"student/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
