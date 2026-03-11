package handlers

import (
	"kms_golang/database"
	"kms_golang/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetSekosakiCreateEstimateList handles GET /dashboard/sekosaki-estimate/list
func GetSekosakiCreateEstimateList(c *gin.Context) {
	sekosakiCd := c.GetString("sekosaki_cd")
	perPage := 10
	if pp := c.Query("per_page"); pp != "" {
		if v, err := strconv.Atoi(pp); err == nil && v > 0 {
			perPage = v
		}
	}
	page := 1
	if p := c.Query("page"); p != "" {
		if v, err := strconv.Atoi(p); err == nil && v > 0 {
			page = v
		}
	}

	db := database.DB.Table("estimate").
		Select("estimate.*, prop_basics.prop_name, reception.room_number").
		Joins("LEFT JOIN prop_basics ON prop_basics.prop_cd = estimate.prop_cd").
		Joins("LEFT JOIN reception ON reception.accept_number = estimate.accept_number").
		Where("estimate.delete_flag = 0")

	if sekosakiCd != "" {
		db = db.Where("estimate.sekosaki_cd = ?", sekosakiCd)
	}

	var total int64
	db.Count(&total)

	var results []map[string]interface{}
	offset := (page - 1) * perPage
	db.Offset(offset).Limit(perPage).Order("estimate.estimate_number DESC").Scan(&results)

	c.JSON(http.StatusOK, gin.H{"estimates": results, "total": total, "page": page, "per_page": perPage})
}

// GetSekosakiCreateEstimateDetail handles GET /dashboard/sekosaki-estimate/:number/:subnumber
func GetSekosakiCreateEstimateDetail(c *gin.Context) {
	number := c.Param("number")
	subnumber := c.Param("subnumber")

	var estimate models.Estimate
	if err := database.DB.Where("estimate_number = ? AND subnumber = ? AND delete_flag = 0", number, subnumber).First(&estimate).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "見積が見つかりません"})
		return
	}

	var estConstructions []models.EstConstruction
	database.DB.Where("estimate_number = ? AND subnumber = ?", number, subnumber).Find(&estConstructions)

	var estConstructionDetails []models.EstConstructionDetail
	for _, ec := range estConstructions {
		var details []models.EstConstructionDetail
		database.DB.Where("est_construction_id = ?", ec.ID).Find(&details)
		estConstructionDetails = append(estConstructionDetails, details...)
	}

	c.JSON(http.StatusOK, gin.H{
		"estimate":                 estimate,
		"est_constructions":        estConstructions,
		"est_construction_details": estConstructionDetails,
	})
}

// PostSekosakiCreateEstimateUpdateDatetime handles POST /dashboard/sekosaki-estimate/update-datetime
func PostSekosakiCreateEstimateUpdateDatetime(c *gin.Context) {
	var req struct {
		EstimateNumber string `json:"estimate_number" binding:"required"`
		Subnumber      string `json:"subnumber" binding:"required"`
		DateField      string `json:"date_field" binding:"required"`
		DateValue      string `json:"date_value"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	database.DB.Model(&models.Estimate{}).
		Where("estimate_number = ? AND subnumber = ?", req.EstimateNumber, req.Subnumber).
		Update(req.DateField, req.DateValue)

	c.JSON(http.StatusOK, gin.H{"message": "更新しました"})
}

// GetSekosakiOrderWorkType1List handles GET /dashboard/sekosaki-estimate/order-work-type1
func GetSekosakiOrderWorkType1List(c *gin.Context) {
	propCd := c.Query("prop_cd")

	var orders []models.Order
	db := database.DB.Where("estimate_type = 1")
	if propCd != "" {
		db = db.Where("prop_cd = ?", propCd)
	}
	db.Find(&orders)

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

// GetSekosakiOrderWorkType2List handles GET /dashboard/sekosaki-estimate/order-work-type2
func GetSekosakiOrderWorkType2List(c *gin.Context) {
	propCd := c.Query("prop_cd")

	var orders []models.Order
	db := database.DB.Where("estimate_type = 2")
	if propCd != "" {
		db = db.Where("prop_cd = ?", propCd)
	}
	db.Find(&orders)

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}
