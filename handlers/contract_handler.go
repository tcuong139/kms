package handlers

import (
	"kms_golang/database"
	"kms_golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetContractList handles GET /dashboard/contract/list
// Returns contracts (construction + building management) for a given estimate
func GetContractList(c *gin.Context) {
	estimateNumber := c.Query("estimate_number")
	propCd := c.Query("prop_cd")

	var constructions []models.ContractConstruction
	db1 := database.DB.Where("delete_flag = 0")
	if propCd != "" {
		db1 = db1.Where("prop_cd = ?", propCd)
	}
	db1.Find(&constructions)

	var buildings []models.ContractBuildingManagement
	db2 := database.DB.Where("delete_flag = 0")
	if propCd != "" {
		db2 = db2.Where("prop_cd = ?", propCd)
	}
	db2.Find(&buildings)

	// Also load the estimate if given
	var estimate models.Estimate
	if estimateNumber != "" {
		database.DB.Where("estimate_number = ? AND delete_flag = 0", estimateNumber).First(&estimate)
	}

	c.JSON(http.StatusOK, gin.H{
		"contract_constructions":        constructions,
		"contract_building_managements": buildings,
		"estimate":                      estimate,
	})
}

// GetContractType1 handles GET /dashboard/contract/type1
func GetContractType1(c *gin.Context) {
	estimateNumber := c.Query("estimate_number")
	if estimateNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "estimate_number is required"})
		return
	}

	var estimate models.Estimate
	if err := database.DB.Where("estimate_number = ? AND delete_flag = 0", estimateNumber).
		First(&estimate).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "見積が見つかりません"})
		return
	}

	var contract models.ContractConstruction
	database.DB.Where("prop_cd = ? AND delete_flag = 0", estimate.PropCd).First(&contract)

	var comments []models.ContractComment
	database.DB.Where("contract_type = 1").Find(&comments)

	c.JSON(http.StatusOK, gin.H{
		"estimate": estimate,
		"contract": contract,
		"comments": comments,
	})
}

// GetContractType2 handles GET /dashboard/contract/type2
func GetContractType2(c *gin.Context) {
	estimateNumber := c.Query("estimate_number")
	if estimateNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "estimate_number is required"})
		return
	}

	var estimate models.Estimate
	if err := database.DB.Where("estimate_number = ? AND delete_flag = 0", estimateNumber).
		First(&estimate).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "見積が見つかりません"})
		return
	}

	var contract models.ContractBuildingManagement
	database.DB.Where("prop_cd = ? AND delete_flag = 0", estimate.PropCd).First(&contract)

	var comments []models.ContractComment
	database.DB.Where("contract_type = 2").Find(&comments)

	c.JSON(http.StatusOK, gin.H{
		"estimate": estimate,
		"contract": contract,
		"comments": comments,
	})
}

// PostContractType1 handles POST /dashboard/contract/type1  (save/update contract type 1)
func PostContractType1(c *gin.Context) {
	var req models.ContractConstruction
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	flag := int16(0)
	req.DeleteFlag = &flag

	if err := database.DB.Save(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"contract": req})
}

// PostContractType2 handles POST /dashboard/contract/type2 (save/update contract type 2)
func PostContractType2(c *gin.Context) {
	var req models.ContractBuildingManagement
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	flag := int16(0)
	req.DeleteFlag = &flag

	if err := database.DB.Save(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"contract": req})
}

// PostContractSendCustomer handles POST /dashboard/contract/send-customer
func PostContractSendCustomer(c *gin.Context) {
	var req struct {
		EstimateNumber string `json:"estimate_number" binding:"required"`
		CustomerCd     string `json:"customer_cd" binding:"required"`
		Message        string `json:"message"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	// Create customer notification
	msg := req.Message
	notif := models.CustomerNotifi{
		CustomerCd: req.CustomerCd,
		Message:    &msg,
	}
	isRead := int16(0)
	notif.IsRead = &isRead

	if err := database.DB.Create(&notif).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "顧客に送信しました"})
}
