package main

import (
	"fmt"
	"testing"
)

func Test_getConfig(t *testing.T) {
	c := getConfig()
	fmt.Println(c.Ratio)
	if c.Ratio != 10.0 {
		t.Error("ratio not 10.0")
	}
}
func Test_getNowAndOld(t *testing.T) {
	now, _ := getNowAndOld()
	if !(now.Bpi.Usd.RateFloat > 1) {
		t.Error("getNowAndOld get data error")
	}
}

func Test_getOld(t *testing.T) {
	p := getOld()
	if p.Bpi.Usd.RateFloat < 1.0 {
		t.Error("getOldPrice not good")
	}
}
func Test_exceed(t *testing.T) {
	if !exceed(200, 300, 10) {
		t.Error("exceed 100?")
	}
}
func Test_makeContent(t *testing.T) {
	a := getOld()
	b := getOld()
	a.Bpi.Usd.Rate = "111.1"
	// a.Bpi.Usd.RateFloat=111.1
	a.Time.Updated = "a"
	b.Bpi.Usd.Rate = "2222.2"
	// b.Bpi.Usd.RateFloat=2222.2
	b.Time.Updated = "b"
	fmt.Println(makeContent(a, b))
	if makeContent(a, b) != "价格从 $111.1 到 $2222.2 <br>(a ~ b)" {
		t.Error("content wrong")
	}
}
