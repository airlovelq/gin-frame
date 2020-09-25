package controller

import (
	"scoremanager/database"
	"scoremanager/middleware/cache"
)

// BaseOp ...定义了各种中间件，数据库操作使用情况
type BaseOp struct {
	CacheOp cache.CacheOp
	DbOp    *database.DatabaseOp
	useDB   bool
}

func NewBaseOp() *BaseOp {
	return &BaseOp{
		DbOp:    database.NewDatabaseOp(),
		CacheOp: cache.NewCacheOp(),
		useDB:   true,
	}
}

type BeginOptions func(*BaseOp)

func NotUseDBOption(c *BaseOp) {
	c.useDB = false
}

func (c *BaseOp) BeginOp(opts ...BeginOptions) {
	for _, opt := range opts {
		opt(c)
	}
	if c.useDB {
		err := c.DbOp.Begin()
		if err != nil {
			panic(err)
		}
	}
}

func (c *BaseOp) EndOp() {
	// fmt.Printf("db close!\n")
	err := recover()
	if err != nil {
		if c.useDB {
			c.DbOp.Rollback()
		}
		panic(err)
	} else {
		if c.useDB {
			c.DbOp.Commit()
		}
	}
}
