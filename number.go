package ckutils

import "math/rand"

// 랜덤 숫자를 구한다.
//
// [PRE]
//	min - 최소값
//	max - 최대값
//
// [POST]
//	랜덤 숫자
func GetRandomNumber(min int, max int) int {
	return min + rand.Intn(max - min)
}
