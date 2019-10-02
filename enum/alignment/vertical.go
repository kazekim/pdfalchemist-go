/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package alignment


import "reflect"

const (
	VerticalUndefined     VerticalAlignment = iota
	TOP
	VerticalCenter
	Bottom
	BaseLine
)

type VerticalAlignment int64

func (a VerticalAlignment) Int() int64 {
	return int64(a)
}

func VerticalAlignmentOf(id int64) VerticalAlignment {

	var a VerticalAlignment
	rv := reflect.Indirect(reflect.ValueOf(&a))
	rv.SetInt(id)

	return rv.Interface().(VerticalAlignment)
}

