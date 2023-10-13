// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"student/internal/model"
)

type (
	IStudent interface {
		AddStudent(ctx context.Context, in model.StudentAddInput) (*model.StudentAddOutput, error)
		DeleteStudent(ctx context.Context, in model.StudentDeleteInput) (*model.StudentDeleteOutput, error)
		QueryStudent(input model.StudentQueryInput) (*model.StudentQueryOutput, error)
	}
)

var (
	localStudent IStudent
)

func Student() IStudent {
	if localStudent == nil {
		panic("implement not found for interface IStudent, forgot register?")
	}
	return localStudent
}

func RegisterStudent(i IStudent) {
	localStudent = i
}

type name struct {
	
}