package factory

import "fmt"

type HistogramChart struct {
}

func (h *HistogramChart) Display() {
	fmt.Println("显示柱状图")
}
