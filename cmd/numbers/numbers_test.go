package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

var in io.Reader = os.Stdin
var out io.Writer = os.Stdout

func Main() {

	var input string
	fmt.Fscanf(in,"%s\n", &input)

	firsNum, secondNum := ReadNumbers(input)
	sumList := AddNumbers(firsNum, secondNum)
	fmt.Fprint(out, sumList)

}

func TestFirst(t *testing.T) {
	buffer := new(bytes.Buffer)
	out = buffer
	in = buffer

	fmt.Fprintln(out, "l1=[2,4,3],l2=[5,6,4]")
	Main()

	result := out.(*bytes.Buffer).String()
	if result != "[7,0,8]" {
		t.Error("AddNumbers result is not correctno: ", result)
	}
}

func TestSecond(t *testing.T) {
	buffer := new(bytes.Buffer)
	out = buffer
	in = buffer

	fmt.Fprintln(out, "l1=[1,2],l2=[1]")
	Main()

	result := out.(*bytes.Buffer).String()
	if result != "[2,2]" {
		t.Error("AddNumbers result is not correctno: ", result)
	}
}

func TestThird(t *testing.T) {
	buffer := new(bytes.Buffer)
	out = buffer
	in = buffer

	fmt.Fprintln(out, "l1=[9,9,9,9,9,9,9],l2=[9,9,9,9]")
	Main()

	result := out.(*bytes.Buffer).String()
	if result != "[8,9,9,9,0,0,0,1]" {
		t.Error("AddNumbers result is not correctno: ", result)
	}
}