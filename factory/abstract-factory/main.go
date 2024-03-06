package main

func main() {
	var factory SkinFactory
	factory = &SpringSkinFactory{}

	button := factory.CreateButton()
	button.Display()
	comboBox := factory.CreateComboBox()
	comboBox.Display()
	textField := factory.CreateTextField()
	textField.Display()

	factory = &SummerSkinFactory{}
	button = factory.CreateButton()
	button.Display()
	comboBox = factory.CreateComboBox()
	comboBox.Display()
	textField = factory.CreateTextField()
	textField.Display()
}
