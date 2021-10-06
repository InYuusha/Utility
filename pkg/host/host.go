package host

import (
	"log"

	"github.com/shirou/gopsutil/host"
)

func GetHost() interface{} {
	tmp, err := host.Info()
	if err != nil {
		log.Printf("%v", err)
	}
	return tmp
}
