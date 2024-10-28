package main

import (
	"fmt"
	"syscall/js"
)

func gcf(u, v int) int {
	if v == 0 {
		return u
	}
	return gcf(v, u%v)
}

func toRadian() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		var degree int = args[0].Int()
		n := gcf(180, degree)
		return fmt.Sprintf("\\frac{%v\\pi}{%v}", degree/n, 180/n)
	})
}

func main() {
	fmt.Println("Code agha")
	js.Global().Set("toRadian", toRadian())
	<-make(chan struct{})
}
