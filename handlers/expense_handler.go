package handlers

import (
	"kms_golang/database"
	"kms_golang/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetExpenseList handles GET /dashboard/expense/list
func GetExpenseList(c *gin.Context) {
	var expenses []models.Expense
	db := database.DB.Where("delete_flag = 0")

	if userID := c.Query("user_id"); userID != "" {
		db = db.Where("user_id = ?", userID)
	}
	if yearMonth := c.Query("year_month"); yearMonth != "" {
		db = db.Where("year_month = ?", yearMonth)
	}

	if err := db.Order("expense_date DESC").Find(&expenses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"expenses": expenses})
}

// PostExpenseCreate handles POST /dashboard/expense/create
func PostExpenseCreate(c *gin.Context) {
	var req models.Expense
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

	c.JSON(http.StatusCreated, gin.H{"expense": req})
}

// PutExpenseUpdate handles PUT /dashboard/expense/:id
func PutExpenseUpdate(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "無効なID"})
		return
	}
	var expense models.Expense
	if err := database.DB.Where("id = ? AND delete_flag = 0", id).First(&expense).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "経費が見つかりません"})
		return
	}

	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	expense.ID = uint(id)
	if err := database.DB.Save(&expense).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"expense": expense})
}

// DeleteExpense handles DELETE /dashboard/expense/:id (soft delete)
func DeleteExpense(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "無効なID"})
		return
	}
	delFlag := int16(1)
	if err := database.DB.Model(&models.Expense{}).Where("id = ?", id).Update("delete_flag", delFlag).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "経費を削除しました"})
}

// GetQuotationList handles GET /dashboard/quotation/list
func GetQuotationList(c *gin.Context) {
	var quotations []models.RequestQuotation
	db := database.DB.Where("delete_flag = 0")

	if propCd := c.Query("prop_cd"); propCd != "" {
		db = db.Where("prop_cd = ?", propCd)
	}
	if sekosakiCd := c.Query("sekosaki_cd"); sekosakiCd != "" {
		db = db.Where("sekosaki_cd = ?", sekosakiCd)
	}

	if err := db.Order("regist_datetime DESC").Find(&quotations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"quotations": quotations})
}

// PostQuotationCreate handles POST /dashboard/quotation/create
func PostQuotationCreate(c *gin.Context) {
	var req models.RequestQuotation
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	delFlag2 := int16(0)
	req.DeleteFlag = &delFlag2
	if err := database.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"quotation": req})
}

// DeleteQuotation handles DELETE /dashboard/quotation/:id (soft delete)
func DeleteQuotation(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "無効なID"})
		return
	}
	delFlag3 := int16(1)
	if err := database.DB.Model(&models.RequestQuotation{}).Where("id = ?", id).Update("delete_flag", delFlag3).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "見積依頼を削除しました"})
}

// PutQuotationUpdate handles PUT /dashboard/quotation/:id
func PutQuotationUpdate(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "無効なID"})
		return
	}
	var q models.RequestQuotation
	if err := database.DB.Where("id = ? AND delete_flag = 0", id).First(&q).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "見つかりません"})
		return
	}
	if err := c.ShouldBindJSON(&q); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	q.ID = uint(id)
	database.DB.Save(&q)
	c.JSON(http.StatusOK, gin.H{"quotation": q})
}

// PostQuotationSendMail handles POST /dashboard/quotation/send-mail
func PostQuotationSendMail(c *gin.Context) {
	var req struct {
		QuotationID uint `json:"quotation_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Model(&models.RequestQuotation{}).Where("id = ?", req.QuotationID).Update("send_mail_flg", 1)
	c.JSON(http.StatusOK, gin.H{"message": "メールを送信しました"})
}

// GetQuotationDropdown handles GET /dashboard/quotation/dropdown
func GetQuotationDropdown(c *gin.Context) {
	var props []models.PropBasic
	database.DB.Where("delete_flag = 0").Select("prop_cd, prop_name").Find(&props)

	var sekosakis []models.Sekosaki
	database.DB.Where("delete_flag = 0").Select("sekosaki_cd, sekosaki_name").Find(&sekosakis)
	for i := range sekosakis {
		sekosakis[i].SekosakiPassword = nil
	}

	c.JSON(http.StatusOK, gin.H{"prop_basics": props, "sekosakis": sekosakis})
}

// GetExpenseDetail handles GET /dashboard/expense/:id
func GetExpenseDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "無効なID"})
		return
	}
	var expense models.Expense
	if err := database.DB.Where("id = ? AND delete_flag = 0", id).First(&expense).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "経費が見つかりません"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"expense": expense})
}

// GetExpenseDropdown handles GET /dashboard/expense/dropdown
func GetExpenseDropdown(c *gin.Context) {
	var users []models.User
	database.DB.Where("delete_flg = 0").Select("user_id, user_name").Find(&users)
	for i := range users {
		users[i].Password = nil
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
