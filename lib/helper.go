package lib

import (
	"reflect"
	"runtime"
)

func GetFuncName(fc interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(fc).Pointer()).Name()
}
