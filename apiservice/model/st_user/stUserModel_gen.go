package st_user

import (
	"context"
	"database/sql"
	"fmt"
	"sectran/apiservice/internal/types"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	stUserFieldNames          = builder.RawFieldNames(&StUser{})
	stUserRows                = strings.Join(stUserFieldNames, ",")
	stUserRowsExpectAutoSet   = strings.Join(stringx.Remove(stUserFieldNames, "`user_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	stUserRowsWithPlaceHolder = strings.Join(stringx.Remove(stUserFieldNames, "`user_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	stUserModel interface {
		Insert(ctx context.Context, data *types.UserAllInfo) (sql.Result, error)
		Find(ctx context.Context, user *types.UserQueryInfo) ([]*StUser, error)
		Update(ctx context.Context, data *types.UserAllInfo) error
		Delete(ctx context.Context, userId int64) error
	}

	defaultStUserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	StUser struct {
		UserId      int64          `db:"user_id"`       // 用户ID
		Account     string         `db:"account"`       // 用户账号
		Password    string         `db:"password"`      // 用户密码
		Username    sql.NullString `db:"username"`      // 用户姓名
		DeptId      int64          `db:"dept_id"`       // 用户所属部门ID
		Disable     int64          `db:"disable"`       // 账号是否禁用
		Description sql.NullString `db:"description"`   // 账号描述
		CreateTime  time.Time      `db:"create_time"`   // 创建时间
		CreateByUid int64          `db:"create_by_uid"` // 创建人
		RoleId      int64          `db:"role_id"`       // 用户角色ID
		Telephone   string         `db:"telephone"`     // 用户手机号码
		Email       string         `db:"email"`         // 用户邮箱
	}
)

func newStUserModel(conn sqlx.SqlConn) *defaultStUserModel {
	return &defaultStUserModel{
		conn:  conn,
		table: "`st_user`",
	}
}

func (m *defaultStUserModel) withSession(session sqlx.Session) *defaultStUserModel {
	return &defaultStUserModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`st_user`",
	}
}

func (m *defaultStUserModel) Delete(ctx context.Context, userId int64) error {
	if _, err := m.FindById(ctx, userId); err != nil {
		return err
	}

	query := fmt.Sprintf("delete from %s where `user_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, userId)
	return err
}

func (m *defaultStUserModel) Find(ctx context.Context, user *types.UserQueryInfo) ([]*StUser, error) {
	query := fmt.Sprintf("select %s from %s", stUserRows, m.table)
	var args []interface{}

	if user != nil {
		query = fmt.Sprintf(" %s where 1=1 ", query)

		if len(user.Username) > 0 {
			query = fmt.Sprintf(" %s and `username` = ? ", query)
			args = append(args, user.Username)
		}
		if user.UserId > 0 {
			query = fmt.Sprintf(" %s and`user_id` = ? ", query)
			args = append(args, user.UserId)
		}
		if len(user.Account) > 0 {
			query = fmt.Sprintf(" %s and `account` = ? ", query)
			args = append(args, user.Account)
		}
		if len(user.Email) > 0 {
			query = fmt.Sprintf(" %s and `email` = ? ", query)
			args = append(args, user.Email)
		}
		if len(user.Telephone) > 0 {
			query = fmt.Sprintf(" %s and `telephone` = ? ", query)
			args = append(args, user.Telephone)
		}
		if user.RoleId > 0 {
			query = fmt.Sprintf(" %s and `role_id` = ? ", query)
			args = append(args, user.RoleId)
		}
		if user.DeptId > 0 {
			query = fmt.Sprintf(" %s and `dept_id` = ?  ", query)
			args = append(args, user.DeptId)
		}

		if user.Disable > 0 {
			query = fmt.Sprintf(" %s and `disable` = ? ", query)
			args = append(args, user.Disable)
		} else {
			query = fmt.Sprintf(" %s and `disable` = ? ", query)
			args = append(args, 0)
		}
		// desc
		// creat_time
	}

	var resp []*StUser
	err := m.conn.QueryRowsCtx(ctx, &resp, query, args...)
	switch err {
	case nil:
		return resp, nil
	default:
		logx.Errorf("error query user by account deu to %s", err)
		return nil, err
	}
}

func (m *defaultStUserModel) Insert(ctx context.Context, data *types.UserAllInfo) (sql.Result, error) {
	user, err := m.FindByAccount(ctx, data.Account)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return nil, fmt.Errorf("duplicate user account")
	}

	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, stUserRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Account, data.Password, data.Username, data.DeptId, data.Disable, data.Description, data.CreateByUid, data.IsDeleted, data.RoleId, data.Telephone, data.Email)
	return ret, err
}

// update by user id
func (m *defaultStUserModel) Update(ctx context.Context, data *types.UserAllInfo) error {
	if _, err := m.FindById(ctx, data.UserId); err != nil {
		return err
	}

	query := fmt.Sprintf("update %s set %s where `user_id` = ?", m.table, stUserRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Account, data.Password, data.Username, data.DeptId, data.Disable, data.Description, data.CreateByUid, data.IsDeleted, data.RoleId, data.Telephone, data.Email, data.UserId)
	return err
}

func (m *defaultStUserModel) FindByAccount(ctx context.Context, account string) (*StUser, error) {
	uq := &types.UserQueryInfo{
		UserVisibleQueryInfo: types.UserVisibleQueryInfo{Account: account},
	}

	u, err := m.Find(ctx, uq)
	if err != nil {
		return nil, err
	}

	if len(u) == 0 {
		return nil, nil
	}
	return u[0], nil
}

func (m *defaultStUserModel) FindById(ctx context.Context, userId int64) (*StUser, error) {
	uq := &types.UserQueryInfo{
		UserVisibleQueryInfo: types.UserVisibleQueryInfo{UserId: userId},
	}

	u, err := m.Find(ctx, uq)
	if err != nil {
		return nil, err
	}

	if len(u) == 0 {
		return nil, nil
	}

	return u[0], nil
}

func (m *defaultStUserModel) tableName() string {
	return m.table
}
