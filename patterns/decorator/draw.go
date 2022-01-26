package decorator

type Drawable interface {
	Draw() string
}

type Square struct{}

func (s *Square) Draw() string {
	return "this is a square"
}

type ColoredSquare struct {
	square Drawable
	color  string
}

func NewColoredSquare(square Drawable, color string) *ColoredSquare {
	return &ColoredSquare{
		square: square,
		color:  color,
	}
}

func (s *ColoredSquare) Draw() string {
	return s.square.Draw() + ", color is " + s.color
}
