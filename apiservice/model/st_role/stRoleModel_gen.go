package st_role

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"sectran/apiservice/internal/types"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	stRoleFieldNames          = builder.RawFieldNames(&StRole{})
	stRoleRows                = strings.Join(stRoleFieldNames, ",")
	stRoleRowsExpectAutoSet   = strings.Join(stringx.Remove(stRoleFieldNames, "`role_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	stRoleRowsWithPlaceHolder = strings.Join(stringx.Remove(stRoleFieldNames, "`role_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	stRoleModel interface {
		Insert(ctx context.Context, data *types.RoleVisibleInfo) (sql.Result, error)
		Find(ctx context.Context, roleId *types.RoleQueryInfo) (*types.DataType, error)
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

// []*StRole
func (m *defaultStRoleModel) Find(ctx context.Context, roleQuery *types.RoleQueryInfo) (*types.DataType, error) {
	where := "1=1"

	if len(roleQuery.Name) > 0 {
		where = where + fmt.Sprintf(" AND name like '%%%s%%'", roleQuery.Name)
	}
	var total int64
	totalQuery := fmt.Sprintf("select count(*) as count from %s where %s", m.table, where)
	totalErr := m.conn.QueryRow(&total, totalQuery)
	if totalErr != nil {
		logx.Errorf("error query user by account deu to %s", totalErr)
		return nil, totalErr
	}

	query := fmt.Sprintf("select %s from %s where %s limit ?,?", stRoleRows, m.table, where)
	var resp []*StRole
	err := m.conn.QueryRows(&resp, query, (roleQuery.PageNum-1)*roleQuery.PageSize, roleQuery.PageSize)
	data := &types.DataType{
		List: resp,
		PageData: types.PageType{
			PageNum:  roleQuery.PageNum,
			PageSize: roleQuery.PageSize,
			Total:    total,
		},
	}
	switch err {
	case nil:
		return data, nil
	default:
		logx.Errorf("error query user by account deu to %s", err)
		return nil, err
	}
}

func (m *defaultStRoleModel) Insert(ctx context.Context, data *types.RoleVisibleInfo) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, stRoleRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Name, data.Description, 1)
	return ret, err
}

func (m *defaultStRoleModel) Update(ctx context.Context, data *StRole) error {
	query := fmt.Sprintf("update %s set %s where `role_id` = ?", m.table, stRoleRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Name, data.Description, data.CreateByUid, data.RoleId)
	return err
}

func (m *defaultStRoleModel) tableName() string {
	return m.table
}
