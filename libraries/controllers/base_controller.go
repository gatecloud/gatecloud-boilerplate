package controllers

import (
	"errors"
	"gatecloud-boilerplate/libraries/models"
	"net/http"
	"strings"
	"utils"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	validator "gopkg.in/go-playground/validator.v8"
)

// Controller is the interface of struct Control
type Controller interface {
	Init(db *gorm.DB, validate *validator.Validate, model interface{}, redis *redis.Client)
	Prepare()
	Post(ctx *gin.Context)
	Patch(ctx *gin.Context)
	Delete(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Head(ctx *gin.Context)
	Options(ctx *gin.Context)
}

// BaseControl representing the structure of a RESTful API handler
type BaseControl struct {
	DB          *gorm.DB
	Validate    *validator.Validate
	Model       interface{}
	RedisClient *redis.Client
}

// Init inits the Control data
func (ctrl *BaseControl) Init(db *gorm.DB, validate *validator.Validate, model interface{}, redis *redis.Client) {
	ctrl.DB = db
	ctrl.Validate = validate
	ctrl.Model = model
	ctrl.RedisClient = redis
}

// Prepare corresponds http Prepare method
func (ctrl *BaseControl) Prepare() {}

// Post corresponds http Post method
func (ctrl *BaseControl) Post(ctx *gin.Context) {
	utils.ErrRespWithJSON(ctx, http.StatusNotFound, "Method Not Allowed")
	return
}

// Patch corresponds http Patch method
func (ctrl *BaseControl) Patch(ctx *gin.Context) {
	utils.ErrRespWithJSON(ctx, http.StatusNotFound, "Method Not Allowed")
	return
}

// Delete corresponds http Delete method
func (ctrl *BaseControl) Delete(ctx *gin.Context) {
	utils.ErrRespWithJSON(ctx, http.StatusNotFound, "Method Not Allowed")
	return
}

// GetByID corresponds http Get :id method
func (ctrl *BaseControl) GetByID(ctx *gin.Context) {
	utils.ErrRespWithJSON(ctx, http.StatusNotFound, "Method Not Allowed")
	return
}

// GetAll corresponds http Get method
func (ctrl *BaseControl) GetAll(ctx *gin.Context) {
	utils.ErrRespWithJSON(ctx, http.StatusNotFound, "Method Not Allowed")
	return
}

// Head corresponds http Head method
func (ctrl *BaseControl) Head(ctx *gin.Context) {
	utils.ErrRespWithJSON(ctx, http.StatusNotFound, "Method Not Allowed")
	return
}

// Options corresponds http Options method
func (ctrl *BaseControl) Options(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, nil)
	ctx.Abort()
	return
}

// GetClientProfile gets the user information
func (ctrl *BaseControl) GetClientProfile(ctx *gin.Context) (models.ClientProfile, error) {
	var clientProfile models.ClientProfile
	token, err := utils.ExtractToken(ctx.Request.Header)
	if err != nil {
		return clientProfile, errors.New("Header token: " + err.Error())
	}

	s := strings.SplitN(token, ".", 3)
	if len(s) == 3 {
		// clientProfile, err = jwtParse(token, configs.GConfig.PublicPemPath)
		// if err != nil {
		// 	return clientProfile, err
		// }
	} else {
		if err := ctrl.DB.Where("access_token = ?", token).Find(&clientProfile).Error; err != nil {
			return clientProfile, err
		}
	}

	return clientProfile, nil
}

//GetQueryID gets id from Request
func (ctrl BaseControl) GetQueryID(ctx *gin.Context) string {
	id := ctx.Params.ByName("id")
	if id == "" {
		if id = ctx.Query("id"); id != "" {
			return id
		}
		return ""
	}
	return id
}

// IDToUUID converts string type to uuid
func (ctrl BaseControl) IDToUUID(idStr string) (uuid.UUID, error) {
	if idStr == "" {
		return uuid.UUID{}, errors.New("UUID Missing")
	}
	id, err := uuid.FromString(idStr)
	if err != nil {
		return uuid.UUID{}, err
	}
	if id == (uuid.UUID{}) {
		return id, errors.New("UUID invalid")
	}
	return id, nil
}
