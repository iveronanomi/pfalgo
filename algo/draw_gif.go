package algo

import (
	"image"
	"image/color"
	"image/gif"
	"os"
)

var (
	defaultDelay = 0
	PointSpace   = color.RGBA{255, 255, 255, 255}
	PointWall    = color.RGBA{0, 0, 102, 255}
	PointStart   = color.RGBA{204, 5, 0, 255}
	PointWalk    = color.RGBA{51, 204, 51, 255}
	PointFinish  = color.RGBA{153, 0, 153, 255}
)

// GifGraph ...
type GifGraph struct {
	width  int
	height int
	result *gif.GIF
	filename string
}

var palette = []color.Color{
	PointSpace,
	PointWall,
	PointStart,
	PointFinish,
	PointWalk,
}

// NewGifGraph create new draw for graph visualisation
func NewGifGraph(w, h int, filename string) *GifGraph {
	gg := &GifGraph{
		height: h,
		width:  w,
		filename: filename,
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
func (gg *GifGraph) Point(x, y int, pointType color.RGBA) {
	gg.result.Image[0].Set(x, y, pointType)
}

// Visit set point as visited
func (gg *GifGraph) Visit(x, y int) {
	//create new image from previous one
	f := *gg.result.Image[len(gg.result.Image)-1]
	pix := make([]uint8, 0, len(f.Pix))
	pix = append(pix, f.Pix...)
	f.Pix = pix

	//set new visited node
	f.Set(x, y, PointWalk)

	//append new image frame, and delay to draw
	gg.result.Delay = append(gg.result.Delay, defaultDelay)
	gg.result.Image = append(gg.result.Image, &f)
}

// Save ...
func (gg *GifGraph) Save() {
	//debug: save each frame
	// for idx, v := range im.result.Image {
	// 	f, _ := os.OpenFile(fmt.Sprintf("out/%d.draw", idx), os.O_WRONLY|os.O_CREATE, 0600)
	// 	defer f.Close()
	// 	err := draw.EncodeAll(f, &draw.GIF{
	// 		Delay: []int{0},
	// 		Image: []*image.Paletted{
	// 			v,
	// 		},
	// 	})
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	f, _ := os.OpenFile(gg.filename, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()

	err := gif.EncodeAll(f, gg.result)

	if err != nil {
		panic(err)
	}
}
