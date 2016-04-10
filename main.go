package main

import (
	"fmt"
	qrcode "github.com/skip2/go-qrcode"
	"os"
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

func printHelp() {
	help := `
USAGE: qr [message]

Example:
qr http://www.zhex.me`

	fmt.Println(help)
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}

	q, err := qrcode.New(os.Args[1], qrcode.Medium)
	if err != nil {
		fmt.Print("Ops, something wrong with the message, qrcode can not be generated.")
		os.Exit(0)
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
