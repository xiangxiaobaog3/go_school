package main

import (
	"./common"
	"fmt"
)

type Stat struct {
	MemTotal  uint64 `json:"mem_total"`
}

func checkError(e error){
	if e != nil{
		fmt.Println(e)
	}
}

func main() {
	stat := Stat{}
	mem_stat, _ := common.GetMemStat()
	stat.MemTotal = mem_stat.MemTotal
	total := stat.MemTotal
	fmt.Printf(string(total))
}

