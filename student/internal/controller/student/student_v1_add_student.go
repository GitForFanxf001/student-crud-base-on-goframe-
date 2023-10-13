package student

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"student/internal/model"
	"student/internal/service"

	"student/api/student/v1"
)

func (c *ControllerV1) AddStudent(ctx context.Context, req *v1.AddStudentReq) (res *v1.AddStudentRes, err error) {
	data := model.StudentAddInput{
		Sname:    req.Sname,
		Sscore:   req.Sscore,
		Ssubject: req.Ssubject,
	}
	output, err := service.Student().AddStudent(ctx, data)
	fmt.Println(output)
	g.RequestFromCtx(ctx).Response.WriteJson(g.Map{
		"status": 200,
		"name":   output.Sname,
	})
	resAdd := v1.AddStudentRes{Sname: output.Sname}
	return &resAdd, nil
}
