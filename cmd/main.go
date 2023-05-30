package main

import (
	"flag"
	"log"

	"github.com/gin-gonic/gin"
	"mybox.com/services/fileservice/config"
	"mybox.com/services/fileservice/internal/gcpclient"
	"mybox.com/services/fileservice/internal/services"
)

var enviroment string

var router *gin.Engine

func main() {
	flag.StringVar(&enviroment, "env", "dev", "enviroment type (dev, test, prod)")
	flag.Parse()

	conf, err := config.Init(enviroment)
	if err != nil {
		log.Println("ERR::0X237")
		panic(err)
	}

	bkt, err := gcpclient.Client(conf)
	if err != nil {
		log.Println("ERR::0X236")
		panic(err)
	}

	router = gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	err = services.Run(router, bkt, conf.GetString("app.server.addr"))
	if err != nil {
		panic(err)
	}
}
