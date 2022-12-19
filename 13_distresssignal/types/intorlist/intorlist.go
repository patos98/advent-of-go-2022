package intorlist

type ValueType int

const (
	INT ValueType = iota
	LIST
)

type IntOrList struct {
	ValueType ValueType
	IntValue  int
	ListValue []*IntOrList
	Parent    *IntOrList
}

func (item *IntOrList) GetListValue() []*IntOrList {
	if item.ValueType == LIST {
		return item.ListValue
	} else {
		return []*IntOrList{item}
	}
}

func NewList(parent *IntOrList) *IntOrList {
	return &IntOrList{
		ValueType: LIST,
		IntValue:  0,
		ListValue: []*IntOrList{},
		Parent:    parent,
	}
}

func NewInt(value int, parent *IntOrList) *IntOrList {
	return &IntOrList{
		ValueType: INT,
		IntValue:  value,
		ListValue: []*IntOrList{},
		Parent:    parent,
	}
}
