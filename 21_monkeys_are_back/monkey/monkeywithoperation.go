package monkey

type Operator func(int, int) int

type MonkeyWithOperation struct {
	name        string
	monkeyLeft  string
	monkeyRight string
	operator    Operator
}

func (m MonkeyWithOperation) GetName() string {
	return m.name
}

func (m MonkeyWithOperation) GetNumber(monkeys map[string]Monkey) int {
	return m.operator(
		monkeys[m.monkeyLeft].GetNumber(monkeys),
		monkeys[m.monkeyRight].GetNumber(monkeys),
	)
}
