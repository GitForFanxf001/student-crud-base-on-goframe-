package student

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"student/internal/model"
	"student/internal/service"

	"student/api/student/v1"
)

func (c *ControllerV1) QueryStudent(ctx context.Context, req *v1.QueryStudentReq) (res *v1.QueryStudentRes, err error) {
	idStruct := model.StudentQueryInput{Id: req.Id}
	output, err := service.Student().QueryStudent(idStruct)
	resQuery := v1.QueryStudentRes{Sname: output.Sname, Sscore: output.Sscore, Ssubject: output.Ssubject}
	g.RequestFromCtx(ctx).Response.WriteJson(g.Map{
		"status":  200,
		"name":    resQuery.Sname,
		"score":   resQuery.Sscore,
		"subject": resQuery.Ssubject,
	})
	return &resQuery, nil
}
