package main

import (
	"fmt"
	_ "student/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "student/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"student/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
	fmt.Println("第二次提交")
	fmt.Println("第3次提交")
	fmt.Println("第3次提交")
	fmt.Println("第3次提交")
	fmt.Println("第3次提交")
	fmt.Println("第3次提交")
	fmt.Println("切换到hot-fix分之后的第1次提交")
	fmt.Println("切换到hot-fix分之后的第1次提交")
	fmt.Println("切换到hot-fix分之后的第1次提交")
	fmt.Println("切换到hot-fix分之后的第1次提交")
	fmt.Println("切换到hot-fix分之后的第1次提交")
	fmt.Println("切换到hot-fix分之后的第1次提交")
	fmt.Println("切换到hot-fix分之后的第1次提交")
	fmt.Println("切换到hot-fix分之后的第1次提交")
	fmt.Println("切换到hot-fix分之后的第1次提交")
	fmt.Println("切换到hot-fix分之后的第1次提交")
	fmt.Println("切换到hot-fix分之后的第1次提交")
}
