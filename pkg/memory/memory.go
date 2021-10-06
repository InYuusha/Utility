package memory

import (
	"log"

	"github.com/shirou/gopsutil/mem"
)

func GetPerc() int {
	memory, err := mem.VirtualMemory()
	if err != nil {
		log.Printf("%v", err)
	}
	perc:=memory.UsedPercent

	return int(perc)

}
