package goP

import (
	"errors"
	"fmt"
	"log"
	"net"
)

func (c *RedisConfig) Connect() (*RedisConfig, error) {
	stream, err := net.Dial("tcp", fmt.Sprintf("%s:%d", c.Adress, c.Port))
	if err != nil {
		return nil, err
	}

	c.Connection = RedisConnection{
		Stream: stream,
	}

	if c.Password != "" {
		_, err := c.Auth()
		if err != nil {
			log.Fatalf("Error while auth: %v", err)
		}
	}

	return c, nil
}

func (c *RedisConfig) Auth() (*RedisConfig, error) {
	buffer := make([]byte, 0, 4096)
	tmp := make([]byte, 256)

	authCommand := fmt.Sprintf("*%d\r\n$4\r\nAUTH\r\n$%d\r\n%s\r\n", 2, len(c.Password), c.Password)

	command := []byte(authCommand)

	_, err := c.Connection.Stream.Write(command)
	if err != nil {
		return nil, err
	}

	_, err = c.Connection.Stream.Read(tmp)
	if err != nil {
		return nil, err
	}

	for _, buffertmp := range tmp {
		buffer = append(buffer, buffertmp)
	}

	result := string(buffer)

	if string(result[0]) == "-" {
		return nil, errors.New("authenticaiton failed")
	}

	return c, nil
}

func (c *RedisConfig) Info() (*RedisConfig, error) {
	buffer := make([]byte, 0, 4096)
	tmp := make([]byte, 256)

	infoCommand := fmt.Sprintf("*%d\r\n$%d\r\n%s\r\n", 2, len(c.Password), c.Password)

	command := []byte(infoCommand)

	_, err := c.Connection.Stream.Write(command)
	if err != nil {
		return nil, err
	}

	_, err = c.Connection.Stream.Read(tmp)
	if err != nil {
		return nil, err
	}

	for _, buffertmp := range tmp {
		buffer = append(buffer, buffertmp)
	}

	result := string(buffer)

	if string(result[0]) == "-" {
		return &RedisResponse{
			Message: result,
		}, errors.New("info command failed")
	}

	return &RedisResponse{
		Message: result,
		Success: true,
	}, nil
}
