package system

type Device struct {
	Id          int
	Name        string
	Address     string
	OsKind      int
	Encoding    string
	DeptId      int
	Description string
}

func (Device) TableName() string {
	return "sys_device"
}

type DeviceAccount struct {
	Id              int
	DeviceId        int
	Username        string
	Password        string
	IsAdministrator int
	Protocol        string
	Port            int
	PrivateKetyPas  string
	PrivateKey      string
}

func (DeviceAccount) TableName() string {
	return "sys_device_account"
}
