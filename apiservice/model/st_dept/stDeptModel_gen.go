package st_dept

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	stDeptFieldNames          = builder.RawFieldNames(&StDept{})
	stDeptRows                = strings.Join(stDeptFieldNames, ",")
	stDeptRowsExpectAutoSet   = strings.Join(stringx.Remove(stDeptFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	stDeptRowsWithPlaceHolder = strings.Join(stringx.Remove(stDeptFieldNames, "`dept_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	stDeptModel interface {
		Insert(ctx context.Context, data *StDept) (sql.Result, error)
		FindOne(ctx context.Context, deptId int64) (*StDept, error)
		Update(ctx context.Context, data *StDept) error
		Delete(ctx context.Context, deptId int64) error
	}

	defaultStDeptModel struct {
		conn  sqlx.SqlConn
		table string
	}

	StDept struct {
		DeptId      int64          `db:"dept_id"`       // 部门ID
		Name        string         `db:"name"`          // 部门名称
		Description sql.NullString `db:"description"`   // 部门描述
		ParentId    int64          `db:"parent_id"`     // 上级部门ID
		ChildIds    string         `db:"child_ids"`     // 下级部门ID集合，用逗号分隔
		CreateByUid sql.NullInt64  `db:"create_by_uid"` // 创建者
		Region      string         `db:"region"`        // 部门所在地区
		IsDeleted   int64          `db:"is_deleted"`    // 是否被删除
		CreateTime  time.Time      `db:"create_time"`   // 创建时间
	}
)

func newStDeptModel(conn sqlx.SqlConn) *defaultStDeptModel {
	return &defaultStDeptModel{
		conn:  conn,
		table: "`st_dept`",
	}
}

func (m *defaultStDeptModel) withSession(session sqlx.Session) *defaultStDeptModel {
	return &defaultStDeptModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`st_dept`",
	}
}

func (m *defaultStDeptModel) Delete(ctx context.Context, deptId int64) error {
	query := fmt.Sprintf("delete from %s where `dept_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, deptId)
	return err
}

func (m *defaultStDeptModel) FindOne(ctx context.Context, deptId int64) (*StDept, error) {
	query := fmt.Sprintf("select %s from %s where `dept_id` = ? limit 1", stDeptRows, m.table)
	var resp StDept
	err := m.conn.QueryRowCtx(ctx, &resp, query, deptId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultStDeptModel) Insert(ctx context.Context, data *StDept) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, stDeptRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.DeptId, data.Name, data.Description, data.ParentId, data.ChildIds, data.CreateByUid, data.Region, data.IsDeleted)
	return ret, err
}

func (m *defaultStDeptModel) Update(ctx context.Context, data *StDept) error {
	query := fmt.Sprintf("update %s set %s where `dept_id` = ?", m.table, stDeptRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Name, data.Description, data.ParentId, data.ChildIds, data.CreateByUid, data.Region, data.IsDeleted, data.DeptId)
	return err
}

func (m *defaultStDeptModel) tableName() string {
	return m.table
}
