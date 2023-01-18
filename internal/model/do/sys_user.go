// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysUser is the golang structure of table sys_user for DAO operations like Where/Data.
type SysUser struct {
	g.Meta    `orm:"table:sys_user, do:true"`
	Id        interface{} //
	Gid       interface{} // 组织id
	GroupRoot interface{} // 是否组织总帐号 0否 1是
	Account   interface{} // 帐号
	Password  interface{} // 密码
	Nickname  interface{} // 昵称
	Avatar    interface{} // 头像
	Status    interface{} // 状态 0禁用 1正常
	Token     interface{} // 登陆标志
}