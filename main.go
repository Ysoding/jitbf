package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"
	"unsafe"

	"github.com/Ysoding/jitbf/jitbf"
	"github.com/Ysoding/jitbf/lexer"
)

type Code func(memory []byte)

func jitCompile(ops *jitbf.Ops) (Code, error) {
	var sb strings.Builder

loop:
	for i := 0; i < ops.Count(); i++ {
		op := ops.Items[i]
		switch op.Kind {
		case jitbf.OpInc:
			sb.WriteString("\xFE\x07") // inc byte [rdi]
			break loop
		case jitbf.OpDec:
			panic("not implemented")
		case jitbf.OpLeft:
			panic("not implemented")
		case jitbf.OpRight:
			panic("not implemented")
		case jitbf.OpInput:
			panic("not implemented")
		case jitbf.OpOutput:
			panic("not implemented")
		case jitbf.OpJumpIfZero:
			panic("not implemented")
		case jitbf.OpJumpIfNonZero:
			panic("not implemented")
		}
	}

	sb.WriteString("\xC3") // ret

	data := []byte(sb.String())
	addr, err := syscall.Mmap(-1, 0, len(data), syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC, syscall.MAP_PRIVATE|syscall.MAP_ANON)
	if err != nil {
		return nil, err
	}

	copy(addr, data)
	// for i := 0; i < len(data); i++ {
	// 	*(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&addr)) + uintptr(i))) = data[i]
	// }

	code := *(*Code)(unsafe.Pointer(&addr))

	// 延迟解除映射，避免非法访问

	return code, nil

}

func interpreter(ops *jitbf.Ops) {

	memory := jitbf.NewMemory()
	memory.Add(0)
	head := 0
	ip := 0

	for ip < ops.Count() {
		op := ops.Items[ip]
		switch op.Kind {
		case jitbf.OpInc:
			memory.Items[head] += byte(op.Operand)
			ip += 1
		case jitbf.OpDec:
			memory.Items[head] -= byte(op.Operand)
			ip += 1
		case jitbf.OpLeft:
			if head < op.Operand {
				log.Fatalln("RUNTIME ERROR: Memory underflow")
			}
			head -= op.Operand
			ip += 1
		case jitbf.OpRight:
			head += op.Operand
			for head >= memory.Count() {
				memory.Add(0)
			}
			ip += 1
		case jitbf.OpInput:
			panic("not implemented")
		case jitbf.OpOutput:
			for i := 0; i < op.Operand; i++ {
				fmt.Printf("%c", memory.Items[head])
			}
			ip += 1
		case jitbf.OpJumpIfZero:
			if memory.Items[head] == 0 {
				ip = op.Operand
			} else {
				ip += 1
			}
		case jitbf.OpJumpIfNonZero:
			if memory.Items[head] != 0 {
				ip = op.Operand
			} else {
				ip += 1
			}
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <input.bf>\n", os.Args[0])
		fmt.Println("No input is provided")
		os.Exit(1)
	}

	filePath := os.Args[1]

	bs, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalln(err)
	}

	ops := jitbf.NewOps()
	stack := jitbf.NewAddrStack()
	l := lexer.New(string(bs))

	for curToken := l.Next(); curToken != jitbf.OpEOF; {
		switch curToken {
		case '+', '-', '>', '<', ',', '.':
			count := 1

			nextToken := l.Next()
			for ; nextToken == curToken; nextToken = l.Next() {
				count += 1
			}

			ops.Add(jitbf.Op{
				Kind:    curToken,
				Operand: count,
			})

			curToken = nextToken
		case '[':
			addr := ops.Count()
			ops.Add(jitbf.Op{
				Kind:    curToken,
				Operand: 0,
			})

			stack.Push(addr)

			curToken = l.Next()
		case ']':
			if stack.IsEmpty() {
				log.Fatalf("%s [%d]: ERROR: Unbalanced loop\n", filePath, l.Pos())
			}

			addr := stack.Pop()

			ops.Add(jitbf.Op{
				Kind:    curToken,
				Operand: addr + 1,
			})
			ops.Items[addr].Operand = ops.Count()

			curToken = l.Next()
		default:
			curToken = l.Next()
		}
	}

	code, err := jitCompile(ops)
	if err != nil {
		log.Fatalln(err)
	}

	memory := make([]byte, 10*1000*1000)
	code(memory)

	addr := unsafe.Pointer(&code)
	syscall.Munmap(*(*[]byte)(addr))
}
