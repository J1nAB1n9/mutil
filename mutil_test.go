package muitl

import (
	"math/rand"
	"testing"
	"time"
)

func RandomMaxSize() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int() % 30
}

func RandPassWord(maxSize int) string {
	passWord := make([]byte,maxSize)

	if maxSize != 0 {
		rand.Seed(time.Now().UnixNano())
		size := rand.Int() % maxSize
		for i := 0 ; i < size;i++ {
			passWord = append(passWord,byte(rand.Int() % 100 + 20))
		}
	}

	return string(passWord)
}

// go test -v mutil_test.go net.go
func TestSaveEquals(t *testing.T) {
	a:= "PassWord"
	b:= "PassWord"

	if ok := SafeEquals(a,b,20);ok == true {
		println("Success")
	} else {
		println("Fail")
	}
}

// go test -bench=SafeEquals
func BenchmarkSafeEquals(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxSize := RandomMaxSize()
		pass := RandPassWord(maxSize)
		input := RandPassWord(maxSize)
		if ok := SafeEquals(input,pass,maxSize);ok == true {
			// println("Success",pass)
		} else {
			// println("Fail",pass)
		}
	}
}