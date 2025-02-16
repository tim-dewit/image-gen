package generate

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/draw"
	"net/http"
	"time"
)

import _ "image/jpeg"
import _ "image/png"

type Image struct {
	Position *Position `json:"position"`
	Bounds   *Bounds   `json:"bounds"`
	Url      string    `json:"url"`
	Image    *image.Image
}

func NewImage(data map[string]interface{}) Layer {
	return &Image{
		Position: &Position{
			X: int(data["x"].(float64)),
			Y: int(data["y"].(float64)),
		},
		Bounds: &Bounds{
			Width:  int(data["width"].(float64)),
			Height: int(data["height"].(float64)),
		},
		Url: data["url"].(string),
	}
}

func (i *Image) Load() error {
	start := time.Now()
	res, err := http.Get(i.Url)
	if err != nil {
		return err
	}

	img, _, err := image.Decode(res.Body)
	if err != nil {
		return err
	}

	i.Image = &img

	err = res.Body.Close()

	fmt.Printf("Loaded image %s in %s\n", i.Url, time.Since(start))

	return err
}

func (i *Image) Render(dest *image.RGBA) error {
	if i.Image == nil {
		return fmt.Errorf("failed to render image layer: image not loaded")
	}

	img := *i.Image
	is := img.Bounds().Size()
	if is.X != i.Bounds.Width || is.Y != i.Bounds.Height {
		img = resize.Resize(uint(i.Bounds.Width), uint(i.Bounds.Height), *i.Image, resize.Lanczos3)
	}

	r := image.Rect(i.Position.X, i.Position.Y, i.Position.X+i.Bounds.Width, i.Position.Y+i.Bounds.Height)

	draw.Draw(dest, r, img, image.Point{}, draw.Src)

	return nil
}
