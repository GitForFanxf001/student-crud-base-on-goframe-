package v1

import "github.com/gogf/gf/v2/frame/g"

// api/demo/v1/demo.go

// 添加学生请求的请求参数，学号姓名性别
type AddStudentReq struct {
	g.Meta   `path:"/add" method:"post" tags:"AddService" summary:"add a student info to the system"`
	Sname    string `p:"sname" v:"required"`
	Sscore   int    `p:"sscore" v:"required"`
	Ssubject string `p:"ssubject" v:"required"`
}

// 添加学生的响应，返回添加学生的学号和姓名
type AddStudentRes struct {
	Sname string `json:"sname"`
}

// 添加学生请求的请求参数，学号姓名性别
type DeleteStudentReq struct {
	g.Meta `path:"/delete" method:"delete" tags:"DeleteService" summary:"delete a student info to the system"`
	Id     int `p:"id" v:"required"`
}

// 添加学生的响应，返回添加学生的学号和姓名
type DeleteStudentRes struct {
	Sname string `json:"sname"`
}

type QueryStudentReq struct {
	g.Meta `path:"/query" method:"get" tags:"QueryService" summary:"query a student info to the system"`
	Id     int `p:"id" v:"required"`
}

// 添加学生的响应，返回添加学生的学号和姓名
type QueryStudentRes struct {
	Sname    string `json:"sname"`
	Sscore   int    `json:"sscore"`
	Ssubject string `json:"ssubject"`
}
