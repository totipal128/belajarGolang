package controllers

import (
	"fmt"
	"golang2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)



func ViewMembers(context *gin.Context){
	var view_member []models.Member
	context.JSON(http.StatusOK, gin.H{
		"tittle" 	: "Members || Data Berhasil di Ambil",
		"data"		: db.Find(&view_member),
	})
}

func InsertMembers(context *gin.Context){
	var member models.Member
	fmt.Println(member)

	if err := context.ShouldBindJSON(&member); err != nil {
		context.SecureJSON(http.StatusUnauthorized, gin.H{
			"message" : err.Error(),
		})
		return
	}

	result := db.Create(&member)

	if result.Error != nil {
		context.SecureJSON(http.StatusOK, gin.H{"message" : result.Error})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"tittle"	: "data berhasil ditambah",
		"data "		:member,
	})
}

func UpdateMember(context *gin.Context){
	var member models.Member

	id := context.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&member).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Data not found!"})
		return
	}

	if err := context.ShouldBindJSON(&member); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "tittle" : "1"})
		return
	}

	result := db.Save(&member)
	if result.Error != nil {
		context	.SecureJSON(http.StatusOK, gin.H{"message" : result.Error, "tittle" : "2"})
	}

	context.JSON(http.StatusOK, gin.H{
		"tittle" : "Member || berhasil di update",
		"data"	: member,
		"status": http.StatusOK,
	})
}

func DeleteMember(context *gin.Context){
	var member models.Member

	id := context.Params.ByName("id")
	if err := db.Where("id = ?", id).Delete(&member).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Data not found!"})
		return
	}else{
		context.JSON(http.StatusOK, gin.H{
			"tittle" : "Member || berhasil di delete",
			"data"	: member,
			"status": http.StatusOK,
		})
	}
}