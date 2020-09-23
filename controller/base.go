package controller

import (
	"scoremanager/database"
	"scoremanager/middleware/cache"
)

type BaseOp struct {
	CacheOp cache.CacheOp
	DbOp    *database.DatabaseOp
}

func (c *BaseOp) EndOp() {
	// fmt.Printf("db close!\n")
	err := recover()
	if err != nil {
		if c.DbOp != nil {
			c.DbOp.Rollback()
		}
		panic(err)
	} else {
		if c.DbOp != nil {
			c.DbOp.Commit()
		}
	}
}
