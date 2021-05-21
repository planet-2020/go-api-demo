/**
 * @Description:
 * @author zhouhongpan
 * @date 2021/5/20 16:36
 */
package user_service

import (
	"go-api-demo/internal/code"
	"go-api-demo/model"
	"go-api-demo/model/user_model"
)

/**
 * @Description: 注册会员
 * @author zhouhongpan
 */
type RegisterService struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required,min=11,max=20" label:"手机号"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=40" label:"密码"`
	Nickname string `form:"nickname" json:"nickname" binding:"max=50" label:"昵称"`
	Avatar string `form:"avatar" json:"avatar" label:"头像"`
	Sex uint `form:"sex" json:"sex" label:"性别"`
}

/**
 * @Description: 注册用户
 * @receiver r
 * @return error
 * @author zhouhongpan
 * @date 2021-05-21 09:35:26
 */
func (r *RegisterService) Register() error {
	user := user_model.User{
		Mobile: r.Mobile,
		Nickname: r.Nickname,
		Avatar: r.Avatar,
		Sex: r.Sex,
	}
	//加密密码
	if err := user.SetPassword(r.Password); err != nil {
		return code.ErrEncrypt
	}
	//创建用户
	result := model.DB.Create(&user)
	if result.Error != nil {
		return code.ErrUserCreate
	}

	return nil
}