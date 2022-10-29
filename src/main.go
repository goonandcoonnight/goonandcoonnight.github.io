// finally. goon game

package main

import (
	"fmt"
	"syscall/js"
)

var htmlString = "<h1>goon game hahahahaha</h1>"

func printHtml() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return htmlString
	})
}

func main() {
	fmt.Printf("goon game\n")

	js.Global().Set("printHtml", printHtml())
}
