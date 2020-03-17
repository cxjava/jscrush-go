package gojav1

import (
	"fmt"
	"github.com/dop251/goja"
)

var (
	vm = goja.New()
)

func init() {
	if _, err := vm.RunString(jscrush_src); err != nil {
		fmt.Printf("jscrush_src init failed :%v \n", err)
	}
}

func Jscrush(js string) string {
	vm.Set("js_to_compress", js)
	value, err := vm.RunString(`(function(js_to_compress) {return jscrush(js_to_compress) })(js_to_compress)`)
	if err != nil {
		fmt.Printf("Jscrush:%v\n", err)
		return ""
	}
	return value.String()
}
