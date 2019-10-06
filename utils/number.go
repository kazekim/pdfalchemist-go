/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package utils

import "strconv"

func ParseFloat(value string) *float64 {
	if v, err := strconv.ParseFloat(value, 64); err == nil {
		return &v
	}

	return nil
}


func ParseInt(value string) *int {

	if v, err := strconv.Atoi(value); err == nil {
		return &v
	}
	return nil
}