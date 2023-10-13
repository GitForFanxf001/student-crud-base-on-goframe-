// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Demo is the golang structure for table demo.
type Demo struct {
	Id        int         `json:"id"         ` // ID
	Fielda    string      `json:"fielda"     ` // Field demo
	Fieldb    string      `json:"fieldb"     ` // Private field demo
	CreatedAt *gtime.Time `json:"created_at" ` // Created Time
	UpdatedAt *gtime.Time `json:"updated_at" ` // Updated Time
}
