package jasmine

import (
	"github.com/Bitnick2002/goa/console"
)

const (
	DESCRIBE = iota
	IT
)

type spec struct {
	level       int
	description string
	feature     int
	errorInfo   []string
	children    []spec
}

var (
	g_rootSpec    spec
	g_level       int
	g_specNumber  int
	g_errorNumber int
)

func (this *spec) addErrorInfo(level int, errorString string) {
	pChildren := &this.children
	var lastSpec *spec
	for level > 1 {
		lastSpec = &(*pChildren)[len(*pChildren)-1]
		pChildren = &((*lastSpec).children)
		level--
	}
	lastSpec = &(*pChildren)[len(*pChildren)-1]
	(*lastSpec).errorInfo = append((*lastSpec).errorInfo, errorString)
}

func (this *spec) addChild(level int, child spec) {
	// console.Color(console.YELLOW).Println("aaaaaaaaaaaaaa = ", this.description)
	pChildren := &this.children
	// console.Color(console.YELLOW).Println("in add child len = ", len(*pChildren))
	for level > 1 {
		// console.Color(console.YELLOW).Println("in for")
		// fmt.Println("level = ", level)
		// fmt.Println("len(pChildren) = ", len(pChildren))
		// fmt.Println("level 1 len(pChildren) = ", level)
		lastSpec := &(*pChildren)[len(*pChildren)-1]
		// console.Color(console.YELLOW).Println(lastSpec.description)
		pChildren = &((*lastSpec).children)
		level--
	}
	*pChildren = append(*pChildren, child)
	//temp := &g_rootSpec.children
	//*temp = append(pChildren, child)
	// console.Color(console.YELLOW).Println("in temp child len = ", len(*temp))
	// console.Color(console.YELLOW).Println("in after child len = ", len(*pChildren))
	// console.Color(console.YELLOW).Println("in aftrer g_rootSpec len = ", g_rootSpec.children[0].description)
	// console.Color(console.YELLOW).Println("in aftrer g_rootSpec len = ", len(g_rootSpec.children[0].children))

	// fmt.Println("after append level = ", level)
	// fmt.Println("after append len(slice) = ", len(slice))
}

func PrintSpec() {
	printSpec(g_rootSpec, "")
	// console.Println(g_rootSpec.description)
}
func printAllError(s spec) {
	if len(s.errorInfo) > 0 {
		for _, err := range s.errorInfo {
			console.Color(console.RED).Println(err)
		}
	}
	for _, child := range s.children {
		if s.level == 0 {
			printAllError(child)
		} else {
			printAllError(child)
		}
	}
}
func printConclusion() {
	if g_errorNumber == 0 {
		console.Color(console.GREEN).Println(g_specNumber, "all passed")
	} else {
		console.Color(console.RED).Println(g_errorNumber, "failed")
	}
}
func printSpec(s spec, space string) {
	if s.level == 0 {
		goto __label
	}
	if s.feature == DESCRIBE {
		console.Color(console.CYAN).Println("\n" + space + s.description)
	} else {
		console.Color(console.GREEN).Println(space + s.description)
	}
	if len(s.errorInfo) > 0 {
		for _, err := range s.errorInfo {
			console.Color(console.RED).Println(space + err)
		}
	}
__label:
	for _, child := range s.children {
		if s.level == 0 {
			printSpec(child, space)
		} else {
			printSpec(child, space+"  ")
		}
	}
}
