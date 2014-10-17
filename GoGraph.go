package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"runtime"
)

var (
	adress string
	port   string
)

func init() {
	flag.StringVar(&adress, "adress", "localhost", "adress")
	flag.StringVar(&port, "port", "3000", "port")
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	r := gin.Default()

	listeningAdrs := adress + ":" + port
	r.Run(listeningAdrs)
}
