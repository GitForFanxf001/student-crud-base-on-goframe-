// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package student

import (
	"context"
	
	"student/api/student/v1"
)

type IStudentV1 interface {
	AddStudent(ctx context.Context, req *v1.AddStudentReq) (res *v1.AddStudentRes, err error)
	DeleteStudent(ctx context.Context, req *v1.DeleteStudentReq) (res *v1.DeleteStudentRes, err error)
	QueryStudent(ctx context.Context, req *v1.QueryStudentReq) (res *v1.QueryStudentRes, err error)
}


