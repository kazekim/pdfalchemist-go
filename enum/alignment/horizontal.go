/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package alignment

import "reflect"

const (
	HorizontalUndefined     HorizontalAlignment = iota
	Left
	HorizonatalCenter
	Right
	Justified
	JustifiedAll
)

type HorizontalAlignment int64

func (a HorizontalAlignment) Int() int64 {
	return int64(a)
}

func HorizontalAlignmentOf(id int64) HorizontalAlignment {

	var a HorizontalAlignment
	rv := reflect.Indirect(reflect.ValueOf(&a))
	rv.SetInt(id)

	return rv.Interface().(HorizontalAlignment)
}
