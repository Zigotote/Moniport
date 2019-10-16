package redis

import (
	"fmt"
	"strconv"

	redigo "github.com/gomodule/redigo/redis"
)

var conn redigo.Conn = nil

func getConnection() redigo.Conn {
	return conn
}

func Connect() {
	c, err := redigo.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
	}
	conn = c
}

func CloseConnection() {
	conn.Close()
}

func SendData(key string, data string) {
	c := getConnection()
	_, err := c.Do("SET", key, data)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Set " + key + ": " + data)
	}
}

func getData(key string) (interface{}, error) {
	c := getConnection()
	reply, err := c.Do("GET", key)
	if err != nil {
		fmt.Println(err)
	}
	return reply, err
}

func GetDataString(key string) string {
	s, _ := redigo.String(getData(key))
	return s
}

func GetDataInt(key string) int {
	i, _ := redigo.Int(getData(key))
	return i
}

func KeyExists(key string) bool {
	c := getConnection()
	reply, err := redigo.Int(c.Do("EXISTS", key))
	if err != nil {
		fmt.Println(err)
	}
	return reply == 1
}

func IncrKey(key string) {
	c := getConnection()
	_, err := c.Do("INCR", key)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Incr " + key)
	}
}

func AddToSet(key string, value string) {
	c := getConnection()
	reply, err := redigo.Int(c.Do("SADD", key, value))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(reply)
		if reply == 1 {
			fmt.Println("Added " + value + " to Set " + key)
		}
	}
}

func AddToOrdSet(key string, value string, score int64) {
	c := getConnection()
	reply, err := redigo.Int(c.Do("ZADD", key, score, value))
	if err != nil {
		fmt.Println(err)
	} else {
		if reply == 1 {
			fmt.Println("Added " + value + " to Ordered Set " + key + " with score " + strconv.FormatInt(score, 10))
		}
	}
}
