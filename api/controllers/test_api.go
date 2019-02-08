package controllers

import (
	"gatecloud-boilerplate/api/models"
	"net/http"

	libController "github.com/gatecloud/webservice-library/controller"
	"github.com/gin-gonic/gin"
)

// TestAPIControl tests the framework
type TestAPIControl struct {
	libController.BaseControl
}

func (ctrl *TestAPIControl) GetAll(ctx *gin.Context) {
	var (
		api  models.TestAPI
		apis []models.TestAPI
	)

	api = models.TestAPI{
		Name:  "api",
		Value: "test",
	}
	apis = append(apis, api)
	ctx.Writer.Header().Set("X-Total-Count", "1")
	ctx.JSON(http.StatusOK, apis)
	ctx.Abort()
	return
}
