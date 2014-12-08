package main

import "fmt"

type Programmer interface { Speak() string }

type CppProgrammer struct {}

type JavaProgrammer struct {}

func (c CppProgrammer) Speak() string { return "damn templates!" }

func (j JavaProgrammer) Speak() string { return "damn factories!" }

func speak(p Programmer) { fmt.Println(p.Speak()) }

func main() {
	jp := JavaProgrammer{}
	speak(jp)
	cpp := CppProgrammer{}
	speak(cpp)
}
