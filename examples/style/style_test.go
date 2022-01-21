package style

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Modifier interface {
	change()
}

type S1 struct {
	Value int
}

func (s S1) change() {
	s.Value = s.Value + 1
}

type S2 struct {
	Value int
}

func (s *S2) change() {
	s.Value = s.Value + 1
}

func TestPointerToInterface(t *testing.T) {
	var s1 = S1{Value: 1}
	var s2 = &S2{Value: 1}
	var m1 Modifier = s1
	var m2 Modifier = s2

	m1.change()
	m2.change()

	var b = s1.Value == s2.Value
	fmt.Printf(
		"S1 Value = %d, S2 Value = %d, (S1 == S2) is %v",
		s1.Value,
		s2.Value,
		b,
	)
	assert.NotEqual(t, s1.Value, s2.Value)
}

type S3 struct {
	Value int
}

//func TestVerifyInterfaceCompliance(t *testing.T) {
//	var _ Modifier = S3{}
//	var _ Modifier = &S3{}
//}
