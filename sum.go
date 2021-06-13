package goimagehash

import "image"

type (
	sum    int
	sumHex int
)

const (
	// Sum .Sum() of default imageHash
	Sum = sum(0)
	// SumHex .SumHex() of default imageHash
	SumHex = sumHex(0)
)

func (s sum) AHash(img image.Image) ([]byte, error) {
	return aHashDefault.Sum(img)
}

func (s sumHex) AHash(img image.Image) (string, error) {
	return aHashDefault.SumHex(img)
}

func (s sum) DHash(img image.Image) ([]byte, error) {
	return dHashDefault.Sum(img)
}

func (s sumHex) DHash(img image.Image) (string, error) {
	return dHashDefault.SumHex(img)
}
