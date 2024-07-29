> 为了简单（保持跟视频中一致），直接使用Rust重写了。

在go中执行fasm原生汇编指令失败[runbin]，执行go自己生成的汇编可以[test]

GPT: 

```sh
以下是一些可能导致在Go中执行失败的原因：

内存安全性：
Go语言对内存访问有严格的限制，特别是在涉及指针操作和内存映射时。Go的内存模型设计旨在提供更高的安全性和避免常见的内存错误。因此，直接操作内存映射可能会受到Go的限制。

系统调用：
Go语言的标准库提供了对系统调用的封装，但是在某些情况下，直接调用系统调用可能会受到限制。在您的代码中，您使用了syscall.Mmap和syscall.Munmap来进行内存映射和解除映射，这可能会受到Go运行时的限制。

指针操作：
Go语言对指针操作有一定的限制，尤其是在涉及不安全操作时。在您的代码中，您使用了unsafe.Pointer和指针解引用操作，这可能会受到Go的安全性检查。
虽然Go语言提供了unsafe包来处理不安全操作，但是在处理系统级操作和直接内存访问时，Go的安全性机制可能会限制某些操作。相比之下，Rust更加灵活，允许更细粒度的控制和更多的系统级操作，同时通过所有权系统和借用检查器提供了更强大的内存安全性保证。

因此，尽管您的代码在Rust中可以成功执行，但在Go中可能会受到一些限制。如果您需要在Go中执行类似的操作，可能需要考虑更多的安全性和Go语言的限制，或者尝试使用更适合Go的方式来实现相同的功能。
```