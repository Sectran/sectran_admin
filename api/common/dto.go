package common

type DeleteDto struct {
	Id string `json:"id" gorm:"type:char(36);primary_key"` //部门ID
}
