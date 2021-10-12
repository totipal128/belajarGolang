package controllers

import (
	"fmt"
	"golang2/conection"
	"golang2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	db = conection.ConnDB()
)

func ViewUsers(context *gin.Context){
	var view_Users []models.User
	context.JSON(http.StatusOK, gin.H{
		"tittle" 	: "Users || Data Berhasil di Ambil",
		"data"		: db.Find(&view_Users),
	})
}

func CreateUsers(context *gin.Context){
	var users models.User
	fmt.Println(users)

	if err := context.ShouldBindJSON(&users); err != nil {
		context.SecureJSON(http.StatusUnauthorized, gin.H{
			"message" : err.Error(),
		})
		return
	}

	result := db.Create(&users)

	if result.Error != nil {
		context.SecureJSON(http.StatusOK, gin.H{"message" : result.Error})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"tittle"	: "data berhasil ditambah",
		"data "		:users,
	})
}

func Updateusers(context *gin.Context){
	var users models.User

	id := context.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&users).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Data not found!"})
		return
	}

	if err := context.ShouldBindJSON(&users); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "tittle" : "1"})
		return
	}

	result := db.Save(&users)
	if result.Error != nil {
		context	.SecureJSON(http.StatusOK, gin.H{"message" : result.Error, "tittle" : "2"})
	}

	context.JSON(http.StatusOK, gin.H{
		"tittle" : "users || berhasil di update",
		"data"	: users,
		"status": http.StatusOK,
	})
}

func DeleteUsers(context *gin.Context){
	var users models.User

	id := context.Params.ByName("id")
	if err := db.Where("id = ?", id).Delete(&users).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Data not found!"})
		return
	}else {
		context.JSON(http.StatusOK, gin.H{
			"title"	 	: "Users | Route Delete Data Success",
			"message"	: "Delete Data Success",
		})
	}
}