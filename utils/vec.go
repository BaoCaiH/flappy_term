package utils

type Vec2 struct {
	X, Y int
}

func (v *Vec2) Add(w *Vec2) Vec2 {
	return Vec2{v.X + w.X, v.Y + w.Y}
}
