package treehouse

type Visibility struct {
	count     int
	maxHeight int
}

func (v *Visibility) next(nextHeight int) bool {
	if nextHeight > v.maxHeight {
		v.maxHeight = nextHeight
		return true
	}
	return false
}

func newVisibility() Visibility {
	return Visibility{
		count:     0,
		maxHeight: -1,
	}
}
