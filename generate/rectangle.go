package generate

import (
	"image"
	"image/draw"
)

type Rectangle struct {
	Bounds   *Bounds
	Position *Position
	Color    string
}

func NewRectangle(data map[string]interface{}) Layer {
	return &Rectangle{
		Bounds: &Bounds{
			Width:  int(data["width"].(float64)),
			Height: int(data["height"].(float64)),
		},
		Position: &Position{
			X: int(data["x"].(float64)),
			Y: int(data["y"].(float64)),
		},
		Color: data["color"].(string),
	}
}

func (r *Rectangle) Load() error {
	// Nothing to load
	return nil
}

func (r *Rectangle) Render(dest *image.RGBA) error {
	c, err := ParseHexColor(r.Color)
	if err != nil {
		return err
	}

	draw.Draw(dest, image.Rect(r.Position.X, r.Position.Y, r.Position.X+r.Bounds.Width, r.Position.Y+r.Bounds.Height), &image.Uniform{c}, image.Point{}, draw.Src)

	return nil
}
