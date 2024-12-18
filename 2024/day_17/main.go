package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type INS int

const (
	ADV INS = iota
	BXL
	BST
	JNZ
	BXC
	OUT
	BDV
	CDV
)

type vm struct {
	IP     int
	A      int
	B      int
	C      int
	halt   bool
	stack  []int
	output []int
}

func main() {
	start1 := time.Now().UnixMicro()
	fmt.Println(SolvePart1("example"))
	end1 := time.Now().UnixMicro()
	fmt.Println("solved in: ", end1-start1, " microseconds")

	start2 := time.Now().UnixMicro()
	fmt.Println(SolvePart2("example"))
	end2 := time.Now().UnixMicro()
	fmt.Println("solved in: ", end2-start2, " microseconds")
}

func SolvePart1(inputPath string) string {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var ans string
	vm := &vm{
		IP:     0,
		A:      0,
		B:      0,
		C:      0,
		halt:   false,
		stack:  []int{},
		output: []int{},
	}

	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "Register A:") {
			vm.A, _ = strconv.Atoi(strings.Trim(strings.Split(text, ":")[1], " "))
		}
		if strings.Contains(text, "Register B:") {
			vm.B, _ = strconv.Atoi(strings.Trim(strings.Split(text, ":")[1], " "))
		}
		if strings.Contains(text, "Register C:") {
			vm.C, _ = strconv.Atoi(strings.Trim(strings.Split(text, ":")[1], " "))
		}
		if strings.Contains(text, "Program:") {
			programStr := strings.Split(strings.Trim(strings.Split(text, ":")[1], " "), ",")
			for _, v := range programStr {
				p, _ := strconv.Atoi(v)
				vm.stack = append(vm.stack, p)
			}
		}
	}
	out := vm.run()
	for i := range out {
		ans = fmt.Sprintf("%s%d,", ans, out[i])
	}
	ans = ans[:len(ans)-1]
	return ans
}

func SolvePart2(inputPath string) int {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	ans := 0
	vm := newVm()
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "Register A:") {
			vm.A, _ = strconv.Atoi(strings.Trim(strings.Split(text, ":")[1], " "))
		}
		if strings.Contains(text, "Register B:") {
			vm.B, _ = strconv.Atoi(strings.Trim(strings.Split(text, ":")[1], " "))
		}
		if strings.Contains(text, "Register C:") {
			vm.C, _ = strconv.Atoi(strings.Trim(strings.Split(text, ":")[1], " "))
		}
		if strings.Contains(text, "Program:") {
			programStr := strings.Split(strings.Trim(strings.Split(text, ":")[1], " "), ",")
			for _, v := range programStr {
				p, _ := strconv.Atoi(v)
				vm.stack = append(vm.stack, p)
			}
		}
	}

	qmap := make(map[int][]int)
	for i := 0; i < 8; i++ {
		b := reverse(i)
		if b == vm.stack[len(vm.stack)-1] {
			qmap[len(vm.stack)-1] = append(qmap[len(vm.stack)-1], i)
		}
	}

	for i := len(vm.stack) - 2; i >= 0; i-- {
		for _, v := range qmap[i+1] {
			for j := 0; j < 8; j++ {
				check := (v << 3) + j
				b := reverse(check)
				if b == vm.stack[i] {
					qmap[i] = append(qmap[i], check)
				}
			}
		}
	}

	lowest := len(vm.stack) - 1
	for k, _ := range qmap {
		if k < lowest {
			lowest = k
		}
	}

	ans = qmap[lowest][0]
	for _, v := range qmap[lowest] {
		if v < ans {
			ans = v
		}
	}
	return ans
}

func reverse(a int) int {
	b := 0
	c := 0
	b = a % 8
	b = b ^ 1
	c = a / pow(2, b)
	b = b ^ 5
	b = (b ^ c) % 8
	a = a / pow(2, 3)
	return b
}

func newVm() *vm {
	vm := &vm{
		IP:     0,
		A:      0,
		B:      0,
		C:      0,
		halt:   false,
		stack:  []int{},
		output: []int{},
	}
	return vm
}

func (vm *vm) run() []int {
	for !vm.halt {
		switch vm.instruction() {
		case ADV:
			vm.adv(vm.operand())
			break
		case BXL:
			vm.bxl(vm.operand())
			break
		case BST:
			vm.bst(vm.operand())
			break
		case JNZ:
			vm.jnz(vm.operand())
			break
		case BXC:
			vm.bxc(vm.operand())
			break
		case OUT:
			vm.out(vm.operand())
			break
		case BDV:
			vm.bdv(vm.operand())
			break
		case CDV:
			vm.cdv(vm.operand())
			break
		}
	}
	return vm.output
}

func (vm *vm) instruction() INS {
	return INS(vm.stack[vm.IP])
}

func (vm *vm) operand() int {
	return vm.stack[vm.IP+1]
}

func (vm *vm) getOperandCombo(i int) int {
	if i <= 3 {
		return i
	}
	if i == 4 {
		return vm.A
	}
	if i == 5 {
		return vm.B
	}
	if i == 6 {
		return vm.C
	}
	return -1
}

func (vm *vm) IncIP() {
	if vm.IP+2 < len(vm.stack) {
		vm.IP += 2
	} else {
		vm.halt = true
	}
}

func (vm *vm) adv(op int) {
	cop := vm.getOperandCombo(op)
	denom := pow(2, cop)
	vm.A = vm.A / denom
	vm.IncIP()
}

func (vm *vm) bxl(op int) {
	vm.B = vm.B ^ op
	vm.IncIP()
}

func (vm *vm) bst(op int) {
	o := vm.getOperandCombo(op)
	vm.B = o % 8
	vm.IncIP()
}

func (vm *vm) jnz(op int) {
	if vm.A != 0 {
		vm.IP = op
	} else {
		vm.IncIP()
	}
}

func (vm *vm) bxc(op int) {
	vm.B = vm.B ^ vm.C
	vm.IncIP()
}

func (vm *vm) out(op int) {
	out := vm.getOperandCombo(op) % 8
	vm.output = append(vm.output, out)
	vm.IncIP()
}

func (vm *vm) bdv(op int) {
	cop := vm.getOperandCombo(op)
	denom := pow(2, cop)
	vm.B = vm.A / denom
	vm.IncIP()
}

func (vm *vm) cdv(op int) {
	cop := vm.getOperandCombo(op)
	denom := pow(2, cop)
	vm.C = vm.A / denom
	vm.IncIP()
}

func pow(n, m int) int {
	if m == 0 {
		return 1
	}

	if m == 1 {
		return n
	}

	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}