package model

type DepartmentModel struct {
	Id                     string `json:"id" gorm:"type:char(36);primary_key"` //部门ID
	Name                   string `json:"name"`                                //部门名称
	Describe               string `json:"describe"`                            //部门描述
	AddTime                string `json:"add_time"`                            //创建时间
	RevampTime             string `json:"revampTime"`                          // 最后一次修改时间
	AddUser                string `json:"add_user"`                            //创建人
	IsDelete               int8   `json:"isDelete"`                            //是否删除
	SuperiorId             string `json:"superior_id"`                         //上级Id
	SubordinateId          string `json:"subordinate_id"`                      //下级Id
	CorrelationUserInt     uint16 `json:"correlation_user_int"`                //部门关联用户数量
	CorrelationResourceInt uint16 `json:"correlation_resource_int"`            //部门关联资源数量
	Location               string `json:"location"`                            //部门所在地区
}
