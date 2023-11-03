package st_role

import (
	"context"
	"database/sql"
	"fmt"
	"sectran/apiservice/internal/types"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	stRoleFieldNames          = builder.RawFieldNames(&StRole{})
	stRoleRows                = strings.Join(stRoleFieldNames, ",")
	stRoleRowsExpectAutoSet   = strings.Join(stringx.Remove(stRoleFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	stRoleRowsWithPlaceHolder = strings.Join(stringx.Remove(stRoleFieldNames, "`role_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	stRoleModel interface {
		Insert(ctx context.Context, data *types.RoleAllInfo) (sql.Result, error)
		Find(ctx context.Context, roleId *types.RoleQueryInfo) (*StRole, error)
		Update(ctx context.Context, data *StRole) error
		Delete(ctx context.Context, roleId int64) error
	}

	defaultStRoleModel struct {
		conn  sqlx.SqlConn
		table string
	}

	StRole struct {
		RoleId      int64          `db:"role_id"`       // 角色ID
		Name        string         `db:"name"`          // 角色名称
		Description sql.NullString `db:"description"`   // 角色描述
		CreateByUid int64          `db:"create_by_uid"` // 创建者
		IsDeleted   int64          `db:"is_deleted"`    // 是否被删除
		CreateTime  time.Time      `db:"create_time"`   // 创建时间
	}
)

func newStRoleModel(conn sqlx.SqlConn) *defaultStRoleModel {
	return &defaultStRoleModel{
		conn:  conn,
		table: "`st_role`",
	}
}

func (m *defaultStRoleModel) withSession(session sqlx.Session) *defaultStRoleModel {
	return &defaultStRoleModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`st_role`",
	}
}

func (m *defaultStRoleModel) Delete(ctx context.Context, roleId int64) error {
	query := fmt.Sprintf("delete from %s where `role_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, roleId)
	return err
}

func (m *defaultStRoleModel) Find(ctx context.Context, roleQuery *types.RoleQueryInfo) (*StRole, error) {
	query := fmt.Sprintf("select %s from %s where `role_id` = ? limit 1", stRoleRows, m.table)

	var args []interface{}

	if roleQuery != nil {
		query = fmt.Sprintf(" %s where 1=1 ", query)
		if len(roleQuery.Name) > 0 {
			query = fmt.Sprintf(" %s and `username` = ? ", query)
			args = append(args, roleQuery.Name)
		}
		if roleQuery.RoleId > 0 {
			query = fmt.Sprintf(" %s and`user_id` = ? ", query)
			args = append(args, roleQuery.RoleId)
		}

	}

	var resp StRole
	err := m.conn.QueryRowCtx(ctx, &resp, query, args...)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultStRoleModel) Insert(ctx context.Context, data *types.RoleAllInfo) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, stRoleRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.RoleId, data.Name, data.Description, data.CreateByUid, data.IsDelete)
	return ret, err
}

func (m *defaultStRoleModel) Update(ctx context.Context, data *StRole) error {
	query := fmt.Sprintf("update %s set %s where `role_id` = ?", m.table, stRoleRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Name, data.Description, data.CreateByUid, data.IsDeleted, data.RoleId)
	return err
}

func (m *defaultStRoleModel) tableName() string {
	return m.table
}
