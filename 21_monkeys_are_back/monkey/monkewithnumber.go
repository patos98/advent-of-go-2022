package monkey

type MonkeyWithNumber struct {
	name   string
	number int
}

func (m MonkeyWithNumber) GetName() string {
	return m.name
}

func (m MonkeyWithNumber) GetNumber(monkeys map[string]Monkey) int {
	return m.number
}
