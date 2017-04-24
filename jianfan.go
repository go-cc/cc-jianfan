// Package ccst provides Conversion between Simplified and Traditional Chinese-Characters.
//
// This is a simple implementation of Chinese one-to-one conversion.
// If you need more advanced converter, please visit
// OpenCC https://github.com/BYVoid/OpenCC
package ccst

import "fmt"

var t2sMapping = make(map[rune]rune)
var s2tMapping = make(map[rune]rune)

func init() {
	if len([]rune(chT)) != len([]rune(chS)) {
		panic("cht and chs data length not equal")
	}

	for index := 0; index < len([]rune(chS)); index++ {
		runeValueS := []rune(chS)[index]
		runeValueT := []rune(chT)[index]
		// if runeValueS == '𬸪' || runeValueS == '𬸪' {
		// 	fmt.Fprintf(os.Stderr, "%c:%c\n", runeValueS, runeValueT)
		// 	continue
		// }
		t2sMapping[runeValueT] = runeValueS
		// Simplified => Traditional is one to many, pick first
		if _, ok := s2tMapping[runeValueS]; !ok {
			s2tMapping[runeValueS] = runeValueT
		}
	}
}

// T2S converts Traditional Chinese to Simplified Chinese
func T2S(s string) string {
	var chs []rune
	for _, runeValue := range s {
		if v, ok := t2sMapping[runeValue]; ok {
			chs = append(chs, v)
		} else {
			chs = append(chs, runeValue)
		}
	}
	return string(chs)
}

// S2T converts Simplified Chinese to Traditional Chinese
func S2T(s string) string {
	var cht []rune
	for _, runeValue := range s {
		if v, ok := s2tMapping[runeValue]; ok {
			cht = append(cht, v)
		} else {
			cht = append(cht, runeValue)
		}
	}
	return string(cht)
}

// STDump dumps Simplified to Traditional Chinese dict
func STDump() {
	for kk, vv := range s2tMapping {
		fmt.Printf("%c:%c\n", kk, vv)
	}
}

// TSDump dumps Traditional to Simplified Chinese dict
func TSDump() {
	for kk, vv := range t2sMapping {
		fmt.Printf("%c:%c\n", vv, kk)
	}
}
