// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Information is the golang structure of table information for DAO operations like Where/Data.
type Information struct {
	g.Meta  `orm:"table:information, do:true"`
	Id      int    //
	Name    string //
	Score   int    //
	Subject string //
}
