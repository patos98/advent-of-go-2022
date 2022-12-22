package monkey

const HUMAN_MONKEY_NAME = "humn"

type MonkeyWithNumber struct {
	name   string
	number int
}

func (m MonkeyWithNumber) GetName() string {
	return m.name
}

func (m MonkeyWithNumber) GetMonkeyLeft() string {
	return ""
}

func (m MonkeyWithNumber) GetMonkeyRight() string {
	return ""
}

func (m MonkeyWithNumber) GetNumber(monkeys map[string]Monkey) (int, bool) {
	return m.number, m.name == HUMAN_MONKEY_NAME
}

func (m MonkeyWithNumber) GetMissingNumberForResult(monkeys map[string]Monkey, result int) int {
	return m.number
}
