package monkey

type Monkey interface {
	GetName() string
	GetMonkeyLeft() string
	GetMonkeyRight() string
	GetNumber(monkeys map[string]Monkey) (int, bool)
	GetMissingNumberForResult(map[string]Monkey, int) int
}
