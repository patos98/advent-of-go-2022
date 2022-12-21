package monkey

type Monkey interface {
	GetName() string
	GetNumber(monkeys map[string]Monkey) int
}
