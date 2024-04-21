package types

import (
	"github.com/zeromicro/go-zero/core/errorx"
)

var (
	ErrAccountHasNoRights error
	ErrInternalError      error
	ErrDataNotFound       error
)

func init() {
	ErrAccountHasNoRights = errorx.NewApiForbiddenError("该账号权限不足")
	ErrInternalError = errorx.NewCodeInternalError("系统内部错误")
	ErrDataNotFound = errorx.NewCodeNotFoundError("数据不存在，可能已被删除")
}
