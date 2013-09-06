package console

import (
	"fmt"
)

const (
	BLACK   = "\x1b[30;3m"
	RED     = "\x1b[31;3m"
	GREEN   = "\x1b[32;3m"
	YELLOW  = "\x1b[33;3m"
	BLUE    = "\x1b[34;3m"
	MAGENTA = "\x1b[35;3m"
	CYAN    = "\x1b[36;3m"
	WHITE   = "\x1b[37;3m"
)

const colorCancel = "\x1b[0m"

type colorType string

func Color(color colorType) colorType {
	return color
}

func PrintColors() {
	for color := 30; color < 38; color++ {
		fmt.Printf("%d: ", color)
		for bold := 0; bold < 10; bold++ {
			fmt.Printf("\x1b[%d;%dmHello! \x1b[0m", color, bold)
		}
		fmt.Println()
	}
	Color(BLACK).Println("BLACK")
	Color(RED).Println("RED")
	Color(GREEN).Println("GREEN")
	Color(YELLOW).Println("YELLOW")
	Color(BLUE).Println("BLUE")
	Color(MAGENTA).Println("MAGENTA")
	Color(CYAN).Println("CYAN")
	Color(WHITE).Println("WHITE")
}

func (this colorType) Println(a ...interface{}) (n int, err error) {
	fmt.Print(this)
	n, err = fmt.Println(a...)
	fmt.Print(colorCancel)
	return n, err
}
