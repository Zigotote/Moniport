package redis

import (
	"fmt"

	redigo "github.com/gomodule/redigo/redis"
)

func getConnection() redigo.Conn {
	c, err := redigo.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
	}
	return c
}

func closeConnection(c redigo.Conn) {
	c.Close()
}

func SendData(key string, data string) {
	c := getConnection()
	reply, err := c.Do("SET", key, data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reply)
	defer closeConnection(c)
}
