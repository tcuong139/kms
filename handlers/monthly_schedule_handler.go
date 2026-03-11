package handlers

import (
	"kms_golang/database"
	"kms_golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetMonthlyScheduleList handles GET /dashboard/monthly-schedule/list
func GetMonthlyScheduleList(c *gin.Context) {
	var schedules []models.WorkSchedule
	db := database.DB.Where("delete_flag = 0")

	if propCd := c.Query("prop_cd"); propCd != "" {
		db = db.Where("prop_cd = ?", propCd)
	}
	if sekosakiCd := c.Query("sekosaki_cd"); sekosakiCd != "" {
		db = db.Where("sekosaki_cd = ?", sekosakiCd)
	}
	if yearMonth := c.Query("year_month"); yearMonth != "" {
		// Filter by year-month prefix (YYYY-MM)
		db = db.Where("work_date LIKE ?", yearMonth+"%")
	}

	if err := db.Order("work_date ASC").Find(&schedules).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"schedules": schedules})
}

// GetMonthlyScheduleDropdownData handles GET /dashboard/monthly-schedule/dropdown-data
func GetMonthlyScheduleDropdownData(c *gin.Context) {
	var customers []models.Customer
	database.DB.Select("customer_cd, customer_name").Where("delete_flag = 0").Order("customer_cd").Find(&customers)

	var sekosakis []models.Sekosaki
	database.DB.Select("sekosaki_cd, sekosaki_name").Where("delete_flag = 0").Order("sekosaki_cd").Find(&sekosakis)

	var props []models.PropBasic
	database.DB.Select("prop_cd, prop_name").Where("delete_flag = 0").Order("prop_cd").Find(&props)

	c.JSON(http.StatusOK, gin.H{
		"customers": customers,
		"sekosakis": sekosakis,
		"props":     props,
	})
}

// PostMonthlyScheduleUpdateDatetime handles POST /dashboard/monthly-schedule/update-datetime
func PostMonthlyScheduleUpdateDatetime(c *gin.Context) {
	var req struct {
		ID       uint   `json:"id" binding:"required"`
		WorkDate string `json:"work_date" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	if err := database.DB.Model(&models.WorkSchedule{}).
		Where("id = ?", req.ID).
		Update("work_date", req.WorkDate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "日時を更新しました"})
}

// GetMonthFollowCustomer handles GET /dashboard/monthly-schedule/month-follow-customer
func GetMonthFollowCustomer(c *gin.Context) {
	customerCd := c.Query("customer_cd")
	yearMonth := c.Query("year_month")

	db := database.DB.Table("work_schedule").
		Select("work_schedule.*, orders.customer_cd, orders.prop_cd").
		Joins("LEFT JOIN orders ON orders.order_id = work_schedule.order_id").
		Where("work_schedule.delete_flag = 0")

	if customerCd != "" {
		db = db.Where("orders.customer_cd = ?", customerCd)
	}
	if yearMonth != "" {
		db = db.Where("work_schedule.work_date LIKE ?", yearMonth+"%")
	}

	var results []map[string]interface{}
	if err := db.Find(&results).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"schedules": results})
}
