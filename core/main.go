package WasmDegreeToRadians

import (
	"fmt"
)

func gcf(u, v int) int {
	if v == 0 {
		return u
	}
	return gcf(v, u%v)
}

func toRadian(degree int) string {
	n := gcf(180, degree)
	return fmt.Sprintf("\\frac{%v\\pi}{%v}", degree/n, 180/n)
}
