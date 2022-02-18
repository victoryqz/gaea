package main

import "sync"

var m sync.Map

func main() {

	m.Load("x")

}
