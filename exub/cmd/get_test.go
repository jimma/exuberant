package cmd

import (
	"fmt"
	"testing"
)

func Test_getUrl(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test", args: args{url: "http://localhost:8080"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getUrl(tt.args.url)
		})
	}
}

func TestSlice(t *testing.T) {
	vals := make([]int, 5)
	vals[0] = 0
	vals[1] = 1
	vals[2] = 2
	vals[3] = 3
	vals[4] = 4
	dest := make([]int, 3)
	copy(dest, vals)
	fmt.Println(dest)
	fmt.Println(vals)
	fmt.Println(vals[1:3])
	fmt.Println(vals[3:5])

	for i := 0; i < 100; i++ {
		vals = append(vals, i)
	}
	fmt.Println(vals)
}

// 1
// 2
// -1
// 4
// -2
// 6
func TestIota(t *testing.T) {
	type codeType int
	const (
		_  = iota
		AA = iota
		BB
		CC = -1
		DD = iota
		EE = -2
		FF = iota
		GG
	)
	fmt.Println(AA)
	fmt.Println(BB)
	fmt.Println(CC)
	fmt.Println(DD)
	fmt.Println(EE)
	fmt.Println(FF)
	type ByteSize float64

	const (
		//_           = iota // ignore first value by assigning to blank identifier
		KB ByteSize = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)
	fmt.Println(PB)
}
