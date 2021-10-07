package memory

import (
	"log"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)
type Memory struct{

}

func GetMemory ()*Memory{
	return &Memory{}
}

func(m*Memory) GetPerc() int {
	memory, err := mem.VirtualMemory()
	if err != nil {
		log.Printf("%v", err)
	}
	perc := memory.UsedPercent

	return int(perc)

}

func(m*Memory) GetDisks() []*disk.UsageStat {
	var list []*disk.UsageStat

	tmp, err := disk.Partitions(false)
	if err != nil {
		log.Printf("%v", err)
	}
	for _, d := range tmp {
		u, _ := disk.Usage(d.Mountpoint)
		list = append(list, u)
	}
	return list

}
