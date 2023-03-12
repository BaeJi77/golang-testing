package main

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("before Run()")
	code := m.Run()
	fmt.Println("After Run()")

	os.Exit(code)
}

func TestFoo(t *testing.T) {
	fmt.Println("TestFoo")
}

func TestBar(t *testing.T) {
	fmt.Println("TestBar")
}

func TestlowerCase(t *testing.T) {
	fmt.Println("TestlowerCase")
}

func TestT(t *testing.T) {
	t.Run("sub-test-1", func(t *testing.T) {
		fmt.Println("sub-test-1")
	})

	t.Error()
	t.Errorf("error string format")

	t.Deadline()

	t.Fail()
	t.FailNow()
	t.Failed()

	t.Skip()
	t.Skipf("skip string format")
	t.SkipNow()
	t.Skipped()

	t.Helper()
}

func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < 1000000; i++ {
		rand.Int()
	}
}
