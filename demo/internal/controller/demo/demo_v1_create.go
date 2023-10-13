package demo

// internal/controller/demo/demo_v1_create.go

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"

	v1 "demo/api/demo/v1"
	"demo/internal/model"
	"demo/internal/service"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("通了！")
	data := model.DemoCreateInput{
		Fielda: req.Fielda,
		Fieldb: req.Fieldb,
	}

	result, err := service.Demo().Create(ctx, data)
	if err != nil {
		return nil, err
	}

	return &v1.CreateRes{ID: result.ID}, err
}
