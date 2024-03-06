package factory

import "fmt"

type LineChart struct {
}

func (h *LineChart) Display() {
	fmt.Println("显示线性图")
}
