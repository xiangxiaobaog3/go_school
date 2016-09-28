package main

import (
	"os/exec"
	"fmt"
)

func main() {
	var ClocksPerSec = float64(128)
	f, err := exec.LookPath("ls")
	if err !=nil {
		fmt.Println(err)
	}
	fmt.Println(f) // /bin/ls
	fmt.Println(ClocksPerSec)
}