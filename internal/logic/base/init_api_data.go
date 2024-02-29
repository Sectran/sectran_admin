package base

import (
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"
)

func (l *InitDatabaseLogic) insertApiData() (err error) {
	// User

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("SectranAdmin"),
		Path:        pointy.GetPointer("/user/create"),
		Description: pointy.GetPointer("apiDesc.createUser"),
		ApiGroup:    pointy.GetPointer("user"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("SectranAdmin"),
		Path:        pointy.GetPointer("/user/update"),
		Description: pointy.GetPointer("apiDesc.updateUser"),
		ApiGroup:    pointy.GetPointer("user"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("SectranAdmin"),
		Path:        pointy.GetPointer("/user/delete"),
		Description: pointy.GetPointer("apiDesc.deleteUser"),
		ApiGroup:    pointy.GetPointer("user"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("SectranAdmin"),
		Path:        pointy.GetPointer("/user/list"),
		Description: pointy.GetPointer("apiDesc.getUserList"),
		ApiGroup:    pointy.GetPointer("user"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("SectranAdmin"),
		Path:        pointy.GetPointer("/user"),
		Description: pointy.GetPointer("apiDesc.getUserById"),
		ApiGroup:    pointy.GetPointer("user"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	// Role

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("SectranAdmin"),
		Path:        pointy.GetPointer("/role/create"),
		Description: pointy.GetPointer("apiDesc.createRole"),
		ApiGroup:    pointy.GetPointer("role"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("SectranAdmin"),
		Path:        pointy.GetPointer("/role/update"),
		Description: pointy.GetPointer("apiDesc.updateRole"),
		ApiGroup:    pointy.GetPointer("role"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("SectranAdmin"),
		Path:        pointy.GetPointer("/role/delete"),
		Description: pointy.GetPointer("apiDesc.deleteRole"),
		ApiGroup:    pointy.GetPointer("role"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("SectranAdmin"),
		Path:        pointy.GetPointer("/role/list"),
		Description: pointy.GetPointer("apiDesc.getRoleList"),
		ApiGroup:    pointy.GetPointer("role"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("SectranAdmin"),
		Path:        pointy.GetPointer("/role"),
		Description: pointy.GetPointer("apiDesc.getRoleById"),
		ApiGroup:    pointy.GetPointer("role"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	// Department

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("SectranAdmin"),
		Path:        pointy.GetPointer("/department/create"),
		Description: pointy.GetPointer("apiDesc.createDepartment"),
		ApiGroup:    pointy.GetPointer("department"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("SectranAdmin"),
		Path:        pointy.GetPointer("/department/update"),
		Description: pointy.GetPointer("apiDesc.updateDepartment"),
		ApiGroup:    pointy.GetPointer("department"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("SectranAdmin"),
		Path:        pointy.GetPointer("/department/delete"),
		Description: pointy.GetPointer("apiDesc.deleteDepartment"),
		ApiGroup:    pointy.GetPointer("department"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("SectranAdmin"),
		Path:        pointy.GetPointer("/department/list"),
		Description: pointy.GetPointer("apiDesc.getDepartmentList"),
		ApiGroup:    pointy.GetPointer("department"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("SectranAdmin"),
		Path:        pointy.GetPointer("/department"),
		Description: pointy.GetPointer("apiDesc.getDepartmentById"),
		ApiGroup:    pointy.GetPointer("department"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	return err
}
