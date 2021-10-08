package memory

import (
	"log"
	"math"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

type Memory struct {
	Path        string
	UsedPercent float64
	Total       uint64
}

func GetMemory() *Memory {
	return &Memory{}
}

func (m *Memory) GetPerc() Memory {
	memory, err := mem.VirtualMemory()
	if err != nil {
		log.Printf("%v", err)
	}
	perc := math.Floor(memory.UsedPercent*10) / 10
	return Memory{Path:"",Total: memory.Total,UsedPercent: perc}

}

func (m *Memory) GetDisks() []Memory {
	var list []Memory

	tmp, err := disk.Partitions(false)
	if err != nil {
		log.Printf("%v", err)
	}
	for _, d := range tmp {
		u, _ := disk.Usage(d.Mountpoint)
		perc := math.Floor(u.UsedPercent*10) / 10

		disk := Memory{Path: u.Path, UsedPercent: perc, Total: u.Total}
		list = append(list, disk)
	}
	return list

}
