/**
 * @Description: 用户
 * @author zhouhongpan
 * @date 2021/5/20 15:45
 */
package user_model

import (
	"go-api-demo/model"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	model.CustomModel
	Mobile string `gorm:"unique"`
	Password string
	Nickname string
	Avatar string
	Sex uint
}

var (
	// PassWordCost 密码加密难度
	PassWordCost = 12
)

/**
 * @Description: 自定义表明
 * @receiver User
 * @return string
 * @author zhouhongpan
 * @date 2021-05-20 15:54:19
 */
func (User) TableName() string {
	return "go_user"
}

/**
 * @Description: 设置密码
 * @receiver user
 * @param password
 * @return error
 * @author zhouhongpan
 * @date 2021-05-20 17:00:42
 */
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return nil
	}
	user.Password = string(bytes)
	return nil
}

/**
 * @Description: 校验密码
 * @receiver user
 * @param password
 * @return bool
 * @author zhouhongpan
 * @date 2021-05-21 15:26:11
 */
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}