package main

import (
	"gatecloud-boilerplate/proxy/configs"
	"gatecloud-boilerplate/proxy/middlewares"
	"log"
	proxy "reverse-proxy"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	// Init Database
	db, err := InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Set the log mode
	if configs.Configuration.Production {
		gin.SetMode(gin.ReleaseMode)
		db.LogMode(false)
	} else {
		gin.SetMode(gin.DebugMode)
		db.LogMode(true)
	}

	// Init go-gin engine
	r := gin.Default()
	r.HandleMethodNotAllowed = true

	// Register middlewares
	r.Use(middlewares.Cors(configs.Configuration.CorsEnabled))
	r.Use(middlewares.AddResponseHeader("Access-Control-Expose-Headers", "X-Total-Count"))

	// Read reverse proxy config file
	servers, err := proxy.Default("route.json")
	if err != nil {
		log.Fatal(err)
	}

	// Start up all APIs
	for _, s := range *servers {
		switch s.Method {
		case "GET":
			r.GET(s.Path, s.ReverseProxy())
			r.GET(s.Path+"/:id", s.ReverseProxy())
		case "POST":
			r.POST(s.Path, s.ReverseProxy())
		case "PUT":
			r.PUT(s.Path, s.ReverseProxy())
		case "DELETE":
			r.DELETE(s.Path, s.ReverseProxy())
		case "PATCH":
			r.PATCH(s.Path, s.ReverseProxy())
		case "OPTIONS":
			r.OPTIONS(s.Path, s.ReverseProxy())
		}
	}
	r.Run(configs.Configuration.Port)
}
