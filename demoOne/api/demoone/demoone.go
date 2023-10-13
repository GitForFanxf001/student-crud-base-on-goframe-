// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package demoone

import (
	"context"
	
	"demoOne/api/demoone/v1"
)

type IDemooneV1 interface {
	Add(ctx context.Context, req *v1.AddReq) (res *v1.AddRes, err error)
}


