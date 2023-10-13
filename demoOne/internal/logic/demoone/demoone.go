package demoone

import (
	"context"
	"demoOne/internal/dao"
	"demoOne/internal/model"
)

type sDemo struct{}

func New() *sDemo {
	return &sDemo{}
}

func (s *sDemo) Create(ctx context.Context, in model.AddInput) (*model.AddOutput, error) {

	id, err := dao.Demo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return nil, err
	}

	return &model.AddOutput{
		ID: uint(id),
	}, nil
}
