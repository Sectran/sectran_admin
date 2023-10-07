package model

type UserModel struct {
	Id         string `json:"id" gorm:"type:char(36);primary_key"` //用户ID
	UserName   string `json:"userName"`                            //用户名
	Password   string `json:"password"`                            //密码
	RevampTime string `json:"revampTime"`                          // 修改时间
	RevampName string `json:"revampName"`                          //修改人
	IsDelete   int8   `json:"isDelete"`                            //是否删除
}
