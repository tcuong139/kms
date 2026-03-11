package handlers

import (
	"kms_golang/database"
	"kms_golang/models"
	"kms_golang/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

// GetReceptionList handles GET /dashboard/reception/list
func GetReceptionList(c *gin.Context) {
	var receptions []models.Reception
	db := database.DB.Where("delete_flag = 0")

	if propCd := c.Query("prop_cd"); propCd != "" {
		db = db.Where("prop_cd = ?", propCd)
	}
	if customerCd := c.Query("customer_cd"); customerCd != "" {
		db = db.Where("customer_cd = ?", customerCd)
	}
	if status := c.Query("status"); status != "" {
		db = db.Where("status = ?", status)
	}

	if err := db.Order("accept_number DESC").Find(&receptions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"receptions": receptions})
}

// GetReceptionDetail handles GET /dashboard/reception/:number
func GetReceptionDetail(c *gin.Context) {
	number := c.Param("number")
	var reception models.Reception
	if err := database.DB.Where("accept_number = ? AND delete_flag = 0", number).First(&reception).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "受付が見つかりません"})
		return
	}

	var imgs []models.ReceptionImg
	database.DB.Where("accept_number = ?", number).Find(&imgs)

	var pdfs []models.ReceptionPdf
	database.DB.Where("accept_number = ?", number).Find(&pdfs)

	c.JSON(http.StatusOK, gin.H{
		"reception": reception,
		"images":    imgs,
		"pdfs":      pdfs,
	})
}

// PostReceptionCreate handles POST /dashboard/reception/create
func PostReceptionCreate(c *gin.Context) {
	var req models.Reception
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	number, err := utils.GenerateReceptionNumber()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	req.AcceptNumber = number
	delFlagRec := int16(0)
	req.DeleteFlag = &delFlagRec

	if err := database.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"reception": req})
}

// PutReceptionUpdate handles PUT /dashboard/reception/:number
func PutReceptionUpdate(c *gin.Context) {
	number := c.Param("number")
	var reception models.Reception
	if err := database.DB.Where("accept_number = ? AND delete_flag = 0", number).First(&reception).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "受付が見つかりません"})
		return
	}

	if err := c.ShouldBindJSON(&reception); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	reception.AcceptNumber = number
	if err := database.DB.Save(&reception).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reception": reception})
}

// DeleteReception handles DELETE /dashboard/reception/:number (soft delete)
func DeleteReception(c *gin.Context) {
	number := c.Param("number")
	if err := database.DB.Model(&models.Reception{}).Where("accept_number = ?", number).Update("delete_flag", 1).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "受付を削除しました"})
}

// PutReceptionStatus handles PUT /dashboard/reception/:number/status
func PutReceptionStatus(c *gin.Context) {
	number := c.Param("number")
	var req struct {
		Status int16 `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	if err := database.DB.Model(&models.Reception{}).Where("accept_number = ?", number).Update("status", req.Status).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ステータスを更新しました"})
}

// GetReceptionSubmittedList handles GET /dashboard/reception/submitted-list
func GetReceptionSubmittedList(c *gin.Context) {
	var receptions []models.Reception
	database.DB.Where("delete_flag = 0 AND status = 1").Order("accept_number DESC").Find(&receptions)
	c.JSON(http.StatusOK, gin.H{"receptions": receptions})
}

// GetReceptionComplaintList handles GET /dashboard/reception/complaint-list
func GetReceptionComplaintList(c *gin.Context) {
	var receptions []models.Reception
	database.DB.Where("delete_flag = 0 AND reception_type = 2").Order("accept_number DESC").Find(&receptions)
	c.JSON(http.StatusOK, gin.H{"receptions": receptions})
}

// GetReceptionRegisterForm handles GET /dashboard/reception/register-form
func GetReceptionRegisterForm(c *gin.Context) {
	var customers []models.Customer
	database.DB.Where("delete_flag = 0").Select("customer_cd, customer_name").Find(&customers)
	for i := range customers {
		customers[i].CustomerPassword = nil
	}

	var users []models.User
	database.DB.Where("delete_flg = 0").Select("user_id, user_name").Find(&users)
	for i := range users {
		users[i].Password = nil
	}

	c.JSON(http.StatusOK, gin.H{"customers": customers, "users": users})
}

// GetReceptionDropdownCustomers handles GET /dashboard/reception/dropdown-customers
func GetReceptionDropdownCustomers(c *gin.Context) {
	var customers []models.Customer
	database.DB.Where("delete_flag = 0").Select("customer_cd, customer_name").Find(&customers)
	for i := range customers {
		customers[i].CustomerPassword = nil
	}
	c.JSON(http.StatusOK, gin.H{"customers": customers})
}

// PostReceptionGuestRegister handles POST /dashboard/reception/guest-register
func PostReceptionGuestRegister(c *gin.Context) {
	var req models.Reception
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	number, err := utils.GenerateReceptionNumber()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	req.AcceptNumber = number
	delFlag := int16(0)
	req.DeleteFlag = &delFlag

	if err := database.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"reception": req})
}

// PostReceptionUpdateCancel handles POST /dashboard/reception/:number/cancel
func PostReceptionUpdateCancel(c *gin.Context) {
	number := c.Param("number")
	cancelStatus := int16(9)
	database.DB.Model(&models.Reception{}).Where("accept_number = ?", number).Update("status", cancelStatus)
	c.JSON(http.StatusOK, gin.H{"message": "キャンセルしました"})
}

// PostReceptionUpdateAction handles POST /dashboard/reception/:number/action
func PostReceptionUpdateAction(c *gin.Context) {
	number := c.Param("number")
	var req struct {
		ActionContent *string `json:"action_content"`
		ActionDate    *string `json:"action_date"`
		ActionUserID  *string `json:"action_user_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	updates := map[string]interface{}{}
	if req.ActionContent != nil {
		updates["action_content"] = *req.ActionContent
	}
	if req.ActionDate != nil {
		updates["action_date"] = *req.ActionDate
	}
	if req.ActionUserID != nil {
		updates["action_user_id"] = *req.ActionUserID
	}

	database.DB.Model(&models.Reception{}).Where("accept_number = ?", number).Updates(updates)
	c.JSON(http.StatusOK, gin.H{"message": "対応を更新しました"})
}

// GetReceptionExportExcel handles GET /dashboard/reception/export-excel
func GetReceptionExportExcel(c *gin.Context) {
	var receptions []models.Reception
	database.DB.Where("delete_flag = 0").Order("accept_number DESC").Find(&receptions)

	f := excelize.NewFile()
	sheet := "Sheet1"
	headers := []string{"受付番号", "物件CD", "顧客CD", "ステータス", "受付日", "受付種別"}
	for i, h := range headers {
		col := string(rune('A' + i))
		f.SetCellValue(sheet, col+"1", h)
	}
	for i, r := range receptions {
		row := strconv.Itoa(i + 2)
		f.SetCellValue(sheet, "A"+row, r.AcceptNumber)
		if r.PropCd != nil {
			f.SetCellValue(sheet, "B"+row, *r.PropCd)
		}
		if r.CustomerCd != nil {
			f.SetCellValue(sheet, "C"+row, *r.CustomerCd)
		}
		if r.MReceptionStatusID != nil {
			f.SetCellValue(sheet, "D"+row, *r.MReceptionStatusID)
		}
		if r.RegistDatetime != nil {
			f.SetCellValue(sheet, "E"+row, *r.RegistDatetime)
		}
		if r.ReceptionType != nil {
			f.SetCellValue(sheet, "F"+row, *r.ReceptionType)
		}
	}

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=receptions.xlsx")
	f.Write(c.Writer)
}

// GetReceptionPropBasic handles GET /dashboard/reception/prop-basic
func GetReceptionPropBasic(c *gin.Context) {
	propCd := c.Query("prop_cd")
	var propBasic models.PropBasic
	if err := database.DB.Where("prop_cd = ? AND delete_flag = 0", propCd).First(&propBasic).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "物件が見つかりません"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"prop_basic": propBasic})
}

// PostReceptionSearchPropBasic handles POST /dashboard/reception/search-prop-basic
func PostReceptionSearchPropBasic(c *gin.Context) {
	var req struct {
		Keyword    string `json:"keyword"`
		CustomerCd string `json:"customer_cd"`
	}
	c.ShouldBindJSON(&req)

	var props []models.PropBasic
	db := database.DB.Where("delete_flag = 0")
	if req.Keyword != "" {
		db = db.Where("prop_name LIKE ? OR prop_cd LIKE ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%")
	}
	if req.CustomerCd != "" {
		db = db.Where("customer_cd = ?", req.CustomerCd)
	}
	db.Find(&props)
	c.JSON(http.StatusOK, gin.H{"prop_basics": props})
}

// GetReceptionPersonnel handles GET /dashboard/reception/personnel
func GetReceptionPersonnel(c *gin.Context) {
	customerCd := c.Query("customer_cd")
	var personnel []models.CustomerPersonnel
	database.DB.Where("customer_cd = ?", customerCd).Find(&personnel)
	c.JSON(http.StatusOK, gin.H{"personnel": personnel})
}

// PostReceptionSearchPersonnel handles POST /dashboard/reception/search-personnel
func PostReceptionSearchPersonnel(c *gin.Context) {
	var req struct {
		Keyword    string `json:"keyword"`
		CustomerCd string `json:"customer_cd"`
	}
	c.ShouldBindJSON(&req)

	var personnel []models.CustomerPersonnel
	db := database.DB.Table("customer_personnel")
	if req.CustomerCd != "" {
		db = db.Where("customer_cd = ?", req.CustomerCd)
	}
	if req.Keyword != "" {
		db = db.Where("person_name LIKE ?", "%"+req.Keyword+"%")
	}
	db.Find(&personnel)
	c.JSON(http.StatusOK, gin.H{"personnel": personnel})
}

// PostReceptionContentDetail handles POST /dashboard/reception/:number/content-detail
func PostReceptionContentDetail(c *gin.Context) {
	number := c.Param("number")
	var req struct {
		ContentDetail *string `json:"content_detail"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Model(&models.Reception{}).Where("accept_number = ?", number).
		Update("content_detail", req.ContentDetail)
	c.JSON(http.StatusOK, gin.H{"message": "更新しました"})
}

// PostReceptionReceivedConfirm handles POST /dashboard/reception/:number/received-confirm
func PostReceptionReceivedConfirm(c *gin.Context) {
	number := c.Param("number")
	database.DB.Model(&models.Reception{}).Where("accept_number = ?", number).
		Update("status", 2)
	c.JSON(http.StatusOK, gin.H{"message": "受付確認しました"})
}

// PostReceptionEditBillingCustomer handles POST /dashboard/reception/:number/billing-customer
func PostReceptionEditBillingCustomer(c *gin.Context) {
	number := c.Param("number")
	var req struct {
		CustomerCd string `json:"customer_cd" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Model(&models.Reception{}).Where("accept_number = ?", number).
		Update("customer_cd", req.CustomerCd)
	c.JSON(http.StatusOK, gin.H{"message": "請求先を更新しました"})
}

// PostReceptionUpdateCompletion handles POST /dashboard/reception/:number/completion
func PostReceptionUpdateCompletion(c *gin.Context) {
	number := c.Param("number")
	completedStatus := int16(3)
	database.DB.Model(&models.Reception{}).Where("accept_number = ?", number).
		Update("status", completedStatus)
	c.JSON(http.StatusOK, gin.H{"message": "完了にしました"})
}

// PostReceptionUpdateNotify handles POST /dashboard/reception/:number/notify
func PostReceptionUpdateNotify(c *gin.Context) {
	number := c.Param("number")
	var req struct {
		NotifyFlg *int16 `json:"notify_flg"`
	}
	c.ShouldBindJSON(&req)
	database.DB.Model(&models.Reception{}).Where("accept_number = ?", number).
		Update("notify_flg", req.NotifyFlg)
	c.JSON(http.StatusOK, gin.H{"message": "通知を更新しました"})
}

// GetReceptionCustomerOfProp handles GET /dashboard/reception/customer-of-prop
func GetReceptionCustomerOfProp(c *gin.Context) {
	propCd := c.Query("prop_cd")
	var prop models.PropBasic
	if err := database.DB.Where("prop_cd = ?", propCd).First(&prop).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "物件が見つかりません"})
		return
	}
	if prop.CustomerCd == nil {
		c.JSON(http.StatusOK, gin.H{"customer": nil})
		return
	}
	var customer models.Customer
	database.DB.Where("customer_cd = ?", *prop.CustomerCd).First(&customer)
	customer.CustomerPassword = nil
	c.JSON(http.StatusOK, gin.H{"customer": customer})
}

// PostReceptionPropBasicRegister handles POST /dashboard/reception/prop-basic-register
func PostReceptionPropBasicRegister(c *gin.Context) {
	var req models.PropBasic
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
	c.JSON(http.StatusCreated, gin.H{"prop_basic": req})
}

// GetReceptionLoadReception handles GET /dashboard/reception/load/:number
func GetReceptionLoadReception(c *gin.Context) {
	number := c.Param("number")
	var reception models.Reception
	if err := database.DB.Where("accept_number = ? AND delete_flag = 0", number).First(&reception).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "受付が見つかりません"})
		return
	}

	var imgs []models.ReceptionImg
	database.DB.Where("accept_number = ?", number).Find(&imgs)
	var pdfs []models.ReceptionPdf
	database.DB.Where("accept_number = ?", number).Find(&pdfs)

	var propBasic models.PropBasic
	if reception.PropCd != nil {
		database.DB.Where("prop_cd = ?", *reception.PropCd).First(&propBasic)
	}

	var customer models.Customer
	if reception.CustomerCd != nil {
		database.DB.Where("customer_cd = ?", *reception.CustomerCd).First(&customer)
		customer.CustomerPassword = nil
	}

	c.JSON(http.StatusOK, gin.H{
		"reception":  reception,
		"images":     imgs,
		"pdfs":       pdfs,
		"prop_basic": propBasic,
		"customer":   customer,
	})
}

// PostReceptionCustomer handles POST /dashboard/reception/customer
func PostReceptionCustomer(c *gin.Context) {
	var req struct {
		AcceptNumber string `json:"accept_number" binding:"required"`
		CustomerCd   string `json:"customer_cd" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Model(&models.Reception{}).Where("accept_number = ?", req.AcceptNumber).
		Update("customer_cd", req.CustomerCd)
	c.JSON(http.StatusOK, gin.H{"message": "顧客を設定しました"})
}

// PostReceptionCustomerPersonnel handles POST /dashboard/reception/customer-personnel
func PostReceptionCustomerPersonnel(c *gin.Context) {
	var req struct {
		AcceptNumber string `json:"accept_number" binding:"required"`
		PersonnelID  uint   `json:"personnel_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Model(&models.Reception{}).Where("accept_number = ?", req.AcceptNumber).
		Update("personnel_id", req.PersonnelID)
	c.JSON(http.StatusOK, gin.H{"message": "担当者を設定しました"})
}

// GetReceptionCheckCustomerName handles GET /dashboard/reception/check-customer-name
func GetReceptionCheckCustomerName(c *gin.Context) {
	name := c.Query("customer_name")
	var count int64
	database.DB.Model(&models.Customer{}).Where("customer_name = ? AND delete_flag = 0", name).Count(&count)
	c.JSON(http.StatusOK, gin.H{"exists": count > 0})
}

// GetReceptionCheckCustomerAddress handles GET /dashboard/reception/check-customer-address
func GetReceptionCheckCustomerAddress(c *gin.Context) {
	address := c.Query("block_name")
	var count int64
	database.DB.Model(&models.Customer{}).Where("block_name = ? AND delete_flag = 0", address).Count(&count)
	c.JSON(http.StatusOK, gin.H{"exists": count > 0})
}
