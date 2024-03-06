package main

import "design-pattern/factory"

type HistogramChartFactory struct {
}

func (*HistogramChartFactory) CreateChart() factory.Chart {
	var chart factory.Chart
	chart = &factory.HistogramChart{}

	// 初始化操作

	return chart
}
