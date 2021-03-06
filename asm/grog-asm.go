package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/martinstraus/grog/asm/parser"
	"github.com/martinstraus/grog/vm"
)

var arithmeticOperators = map[string]byte{
	"+": vm.ADD,
	"-": vm.SUBTRACT,
	"*": vm.MULTIPLY,
	"/": vm.DIVIDE,
}

var booleanOperators = map[string]byte{
	"AND": vm.AND,
	"OR":  vm.OR,
	"XOR": vm.XOR,
}

type listener struct {
	*parser.BaseGrogListener
	Output   *bufio.Writer
	Contants map[string]byte
}

func (l *listener) ExitLoad(c *parser.LoadContext) {
	if c.Value != nil {
		l.Output.WriteByte(vm.LOAD_BYTE)
		l.Output.WriteByte(valueByte(c.Value.GetText()))
	} else if c.Address != nil {
		l.Output.WriteByte(vm.LOAD_ADDRESS)
		l.Output.Write(absoluteAddressBytes(c.Address.GetText()))
	} else if c.Offset != nil {
		l.Output.WriteByte(vm.LOAD_OFFSET)
		l.Output.Write(offsetAddressBytes(c.Offset.GetText()))
	} else if c.Pointer != nil {
		l.Output.WriteByte(vm.LOAD_POINTER)
		l.Output.Write(pointerAddressBytes(c.Pointer.GetText()))
	} else if c.Constant != nil {
		l.Output.WriteByte(vm.LOAD_BYTE)
		value := l.Contants[c.Constant.GetText()]
		l.Output.WriteByte(value)

	}
	l.Output.WriteByte(registerByte(c.Register.GetText()))
}

func (l *listener) ExitStore(c *parser.StoreContext) {
	operation := byte(0)
	destination := []byte(nil)
	if c.Address != nil {
		operation = vm.STORE_BYTE_ADDRESS
		destination = absoluteAddressBytes(c.Address.GetText())
	} else if c.Offset != nil {
		operation = vm.STORE_BYTE_OFFSET
		destination = offsetAddressBytes(c.Offset.GetText())
	} else if c.Pointer != nil {
		operation = vm.STORE_BYTE_POINTER
		destination = pointerAddressBytes(c.Pointer.GetText())
	}
	value := byte(0)
	if c.Value != nil {
		value = valueByte(c.Value.GetText())
	} else if c.Register != nil {
		// This is possible because store operations are ordered
		operation = operation + 3
		value = registerByte(c.Register.GetText())
	}
	l.Output.WriteByte(operation)
	l.Output.Write(destination)
	l.Output.WriteByte(value)
}

func (l *listener) ExitCopyRegister(c *parser.CopyRegisterContext) {
	destination := registerByte(c.DestinationRegister.GetText())
	source := registerByte(c.SourceRegister.GetText())
	l.Output.WriteByte(vm.COPY_REGISTER)
	l.Output.WriteByte(destination)
	l.Output.WriteByte(source)
}

func (l *listener) ExitCopyAbsoluteToAbsolute(c *parser.CopyAbsoluteToAbsoluteContext) {
	l.Output.WriteByte(vm.COPY_ADDRESS_ADDRESS)
	l.Output.Write(addressBytes(c.Destination.GetText()))
	l.Output.Write(addressBytes(c.Source.GetText()))
}

func (l *listener) ExitCopyOffsetToAbsolute(c *parser.CopyOffsetToAbsoluteContext) {
	l.Output.WriteByte(vm.COPY_ADDRESS_OFFSET)
	l.Output.Write(addressBytes(c.Destination.GetText()))
	l.Output.Write(addressBytes(c.Source.GetText()))
}

func (l *listener) ExitCopyPointerToAbsolute(c *parser.CopyPointerToAbsoluteContext) {
	l.Output.WriteByte(vm.COPY_POINTER_ADDRESS)
	l.Output.Write(addressBytes(c.Destination.GetText()))
	l.Output.Write(addressBytes(c.Source.GetText()))
}

func (l *listener) ExitCopyAbsoluteToOffset(c *parser.CopyAbsoluteToOffsetContext) {
	l.Output.WriteByte(vm.COPY_OFFSET_ADDRESS)
	l.Output.Write(addressBytes(c.Destination.GetText()))
	l.Output.Write(addressBytes(c.Source.GetText()))
}

func (l *listener) ExitCopyOffsetToOffset(c *parser.CopyOffsetToOffsetContext) {
	l.Output.WriteByte(vm.COPY_OFFSET_OFFSET)
	l.Output.Write(addressBytes(c.Destination.GetText()))
	l.Output.Write(addressBytes(c.Source.GetText()))
}

func (l *listener) ExitCopyPointerToOffset(c *parser.CopyPointerToOffsetContext) {
	l.Output.WriteByte(vm.COPY_OFFSET_POINTER)
	l.Output.Write(addressBytes(c.Destination.GetText()))
	l.Output.Write(addressBytes(c.Source.GetText()))
}

func (l *listener) ExitCopyAbsoluteToPointer(c *parser.CopyAbsoluteToPointerContext) {
	l.Output.WriteByte(vm.COPY_OFFSET_ADDRESS)
	l.Output.Write(addressBytes(c.Destination.GetText()))
	l.Output.Write(addressBytes(c.Source.GetText()))
}

func (l *listener) ExitCopyOffsetToPointer(c *parser.CopyOffsetToPointerContext) {
	l.Output.WriteByte(vm.COPY_POINTER_OFFSET)
	l.Output.Write(addressBytes(c.Destination.GetText()))
	l.Output.Write(addressBytes(c.Source.GetText()))
}

func (l *listener) ExitCopyPointerToPointer(c *parser.CopyPointerToPointerContext) {
	l.Output.WriteByte(vm.COPY_POINTER_POINTER)
	l.Output.Write(addressBytes(c.Destination.GetText()))
	l.Output.Write(addressBytes(c.Source.GetText()))
}

func (l *listener) ExitIncrement(c *parser.IncrementContext) {
	l.Output.WriteByte(vm.INCREMENT)
	l.Output.WriteByte(registerByte(c.Register.GetText()))
}

func (l *listener) ExitDecrement(c *parser.DecrementContext) {
	l.Output.WriteByte(vm.DECREMENT)
	l.Output.WriteByte(registerByte(c.Register.GetText()))
}

func (l *listener) ExitArithmeticOperation(c *parser.ArithmeticOperationContext) {
	l.Output.WriteByte(arithmeticOperators[c.Operator.GetText()])
	l.Output.WriteByte(registerByte(c.Destination.GetText()))
	l.Output.WriteByte(registerByte(c.Source.GetText()))
}

func (l *listener) ExitBinaryBooleanOperation(c *parser.BinaryBooleanOperationContext) {
	l.Output.WriteByte(booleanOperators[c.Operator.GetText()])
	l.Output.WriteByte(registerByte(c.Destination.GetText()))
	l.Output.WriteByte(registerByte(c.Source.GetText()))
}

func (l *listener) ExitUnaryBooleanOperation(c *parser.UnaryBooleanOperationContext) {
	l.Output.WriteByte(vm.NOT)
	l.Output.WriteByte(registerByte(c.Destination.GetText()))
}

func (l *listener) ExitJump(c *parser.JumpContext) {
	// Since all jump operations codes are sequential, we can calculate the opcodes
	operation := vm.JUMP_ADDRESS
	switch c.Operator.GetText() {
	case "je":
		operation = vm.JUMP_EQUAL_ADDRESS
	case "jne":
		operation = vm.JUMP_NOT_EQUAL_ADDRESS
	}
	var destination []byte = nil
	if c.Address != nil {
		destination = absoluteAddressBytes(c.Address.GetText())
	} else if c.Offset != nil {
		operation = operation + 1
		destination = offsetAddressBytes(c.Offset.GetText())
	} else if c.Pointer != nil {
		operation = operation + 2
		destination = pointerAddressBytes(c.Pointer.GetText())
	}
	l.Output.WriteByte(operation)
	l.Output.Write(destination)
}

func (l *listener) ExitIo(c *parser.IoContext) {
	register := registerByte(c.Destination.GetText())
	device := deviceByte(c.Source.GetText())
	operation := byte(0)
	switch c.Operation.GetText() {
	case "input":
		operation = vm.INPUT
	case "output":
		operation = vm.OUTPUT
	}
	l.Output.Write([]byte{operation, device, register})
}

func (l *listener) ExitStop(c *parser.StopContext) {
	l.Output.WriteByte(vm.STOP)
}

func (l *listener) ExitWait(c *parser.WaitContext) {
	l.Output.WriteByte(vm.WAIT)
}

func (l *listener) ExitConstant(c *parser.ConstantContext) {
	name := c.GetName().GetText()
	_, exists := l.Contants[name]
	if exists {
		panic("Constant " + name + " already defined.")
	}
	l.Contants[name] = valueByte(c.GetByteValue().GetText())
}

func (l *listener) ExitCompare(c *parser.CompareContext) {
	register := registerByte(c.DestinationRegister.GetText())
	if c.SourceRegister != nil {
		l.Output.WriteByte(vm.COMPARE_REGISTER_REGISTER)
		l.Output.WriteByte(register)
		l.Output.WriteByte(registerByte(c.SourceRegister.GetText()))
	} else if c.SourceValue != nil {
		l.Output.WriteByte(vm.COMPARE_REGISTER_VALUE)
		l.Output.WriteByte(register)
		l.Output.WriteByte(valueByte(c.SourceValue.GetText()))
	} else if c.SourceMemoryAbsolute != nil {
		l.Output.WriteByte(vm.COMPARE_REGISTER_ABSOLUTE)
		l.Output.WriteByte(register)
		l.Output.Write(addressBytes(c.SourceMemoryAbsolute.GetText()))
	} else if c.SourceMemoryOffset != nil {
		l.Output.WriteByte(vm.COMPARE_REGISTER_OFFSET)
		l.Output.WriteByte(register)
		l.Output.Write(offsetAddressBytes(c.SourceMemoryOffset.GetText()))
	} else if c.SourceMemoryPointer != nil {
		l.Output.WriteByte(vm.COMPARE_REGISTER_POINTER)
		l.Output.WriteByte(register)
		l.Output.Write(pointerAddressBytes(c.SourceMemoryPointer.GetText()))
	}
}

func absoluteAddressBytes(value string) []byte {
	return prefixedAddressBytes(value, "@")
}

func offsetAddressBytes(value string) []byte {
	return prefixedAddressBytes(value, "#")
}

func pointerAddressBytes(value string) []byte {
	return prefixedAddressBytes(value, "*")
}

func prefixedAddressBytes(value string, prefix string) []byte {
	return addressBytes(strings.Split(value, prefix)[1])
}

func registerByte(registerName string) byte {
	register := strings.TrimLeft(registerName, "R")
	registerBytes, _ := hex.DecodeString("0" + register) // We need to prepend the "0" to complete the byte.
	return registerBytes[0]
}

func deviceByte(deviceName string) byte {
	value := strings.TrimLeft(deviceName, "D")
	bytes, _ := hex.DecodeString(value)
	return bytes[0]
}

func valueByte(value string) byte {
	valueBytes, _ := hex.DecodeString(value)
	return valueBytes[0]
}

func addressBytes(address string) []byte {
	addressBytes, _ := hex.DecodeString(address)
	return reversed(addressBytes) // we reverse it because we use big-endian format
}

func reversed(array []byte) []byte {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
	return array
}

func main() {
	fileName := os.Args[1]
	fmt.Printf("Reading file %s\n", fileName)
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Read %d bytes\n", len(data))
	input := antlr.NewInputStream(string(data))
	lexer := parser.NewGrogLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := parser.NewGrogParser(stream)
	outputFile, err := os.Create(fileName + ".grog")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
	listener := listener{
		Output:   bufio.NewWriter(outputFile),
		Contants: make(map[string]byte),
	}
	antlr.ParseTreeWalkerDefault.Walk(&listener, parser.Program())
	listener.Output.Flush()
}
