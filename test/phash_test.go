package test

import (
	"encoding/hex"
	"fmt"
	"testing"

	goimagehash "github.com/M-Quadra/go-imagehash"
	"github.com/M-Quadra/kazaana/v2"
	goimagehashCorona10 "github.com/corona10/goimagehash"
	"golang.org/x/image/draw"
)

func TestPHash8x8(t *testing.T) {
	// JohannesBuchner/imagehash
	// imagehash.phash(imgAria)
	pHashDataPy, err := hex.DecodeString("af2cd5d23e31b1c0")
	if kazaana.HasError(err) {
		t.Fail()
		return
	}

	// corona10/goimagehash
	var pHashDataCorona10 []byte
	{
		pHash, err := goimagehashCorona10.PerceptionHash(imgAria)
		if kazaana.HasError(err) {
			t.Fail()
			return
		}

		fmt.Println(pHash.ToString()[2:])
		pHashDataCorona10, err = hex.DecodeString(pHash.ToString()[2:])
		if kazaana.HasError(err) {
			t.Fail()
			return
		}
	}

	pHashData, err := goimagehash.Sum.PHash(imgAria)
	if kazaana.HasError(err) {
		t.Fail()
		return
	}

	rate := float64(8*8-goimagehash.HammingDistance(pHashData, pHashDataPy)) / float64(8*8)
	fmt.Println("[JohannesBuchner/imagehash] 8x8 similarity:", rate)
	rate = float64(8*8-goimagehash.HammingDistance(pHashData, pHashDataCorona10)) / float64(8*8)
	fmt.Println("[corona10/goimagehash] 8x8 similarity:", rate)
}

func TestPHash16x16(t *testing.T) {
	// JohannesBuchner/imagehash
	// imagehash.phash(img, hash_size=16)
	pHashDataPy, err := hex.DecodeString("aff92c13d580d2b43e4b3339b1afc4d4cc5a2c3333343b983dcfc8c9e49c7070")
	if kazaana.HasError(err) {
		t.Fail()
		return
	}

	// corona10/goimagehash
	var pHashDataCorona10 []byte
	{
		pHash, err := goimagehashCorona10.ExtPerceptionHash(imgAria, 16, 16)
		if kazaana.HasError(err) {
			t.Fail()
			return
		}

		fmt.Println(pHash.ToString()[2:])
		pHashDataCorona10, err = hex.DecodeString(pHash.ToString()[2:])
		if kazaana.HasError(err) {
			t.Fail()
			return
		}
	}

	pHashData, err := goimagehash.NewPerceptualHash(16, 16, 4, draw.BiLinear).Sum(imgAria)
	if kazaana.HasError(err) {
		t.Fail()
		return
	}

	rate := float64(16*16-goimagehash.HammingDistance(pHashData, pHashDataPy)) / float64(16*16)
	fmt.Println("[JohannesBuchner/imagehash] 16x16 similarity:", rate)
	rate = float64(16*16-goimagehash.HammingDistance(pHashData, pHashDataCorona10)) / float64(16*16)
	fmt.Println("[corona10/goimagehash] 16x16 similarity:", rate)
}
