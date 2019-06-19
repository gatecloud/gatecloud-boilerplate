package main

import (
	"gatecloud-boilerplate/api/configs"
	"gatecloud-boilerplate/api/routes"
	"gatecloud-boilerplate/api/validations"
	"log"

	libRoute "github.com/gatecloud/webservice-library/route"
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

	// Create shared resources
	sr := &libRoute.Resource{
		DB:        db,
		Validator: validations.InitValidation(),
	}

	// Init go-gin engine
	r := gin.Default()
	r.HandleMethodNotAllowed = true
	apiRouter := r.Group("/api")
	distributeRouters(apiRouter, routes.RouteMap["api"], sr)

	testRouter := r.Group("/test")
	distributeRouters(testRouter, routes.RouteMap["test"], sr)

	r.Run(configs.Configuration.Port)
}

func distributeRouters(r *gin.RouterGroup, routeGroup []libRoute.Route, sr libRoute.Resourcer) {
	for _, v := range routeGroup {
		r.GET("/"+v.Name, v.Register("GetAll", sr))
		r.POST("/"+v.Name, v.Register("Post", sr))
		r.PATCH("/"+v.Name, v.Register("Patch", sr))
		r.PUT("/"+v.Name, v.Register("Put", sr))
		r.DELETE("/"+v.Name, v.Register("Delete", sr))
		r.GET("/"+v.Name+"/:id", v.Register("GetByID", sr))
		r.PATCH("/"+v.Name+"/:id", v.Register("Patch", sr))
		r.DELETE("/"+v.Name+"/:id", v.Register("Delete", sr))
		r.OPTIONS("/"+v.Name, v.Register("Options", sr))
	}
}
