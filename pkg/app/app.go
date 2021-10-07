package app

import (
	"time"
	"utils/pkg/cpu"
	"utils/pkg/memory"

	"github.com/wailsapp/wails"
)

type MyStruct struct {
	log *wails.CustomLogger
}

func (m *MyStruct) WailsInit(runtime *wails.Runtime) error {

	c := &cpu.Info{}      //cpu Info obj
	s := &memory.Memory{} //cpu Info obj

	go func() {
		for {
			time.Sleep(time.Second * 2)
			runtime.Events.Emit("memory", s.GetPerc())
		}
	}()
	go func() {
		for {
			time.Sleep(time.Second * 2)
			runtime.Events.Emit("cpu", c.GetCPU())
		}
	}()
	return nil
}
