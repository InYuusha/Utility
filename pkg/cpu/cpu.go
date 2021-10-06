package cpu

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func GetCPU() int {
	per, err := cpu.Percent(time.Second, false)
	if err != nil {
		fmt.Printf("%v",err)
	}
	return int(per[0])
}
