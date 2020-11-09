package main

func main() {
	f()
}

type T bool

func f() {
	var t T = false
	panic(t)
}
