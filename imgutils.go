// Package imgutils provides a unified SDK for image manipulation.
// It re-exports all functionality from the individual imgutils packages.
package imgutils

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

// Image wraps an image.Image with convenient methods.
type Image struct {
	img image.Image
}

// New creates a new Image wrapper.
func New(img image.Image) *Image {
	return &Image{img: img}
}

// Open loads an image from a file.
func Open(path string) (*Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return &Image{img: img}, nil
}

// Image returns the underlying image.Image.
func (i *Image) Image() image.Image {
	return i.img
}

// Bounds returns the image bounds.
func (i *Image) Bounds() image.Rectangle {
	return i.img.Bounds()
}

// Width returns the image width.
func (i *Image) Width() int {
	return i.img.Bounds().Dx()
}

// Height returns the image height.
func (i *Image) Height() int {
	return i.img.Bounds().Dy()
}

// SaveJPEG saves the image as a JPEG file.
func (i *Image) SaveJPEG(path string, quality int) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	if quality <= 0 || quality > 100 {
		quality = 85
	}
	return jpeg.Encode(f, i.img, &jpeg.Options{Quality: quality})
}

// SavePNG saves the image as a PNG file.
func (i *Image) SavePNG(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return png.Encode(f, i.img)
}

// EncodeJPEG encodes the image as JPEG to a writer.
func (i *Image) EncodeJPEG(w io.Writer, quality int) error {
	if quality <= 0 || quality > 100 {
		quality = 85
	}
	return jpeg.Encode(w, i.img, &jpeg.Options{Quality: quality})
}

// EncodePNG encodes the image as PNG to a writer.
func (i *Image) EncodePNG(w io.Writer) error {
	return png.Encode(w, i.img)
}

// Clone creates a copy of the image.
func (i *Image) Clone() *Image {
	return &Image{img: i.img}
}
