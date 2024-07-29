package main

import (
	"os"
	"syscall"
	"unsafe"
)

type asmFuncType func(unsafe.Pointer, uint64)

func main() {
	code, err := os.ReadFile("hello.bin")
	if err != nil {
		panic(err)
	}

	codeSize := len(code)

	mmap, err := syscall.Mmap(-1, 0, codeSize, syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC, syscall.MAP_ANON|syscall.MAP_PRIVATE)
	if err != nil {
		panic(err)
	}
	defer syscall.Munmap(mmap)

	copy(mmap, code)

	str := "hello, world!\n"
	length := len(str)
	cstr := []byte(str)

	funcPtr := unsafe.Pointer(&mmap[0])
	execFunc := *(*asmFuncType)(funcPtr)

	execFunc(unsafe.Pointer(&cstr[0]), uint64(length))
}
