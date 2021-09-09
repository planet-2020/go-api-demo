/**
 * @author zhouhongpan
 * @date 2021/5/19 9:56
 */
package code

var (

	// 系统
	OK = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}

	// 鉴权
	ErrToken = &Errno{Code: 30001, Message: "Token错误"}
	ErrTokenNeed = &Errno{Code: 30002, Message: "Token不能为空"}

	//公共
	ErrParam = &Errno{Code: 20001, Message: "请求参数错误"}

	// 用户
	ErrEncrypt = &Errno{Code: 20101, Message: "密码加密错误"}
	ErrUserCreate = &Errno{Code: 20102, Message: "用户创建错误"}
	ErrUserRegisterMobileExist = &Errno{Code: 20103, Message: "手机号已存在"}
	ErrUserNotExist = &Errno{Code: 20104, Message: "用户不存在"}
	ErrUserPasswordError = &Errno{Code: 20105, Message: "密码不正确"}
	ErrUserUpdateParams = &Errno{Code: 20106, Message: "用户更新参数错误"}
	ErrUserUpdate = &Errno{Code: 20107, Message: "用户更新失败"}
	ErrUserDelete = &Errno{Code: 20108, Message: "用户注销失败"}

	// 消息队列
	ErrMqSendNotTopic = &Errno{Code: 20201, Message: "topic 不能为空"}
	ErrMqSendNotMessage = &Errno{Code: 20202, Message: "message 不能为空"}
	ErrMqSendFail = &Errno{Code: 20203, Message: "消息队列发送失败"}

)

/**
 * @Description: 错误
 * @author zhouhongpan
 */
type Errno struct {
	Code int
	Message string
}

func (err Errno) Error() string {
	return err.Message
}

/**
 * @Description: 解析错误
 * @param err 错误
 * @return code 错误码
 * @return message 错误信息
 * @author zhouhongpan
 * @date 2021-05-19 11:06:15
 */
func DecodeErr(err error) (code int, message string) {
	if err == nil {
		return OK.Code, OK.Message
	}
	switch value := err.(type) {
		case *Errno:
			return value.Code, value.Message
	default:
	}
	return InternalServerError.Code, InternalServerError.Message
}