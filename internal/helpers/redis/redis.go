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

	command := []string{"SET", key, data}

	request(command, func(reply interface{}) {
		fmt.Println("Set " + key + ": " + data)
	})
}

func getData(key string) interface{} {

	command := []string{"GET", key}

	var data interface{}

	request(command, func(reply interface{}) {
		data = reply
	})

	return data
}

func GetDataString(key string) string {
	s, _ := redigo.String(getData(key), nil)
	return s
}

func GetDataInt(key string) int {
	i, _ := redigo.Int(getData(key), nil)
	return i
}

func KeyExists(key string) bool {

	command := []string{"EXISTS", key}

	exists := false

	request(command, func(reply interface{}) {
		replyCode, _ := redigo.Int(reply, nil)
		exists = replyCode == 1
	})

	return exists
}

func IncrKey(key string) {

	command := []string{"INCR", key}

	request(command, func(reply interface{}) {
		fmt.Println("Incr " + key)
	})

}

func AddToSet(key string, value string) {

	command := []string{"SADD", key, value}

	request(command, func(reply interface{}) {
		replyCode, _ := redigo.Int(reply, nil)
		if replyCode == 1 {
			fmt.Println("Added " + value + " to Set " + key)
		}
	})
}

func AddToOrdSet(key string, value string, score int64) {

	command := []string{"ZADD", key, strconv.FormatInt(score, 10), value}

	request(command, func(reply interface{}) {
		replyCode, _ := redigo.Int(reply, nil)
		if replyCode == 1 {
			fmt.Println("Added " + value + " to Ordered Set " + key + " with score " + strconv.FormatInt(score, 10))
		}
	})

}

func GetSet(key string) []string {
	command := []string{"SMEMBERS", key}

	var set []string

	request(command, func(reply interface{}) {
		set, _ = redigo.Strings(reply, nil)
	})

	return set
}

func getFromOrderedSet(key string, min, max string) map[int64]string {
	command := []string{"ZRANGEBYSCORE", key, min, max, "WITHSCORES"}

	orderedSet := make(map[int64]string)

	request(command, func(reply interface{}) {
		set, _ := redigo.Strings(reply, nil)

		var value string
		for i, v := range set {
			if i%2 == 0 {
				value = v
			} else {
				keyInt, _ := strconv.ParseInt(v, 10, 64)
				orderedSet[keyInt] = value
			}
		}
	})

	return orderedSet
}

func GetRangeFromOrderedSet(key string, min, max int64) map[int64]string {
	return getFromOrderedSet(key, strconv.FormatInt(min, 10), strconv.FormatInt(max, 10))
}

func GetAllFromOrderedSet(key string) map[int64]string {
	return getFromOrderedSet(key, "-inf", "+inf")
}

func request(command []string, callback func(interface{})) error {
	c := getConnection()

	commandName := command[0]
	arguments := command[1:]

	// Cast from []string to []interface{}
	argumentsItf := make([]interface{}, len(arguments))
	for i, v := range arguments {
		argumentsItf[i] = v
	}

	reply, err := c.Do(commandName, argumentsItf...)
	if err != nil {
		fmt.Println(err)
	} else {
		callback(reply)
	}
	return err
}
