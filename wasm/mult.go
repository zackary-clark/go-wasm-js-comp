//go:build js && wasm

package main

import (
	"syscall/js"
)

func Mult(this js.Value, args []js.Value) any {
	form := GetElementById("mult-form")
	a := GetFloatFromForm(form, "a")
	b := GetFloatFromForm(form, "b")

	c := a * b

	SetOutput(form, "wasm-out", c)
	return ""
}
