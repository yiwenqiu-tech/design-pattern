package main

func main() {
	f := Factory{}
	chart := f.GetChart("line")
	chart.Display()

	chart = f.GetChart("histogram")
	chart.Display()
}
