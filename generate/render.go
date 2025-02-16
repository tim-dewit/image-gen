package generate

import (
	"fmt"
	"image"
	"sync"
	"time"
)

type GenerateRequest struct {
	Bounds Bounds        `json:"bounds"`
	Layers []interface{} `json:"layers"`
}

func Render(r *GenerateRequest) (*image.RGBA, error) {
	dest := image.NewRGBA(image.Rect(0, 0, r.Bounds.Width, r.Bounds.Height))
	layers := BuildLayers(r.Layers)

	ls := time.Now()
	c := make(chan error)
	wg := sync.WaitGroup{}
	wg.Add(len(layers))
	for _, l := range layers {
		go func(l Layer) {
			defer wg.Done()
			err := l.Load()
			if err != nil {
				fmt.Printf("failed to load layer: %s\n", err.Error())

				c <- err
			}
		}(l)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for err := range c {
		return nil, err
	}
	fmt.Printf("Loaded layers in %s\n", time.Since(ls))

	rs := time.Now()
	for _, l := range layers {
		err := l.Render(dest)
		if err != nil {
			return nil, err
		}
	}
	fmt.Printf("Rendered layers in %s\n", time.Since(rs))

	return dest, nil
}
