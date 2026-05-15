/*
  - @Author: it.ww@qq.com
  - @Date: 2026/05/14
*/

package xerror

import "github.com/tangyong230/gox/i18n"

type CustomError struct {
	ErrNo  int
	ErrKey string
	Lang   string
	RawErr error
}

func (e *CustomError) Error() string {
	lang := e.Lang
	if lang == "" {
		lang = "zh-CN"
	}
	return i18n.GetMessage(e.ErrKey, lang)
}

func (e *CustomError) ErrorWithLang(lang string) string {
	return i18n.GetMessage(e.ErrKey, lang)
}

func (e *CustomError) Unwrap() error {
	return e.RawErr
}

func New(errNo int, errKey string) *CustomError {
	return &CustomError{
		ErrNo:  errNo,
		ErrKey: errKey,
	}
}

func NewWithLang(errNo int, errKey string, lang string) *CustomError {
	return &CustomError{
		ErrNo:  errNo,
		ErrKey: errKey,
		Lang:   lang,
	}
}

func Wrap(errNo int, errKey string, err error) *CustomError {
	if err == nil {
		return New(errNo, errKey)
	}
	return &CustomError{
		ErrNo:  errNo,
		ErrKey: errKey,
		RawErr: err,
	}
}

func WrapWithLang(errNo int, errKey string, lang string, err error) *CustomError {
	if err == nil {
		return NewWithLang(errNo, errKey, lang)
	}
	return &CustomError{
		ErrNo:  errNo,
		ErrKey: errKey,
		Lang:   lang,
		RawErr: err,
	}
}
