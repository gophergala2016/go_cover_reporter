// Code coverage animated gif cover_badge.go
package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"io"
	"io/ioutil"
	"log"
	"math"
	"strconv"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

func toText(frameNumber int, percentage float64) string {
	if percentage == 0.00 {
		return "cover:0.00%"
	}
	finalCount := int(math.Floor(percentage))
	increaseFloatCount, _ := strconv.ParseFloat(strconv.FormatFloat(percentage/100, 'f', 2, 64), 64) // single frame increase
	if finalCount > frameNumber {
		return "cover:" + strconv.FormatFloat(increaseFloatCount*float64(frameNumber), 'f', 2, 32) + "%"
	} else {
		return "cover:" + strconv.FormatFloat(percentage, 'f', 2, 32) + "%"
	}
}

func coverBadge(out io.Writer, percent float64) {

	const (
		fontfile = "AnonymousProB.ttf"
		fontsize = 13
		dpi      = 75

		grayIndex  = 0
		greenIndex = 1
		whiteIndex = 2

		badgeLength = 90
		badgeHeight = 20

		numberOfFrames     = 100
		delayBetweenFrames = 5
		lastFrameDelay     = 15
	)

	var palette = []color.Color{color.RGBA{0x5b, 0x5b, 0x5b, 0xff}, color.RGBA{0x4b, 0xc5, 0x1d, 0xff}, color.White}

	fontBytes, err := ioutil.ReadFile(fontfile)
	if err != nil {
		log.Println(err)
		return
	}

	font, err := truetype.Parse(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}

	fontoptions := truetype.Options{}
	fontoptions.Size = fontsize
	fontface := truetype.NewFace(font, &fontoptions)

	fg := image.White
	c := freetype.NewContext()
	c.SetDPI(dpi)
	c.SetFont(font)
	c.SetFontSize(fontsize)

	c.SetSrc(fg)

	anim := gif.GIF{LoopCount: numberOfFrames}

	for i := 0; i < numberOfFrames; i++ {
		rect := image.Rect(0, 0, badgeLength, badgeHeight)
		img := image.NewPaletted(rect, palette)

		for verticalPosition := 0; verticalPosition < badgeHeight; verticalPosition++ {
			var line int
			if i < int(math.Floor(percent)) {
				line = i
			} else {
				line = int(math.Floor(percent))
			}
			for horisontalPosition := 0; horisontalPosition < line; horisontalPosition++ {
				img.SetColorIndex(horisontalPosition, verticalPosition, greenIndex)
			}

		}

		c.SetDst(img)
		c.SetClip(img.Bounds())
		draw.Draw(img, img.Bounds(), img.SubImage(img.Bounds()), image.ZP, draw.Src)

		for j, x := range toText(i, percent) {
			_, ok := fontface.GlyphAdvance(rune(x))
			if ok != true {
				log.Println(err)
				return
			}

			c.DrawString(string(x), freetype.Pt(j*7+5, 15))

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
