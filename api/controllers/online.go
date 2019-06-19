package controllers

import (
	"net/http"
	"timingniao_wlx_api/libraries/models"

	"github.com/gatecloud/utils"
	"github.com/gin-gonic/gin"
)

type OnlineController struct {
	Controller
}

func (ctrl *OnlineController) Patch(ctx *gin.Context) {
	var (
		entities   []models.Online
		tableModel interface{}
	)

	clientProfile, err := ctrl.GetClientProfile(ctx)
	if err != nil {
		utils.ErrRespWithJSON(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	if err := ctx.BindJSON(&entities); err != nil {
		utils.ErrRespWithJSON(ctx, http.StatusBadRequest, err.Error())
		return
	}

	tx := ctrl.DB.Begin()
	for i, bindEntity := range entities {
		if err := ctrl.Validator.Struct(bindEntity); err != nil {
			tx.Rollback()
			utils.ErrRespWithJSON(ctx, http.StatusBadRequest, err.Error())
			return
		}

		switch bindEntity.Type {
		// case types.OnlineProduct:
		// 	if !ctrl.RolePermit(clientProfile) {
		// 		tx.Rollback()
		// 		utils.ErrRespWithJSON(ctx, http.StatusForbidden, "permission forbidden")
		// 		return
		// 	}

		// 	tableModel = &models.Product{}
		// 	var chkEntity models.Product
		// 	if tx.Where("id = ?", bindEntity.ID).Find(&chkEntity).RecordNotFound() {
		// 		utils.ErrRespWithJSON(ctx, http.StatusBadRequest, "Product not found")
		// 		return
		// 	}

		// t := tx.NewScope(&models.Product{}).TableName()
		// if err := tx.Exec("UPDATE "+t+" SET online = ? WHERE id = ?",
		// 	bindEntity.Online, bindEntity.ID).Error; err != nil {
		// 	utils.ErrRespWithJSON(ctx, http.StatusInternalServerError, err.Error())
		// 	return
		// }

		// case types.OnlineInventory:
		// 	tableModel = &models.ProductInventory{}
		// 	var chkEntity models.ProductInventory
		// 	if tx.Where("id = ?", bindEntity.ID).Find(&chkEntity).RecordNotFound() {
		// 		utils.ErrRespWithJSON(ctx, http.StatusBadRequest, "Inventory Category not found")
		// 		return
		// 	}
		default:
			utils.ErrRespWithJSON(ctx, http.StatusBadRequest, "The current type is not supported")
			return
		}

		tableName := tx.NewScope(tableModel).TableName()
		if err := tx.Exec("UPDATE "+tableName+" SET online=? WHERE id = ?",
			bindEntity.Online, bindEntity.ID).Error; err != nil {
			utils.ErrRespWithJSON(ctx, http.StatusInternalServerError, err.Error())
			return
		}
		entities[i] = bindEntity
	}
	tx.Commit()
	ctx.JSON(http.StatusOK, entities)
	ctx.Abort()
	return
}
