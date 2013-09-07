package main

import (
	"fmt"
)

type Code struct {
	ss string
}

var OnBeforeSweep func()

func main() {
	mp := make(map[string]*Code)
	OnBeforeSweep = func() {}
	OnBeforeSweep = nil
	delete(mp, "sdf")
	fmt.Println(mp["asdfaf"] == nil)
	fmt.Println(OnBeforeSweep == nil)
}
