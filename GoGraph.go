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
	r.Use(CORSMiddleware())

	public := r.Group("/api")

	public.POST("/login", func(c *gin.Context) {

	})

	private := r.Group("/api/auth", tokenMiddleWare("unicornsareawesome"))

	private.POST("/dijkstras", func(c *gin.Context) {

	})
	listeningAdrs := adress + ":" + port
	r.Run(listeningAdrs)
}
