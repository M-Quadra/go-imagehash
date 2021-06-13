package goimagehash

import (
	"errors"
	"image"

	"golang.org/x/image/draw"
)

// AHash aHash
type aHash struct {
	w, h   int
	scaler draw.Scaler
}

var aHashDefault = aHash{
	w: 8, h: 8,
	scaler: draw.BiLinear,
}

func (a *aHash) checkValue() {
	if a.w <= 0 {
		a.w = aHashDefault.w
	}
	if a.h <= 0 {
		a.h = aHashDefault.h
	}

	if a.w*a.h < 2 {
		if a.w <= 1 {
			a.w = 2
		} else if a.h <= 1 {
			a.h = 2
		}
	}

	if a.scaler == nil {
		a.scaler = aHashDefault.scaler
	}
}

// NewAverageHash aHash
//  recommend w*h = 8*n n ∈ N+
func NewAverageHash(w, h int, scaler draw.Scaler) aHash {
	aHash := aHash{
		w:      w,
		h:      h,
		scaler: scaler,
	}
	aHash.checkValue()
	return aHash
}

func (a aHash) Bits() int {
	return a.w * a.h
}

func (a aHash) Hexs() int {
	return (a.Bits() + hexBits - 1) / hexBits
}

// Size lenght of []byte
func (a aHash) Size() int {
	return (a.Bits() + byteBits - 1) / byteBits
}

// Sum aHash result of img
func (a aHash) Sum(img image.Image) ([]byte, error) {
	if img == nil {
		return nil, errors.New("image is nil")
	}

	grayImg := image.NewGray(image.Rect(0, 0, a.w, a.h))
	a.scaler.Scale(grayImg, grayImg.Bounds(), img, img.Bounds(), draw.Over, nil)

	avg := 0.0
	for _, v := range grayImg.Pix {
		avg += float64(v)
	}
	avg /= float64(len(grayImg.Pix))

	opt := make([]byte, a.Size())
	for i, v := range grayImg.Pix {
		if float64(v) > avg {
			opt[i/byteBits] |= 1 << (byteBits - 1 - (i % byteBits))
		}
	}

	return opt, nil
}

func (a aHash) SumHex(img image.Image) (string, error) {
	data, err := a.Sum(img)
	if err != nil {
		return "", err
	}

	return EncodeToString(data, a.Hexs()), err
}
