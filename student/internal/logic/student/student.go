package student

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"student/internal/model"
	"student/internal/service"
)

type sStudent struct {
}

func init() {
	service.RegisterStudent(New())
}

func New() *sStudent {
	return &sStudent{}
}

func (student *sStudent) AddStudent(ctx context.Context, in model.StudentAddInput) (*model.StudentAddOutput, error) {
	res, err := g.Model("information").Data(g.Map{"name": in.Sname, "score": in.Sscore, "subject": in.Ssubject}).Insert()
	if err != nil {
		fmt.Println("数据插入失败！")
		glog.Error(ctx, err)
	}
	fmt.Println("res = ", res)
	outPut := model.StudentAddOutput{
		Sname: in.Sname,
	}
	return &outPut, nil
}

func (student *sStudent) DeleteStudent(ctx context.Context, in model.StudentDeleteInput) (*model.StudentDeleteOutput, error) {
	sname, err := g.Model("information").Fields("name").Where("id", in.Id).Value()
	if err != nil {
		fmt.Println("字段不存在或者查询语句错误")
	}
	res, err := g.Model("information").Where("id", in.Id).Delete()
	if err != nil {
		glog.Error(ctx, err)
	}
	fmt.Println("delete res = ", res)
	resName := model.StudentDeleteOutput{
		Sname: sname.String(),
	}
	return &resName, nil
}

func (student *sStudent) QueryStudent(input model.StudentQueryInput) (*model.StudentQueryOutput, error) {
	res, err := g.Model("information").Where("id=?", 9).One()
	output := model.StudentQueryOutput{}
	err = res.Struct(&output)
	if err != nil {
		fmt.Println("map转结构体失败！")
	}
	fmt.Println("output = ", output)
	if err != nil {
		fmt.Println("err = ", err)
	}
	return &output, nil
}
