package main

import "design-pattern/factory"

type LineChartFactory struct {
}

func (*LineChartFactory) CreateChart() factory.Chart {
	var chart factory.Chart
	chart = &factory.LineChart{}

	// 初始化操作

	return chart
}
