package base

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/msg/logmsg"
	"github.com/zeromicro/go-zero/core/errorx"

	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func initDept(ctx *svc.ServiceContext) {
	srvCtx := context.Background()
	c, err := ctx.DB.Department.Query().Count(srvCtx)

	if err != nil {
		fmt.Printf("项目初始化失败:%s\n", err)
		return
	}

	if c == 0 {
		_, err := ctx.DB.Department.Create().
			SetName("山川科技").
			SetArea("北京").
			SetDescription("北京山川科技股份有限公司根部门").
			SetParentDepartmentID(math.MaxInt - 1).
			SetParentDepartments("1").
			Save(context.Background())
		if err != nil {
			fmt.Printf("项目初始化失败:%s\n", err)
		}
	}
}

func initRole(ctx *svc.ServiceContext) {
	srvCtx := context.Background()
	c, err := ctx.DB.Role.Query().Count(srvCtx)

	if err != nil {
		fmt.Printf("项目初始化失败:%s\n", err)
		return
	}

	if c == 0 {
		_, err := ctx.DB.Role.Create().
			SetName("开发者").
			SetWeight(0).Save(srvCtx)
		if err != nil {
			fmt.Printf("项目初始化失败:%s\n", err)
		}
	}
}

func initUser(ctx *svc.ServiceContext) {
	srvCtx := context.Background()
	c, err := ctx.DB.User.Query().Count(srvCtx)

	if err != nil {
		fmt.Printf("项目初始化失败:%s\n", err)
		return
	}

	if c == 0 {
		_, err := ctx.DB.User.Create().
			SetAccount("administrator").
			SetName("开发者管理员").
			SetDepartmentID(1).
			SetRoleID(1).
			SetPassword("Passwordryan@0").
			SetStatus(true).
			Save(srvCtx)
		if err != nil {
			fmt.Printf("项目初始化失败:%s\n", err)
		}
	}
}

func (l *InitDatabaseLogic) InitDatabase() (resp *types.BaseMsgResp, err error) {
	if err := l.svcCtx.DB.Schema.Create(l.ctx, schema.WithForeignKeys(false)); err != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		return nil, errorx.NewInternalError(err.Error())
	}

	initDept(l.svcCtx)
	initRole(l.svcCtx)
	initUser(l.svcCtx)
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.Success)}, nil
}
