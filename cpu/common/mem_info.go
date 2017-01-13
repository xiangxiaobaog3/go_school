package common

import (
	"io/ioutil"
	"strings"
	"strconv"


)

type MemStat struct {
	MemTotal  uint64 `json:"mem_total"`
	MemUsed   uint64 `json:"mem_used"`
	MemFree   uint64 `json:"mem_free"`
	Buffers   uint64 `json:"buffers"`
	Cached    uint64 `json:"cached"`
	SwapTotal uint64 `json:"swap_total"`
	SwapUsed  uint64 `json:"swap_used"`
	SwapFree  uint64 `json:"swap_free"`
	MemRate   uint64 `json:"mem_rate"`
	SwapRate  uint64 `json:"swap_rate"`
}

func GetMemStat() (stat MemStat, err error) {
	text, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		return
	}

	stat = MemStat{}
	lines := strings.Split(string(text), "\n")
	for i := 0; i < len(lines); i++ {
		if !strings.Contains(lines[i], ":") {
			continue
		}

		maps := strings.Split(lines[i], ":")
		key := strings.Trim(maps[0], " ")
		if key == "MemTotal" {
			stat.MemTotal = parse_mem_value(maps[1])
		}
	}
	return
}

func parse_mem_value(value string) (mem uint64) {
	data := strings.Fields(value)
	mem, _ = strconv.ParseUint(strings.Trim(data[0], "" ), 10, strconv.IntSize)
	mem = mem / 1024

	return
}