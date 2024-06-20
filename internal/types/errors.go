package types

import (
	"github.com/zeromicro/go-zero/core/errorx"
)

var (
	//	//	20000 custom error
	ErrAccountHasNoRights error = &errorx.CodeError{Code: 20001, Msg: "该账号权限不足，继续操作将被认定为攻击行为"}
	ErrInternalError            = &errorx.CodeError{Code: 20002, Msg: "系统内部错误,请联系开发者"}
	ErrDataNotFound             = &errorx.CodeError{Code: 20003, Msg: "操作数据不存在,可能已经被删除或者被移动"}
	ErrInvalidToken             = &errorx.CodeError{Code: 20004, Msg: "用户还未登录或者回话已过期，请重新登录"}
	ErrForceLoginOut            = &errorx.CodeError{Code: 20005, Msg: "登录的用户主体已变更，强制下线"}
	ErrRedis                    = &errorx.CodeError{Code: 20006, Msg: "Redis 服务器异常，请联系开发者"}
)

func CustomError(msg string) error {
	return &errorx.CodeError{Code: 20000, Msg: msg}
}
