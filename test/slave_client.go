package main

import "hqc/ipc"

func main() {
	register := ipc.Register()
	ipc.Heartbeat(register)
}
