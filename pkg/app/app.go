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
	go func() {
		for {
			time.Sleep(time.Second*2)
			runtime.Events.Emit("memory", memory.GetPerc())
		}
	}()
	go func(){
		for{
			time.Sleep(time.Second*2)
			runtime.Events.Emit("cpu",cpu.GetCPU())
		}
	}()
	return nil
}

