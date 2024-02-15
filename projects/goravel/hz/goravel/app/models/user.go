package models

import (
	"github.com/goravel/framework/database/orm"
)

type User1 struct {
	orm.Model
	Name   string
	Avatar string
	orm.SoftDeletes
}
