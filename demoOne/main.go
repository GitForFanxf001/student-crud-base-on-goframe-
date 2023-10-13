package main

import (
	_ "demoOne/internal/packed"

	_ "demoOne/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"demoOne/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
