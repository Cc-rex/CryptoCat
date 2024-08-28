package utils

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

// GetValidMsg 返回结构体中的msg参数
func GetValidMsg(err error, obj any) string {
	//使用的时候需要传object指针
	getObj := reflect.TypeOf(obj)
	//将err接口断言为指定类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		//断言成功
		for _, e := range errs {
			//循环每一个错误信息
			//根据报错字段名，获取结构体具体字段
			if f, exist := getObj.Elem().FieldByName(e.Field()); exist {
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}
	return err.Error()
}
