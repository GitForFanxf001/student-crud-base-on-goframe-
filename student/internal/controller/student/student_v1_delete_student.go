package student

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"student/internal/model"
	"student/internal/service"

	"student/api/student/v1"
)

func (c *ControllerV1) DeleteStudent(ctx context.Context, req *v1.DeleteStudentReq) (res *v1.DeleteStudentRes, err error) {
	idStruct := model.StudentDeleteInput{Id: req.Id}
	output, err := service.Student().DeleteStudent(ctx, idStruct)
	resDelete := v1.DeleteStudentRes{Sname: output.Sname}
	g.RequestFromCtx(ctx).Response.WriteJson(g.Map{
		"status": 200,
		"name":   resDelete.Sname,
	})

	return &resDelete, nil
}
