package gogen

import (
	"fmt"
	"reflect"
	"strings"
)

type ControllerStarter interface {
	Start()
}

type UsecaseRegisterer interface {
	AddUsecase(inports ...any)
	GetUsecase(nameStructType any) (any, error)
}

type ControllerRegisterer interface {
	ControllerStarter
	UsecaseRegisterer
	RegisterRouter()
}

type BaseController struct {
	inportObjs map[any]any
}

func NewBaseController() UsecaseRegisterer {
	return &BaseController{
		inportObjs: map[any]any{},
	}
}

func (r *BaseController) GetUsecase(nameStructType any) (any, error) {
	x := reflect.TypeOf(nameStructType).String()
	packageName := x[:strings.Index(x, ".")]
	uc, ok := r.inportObjs[packageName]
	if !ok {

		msg := "usecase with package \"%s\" is not registered yet in application. Please call 'gogen application' manually"

		return nil, fmt.Errorf(msg, packageName)
	}
	return uc, nil
}

func (r *BaseController) AddUsecase(inports ...any) {
	for _, inport := range inports {
		x := reflect.ValueOf(inport).Elem().Type().String()
		packagePath := x[:strings.Index(x, ".")]
		r.inportObjs[packagePath] = inport
	}
}
