package controllers

import (
	"net/http"
	"utils"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	validator "gopkg.in/go-playground/validator.v8"
)

// Controller is the interface of struct Control
type Controller interface {
	Init()
	Prepare()
	Post(ctx *gin.Context)
	Patch(ctx *gin.Context)
	Delete(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Head(ctx *gin.Context)
	Options(ctx *gin.Context)
}

// Control representing the structure of a RESTful API handler
type Control struct {
	DB          *gorm.DB
	Name        string
	Model       interface{}
	Validate    *validator.Validate
	RedisClient *redis.Client
}

// Init inits the Control data
func (ctrl *Control) Init(db *gorm.DB, resourceName, languageCode string, validate *validator.Validate, model interface{}, redisClient *redis.Client) {
	ctrl.DB = db
	ctrl.Name = resourceName
	ctrl.Validate = validate
	ctrl.Model = model
	ctrl.RedisClient = redisClient
}

// Prepare corresponds http Prepare method
func (ctrl *Control) Prepare() {}

// Post corresponds http Post method
func (ctrl *Control) Post(ctx *gin.Context) {
	utils.ErrRespWithJSON(ctx, http.StatusNotFound, "Method Not Allowed")
	return
}

// Patch corresponds http Patch method
func (ctrl *Control) Patch(ctx *gin.Context) {
	utils.ErrRespWithJSON(ctx, http.StatusNotFound, "Method Not Allowed")
	return
}

// Delete corresponds http Delete method
func (ctrl *Control) Delete(ctx *gin.Context) {
	utils.ErrRespWithJSON(ctx, http.StatusNotFound, "Method Not Allowed")
	return
}

// GetByID corresponds http Get :id method
func (ctrl *Control) GetByID(ctx *gin.Context) {
	utils.ErrRespWithJSON(ctx, http.StatusNotFound, "Method Not Allowed")
	return
}

// GetAll corresponds http Get method
func (ctrl *Control) GetAll(ctx *gin.Context) {
	utils.ErrRespWithJSON(ctx, http.StatusNotFound, "Method Not Allowed")
	return
}

// Head corresponds http Head method
func (ctrl *Control) Head(ctx *gin.Context) {
	utils.ErrRespWithJSON(ctx, http.StatusNotFound, "Method Not Allowed")
	return
}

// Options corresponds http Options method
func (ctrl *Control) Options(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, nil)
	ctx.Abort()
	return
}
