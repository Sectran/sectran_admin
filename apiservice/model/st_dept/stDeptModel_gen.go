package st_dept

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
	stDeptFieldNames          = builder.RawFieldNames(&StDept{})
	stDeptRows                = strings.Join(stDeptFieldNames, ",")
	stDeptRowsExpectAutoSet   = strings.Join(stringx.Remove(stDeptFieldNames, "`dept_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	stDeptRowsWithPlaceHolder = strings.Join(stringx.Remove(stDeptFieldNames, "`dept_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	stDeptModel interface {
		Insert(ctx context.Context, data *types.DeptAddRequest) (sql.Result, error)
		Find(ctx context.Context, data *types.DeptQueryInfo) (*types.PageListVisibleInfo, error)
		Update(ctx context.Context, data *types.DeptEditInfo) error
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

func (m *defaultStDeptModel) Find(ctx context.Context, deptQuery *types.DeptQueryInfo) (*types.PageListVisibleInfo, error) {
	where := "1=1"
	if len(deptQuery.Name) > 0 {
		where = where + fmt.Sprintf(" AND name = '%s'", deptQuery.Name)
	}
	if deptQuery.DeptId > 0 {
		where = where + fmt.Sprintf(" AND dept_id = '%d'", deptQuery.DeptId)
	}
	if deptQuery.ParentId > 0 {
		where = where + fmt.Sprintf(" AND parent_id = '%d'", deptQuery.ParentId)
	}

	if len(deptQuery.Region) > 0 {
		where = where + fmt.Sprintf(" AND region = '%s'", deptQuery.Region)
	}

	var total int64
	totalQuery := fmt.Sprintf("select count(*) as count from %s where %s", m.table, where)
	totalErr := m.conn.QueryRow(&total, totalQuery)
	if totalErr != nil {
		logx.Errorf("error query dept total fail  %s", totalErr)
		return nil, totalErr
	}

	query := fmt.Sprintf("select %s from %s where %s limit ?,?", stDeptRows, m.table, where)
	var resp []*StDept
	err := m.conn.QueryRows(&resp, query, (deptQuery.PageNum-1)*deptQuery.PageSize, deptQuery.PageSize)
	data := &types.PageListVisibleInfo{
		List: resp,
		PageInfo: types.PageVisibleInfo{
			PageNum:  deptQuery.PageNum,
			PageSize: deptQuery.PageSize,
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

func (m *defaultStDeptModel) Insert(ctx context.Context, data *types.DeptAddRequest) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, stDeptRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Name, data.Description, data.ParentId, data.ChildIds, data.CreateByUid, data.Region)
	return ret, err
}

func (m *defaultStDeptModel) Update(ctx context.Context, data *types.DeptEditInfo) error {
	query := fmt.Sprintf("update %s set %s where `dept_id` = ?", m.table, stDeptRowsWithPlaceHolder)
	//data.ChildIds, data.CreateByUid,
	_, err := m.conn.ExecCtx(ctx, query, data.Name, data.Description, data.ParentId, data.Region, data.DeptId)
	return err
}

func (m *defaultStDeptModel) tableName() string {
	return m.table
}
