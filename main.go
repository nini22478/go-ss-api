package main

import (
	"jdudp/routers"
)

func main() {
	router := routers.InitRouter()
	//静态资源
	router.Run(":8082")
}
