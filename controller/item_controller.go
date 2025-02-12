package controller

import (
	"mall/global"
	models "mall/model"
	"mall/utils"

	"github.com/gin-gonic/gin"
)

func ShowItems(ctx *gin.Context) {
	// ctx.JSON(200, gin.H{"message": "ShowItems"})
	var ItemList []models.Item
	if err := global.Db.Find(&ItemList).Error; err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "ShowItems", "items": ItemList})
}

func SearchItem(ctx *gin.Context) {
	// ctx.JSON(200, gin.H{"message": "SearchItem"})
	type search struct {
		Search string `json:"search"`
	}
	var s search
	if err := ctx.ShouldBindJSON(&s); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var ItemList []models.Item
	if err := global.Db.Where("name LIKE ?", "%"+s.Search+"%").Find(&ItemList).Error; err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "SearchItem", "items": ItemList})
}

func CreateItem(ctx *gin.Context) {
	// ctx.JSON(200, gin.H{"message": "CreateItem"})
	var item models.Item

	if err := global.Db.AutoMigrate(&item); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if err := ctx.ShouldBindJSON(&item); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	token := ctx.GetHeader("Authorization")
	username, err := utils.ParseJWT(token)

	if err != nil {
		ctx.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	item.Owner = username

	if err := global.Db.Create(&item).Error; err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "CreateItem", "item": item}) //到这一步符合预期
}
