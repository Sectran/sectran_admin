package st_dept

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ StDeptModel = (*customStDeptModel)(nil)

type (
	// StDeptModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStDeptModel.
	StDeptModel interface {
		stDeptModel
	}

	customStDeptModel struct {
		*defaultStDeptModel
	}
)

// NewStDeptModel returns a model for the database table.
func NewStDeptModel(conn sqlx.SqlConn) StDeptModel {
	return &customStDeptModel{
		defaultStDeptModel: newStDeptModel(conn),
	}
}
