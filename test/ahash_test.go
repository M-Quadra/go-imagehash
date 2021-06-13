package test

import (
	"encoding/hex"
	"fmt"
	"testing"

	goimagehash "github.com/M-Quadra/go-imagehash"
	"github.com/M-Quadra/kazaana/v2"
	goimagehashCorona10 "github.com/corona10/goimagehash"
)

func TestAHash(t *testing.T) {
	testAry := []struct {
		give int
		want string
	}{
		{
			give: 2,
			want: "3",
		},
		{
			give: 3,
			want: "a30",
		},
		{
			give: 4,
			want: "9166",
		},
		{
			give: 5,
			want: "8c42e70",
		},
		{
			give: 6,
			want: "8e105d79e",
		},
		{
			give: 7,
			want: "c78e08178f9f0",
		},
		{
			give: 8,
			want: "c3c38301157c7c3e",
		},
	}
	for _, v := range testAry {
		aHash := goimagehash.NewAverageHash(v.give, v.give, nil)
		aHashHex, err := aHash.SumHex(imgAria)
		if kazaana.HasError(err) {
			t.Fail()
			return
		}

		if aHashHex != v.want {
			fmt.Println(aHashHex, v.want)
			t.Fail()
			return
		}
	}
}

func TestAHashCorona10(t *testing.T) {
	aHashData, err := goimagehash.Sum.AHash(imgAria)
	if kazaana.HasError(err) {
		t.Fail()
		return
	}
	fmt.Println(hex.EncodeToString(aHashData))

	aHashC10, err := goimagehashCorona10.AverageHash(imgAria)
	if kazaana.HasError(err) {
		t.Fail()
		return
	}
	fmt.Println(aHashC10.ToString()[2:])

	outData, err := hex.DecodeString(aHashC10.ToString()[2:])
	if kazaana.HasError(err) {
		t.Fail()
		return
	}

	dif := goimagehash.HammingDistance(aHashData, outData)
	fmt.Println(float64(dif) / float64(len(outData)*8))
}
