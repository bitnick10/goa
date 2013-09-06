package jasmine

import (
	"fmt"
)

func used() {
	fmt.Println("...")
}

func Describe(description string, f func()) {
	g_level++
	child := spec{g_level, description, DESCRIBE, make([]string, 0), make([]spec, 0)}
	g_rootSpec.addChild(g_level, child)
	f()
	g_level--
	if g_level == 0 {
		PrintSpec()
		fmt.Println("")
		printAllError(g_rootSpec)
		printConclusion()
	}
}

func It(description string, f func()) {
	g_level++
	child := spec{g_level, description, IT, make([]string, 0), make([]spec, 0)}
	g_rootSpec.addChild(g_level, child)
	f()
	g_level--
	// if g_level == 0 {
	// 	PrintSpec()
	// }
}
