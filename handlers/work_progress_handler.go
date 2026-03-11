package handlers

import (
	"kms_golang/database"
	"kms_golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetWorkProgressList handles GET /dashboard/work-progress/list
func GetWorkProgressList(c *gin.Context) {
	var customers []models.Customer
	database.DB.Where("delete_flag = 0").Select("customer_cd, customer_name").Find(&customers)

	var propBasics []models.PropBasic
	database.DB.Where("delete_flag = 0").Select("prop_cd, prop_name").Find(&propBasics)

	var personnel []models.CustomerPersonnel
	database.DB.Select("personnel_code, personnel_name, customer_cd").Find(&personnel)

	c.JSON(http.StatusOK, gin.H{
		"customers":   customers,
		"prop_basics": propBasics,
		"personnels":  personnel,
	})
}

// GetWorkYearScheduleList handles GET /dashboard/work-year-schedule/list
func GetWorkYearScheduleList(c *gin.Context) {
	var customers []models.Customer
	database.DB.Where("delete_flag = 0").Select("customer_cd, customer_name").Find(&customers)

	c.JSON(http.StatusOK, gin.H{"customers": customers})
}
