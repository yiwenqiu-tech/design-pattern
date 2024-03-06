package main

type SkinFactory interface {
	CreateButton() Button
	CreateTextField() TextField
	CreateComboBox() ComboBox
}
