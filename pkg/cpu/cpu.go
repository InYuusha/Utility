package cpu

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/process"
)

type Info struct {
	Pid        *process.Process
	Cpu        float64
	Exe        string
	Background bool
}

//binding func
func CPU() *Info {
	return &Info{}
}

//cpu usage
func (t *Info) GetCPU() int {
	per, err := cpu.Percent(time.Second, false)
	if err != nil {
		fmt.Printf("%v", err)
	}
	return int(per[0])
}

func (t *Info) GetProcesses() []Info {

	stats := []Info{}
	processes, err := process.Processes()
	if err != nil {
		log.Println(err)
	}
	for _, p := range processes {

		perc, err := p.CPUPercent()
		if err != nil {
			log.Println(err)
		}
		perc = math.Floor(perc*1000) / 1000
		background, err := p.Background()
		if err != nil {
			log.Println(err)
		}
		exe, err := p.Exe()
		if err != nil {
			log.Println(err)
		}
		s := Info{Pid: p, Cpu: perc, Exe: exe, Background: background}
		stats = append(stats, s)

	}
	return stats
}

func (t *Info) KillP(p int32) {
	pids, err := process.Processes()
	if err != nil {
		log.Println(err)
	}
	for _, pid := range pids {
		if pid.Pid == p {
			err := pid.Kill()
			if err != nil {
				log.Printf("Cannot Kill %v", err)
			}
		}
	}
	log.Println("Pid does not exists")
}

func(t*Info)GetCpuDetails() []cpu.InfoStat{
	info, err := cpu.Info()
	if err != nil {
		log.Printf("%v", err)
	}
	return info
}
