package main

func main() {
	var factory Factory
	factory = &HistogramChartFactory{} // 该操作可以放到配置文件，再通过反射来创建。
	chart := factory.CreateChart()
	chart.Display()

	factory = &LineChartFactory{}
	chart = factory.CreateChart()
	chart.Display()
}
