// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountsColumns holds the columns for the "accounts" table.
	AccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Comment: "Create Time | 创建日期"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "Update Time | 修改日期"},
		{Name: "username", Type: field.TypeString, Size: 16, Comment: "account username|账号名称"},
		{Name: "port", Type: field.TypeUint32, Comment: "account port|端口"},
		{Name: "protocol", Type: field.TypeUint8, Comment: "protocol of the this account.|账号协议"},
		{Name: "password", Type: field.TypeString, Size: 128, Comment: "account password|账号密码"},
		{Name: "private_key", Type: field.TypeString, Size: 4096, Comment: "private_key of the this account.|账号私钥"},
		{Name: "device_id", Type: field.TypeUint64, Nullable: true, Comment: "account belong to|账号所属设备"},
	}
	// AccountsTable holds the schema information for the "accounts" table.
	AccountsTable = &schema.Table{
		Name:       "accounts",
		Columns:    AccountsColumns,
		PrimaryKey: []*schema.Column{AccountsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "accounts_devices_devices",
				Columns:    []*schema.Column{AccountsColumns[8]},
				RefColumns: []*schema.Column{DevicesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DepartmentsColumns holds the columns for the "departments" table.
	DepartmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Comment: "Create Time | 创建日期"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "Update Time | 修改日期"},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 32, Comment: "The name of the department.|部门名称"},
		{Name: "area", Type: field.TypeString, Size: 128, Comment: "The area where the department is located.|部门所在地区"},
		{Name: "description", Type: field.TypeString, Size: 128, Comment: "Description of the department.|部门描述"},
		{Name: "parent_departments", Type: field.TypeString, Comment: "Comma-separated list of parent department IDs in ascending order.|上级部门集合逗号分隔升序排列"},
	}
	// DepartmentsTable holds the schema information for the "departments" table.
	DepartmentsTable = &schema.Table{
		Name:       "departments",
		Columns:    DepartmentsColumns,
		PrimaryKey: []*schema.Column{DepartmentsColumns[0]},
	}
	// DevicesColumns holds the columns for the "devices" table.
	DevicesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Comment: "Create Time | 创建日期"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "Update Time | 修改日期"},
		{Name: "name", Type: field.TypeString, Size: 128, Comment: "The name of the device.|设备名称"},
		{Name: "department_id", Type: field.TypeUint64, Nullable: true, Comment: "ID of the device's department.|设备所属部门"},
		{Name: "host", Type: field.TypeString, Unique: true, Comment: "login host|设备地址"},
		{Name: "description", Type: field.TypeString, Comment: "Description of the device.|设备描述"},
	}
	// DevicesTable holds the schema information for the "devices" table.
	DevicesTable = &schema.Table{
		Name:       "devices",
		Columns:    DevicesColumns,
		PrimaryKey: []*schema.Column{DevicesColumns[0]},
	}
	// PolicyAuthsColumns holds the columns for the "policy_auths" table.
	PolicyAuthsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Comment: "Create Time | 创建日期"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "Update Time | 修改日期"},
		{Name: "name", Type: field.TypeString, Size: 64, Comment: "policy name|策略名称"},
		{Name: "power", Type: field.TypeInt32, Comment: "policy power|策略优先级"},
		{Name: "department_id", Type: field.TypeUint64, Nullable: true, Comment: "ID of the policy's department.|策略所属部门"},
		{Name: "users", Type: field.TypeString, Comment: "策略关联用户"},
		{Name: "accounts", Type: field.TypeString, Comment: "策略关联账号"},
		{Name: "direction", Type: field.TypeBool, Comment: "策略相关性方向,默认正向，即断言正向用户与账号", Default: true},
	}
	// PolicyAuthsTable holds the schema information for the "policy_auths" table.
	PolicyAuthsTable = &schema.Table{
		Name:       "policy_auths",
		Columns:    PolicyAuthsColumns,
		PrimaryKey: []*schema.Column{PolicyAuthsColumns[0]},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Comment: "Create Time | 创建日期"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "Update Time | 修改日期"},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 32, Comment: "The name of the role.|角色名称"},
		{Name: "weight", Type: field.TypeInt, Comment: "The weight of the role. Smaller values indicate higher priority.|角色优先级，值越小优先级越高"},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:       "roles",
		Columns:    RolesColumns,
		PrimaryKey: []*schema.Column{RolesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Comment: "Create Time | 创建日期"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "Update Time | 修改日期"},
		{Name: "account", Type: field.TypeString, Unique: true, Size: 64, Comment: "User account.|用户账号"},
		{Name: "name", Type: field.TypeString, Size: 64, Comment: "User name.|用户姓名"},
		{Name: "password", Type: field.TypeString, Size: 128, Comment: "User password.|用户密码"},
		{Name: "status", Type: field.TypeBool, Comment: "User status (enabled(true) or disabled(false)).|用户账号状态", Default: true},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 128, Comment: "User description.|用户账号描述"},
		{Name: "email", Type: field.TypeString, Nullable: true, Size: 64, Comment: "User email.|用户邮箱"},
		{Name: "phone_number", Type: field.TypeString, Nullable: true, Size: 32, Comment: "User phone number.|用户手机号码"},
		{Name: "department_id", Type: field.TypeUint64, Nullable: true, Comment: "ID of the user's department.|用户所属部门"},
		{Name: "role_id", Type: field.TypeUint64, Nullable: true, Comment: "ID of the user's role.|用户所属角色"},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_departments_departments",
				Columns:    []*schema.Column{UsersColumns[10]},
				RefColumns: []*schema.Column{DepartmentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "users_roles_roles",
				Columns:    []*schema.Column{UsersColumns[11]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountsTable,
		DepartmentsTable,
		DevicesTable,
		PolicyAuthsTable,
		RolesTable,
		UsersTable,
	}
)

func init() {
	AccountsTable.ForeignKeys[0].RefTable = DevicesTable
	UsersTable.ForeignKeys[0].RefTable = DepartmentsTable
	UsersTable.ForeignKeys[1].RefTable = RolesTable
}
