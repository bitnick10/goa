package jasmine

import (
	"fmt"
	// "github.com/Bitnick2002/forest/console"
	"reflect"
	"runtime/debug"
	"strings"
	// "testing"
)

type expect struct {
	hold interface{}
}

func (this expect) isEqual(i interface{}) bool {
	holdType := reflect.TypeOf(this.hold)
	iType := reflect.TypeOf(i)
	holdValue := reflect.ValueOf(this.hold)
	iValue := reflect.ValueOf(i)
	if holdType != iType {
		return false
	}
	switch holdType.Kind() {
	case reflect.Bool:
		return holdValue.Bool() == iValue.Bool()
	case reflect.Int:
		return holdValue.Int() == iValue.Int()
	case reflect.String:
		return holdValue.String() == iValue.String()
	default:
		fmt.Println("type not define please contact author ", holdType.Kind())
		return false
	}
	return false
}
func getValueString(value reflect.Value) string {
	switch value.Kind() {
	case reflect.Bool:
		return fmt.Sprint(value.Bool())
	case reflect.Int:
		return fmt.Sprint(value.Int())
	case reflect.String:
		return fmt.Sprint(value.String())
	default:
		return fmt.Sprintln("type not define please contact author ", value.Kind())
	}
	return fmt.Sprintln("type not define please contact author ", value.Kind())
}

func (this expect) toBe(i interface{}) {
	g_specNumber++
	if this.isEqual(i) {
		// fmt.Println("this == tv")
		return
	} else {
		g_errorNumber++
		str := string(debug.Stack())
		strs := strings.Split(str, "\n")
		g_rootSpec.addErrorInfo(g_level, strs[4])
		g_rootSpec.addErrorInfo(g_level, strs[5])
		exp := fmt.Sprint("expect ", getValueString(reflect.ValueOf(this.hold)), " tobe ", getValueString(reflect.ValueOf(i)))
		g_rootSpec.addErrorInfo(g_level, exp)
		// console.Color(console.RED).Println(strs[4])
		// console.Color(console.RED).Println(strs[5])
		// console.Color(console.RED).Println("expect is ", reflect.ValueOf(this.hold), "  fact : ", reflect.ValueOf(i))
		// fmt.Println(tv.v)
		// log.Fatal("...")
		// gt.Error(" error")
	}
}
func (this expect) ToEqual(i interface{}) {
	// fmt.Println("expect is : ", reflect.ValueOf(this.hold))
	this.toBe(i)
}
func (this expect) ToBeTruthy() {
	this.toBe(true)
}
func (this expect) ToBeFalsy() {
	this.toBe(false)
}
func (this expect) ToBe(i interface{}) {
	this.toBe(i)
}

func Expect(i interface{}) expect {
	// fmt.Println("in EXPECTexpect is : ", reflect.ValueOf(i))
	return expect{i}
}
