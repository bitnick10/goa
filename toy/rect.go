package toy

type Rect struct {
	Height, Width int
}

func (rect *Rect) Area() int {
	return rect.Height * rect.Width
}
