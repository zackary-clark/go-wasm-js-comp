//go:build js && wasm

package main

import (
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("addWASM", js.FuncOf(add))
	js.Global().Set("multWASM", js.FuncOf(mult))
	<-c
}

func mult(this js.Value, args []js.Value) any {
	form := js.Global().Get("document").Call("getElementById", "mult-form")
	a := form.Get("elements").Get("a").Get("valueAsNumber").Float()
	b := form.Get("elements").Get("b").Get("valueAsNumber").Float()
	c := a * b
	output := form.Get("elements").Get("wasm-out")
	output.Set("value", js.ValueOf(c))
	return ""
}

func add(this js.Value, args []js.Value) any {
	form := js.Global().Get("document").Call("getElementById", "add-form")
	a := form.Get("elements").Get("a").Get("valueAsNumber").Float()
	b := form.Get("elements").Get("b").Get("valueAsNumber").Float()
	c := a + b
	output := form.Get("elements").Get("wasm-out")
	output.Set("value", js.ValueOf(c))
	return ""
}
