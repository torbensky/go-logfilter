package main

import "github.com/torbensky/gofilelogger/example"

func main() {
	example.Run()
	/*
		Output:
			Setting file1.go to log at debug
			Setting file2.go to log at warning
			Setting github.com/torbensky/gofilelogger to log at panic
			debug:File1
			info:File1
			warning:File1
			error:File1
			warning:File2
			error:File2
	*/
}
