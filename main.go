package main

import (
	"fmt"
	qrcode "github.com/skip2/go-qrcode"
)

const (
	BLACK = "\033[40m  \033[0m"
	WHITE = "\033[47m  \033[0m"
)

func stripBorder(bitmap [][]bool, borderWidth int) [][]bool {
	var m [][]bool

	for i := borderWidth; i < len(bitmap)-borderWidth; i++ {
		row := bitmap[i]
		m = append(m, row[borderWidth:len(row)-borderWidth])
	}
	return m
}

func main() {
	q, err := qrcode.New("http://www.baidu.com", qrcode.Medium)
	if err != nil {
		fmt.Print(err)
	}

	out := ""
	bitmap := stripBorder(q.Bitmap(), 3)

	for _, row := range bitmap {
		for _, cell := range row {
			if cell {
				out += BLACK
			} else {
				out += WHITE
			}
		}
		out += "\n"
	}
	fmt.Print(out)
}
