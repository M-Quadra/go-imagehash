package goimagehash

import (
	"errors"
	"image"

	"golang.org/x/image/draw"
)

type dHash struct {
	w, h   int
	scaler draw.Scaler
}

var dHashDefault = dHash{
	w: 9, h: 8,
	scaler: draw.BiLinear,
}

func (d *dHash) checkValue() {
	if d.w <= 0 {
		d.w = dHashDefault.w
	} else if d.w == 1 {
		d.w = 2
	}

	if d.h <= 0 {
		d.h = dHashDefault.h
	}

	if d.scaler == nil {
		d.scaler = dHashDefault.scaler
	}
}

// NewDifferenceHash dHash
//  recommend (w-1)*h = 8*n n ∈ N+
//  w>1
func NewDifferenceHash(w, h int, scaler draw.Scaler) dHash {
	dHash := dHash{
		w:      w,
		h:      h,
		scaler: scaler,
	}
	dHash.checkValue()
	return dHash
}

func (d dHash) Bits() int {
	return (d.w - 1) * d.h
}

func (d dHash) Hexs() int {
	return (d.Bits() + hexBits - 1) / hexBits
}

// Size lenght of []byte
func (d dHash) Size() int {
	return ((d.w-1)*d.h + byteBits - 1) / byteBits
}

func (d dHash) Sum(img image.Image) ([]byte, error) {
	if img == nil {
		return nil, errors.New("image is nil")
	}

	grayImg := image.NewGray(image.Rect(0, 0, d.w, d.h))
	d.scaler.Scale(grayImg, grayImg.Bounds(), img, img.Bounds(), draw.Over, nil)

	var (
		opt    = make([]byte, d.Size())
		offset = 0
	)

	for y := 0; y < grayImg.Bounds().Dy(); y++ {
		for x := 0; x < grayImg.Bounds().Dx()-1; x++ {
			if grayImg.GrayAt(x, y).Y < grayImg.GrayAt(x+1, y).Y {
				opt[offset/byteBits] |= 1 << (byteBits - 1 - (offset % byteBits))
			}
			offset++
		}
	}

	return opt, nil
}

func (d dHash) SumHex(img image.Image) (string, error) {
	data, err := d.Sum(img)
	if err != nil {
		return "", err
	}

	return EncodeToString(data, d.Hexs()), err
}
