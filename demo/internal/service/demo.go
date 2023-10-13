// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"demo/internal/model"
)

type (
	IDemo interface {
		Create(ctx context.Context, in model.DemoCreateInput) (*model.DemoCreateOutput, error)
	}
)

var (
	localDemo IDemo
)

func Demo() IDemo {
	if localDemo == nil {
		panic("implement not found for interface IDemo, forgot register?")
	}
	return localDemo
}

func RegisterDemo(i IDemo) {
	localDemo = i
}
