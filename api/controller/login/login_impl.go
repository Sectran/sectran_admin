package loginController

import (
	"golang.org/x/crypto/bcrypt"
	"sectran/api/common"
	"sectran/api/model"
)

func LoginImpl(p loginParameter) (error, string) {
	Db := common.Db
	userModel := model.UserModel{}
	if err := Db.Where("user_name = ?", p.UserName).First(&userModel).Error; err != nil {
		return err, "账号不存在"
	} else {
		//解密判断
		err = bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(p.Password)) //验证（对比）
		if err != nil {
			return err, "密码错误"
		} else {
			return err, "登录成功"
		}
	}
}
