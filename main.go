//go:build js && wasm

package main

import (
// "syscall/js"
)

func main() {
	c := make(chan struct{}, 0)
	println("Go WebAssembly Initialized")
	<-c
}
