package main

import "runtime"

func init() {
	runtime.GOMAXPROCS(1)
	runtime.LockOSThread()
}
