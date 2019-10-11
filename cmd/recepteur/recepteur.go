package main

import (
	redis "moniport/cmd/recepteur/redis"
)

func main() {

	redis.SendData("test", "yes")

}
