package docker

import (
	"log"

	"github.com/shirou/gopsutil/docker"
)


func GetDocker() interface{} {
	list, err := docker.GetDockerStat()
	if err != nil {
		log.Println(err)
	}
	return list
}
