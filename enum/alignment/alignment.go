/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package alignment

import "reflect"

const (
	Undefined     Alignment = iota
	Left
	Center
	Right
	Justified
	JustifiedAll
)

type Alignment int64

func (a Alignment) Int() int64 {
	return int64(a)
}

func AlignmentOf(id int64) Alignment {

	var a Alignment
	rv := reflect.Indirect(reflect.ValueOf(&a))
	rv.SetInt(id)

	return rv.Interface().(Alignment)
}
