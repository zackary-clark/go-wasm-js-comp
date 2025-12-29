//go:build js && wasm

package main

import (
	"syscall/js"
)

func Add(this js.Value, args []js.Value) any {
	form := GetElementById("add-form")
	a := GetFloatFromForm(form, "a")
	b := GetFloatFromForm(form, "b")

	c := a + b

	SetOutput(form, "wasm-out", c)
	return ""
}
