// Code coverage animated gif
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
)

var palette = []color.Color{color.RGBA{0x5b, 0x5b, 0x5b, 0xff}, color.RGBA{0x4b, 0xc5, 0x1d, 0xff}, color.White}

const (
	grayIndex  = 0
	greenIndex = 1
	whiteIndex = 2
)

func coverBadge(out io.Writer, percent int) {
	const (
		badgeLength        = 90
		badgeHeight        = 20
		numberOfFrames     = 100
		delayBetweenFrames = 5
		lastFrameDelay     = 15
	)

	anim := gif.GIF{LoopCount: numberOfFrames}

	for i := 0; i < numberOfFrames; i++ {
		rect := image.Rect(0, 0, badgeLength, badgeHeight)
		img := image.NewPaletted(rect, palette)

		for verticalPosition := 0; verticalPosition < badgeHeight; verticalPosition++ {
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
