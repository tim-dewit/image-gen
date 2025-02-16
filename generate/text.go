package generate

import (
	"fmt"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"image"
	"io"
	"net/http"
	"time"
)

type Text struct {
	FontUrl  string   `json:"fontUrl"`
	Content  string   `json:"content"`
	Size     int      `json:"size"`
	Color    string   `json:"color"`
	Position Position `json:"position"`
	Align    string   `json:"align"`
	Font     *truetype.Font
}

func NewText(data map[string]interface{}) Layer {
	return &Text{
		FontUrl: data["fontUrl"].(string),
		Content: data["content"].(string),
		Size:    int(data["size"].(float64)),
		Color:   data["color"].(string),
		Align:   data["align"].(string),
		Position: Position{
			X: int(data["x"].(float64)),
			Y: int(data["y"].(float64)),
		},
	}
}

func (t *Text) Load() error {
	start := time.Now()

	res, err := http.Get(t.FontUrl)
	if err != nil {
		return err
	}

	fb, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	f, err := freetype.ParseFont(fb)
	if err != nil {
		return err
	}

	t.Font = f

	fmt.Printf("Loaded font %s in %s\n", t.FontUrl, time.Since(start))

	return nil
}

func (t *Text) Render(dest *image.RGBA) error {
	fc, err := ParseHexColor(t.Color)
	if err != nil {
		return err
	}

	c := freetype.NewContext()
	c.SetFont(t.Font)
	c.SetFontSize(float64(t.Size))
	c.SetClip(dest.Bounds())
	c.SetDst(dest)
	c.SetSrc(image.NewUniform(fc))

	_, err = c.DrawString(t.Content, freetype.Pt(t.Position.X+t.resolveHorizontalOffset(), t.Position.Y))

	return err
}

func (t *Text) resolveHorizontalOffset() int {
	ff := truetype.NewFace(t.Font, &truetype.Options{
		Size: float64(t.Size),
	})
	fb := font.MeasureString(ff, t.Content)

	switch t.Align {
	case "center":
		return -fb.Floor() / 2
	case "right":
		return -fb.Floor()
	default:
		return 0
	}
}
