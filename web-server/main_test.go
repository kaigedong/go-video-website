package main

import (
	"fmt"
	"testing"
)

func testPrint(t *testing.T) {
	res := Print1to20()
	fmt.Println("hey")
	if res != 210 {
		t.Errorf("Wrong result of Print1to20")
	}
}

func testPrint2(t *testing.T) {
	res := Print1to20()
	res++
	if res != 211 {
		t.Errorf("Test Print2 failed")
	}
}

func TestAll(t *testing.T) {
	t.Run("TestPrint", testPrint)
	t.Run("TestPrint2", testPrint2)
}

func TestMain(m *testing.M) {
	fmt.Println("Testing begins...")
	m.Run() // 如果注销这一行，则除了这个TestMain之外都不会执行
}

// func TestPrint(t *testing.T) {
//	t.Run("a1", func(t *testing.T) {fmt.Println("a1")})
//	t.Run("a2", func(t *testing.T) {fmt.Println("a2")})
//	t.Run("a3", func(t *testing.T) {fmt.Println("a3")})
//}

func BenchmarkAll(b *testing.B) {
	for n := 0; n < b.N; n++ { // b.N是默认的次数，跑的时候回测试到稳定的次数以及消耗的稳定的时间
		Print1to20()
	}
}
