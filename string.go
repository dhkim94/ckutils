package ckutils

import (
	"math/rand"
	"time"
)

var lowerAlphabet []string = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
	"k", "l", "m", "n", "o", "p", "q", "r", "s", "t",
	"u", "v", "w", "x", "y", "z",
}

var upperAlphabet []string = []string{
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
	"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T",
	"U", "V", "W", "X", "Y", "Z",
}

var numberAndAlphabet []string = []string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
	"k", "l", "m", "n", "o", "p", "q", "r", "s", "t",
	"u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
	"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T",
	"U", "V", "W", "X", "Y", "Z",
}

var numberAlphabetAndSymbol []string = []string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
	"k", "l", "m", "n", "o", "p", "q", "r", "s", "t",
	"u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
	"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T",
	"U", "V", "W", "X", "Y", "Z",
	"~", "!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "+", "=",
}

// 랜덤한 문자열을 만든다.
//
// [PRE]
//	length - 랜덤 스트링 길이
//	charSet - 랜덤 스트링에 포함할 문자열 세트
//
// [POST]
//	 원하는 길이의 랜덤 스트링을 리턴 한다.
func GetRandomString(length uint, charSet []string) (str string) {
	rand.Seed(time.Now().UTC().UnixNano())

	max := len(charSet) - 1
	len2 := int(length)
	
	for i := 0; i < len2; i++ {
		str += charSet[GetRandomNumber(0, max)]
	}

	return
}

// 숫자, 영문자, 심볼로 이루어진 랜덤 스트링을 구한다.
//
// [PRE]
//	length - 랜덤 스트링 길이
//
// [POST]
//	랜덤 하게 생성한 스트링
func GetRandomNumberAlphabetAndSymbolString(length uint) string {
	return GetRandomString(length, numberAlphabetAndSymbol)
}