package main

import "github.com/torbensky/gofilelogger/example"

func main() {
	example.Run()
	/*
		Output:
			debug:File1
			info:File1
			warning:File1
			error:File1
			warning:File2
			error:File2
	*/
}
