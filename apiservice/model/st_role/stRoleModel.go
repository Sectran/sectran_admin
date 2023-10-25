package st_role

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ StRoleModel = (*customStRoleModel)(nil)

type (
	// StRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStRoleModel.
	StRoleModel interface {
		stRoleModel
	}

	customStRoleModel struct {
		*defaultStRoleModel
	}
)

// NewStRoleModel returns a model for the database table.
func NewStRoleModel(conn sqlx.SqlConn) StRoleModel {
	return &customStRoleModel{
		defaultStRoleModel: newStRoleModel(conn),
	}
}
