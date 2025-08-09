package main

import (
	"crypto/sha256"
	"fmt"
	"net"
	"strconv"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func pop() {
	fmt.Println(PopCount(123123))
	fmt.Println(strconv.FormatInt(int64(26), 2))
	fmt.Println(strconv.FormatInt(int64(25), 3))
	fmt.Println(strconv.FormatInt(int64(984), 6))
	fmt.Println(fmt.Sprintf("x=%b", 982))

	var a [3]int
	fmt.Println(a)

	var q = [...]int{1, 2, 3}
	fmt.Println(q)

	symbol := [...]string{USD: "$", EUR: "$", GBP: "$", RMB: "$"}
	fmt.Println(symbol[USD])

	r := [...]int{99: -1}
	fmt.Println(r[99])

	c := sha256.Sum256([]byte("x"))
	fmt.Println(c)
}

func IsUp(v net.Flags) bool {
	return v&net.FlagUp == net.FlagUp
}
