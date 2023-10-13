package demo

import (
	"context"
	"demo/internal/dao"
	"demo/internal/model"
	_ "demo/internal/model/do"
	_ "demo/internal/model/entity"
	"demo/internal/service"
)

type sDemo struct{}

func init() {
	service.RegisterDemo(New())
}

func New() *sDemo {
	return &sDemo{}
}

func (s *sDemo) Create(ctx context.Context, in model.DemoCreateInput) (*model.DemoCreateOutput, error) {

	id, err := dao.Demo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return nil, err
	}

	return &model.DemoCreateOutput{
		ID: uint(id),
	}, nil
}
