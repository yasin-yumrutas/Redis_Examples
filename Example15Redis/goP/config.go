package goP

import (
	"net"
)

type RedisConfig struct {
	Adress     string
	Password   string
	Port       int
	Connection RedisConnection
}

type RedisConnection struct {
	Stream net.Conn
}

type RedisResponse struct {
	Message string
	Success bool
}
