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

	// 业务
	// 社区


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