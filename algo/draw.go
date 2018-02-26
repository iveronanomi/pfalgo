package algo

import (
	"image"
	"image/color"
	"image/gif"
	"os"
)

var (
	defaultDelay = 0
	pixSpace     = color.RGBA{255, 255, 255, 255}
	pixWall      = color.RGBA{0, 0, 102, 255}
	pixStart     = color.RGBA{204, 5, 0, 255}
	pixWalk      = color.RGBA{51, 204, 51, 255}
	pixFinish    = color.RGBA{153, 0, 153, 255}
)

// GifGraph ...
type GifGraph struct {
	width  int
	height int
	result *gif.GIF
}

var palette = []color.Color{
	pixSpace,
	pixWall,
	pixStart,
	pixFinish,
	pixWalk,
}

// NewGifGraph create new gif for graph visualisation
func NewGifGraph(w, h int) *GifGraph {
	gg := &GifGraph{
		height: h,
		width:  w,
	}
	gg.result = &gif.GIF{
		Image: []*image.Paletted{
			image.NewPaletted(image.Rect(0, 0, w, h), palette),
		},
		Delay: []int{defaultDelay},
	}
	return gg
}

// SetWall ...
func (gg *GifGraph) SetWall(x, y int) {
	gg.result.Image[0].Set(x, y, pixWall)
}

// SetStart ...
func (gg *GifGraph) SetStart(x, y int) {
	gg.result.Image[0].Set(x, y, pixStart)
}

// SetFinish ...
func (gg *GifGraph) SetFinish(x, y int) {
	gg.result.Image[0].Set(x, y, pixFinish)
}

// Visit set point as visited
func (gg *GifGraph) Visit(x, y int) {
	//create new image from previous one
	f := *gg.result.Image[len(gg.result.Image)-1]
	pix := make([]uint8, 0, len(f.Pix))
	pix = append(pix, f.Pix...)
	f.Pix = pix

	//set new visited node
	f.Set(x, y, pixWalk)

	//append new image frame, and delay to gif
	gg.result.Delay = append(gg.result.Delay, defaultDelay)
	gg.result.Image = append(gg.result.Image, &f)
}

// Save ...
func (gg *GifGraph) Save(filename string) {
	//debug: save each frame
	// for idx, v := range im.result.Image {
	// 	f, _ := os.OpenFile(fmt.Sprintf("out/%d.gif", idx), os.O_WRONLY|os.O_CREATE, 0600)
	// 	defer f.Close()
	// 	err := gif.EncodeAll(f, &gif.GIF{
	// 		Delay: []int{0},
	// 		Image: []*image.Paletted{
	// 			v,
	// 		},
	// 	})
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	f, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()

	err := gif.EncodeAll(f, gg.result)

	if err != nil {
		panic(err)
	}
}
