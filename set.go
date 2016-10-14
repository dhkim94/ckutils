package ckutils

import "reflect"

// slice 가 특정 요소를 포함하고 있는지 체크 한다.
//
// [PRE]
//	set - slice
//	val - 포함하고 있는지 체크할 value
//
// [POST]
//	특정 요소를 포함하고 있다면 true 아니면 false 를 리턴 한다.
func IsContain(set interface{}, val interface{}) bool {
	arrV := reflect.ValueOf(set)

	if arrV.Kind() == reflect.Slice {
		for i := 0; i < arrV.Len(); i++ {

			// XXX - panics if slice element points to an unexported struct field
			// see https://golang.org/pkg/reflect/#Value.Interface
			if arrV.Index(i).Interface() == val {
				return true
			}
		}
	}

	return false
}