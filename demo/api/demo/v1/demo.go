package v1

import "github.com/gogf/gf/v2/frame/g"

// api/demo/v1/demo.go

type CreateReq struct {
	g.Meta `path:"/create" method:"post" tags:"DemoService" summary:"Create a demo record"`
	Fielda string `p:"fileda" v:"required"`
	Fieldb string `p:"filedb" v:"required"`
}

type CreateRes struct {
	ID uint `json:"id"`
}
