
# cc-jianfan

[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/go-cc/cc-jianfan?status.svg)](http://godoc.org/github.com/go-cc/cc-jianfan)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-cc/cc-jianfan)](https://goreportcard.com/report/github.com/go-cc/cc-jianfan)
[![travis Status](https://travis-ci.org/go-cc/cc-jianfan.svg?branch=master)](https://travis-ci.org/go-cc/cc-jianfan)

## TOC
- [API](#api)
  - [> jianfan_test.go](#-jianfan_testgo)

中文简繁互转  
Chinese-Character Jian<=>Fan converter


# API

#### > jianfan_test.go
```go
package ccst_test

import (
	"fmt"

	ccst "github.com/go-cc/cc-jianfan"
)

// for standalone test, change package to main and the next func def to,
// func main() {
func Example_output() {
	sj := "看书学习。\n名著：《红楼梦》〖清〗曹雪芹 著、高鹗 续／『人民文学』出版社／1996—9月30日／59.70【元】，《三国演义》〖明〗罗贯中。"

	sf := ccst.S2T(sj)
	fmt.Println(sf)
	// Output:
	//   not here!
	fmt.Println(ccst.T2S(sf))
	sj = ccst.T2S("增補臺灣標準IT詞彙。")
	fmt.Println(sj);
	fmt.Println(ccst.S2T(sj))

	// Output:
	// 看書學習。
	// 名著：《紅樓夢》〖清〗曹雪芹 著、高鶚 續／『人民文學』齣版社／1996—9月30日／59.70【元】，《三國演義》〖明〗羅貫中。
	// 看书学习。
	// 名著：《红楼梦》〖清〗曹雪芹 著、高鹗 续／『人民文学』出版社／1996—9月30日／59.70【元】，《三国演义》〖明〗罗贯中。
	// 增补台湾标准IT词汇。
	// 增補臺灣標準IT詞匯。
}
```

All patches welcome.
