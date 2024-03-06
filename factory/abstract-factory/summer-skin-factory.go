package main

type SummerSkinFactory struct {
}

func (s *SummerSkinFactory) CreateButton() Button {
	return &SummerButton{}
}

func (s *SummerSkinFactory) CreateTextField() TextField {
	return &SummerTextField{}
}

func (s *SummerSkinFactory) CreateComboBox() ComboBox {
	return &SummerComboBox{}
}
