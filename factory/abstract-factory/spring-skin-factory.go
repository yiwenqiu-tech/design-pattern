package main

type SpringSkinFactory struct {
}

func (s *SpringSkinFactory) CreateButton() Button {
	return &SpringButton{}
}

func (s *SpringSkinFactory) CreateTextField() TextField {
	return &SpringTextField{}
}

func (s *SpringSkinFactory) CreateComboBox() ComboBox {
	return &SpringComboBox{}
}
