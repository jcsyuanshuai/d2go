package structs

type Modifier interface {
	change()
}

type S1 struct {
	Value int
}

func NewS1(v int) S1 {
	return S1{Value: v}
}

func (s S1) change() {
	s.Value = s.Value + 1
}

type S2 struct {
	Value int
}

func NewS2(v int) *S2 {
	return &S2{Value: v}
}

func (s *S2) change() {
	s.Value = s.Value + 1
}
