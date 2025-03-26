package io_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/bytedance/sonic"
	gojson "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
)

type User struct {
	Id       int
	Name     string
	Weight   float32
	BodyTall float32
	Age      int
	Sex      byte `json:"gender"`
	Ignore   int  `json:"-" gorm:"-"`
}

type Book struct {
	ISBN     string `json:"isbn"`
	Name     string
	Price    float32  `json:"price"`
	Author   *User    `json:"author"`
	Keywords []string `json:"kws"`
	Local    map[int]bool
}

var (
	user = User{
		Name:   "钱钟书",
		Age:    57,
		Sex:    1,
		Ignore: 1,
	}
	book = Book{
		ISBN:     "4243547567",
		Name:     "围城",
		Price:    34.8,
		Author:   &user,
		Keywords: []string{"爱情", "民国", "留学"},
		Local:    map[int]bool{2: true, 3: false},
	}
)

func init() {
	bs, _ := json.Marshal(book)
	fmt.Printf("序列化之后的字节数%d\n", len(bs)) //195个字节
}

// 标准库
func BenchmarkJsonStd(b *testing.B) {
	var inst Book
	for b.Loop() {
		bs, _ := json.Marshal(book)
		json.Unmarshal(bs, &inst)
	}
}

// 注意看sonic的官方文档，跟go的哪些版本适配。https://github.com/bytedance/sonic
func BenchmarkJsonSonic(b *testing.B) {
	var inst Book
	for b.Loop() {
		bs, _ := sonic.Marshal(book)
		sonic.Unmarshal(bs, &inst)
	}
}

// json-iterator
func BenchmarkJsoniter(b *testing.B) {
	var inst Book
	for b.Loop() {
		bs, _ := jsoniter.Marshal(book)
		jsoniter.Unmarshal(bs, &inst)
	}
}

func BenchmarkJsonGo(b *testing.B) {
	var inst Book
	for b.Loop() {
		bs, _ := gojson.Marshal(book)
		gojson.Unmarshal(bs, &inst)
	}
}

// go test ./io -bench=^BenchmarkJson -run=^$ --ldflags="-checklinkname=0"
/**
goos: windows
goarch: amd64
pkg: dqq/go/frame/io
cpu: 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz
BenchmarkJsonStd-8        266196              4653 ns/op
BenchmarkJsonSonic-8      891794              1332 ns/op
BenchmarkJsoniter-8       490380              2091 ns/op
BenchmarkJsonGo-8         787412              1525 ns/op
*/
