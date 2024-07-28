package main

import (
	"fmt"
	"io/ioutil"
	"syscall"
	"unsafe"
)

func main() {
	// 读取二进制文件
	code, err := ioutil.ReadFile("simple.bin")
	if err != nil {
		panic(err)
	}

	// 获取二进制代码的长度
	codeSize := len(code)

	// 使用mmap分配内存
	mmap, err := syscall.Mmap(-1, 0, codeSize, syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC, syscall.MAP_ANON|syscall.MAP_PRIVATE)
	if err != nil {
		panic(err)
	}

	// 将二进制代码复制到mmap内存中
	copy(mmap, code)

	// 将mmap指针转换为函数指针
	funcPtr := unsafe.Pointer(&mmap[0])
	execFunc := *(*func() int)(unsafe.Pointer(&funcPtr))

	// 执行函数并打印结果
	result := execFunc()
	fmt.Printf("Result: %d\n", result)

	// 解除mmap映射
	syscall.Munmap(mmap)
}
