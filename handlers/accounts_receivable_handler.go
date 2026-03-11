package handlers

import (
	"kms_golang/database"
	"kms_golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAccountsReceivablePage handles GET /dashboard/accounts-receivable
func GetAccountsReceivablePage(c *gin.Context) {
	var totalizationLists []models.TotalizationList
	database.DB.Order("id ASC").Find(&totalizationLists)

	var totalizationUnits []models.TotalizationUnit
	database.DB.Order("id ASC").Find(&totalizationUnits)

	var company []struct {
		CoCode string  `json:"co_code"`
		CoName *string `json:"co_name"`
	}
	database.DB.Table("company_info").Select("co_code, co_name").Find(&company)

	c.JSON(http.StatusOK, gin.H{
		"totalization_lists": totalizationLists,
		"totalization_units": totalizationUnits,
		"company":            company,
	})
}

// GetAccountsReceivableList handles GET /dashboard/accounts-receivable/list
func GetAccountsReceivableList(c *gin.Context) {
	var lists []models.TotalizationList

	db := database.DB.Where("delete_flag = 0")
	if propCd := c.Query("prop_cd"); propCd != "" {
		db = db.Where("prop_cd = ?", propCd)
	}
	if targetMonth := c.Query("target_month"); targetMonth != "" {
		db = db.Where("target_month = ?", targetMonth)
	}

	if err := db.Order("id ASC").Find(&lists).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Load units for each list
	var units []models.TotalizationUnit
	var ids []uint
	for _, l := range lists {
		ids = append(ids, l.ID)
	}
	if len(ids) > 0 {
		database.DB.Where("totalization_list_id IN ?", ids).Find(&units)
	}

	c.JSON(http.StatusOK, gin.H{
		"totalization_lists": lists,
		"totalization_units": units,
	})
}
