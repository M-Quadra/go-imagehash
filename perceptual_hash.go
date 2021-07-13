package goimagehash

import (
	"errors"
	"image"
	"math"
	"sort"
	"sync"

	"golang.org/x/image/draw"
)

type pHash struct {
	w, h   int
	factor int
	scaler draw.Scaler
}

var pHashDefault = pHash{
	w: 8, h: 8,
	factor: 4,
	scaler: draw.BiLinear,
}

func (p *pHash) checkValue() {
	if p.w <= 0 {
		p.w = pHashDefault.w
	}
	if p.h <= 0 {
		p.h = pHashDefault.h
	}
	if p.w*p.h < 2 {
		p.w = pHashDefault.w
		p.h = pHashDefault.h
	}

	if p.scaler == nil {
		p.scaler = aHashDefault.scaler
	}
}

// NewPerceptualHash pHash
//  recommend w*h = 8*n n ∈ N+, factor = 4
func NewPerceptualHash(w, h, factor int, scaler draw.Scaler) pHash {
	pHash := pHash{
		w:      w,
		h:      h,
		factor: factor,
		scaler: scaler,
	}
	pHash.checkValue()
	return pHash
}

func (p pHash) Bits() int {
	return p.w * p.h
}

func (p pHash) Hexs() int {
	return (p.Bits() + hexBits - 1) / hexBits
}

// Size lenght of []byte
func (p pHash) Size() int {
	return (p.Bits() + byteBits - 1) / byteBits
}

func (p pHash) Sum(img image.Image) ([]byte, error) {
	if img == nil {
		return nil, errors.New("image is nil")
	}

	w, h := p.w*p.factor, p.h*p.factor
	grayImg := image.NewGray(image.Rect(0, 0, w, h))
	p.scaler.Scale(grayImg, grayImg.Bounds(), img, img.Bounds(), draw.Over, nil)

	pixAry := make([][]float64, 0, h)
	for y := 0; y < h; y++ {
		row := make([]float64, 0, w)
		for x := 0; x < w; x++ {
			row = append(row, float64(grayImg.GrayAt(x, y).Y))
		}
		pixAry = append(pixAry, row)
	}

	dctAry := dct2D(pixAry)
	ary := make([]float64, 0, p.w*p.h)
	for y := 0; y < p.h; y++ {
		ary = append(ary, dctAry[y][:p.w]...)
	}
	med := median(ary)

	optAry := make([]byte, p.Size())
	for i, v := range ary {
		if v > med {
			optAry[i/byteBits] |= 1 << (byteBits - 1 - (i % byteBits))
		}
	}

	return optAry, nil
}

func (p pHash) SumHex(img image.Image) (string, error) {
	data, err := p.Sum(img)
	if err != nil {
		return "", err
	}

	return EncodeToString(data, p.Hexs()), err
}

// 1-D DCT-II
//   scipy.fftpack.dct
func dct1D(x []float64) []float64 {
	N := len(x)
	X := make([]float64, N)

	for k := 0; k < N; k++ {
		for n := 0; n < N; n++ {
			X[k] += x[n] * math.Cos((math.Pi/float64(N))*(float64(n)+0.5)*float64(k))
		}
		X[k] *= 2
	}
	return X
}

// 2-D DCT-II
//   scipy.fftpack.dct(scipy.fftpack.dct(ipt, axis=0), axis=1)
func dct2D(ipt [][]float64) [][]float64 {
	h, w := len(ipt), len(ipt[0])
	optAry := make([][]float64, h)
	for y := 0; y < h; y++ {
		optAry[y] = make([]float64, w)
	}

	wg := sync.WaitGroup{}

	wg.Add(w)
	for x := 0; x < w; x++ {
		go func(x int) {
			defer wg.Done()

			ary := make([]float64, 0, h)
			for y := 0; y < h; y++ {
				ary = append(ary, ipt[y][x])
			}
			ary = dct1D(ary)
			for y := 0; y < h; y++ {
				optAry[y][x] = ary[y]
			}
		}(x)
	}
	wg.Wait()

	wg.Add(h)
	for y := 0; y < h; y++ {
		go func(y int) {
			defer wg.Done()

			optAry[y] = dct1D(optAry[y])
		}(y)
	}
	wg.Wait()

	return optAry
}

// Quick Median soon?
func median(ary []float64) float64 {
	tmpAry := make([]float64, len(ary))
	tmpAry = append(tmpAry, ary...)
	sort.Float64s(tmpAry)
	return tmpAry[len(ary)/2]
}
