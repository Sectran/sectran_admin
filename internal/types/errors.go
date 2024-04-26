package types

import (
	"github.com/zeromicro/go-zero/core/errorx"
)

var (
	//	//	20000 custom error
	ErrAccountHasNoRights error = &errorx.CodeError{Code: 20001, Msg: "改账号没有权限不足"}
	ErrInternalError            = &errorx.CodeError{Code: 20002, Msg: "系统内部错误,请联系开发者"}
	ErrDataNotFound             = &errorx.CodeError{Code: 20003, Msg: "操作数据不存在,可能已经被删除或者被移动"}
	ErrInvalidToken             = &errorx.CodeError{Code: 20004, Msg: "用户还未登录或者回话已过期，请重新登录"}
)

func CustomError(msg string) error {
	return &errorx.CodeError{Code: 20000, Msg: msg}
}
