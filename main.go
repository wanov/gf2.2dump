package main

import (
	"gftest/internal/dao"
	_ "gftest/internal/packed"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	type PaginateIn struct {
		Model *gdb.Model
	}
	in := &PaginateIn{
		Model: dao.SysUser.Ctx(gctx.New()).OrderDesc("id"),
	}
	g.Dump(in)
}
