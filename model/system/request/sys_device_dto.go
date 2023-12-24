package request

type DeviceListDTO struct {
	List
	Name string
}

type DeviceAccountListDTO struct {
	List
	DeviceId int
	Name     string
}
