package handlers

import (
	"kms_golang/database"
	"kms_golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCrewList handles GET /dashboard/crew/list
func GetCrewList(c *gin.Context) {
	var crews []models.Crew
	db := database.DB.Where("delete_flag = 0")

	if name := c.Query("crew_name"); name != "" {
		db = db.Where("crew_name LIKE ?", "%"+name+"%")
	}
	if userID := c.Query("user_id"); userID != "" {
		db = db.Where("user_id = ?", userID)
	}

	if err := db.Order("crew_code ASC").Find(&crews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"crews": crews})
}

// GetCrewDetail handles GET /dashboard/crew/:id
func GetCrewDetail(c *gin.Context) {
	id := c.Param("id")
	var crew models.Crew
	if err := database.DB.Where("crew_code = ? AND delete_flag = 0", id).First(&crew).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "クルーが見つかりません"})
		return
	}

	var workplaces []models.CrewWorkplace
	database.DB.Where("crew_code = ?", id).Find(&workplaces)

	c.JSON(http.StatusOK, gin.H{"crew": crew, "workplaces": workplaces})
}

// PostCrewCreate handles POST /dashboard/crew/create
func PostCrewCreate(c *gin.Context) {
	var req models.Crew
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	delFlag := int16(0)
	req.DeleteFlag = &delFlag
	if err := database.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"crew": req})
}

// PutCrewUpdate handles PUT /dashboard/crew/:id
func PutCrewUpdate(c *gin.Context) {
	id := c.Param("id")
	var crew models.Crew
	if err := database.DB.Where("crew_code = ? AND delete_flag = 0", id).First(&crew).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "クルーが見つかりません"})
		return
	}

	if err := c.ShouldBindJSON(&crew); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	crew.CrewCode = id
	if err := database.DB.Save(&crew).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"crew": crew})
}

// DeleteCrew handles DELETE /dashboard/crew/:id (soft delete)
func DeleteCrew(c *gin.Context) {
	id := c.Param("id")
	delFlag := int16(1)
	if err := database.DB.Model(&models.Crew{}).Where("crew_code = ?", id).Update("delete_flag", delFlag).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "クルーを削除しました"})
}

// GetMonthlyReportList handles GET /dashboard/monthly-report/list
func GetMonthlyReportList(c *gin.Context) {
	var reports []models.MonthlyReportNotes
	db := database.DB

	if userID := c.Query("user_id"); userID != "" {
		db = db.Where("user_id = ?", userID)
	}
	if yearMonth := c.Query("year_month"); yearMonth != "" {
		db = db.Where("year_month = ?", yearMonth)
	}

	if err := db.Order("year_month DESC").Find(&reports).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"monthly_reports": reports})
}

// PostMonthlyReportCreate handles POST /dashboard/monthly-report/create
func PostMonthlyReportCreate(c *gin.Context) {
	var req models.MonthlyReportNotes
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	if err := database.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"monthly_report": req})
}

// GetMonthlyReportDetail handles GET /dashboard/monthly-report/:id
func GetMonthlyReportDetail(c *gin.Context) {
	id := c.Param("id")
	var report models.MonthlyReportNotes
	if err := database.DB.First(&report, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "月報が見つかりません"})
		return
	}

	var imgs []models.MonthlyReportNotesImg
	database.DB.Where("monthly_report_notes_id = ?", id).Find(&imgs)

	c.JSON(http.StatusOK, gin.H{"monthly_report": report, "images": imgs})
}

// PutMonthlyReportUpdate handles PUT /dashboard/monthly-report/:id
func PutMonthlyReportUpdate(c *gin.Context) {
	id := c.Param("id")
	var report models.MonthlyReportNotes
	if err := database.DB.First(&report, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "月報が見つかりません"})
		return
	}
	if err := c.ShouldBindJSON(&report); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Save(&report)
	c.JSON(http.StatusOK, gin.H{"monthly_report": report})
}

// GetMonthlyReportDropdown handles GET /dashboard/monthly-report/dropdown
func GetMonthlyReportDropdown(c *gin.Context) {
	var users []models.User
	database.DB.Where("delete_flg = 0").Select("user_id, user_name").Find(&users)
	for i := range users {
		users[i].Password = nil
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

// PostMonthlyReportSearch handles POST /dashboard/monthly-report/search
func PostMonthlyReportSearch(c *gin.Context) {
	var req struct {
		UserID    string `json:"user_id"`
		YearMonth string `json:"year_month"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	db := database.DB
	if req.UserID != "" {
		db = db.Where("user_id = ?", req.UserID)
	}
	if req.YearMonth != "" {
		db = db.Where("year_month = ?", req.YearMonth)
	}

	var reports []models.MonthlyReportNotes
	db.Order("year_month DESC").Find(&reports)
	c.JSON(http.StatusOK, gin.H{"monthly_reports": reports})
}

// PostMonthlyReportCancelSend handles POST /dashboard/monthly-report/cancel-send
func PostMonthlyReportCancelSend(c *gin.Context) {
	var req struct {
		ID uint `json:"id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Model(&models.MonthlyReportNotes{}).Where("id = ?", req.ID).Update("send_flg", 0)
	c.JSON(http.StatusOK, gin.H{"message": "送信をキャンセルしました"})
}

// GetCrewWorkplaceDetail handles GET /dashboard/crew/:id/workplace-detail
func GetCrewWorkplaceDetail(c *gin.Context) {
	id := c.Param("id")
	var details []models.CrewWorkplaceDetail
	database.DB.Where("crew_code = ?", id).Find(&details)
	c.JSON(http.StatusOK, gin.H{"workplace_details": details})
}

// PostCrewWorkplaceDetailCreate handles POST /dashboard/crew/workplace-detail/create
func PostCrewWorkplaceDetailCreate(c *gin.Context) {
	var req models.CrewWorkplaceDetail
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Create(&req)
	c.JSON(http.StatusCreated, gin.H{"workplace_detail": req})
}
