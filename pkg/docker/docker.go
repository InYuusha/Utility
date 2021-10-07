package docker

import (
	"log"

	"github.com/shirou/gopsutil/docker"
)
type Docker struct{}

func NewDocker()*Docker{
	return &Docker{}
}

func (d*Docker)GetAllDocker() interface{} {
	list, err := docker.GetDockerStat()
	if err != nil {
		log.Println(err)
	}
	return list
}
func (d*Docker)GetActiveDocker() interface{} {
	list, err := docker.GetDockerIDList()
	if err != nil {
		log.Println(err)
	}
	return list
}
