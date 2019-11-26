package main

import (
	"github.com/gin-gonic/gin"
	"github.com/suzuito/geolocation-sandbox-go/web/router"
)

// ApplicationImpl ...
type ApplicationImpl struct{}

func main() {
	app := ApplicationImpl{}
	root := gin.Default()
	root.GET("/", router.PostBeaconLocation(&app))
	root.Run()
}
