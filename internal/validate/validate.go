/**
 * @Description: 参数校验
 * @author zhouhongpan
 * @date 2021/5/21 14:00
 */
package validate

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"go-api-demo/internal/code"
	"reflect"
)

var Trans ut.Translator

/**
 * @Description: 注册中文翻译
 * @author zhouhongpan
 * @date 2021-05-21 14:22:41
 */
func InitValidate()  {
	uni := ut.New(zh.New())
	Trans, _ = uni.GetTranslator("zh")

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册翻译器
		_= zhTranslations.RegisterDefaultTranslations(v, Trans)
		//注册一个函数，获取struct tag里自定义的label作为字段名
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := fld.Tag.Get("label")
			return name
		})
	}
}

/**
 * @Description: 获取校验错误
 * @param err
 * @return *code.Errno
 * @author zhouhongpan
 * @date 2021-05-21 14:23:00
 */
func GetValidateErr(err error) *code.Errno {
	errs := err.(validator.ValidationErrors)
	var errFirst string
	for _, e := range errs {
		errFirst = e.Translate(Trans)
		break
	}
	return &code.Errno{
		Code:    code.ErrParam.Code,
		Message: errFirst,
	}
}
