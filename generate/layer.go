package generate

import "image"

type Bounds struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Layer interface {
	Load() error
	Render(dest *image.RGBA) error
}

func BuildLayers(data []interface{}) []Layer {
	layers := make([]Layer, len(data))
	for i, d := range data {
		layer := d.(map[string]interface{})
		switch layer["type"] {
		case "text":
			layers[i] = NewText(layer)
		case "image":
			layers[i] = NewImage(layer)
		case "rectangle":
			layers[i] = NewRectangle(layer)
		}
	}

	return layers
}
