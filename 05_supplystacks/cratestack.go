package supplystacks

type CrateStack struct {
	crates []string
}

func newCrateStack() CrateStack {
	return CrateStack{crates: []string{}}
}

func (s *CrateStack) pop() string {
	crate := s.lastElement()
	s.crates = s.crates[:len(s.crates)-1]
	return crate
}

func (s *CrateStack) push(c string) {
	s.crates = append(s.crates, c)
}

func (s *CrateStack) lastElement() string {
	return s.crates[len(s.crates)-1]
}
