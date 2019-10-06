/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package xml

type Attributes interface {
	Length() int
	Uri(var1 int) string
	LocalName(var1 int) string
	QName(var1 int) string
	TypeFromInt(var1 int) string
	Value(var1 int) string
	Index(var1 string, var2 *string) int
	TypeFromString(var1 string, var2 *string)
	ValueFromString(var1 string, var2 *string)
}
