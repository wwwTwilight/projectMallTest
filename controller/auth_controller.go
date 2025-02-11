package controller

import (
	"mall/global"
	models "mall/model"

	"github.com/gin-gonic/gin"
)

type userInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil { // 把json的内容绑定到结构体
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := global.Db.AutoMigrate(&user); err != nil { //自动建表，根据user的结构体，不会重复创建
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if err := global.Db.Create(&user).Error; err != nil { //创建用户
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "success"})

}

func Login(ctx *gin.Context) {
	input := userInfo{}

	if err := ctx.ShouldBindJSON(&input); err != nil { // 把json的内容绑定到结构体
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	if err := global.Db.Where("username = ?", input.Username).First(&user).Error; err != nil { //查询用户，查询的结果会绑定到user
		ctx.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	if user.Password != input.Password {
		ctx.JSON(401, gin.H{"error": "Username or Password incorrect"})
		return
	}

	ctx.JSON(200, gin.H{"message": "Login successfully"})
}
