package test

import (
	"encoding/hex"
	"fmt"
	"testing"

	goimagehash "github.com/M-Quadra/go-imagehash"
	"github.com/M-Quadra/kazaana/v2"
	goimagehashCorona10 "github.com/corona10/goimagehash"
)

func TestDHash(t *testing.T) {
	testAry := []struct {
		give int
		want string
	}{
		{
			give: 2,
			want: "6",
		},
		{
			give: 3,
			want: "260",
		},
		{
			give: 4,
			want: "33aa",
		},
		{
			give: 5,
			want: "18d3490",
		},
		{
			give: 6,
			want: "0c3575d34",
		},
		{
			give: 7,
			want: "0e0c3b5cbb320",
		},
		{
			give: 8,
			want: "0707072ba9c9cc44",
		},
	}
	for _, v := range testAry {
		dHash := goimagehash.NewDifferenceHash(v.give+1, v.give, nil)
		dHashHex, err := dHash.SumHex(imgAria)
		if kazaana.HasError(err) {
			t.Fail()
			return
		}

		if dHashHex != v.want {
			fmt.Println(dHashHex, v.want)
			t.Fail()
			return
		}
	}
}

func TestDHashCorona10(t *testing.T) {
	dHashData, err := goimagehash.Sum.DHash(imgAria)
	if kazaana.HasError(err) {
		t.Fail()
		return
	}
	fmt.Println(hex.EncodeToString(dHashData))

	aHashC10, err := goimagehashCorona10.DifferenceHash(imgAria)
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

	dif := goimagehash.HammingDistance(dHashData, outData)
	fmt.Println(float64(dif) / float64(len(outData)*8))
}
