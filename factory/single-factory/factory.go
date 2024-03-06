package main

import "design-pattern/factory"

type Factory struct {
}

func (f *Factory) GetChart(t string) factory.Chart {
	if t == "histogram" {
		return &factory.HistogramChart{}
	} else if t == "line" {
		return &factory.LineChart{}
	} else {
		panic("unSupport")
	}
}
