package image

import (
	"github.com/anthonynsimon/bild/effect"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
)

func Resize(input, output string, h, w int) error {
	img, e := imgio.Open(input)
	if e != nil {
		return e
	}
	inverted := effect.Invert(img)
	resized := transform.Resize(inverted, w, h, transform.Linear)
	if err := imgio.Save(output, resized, imgio.PNGEncoder()); err != nil {
		return e
	}
	return nil
}

func ResizeOptional(input, output string, h, w, maxSize int) error {
	//img, e := imgio.Open(input)
	//if e != nil {
	//	return e
	//}
	//img.Bounds().Size()
	return Resize(input, output, h, w)
}
