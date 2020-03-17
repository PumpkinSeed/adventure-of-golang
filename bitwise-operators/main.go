package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// https://medium.com/learning-the-go-programming-language/bit-hacking-with-go-e0acee258827

	// ----------------
	// & operator (AND)
	//and()

	// ---------------
	// | operator (OR)
	//or()

	// ---------------
	// ^ operator (XOR)
	//xor()

	// ---------------
	// ^ as bitwise complement (NOT)
	//not()

	// ---------------
	// &^ operator (AND NOT)
	//andnot()

	// ---------------
	// << >> operators (SHIFT)
	shift()
}

func shift() {
	var a int8 = 3
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", a<<1)
	fmt.Printf("%08b\n", a<<2)
	fmt.Printf("%08b\n", a<<3)
	fmt.Println("--------")

	var b uint8 = 120
	fmt.Printf("%08b\n", b)
	fmt.Printf("%08b\n", b>>1)
	fmt.Printf("%08b\n", b>>2)
}

func andnot() {
	var x uint8 = 0xAC // x = 10101100
	var y uint8 = 0xF0 // y = 11110000
	var r = x ^ y      // r = 01011100
	fmt.Printf("%x\n", r)
	fmt.Printf("%b\n", r)
}

func not() {
	var a byte = 0x0F
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", ^a)
}

func xor() {
	var x uint8 = 0xAC // x = 10101100
	var y uint8 = 0xF0 // y = 11110000
	var r = x ^ y      // r = 01011100
	fmt.Printf("%x\n", r)
	fmt.Printf("%b\n", r)

	// toggle bits from one value to another

	// check the sign is equal
	a, b := 12, 25
	fmt.Println("a and b have same sign?", (a^b) >= 0)

	a, b = -12, 25
	fmt.Println("a and b have same sign?", (a^b) >= 0)
	fmt.Printf("%b\n", a)   // -01100
	fmt.Printf("%b\n", b)   // +11001
	fmt.Printf("%b\n", a^b) // -10011 -> so the signs are different so it's smaller than 0
}

func or() {
	var x uint8 = 0xAC // x = 10101100
	var y uint8 = 0xF0 // y = 11110000
	var r = x | y      // r = 11111100
	fmt.Printf("%x\n", r)
	fmt.Printf("%b\n", r)

	// nice side effect: selectively set bits

	// using bitmasks
	const (
		Card          = 1 << iota // 1
		ACH                       // 10
		Terminal                  // 100
		Cash                      // 1000
		CustomerToken             // 10000
		Token                     // 100000
	)

	var supportedMethod = 0
	supportedMethod |= ACH
	supportedMethod |= Cash
	fmt.Printf("%b\n", supportedMethod) // 1010
	supportedMethod |= Token
	fmt.Printf("%b\n", supportedMethod) // 101010

	if supportedMethod&Cash == Cash {
		fmt.Println("Support Cash")
	}
	if supportedMethod&Terminal != Terminal {
		fmt.Println("Not support Terminal")
	}
}

func and() {
	var x uint8 = 0xAC // x = 10101100
	var y uint8 = 0xF0 // y = 11110000
	var r = x & y      // r = 10100000
	fmt.Printf("%x\n", r)
	fmt.Printf("%b\n", r)
	// can be used to clear LSB (least significant bit)
	// LSB is the bit position in a binary integer determining whether the number is even or odd

	// nice side effect: selectively clearing bits

	// short-hand
	var x2 uint8 = 0xAC
	x2 &= 0xF0

	// check whether the number is even or odd
	//num := rand.Int()
	//fmt.Printf("%d - %b => %t\n", num, num, num&1 == 1)
	//
	//n := int8(-127)
	//signCheck(n)
	//n = 127
	//signCheck(n)
}

func signCheck(n int8) {
	if n&1<<7 == 1 {
		fmt.Printf("%d is odd - %b\n", n, n)
	} else {
		fmt.Printf("%d is even - %b\n", n, n)
	}
}
