//go:build js && wasm

package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)
	fmt.Println("from Go+WASM, edited!")
	js.Global().Set("addWASM", js.FuncOf(add))
	<-c
}

func add(this js.Value, args []js.Value) any {
	a, _ := getNumberFromInput("a")
	b, _ := getNumberFromInput("b")
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	c := a + b
	outputNumberToInput("out", c)
	fmt.Println("c: ", c)
	return ""
}

func getNumberFromInput(className string) (float64, error) {
	value := getElementByClassName(className).Get("value").String()
	return strconv.ParseFloat(value, 64)
}

func outputNumberToInput(className string, output float64) {
	getElementByClassName("wasm-output").Call("getElementsByClassName", className).Index(0).Set("value", js.ValueOf(output))
}

func getElementByClassName(className string) js.Value {
	elements := getAddArticle().Call("getElementsByClassName", className)
	return elements.Index(0)
}

func getAddArticle() js.Value {
	return js.Global().Get("document").Call("getElementById", "add")
}
