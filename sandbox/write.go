package main

import (
	"syscall"
)

func main() {
	syscall.Write(syscall.Stderr, []byte("test]n"))
}
