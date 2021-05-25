/**
 * @Description:
 * @author zhouhongpan
 * @date 2021/5/20 16:36
 */
package user_service

import (
	"errors"
	"go-api-demo/internal/code"
	"go-api-demo/internal/config"
	"go-api-demo/model"
	"go-api-demo/model/user_model"
	"go-api-demo/pkg/token"
	"gorm.io/gorm"
	"time"
)

/**
 * @Description: 注册用户
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
 * @Description: 校验参数
 * @receiver r
 * @return error
 * @author zhouhongpan
 * @date 2021-05-21 15:14:05
 */
func (r *RegisterService) valid() error {
	//手机号是否存在
	query := map[string]interface{} {
		"mobile": r.Mobile,
	}
	var user user_model.User
	result := model.DB.Where(query).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return result.Error
	}
	return code.ErrUserRegisterMobileExist
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

	//表单验证
	if err := r.valid(); err != nil {
		return err
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

/**
 * @Description: 用户登录
 * @author zhouhongpan
 */
type LoginService struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required,min=11,max=20" label:"手机号"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=40" label:"密码"`
}

/**
 * @Description: 用户登录
 * @receiver l
 * @return string token字符串
 * @return error
 * @author zhouhongpan
 * @date 2021-05-21 15:35:46
 */
func (l *LoginService) Login() (string, error) {
	query := map[string]interface{} {
		"mobile": l.Mobile,
		"delete_time": 0,
	}
	var user user_model.User
	result := model.DB.Where(query).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound){
			return "", code.ErrUserNotExist
		}
		return "", result.Error
	}
	if !user.CheckPassword(l.Password) {
		return "", code.ErrUserPasswordError
	}
	tokenString, err := token.GenerateToken(config.Conf.App.JwtSubject,config.Conf.App.JwtSecret,config.Conf.App.JwtExpireTime,user.ID,user.Mobile)

	return tokenString, err
}

type UserInfo struct {
	Uid uint `json:"uid"`
	Mobile string `json:"mobile"`
	Nickname string `json:"nickname"`
	Avatar string `json:"avatar"`
	Sex uint `json:"sex"`
	CreateTime uint `json:"create_time"`
}

/**
 * @Description: 获取用户信息
 * @receiver userInfo
 * @param uid
 * @return *UserInfo
 * @return error
 * @author zhouhongpan
 * @date 2021-05-25 11:20:18
 */
func (userInfo *UserInfo) Get(uid interface{}) error {
	query := map[string]interface{}{
		"id": uid,
		"delete_time": 0,
	}
	var user user_model.User
	res := model.DB.Where(query).First(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return code.ErrUserNotExist
		}
		return res.Error
	}
	userInfo.Uid = user.ID
	userInfo.Mobile = user.Mobile
	userInfo.Nickname = user.Nickname
	userInfo.Avatar = user.Avatar
	userInfo.Sex = user.Sex
	userInfo.CreateTime = user.CreateTime
	return nil
}

type UpdateService struct {
	Nickname string `form:"nickname" json:"nickname"`
	Avatar string `form:"avatar" json:"avatar"`
	Sex uint `form:"sex" json:"sex"`
}

/**
 * @Description: 更新用户信息
 * @param uid
 * @param nickname
 * @param avatar
 * @param sex
 * @return *UserInfo
 * @return error
 * @author zhouhongpan
 * @date 2021-05-25 12:04:42
 */
func (u *UpdateService) Update(uid interface{}) (*UserInfo, error) {
	query := map[string]interface{}{
		"id": uid,
		"delete_time": 0,
	}
	params := make(map[string]interface{})
	if len(u.Nickname) > 0 {
		params["nickname"] = u.Nickname
	}
	if len(u.Avatar) > 0 {
		params["avatar"] = u.Avatar
	}
	if u.Sex > 0 && u.Sex <= 2 {
		params["sex"] = u.Sex
	}
	if len(params) == 0 {
		return nil, code.ErrUserUpdateParams
	}
	res := model.DB.Model(&user_model.User{}).Where(query).Updates(params)
	if res.Error != nil {
		return nil, code.ErrUserUpdate
	}
	var userInfo UserInfo
	err := userInfo.Get(uid)
	if err != nil {
		return nil, err
	}
	return &userInfo, nil
}

/**
 * @Description: 注销用户
 * @param uid
 * @return error
 * @author zhouhongpan
 * @date 2021-05-25 14:20:35
 */
func Delete(uid interface{}) error {
	query := map[string]interface{}{
		"id": uid,
	}
	nowTime := time.Now().Unix()
	params := map[string]interface{}{
		"mobile": nowTime,
		"delete_time": nowTime,
	}
	res := model.DB.Model(&user_model.User{}).Where(query).Updates(params)
	if res.Error != nil {
		return code.ErrUserDelete
	}
	return nil
}