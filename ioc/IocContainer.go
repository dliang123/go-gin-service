package ioc

import (
	"github.com/codegangsta/inject"
	"go-gin-demo/util"
	"reflect"
)

var injector inject.Injector

type IocContainer struct {
}

func (*IocContainer) GetInjector() (inject.Injector) {

	if injector == nil {
		injector = inject.New()
	}
	return injector
}

func (*IocContainer) Register(injector inject.Injector) {
	ResponseUtil := &util.ResponseUtil{}
	injector.Set(reflect.TypeOf(ResponseUtil), reflect.ValueOf(ResponseUtil))

}
