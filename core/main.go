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

func SPrintFractionStyleWithPI(on, under int, onVar string) string {
	onStr := fmt.Sprintf("%d%v", on, onVar)
	if on == 1 {
		onStr = onVar
	}
	underStr := fmt.Sprintf("%d", under)
	if under != 1 {
		underStr = fmt.Sprintf("\\over %v}", underStr)
	} else {
		underStr = "}"
	}

	return fmt.Sprintf("{%v %v", onStr, underStr)
}

func toJsRadian() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		var degree int = args[0].Int()
		n := gcf(180, degree)
		return SPrintFractionStyleWithPI((degree / n), (180 / n), "\\pi")
	})
}

func main() {
	fmt.Println("Code agha")
	js.Global().Set("toRadian", toJsRadian())
	<-make(chan struct{})
}
