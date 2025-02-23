package func_test

import (
	"fmt"
	"testing"
	"time"
)

func TestFuncInner(t *testing.T) {
	t.Log("TestFuncInner")
	timeCost(slowFunc)(100)
}

func timeCost(f func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := f(n)
		fmt.Println("time cost:", time.Since(start))
		return ret
	}
}
func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestDefer(t *testing.T) {
	t.Log("TestDefer")
	defer func() {
		fmt.Println("Defer")
	}()
	fmt.Println("TestDefer")
	panic("error")
}
