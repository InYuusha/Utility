package net

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"

	"github.com/shirou/gopsutil/net"
)

type Connection struct {
}

func GetConnection() *Connection {
	return &Connection{}
}

//get running connections
func (c *Connection) GetConnections(kind string) []net.ConnectionStat {
	conn, err := net.Connections(kind)
	if err != nil {
		log.Println(err)
	}
	return conn
}

type Operation  struct{
	Msg string
	Status int32
}
func(c*Connection) killPort(port int32)Operation{

	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		command := fmt.Sprintf("(Get-NetTCPConnection -LocalPort %v).OwningProcess -Force", port)
		cmd=exec.Command("Stop-Process", "-Id", command)
	} else {
		command := fmt.Sprintf("lsof -i tcp:%v | grep LISTEN | awk '{print $2}' | xargs kill -9", port)
		cmd=exec.Command("bash", "-c", command)

		
	}
	err:=cmd.Run()

	if err!=nil{
		log.Print(err)
		return Operation{Msg:err.Error(),Status: 400}
	}else{
		log.Printf("Killed port %d",port)
		return Operation{Msg:"Port freed successfully",Status: 200}
	} 
}
