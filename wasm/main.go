//go:build js && wasm

package main

import (
	"syscall/js"
)

func main() {
	randomChannel := make(chan struct{}, 0)
	js.Global().Set("addWASM", js.FuncOf(Add))
	js.Global().Set("multWASM", js.FuncOf(Mult))
	<-randomChannel
}

func SetOutput(form js.Value, fieldName string, value float64) {
	output := GetFieldFromForm(form, "wasm-out")
	output.Set("value", js.ValueOf(value))
}

func GetFloatFromForm(form js.Value, fieldName string) float64 {
	return GetFieldFromForm(form, fieldName).Get("valueAsNumber").Float()
}

func GetFieldFromForm(form js.Value, fieldName string) js.Value {
	return form.Get("elements").Get(fieldName)
}

func GetElementById(id string) js.Value {
	return js.Global().Get("document").Call("getElementById", id)
}
