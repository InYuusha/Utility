package net

import (
	"log"

	"github.com/shirou/gopsutil/net"
)

type Connection struct {
}

func GetConnection() *Connection{
	return &Connection{}
}

func (c *Connection)GetConnections(kind string) []net.ConnectionStat {
	conn, err := net.Connections(kind)
	if err != nil {
		log.Println(err)
	}
	return conn
}
