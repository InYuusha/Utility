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
// running processes
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
// process operations
type Operation struct{
	Msg string
	Status int32
}
// kill process
func (t *Info) KillP(p int32)Operation {
	pids, err := process.Processes()
	if err != nil {
		log.Println(err)
	}
	for _, pid := range pids {
		if pid.Pid == p {
			err := pid.Kill()
			if err != nil {
				log.Printf("Cannot Kill %v", err)
				return Operation{Msg:err.Error(),Status:400}
			}else{
				log.Printf("Killed %d",p)
				return Operation{Msg:"Process killed successfully",Status:200}
			}
		}
	}
	log.Println("Pid does not exists")
	return Operation{Msg:"Process is already dead",Status:404}
}
// stop process
func (t *Info) stopP(p int32) {
	pids, err := process.Processes()
	if err != nil {
		log.Println(err)
	}
	for _, pid := range pids {
		if pid.Pid == p {
			err := pid.Suspend()
			if err != nil {
				log.Printf("Cannot Stop %v", err)
			}else{
				log.Printf("Stopped %d",p)
			}
		}
	}
	log.Println("Pid does not exists")
}
//continue process
func (t *Info) contP(p int32) {
	pids, err := process.Processes()
	if err != nil {
		log.Println(err)
	}
	for _, pid := range pids {
		if pid.Pid == p {
			err := pid.Resume()
			if err != nil {
				log.Printf("Cannot Resume %v", err)
			}else{
				log.Printf("Resumed %d",p)
			}
		}
	}
	log.Println("Pid does not exists")
}

// cpu info
func(t*Info)GetCpuDetails() []cpu.InfoStat{
	info, err := cpu.Info()
	if err != nil {
		log.Printf("%v", err)
	}
	return info
}
