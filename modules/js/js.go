package js

import (
	"strings"

	"github.com/dop251/goja"
)

func CreateJSEngine() *goja.Runtime {

	//************************************************
	//Setup the js engine with all the passed function
	//************************************************
	vm := goja.New()

	vm.Set("clean_space", strings.ReplaceAll)

	return vm
}

func EvaluateJSCode(code string, model interface{}) (interface{}, error) {

	//************************************************
	//Evaluate the code
	//************************************************
	vm := CreateJSEngine()
	vm.Set("model", model)
	val, err := vm.RunString(code)
	if err != nil {
		return nil, err
	}
	return val.Export(), nil
}
