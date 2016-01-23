// Code coverage animated gif
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
)

var palette = []color.Color{color.Black, color.RGBA{0xff, 0x00, 0x00, 0xff}, color.RGBA{0x00, 0xff, 0x00, 0xff}, color.White}

const (
	blackIndex = 0
	redIndex   = 1
	greenIndex = 2
	whiteIndex = 3
)

func coverBadge(out io.Writer, percent int) {
	const (
		bedgeLength        = 50
		bedgeHeight        = 25
		numberOfFrames     = 100
		delayBetweenFrames = 5
		lastFrameDelay     = 15
	)

	anim := gif.GIF{LoopCount: numberOfFrames}

	for i := 0; i < numberOfFrames; i++ {
		rect := image.Rect(0, 0, bedgeLength, bedgeHeight)
		img := image.NewPaletted(rect, palette)

		for verticalPosition := 0; verticalPosition < bedgeHeight; verticalPosition++ {
			for horisontalPosition := 0; horisontalPosition < i; horisontalPosition++ {
				img.SetColorIndex(horisontalPosition, verticalPosition, greenIndex)
			}
		}

		switch {
		case i == numberOfFrames-1:
			anim.Delay = append(anim.Delay, lastFrameDelay)
		default:
			anim.Delay = append(anim.Delay, delayBetweenFrames)
		}

		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
