package handlers

import (
	"kms_golang/database"
	"kms_golang/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

// GetInvoiceList handles GET /dashboard/invoice/list
func GetInvoiceList(c *gin.Context) {
	var invoices []models.CusInvoiceDetail
	db := database.DB.Where("delete_flag = 0")

	if customerCd := c.Query("customer_cd"); customerCd != "" {
		db = db.Where("customer_cd = ?", customerCd)
	}
	if invoiceMonth := c.Query("invoice_month"); invoiceMonth != "" {
		db = db.Where("invoice_month = ?", invoiceMonth)
	}

	if err := db.Order("id DESC").Find(&invoices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"invoices": invoices})
}

// GetInvoiceDetail handles GET /dashboard/invoice/:id
func GetInvoiceDetail(c *gin.Context) {
	idStr := c.Param("number")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "無効なID"})
		return
	}
	var invoice models.CusInvoiceDetail
	if err := database.DB.Where("id = ? AND delete_flag = 0", id).First(&invoice).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "請求書が見つかりません"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"invoice": invoice})
}

// PostInvoiceCreate handles POST /dashboard/invoice/create
func PostInvoiceCreate(c *gin.Context) {
	var req models.CusInvoiceDetail
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

	c.JSON(http.StatusCreated, gin.H{"invoice": req})
}

// PutInvoiceUpdate handles PUT /dashboard/invoice/:id
func PutInvoiceUpdate(c *gin.Context) {
	idStr := c.Param("number")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "無効なID"})
		return
	}
	var invoice models.CusInvoiceDetail
	if err := database.DB.Where("id = ? AND delete_flag = 0", id).First(&invoice).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "請求書が見つかりません"})
		return
	}

	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	invoice.ID = uint(id)
	if err := database.DB.Save(&invoice).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"invoice": invoice})
}

// DeleteInvoice handles DELETE /dashboard/invoice/:id (soft delete)
func DeleteInvoice(c *gin.Context) {
	idStr := c.Param("number")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "無効なID"})
		return
	}
	delFlag := int16(1)
	if err := database.DB.Model(&models.CusInvoiceDetail{}).Where("id = ?", id).Update("delete_flag", delFlag).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "請求書を削除しました"})
}

// PutInvoiceApprove handles PUT /dashboard/invoice/:id/approve
func PutInvoiceApprove(c *gin.Context) {
	idStr := c.Param("number")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "無効なID"})
		return
	}
	approvedType := int16(2)
	if err := database.DB.Model(&models.CusInvoiceDetail{}).
		Where("id = ?", id).
		Update("invoice_type", approvedType).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "請求書を承認しました"})
}

// GetInvoiceEstimateType1 handles GET /dashboard/invoice/estimate-type1
func GetInvoiceEstimateType1(c *gin.Context) {
	propCd := c.Query("prop_cd")
	var estimates []models.Estimate
	database.DB.Where("prop_cd = ? AND delete_flag = 0 AND estimate_type = 1", propCd).Find(&estimates)
	c.JSON(http.StatusOK, gin.H{"estimates": estimates})
}

// GetInvoiceEstimateType2 handles GET /dashboard/invoice/estimate-type2
func GetInvoiceEstimateType2(c *gin.Context) {
	propCd := c.Query("prop_cd")
	var estimates []models.Estimate
	database.DB.Where("prop_cd = ? AND delete_flag = 0 AND estimate_type = 2", propCd).Find(&estimates)
	c.JSON(http.StatusOK, gin.H{"estimates": estimates})
}

// GetInvoiceResultInvoices handles GET /dashboard/invoice/result-invoices
func GetInvoiceResultInvoices(c *gin.Context) {
	customerCd := c.Query("customer_cd")
	yearMonth := c.Query("year_month")

	var invoices []models.CusInvoiceDetail
	db := database.DB.Where("delete_flag = 0")
	if customerCd != "" {
		db = db.Where("customer_cd = ?", customerCd)
	}
	if yearMonth != "" {
		db = db.Where("invoice_month = ?", yearMonth)
	}
	db.Order("id DESC").Find(&invoices)
	c.JSON(http.StatusOK, gin.H{"invoices": invoices})
}

// GetInvoiceResultDeposit handles GET /dashboard/invoice/result-deposit
func GetInvoiceResultDeposit(c *gin.Context) {
	customerCd := c.Query("customer_cd")
	yearMonth := c.Query("year_month")

	var deposits []models.DepositSlip
	db := database.DB.Where("delete_flag = 0")
	if customerCd != "" {
		db = db.Where("customer_cd = ?", customerCd)
	}
	if yearMonth != "" {
		db = db.Where("year_month = ?", yearMonth)
	}
	db.Order("deposit_number DESC").Find(&deposits)
	c.JSON(http.StatusOK, gin.H{"deposits": deposits})
}

// GetInvoiceDropdownData handles GET /dashboard/invoice/dropdown-data
func GetInvoiceDropdownData(c *gin.Context) {
	var customers []models.Customer
	database.DB.Where("delete_flag = 0").Select("customer_cd, customer_name").Find(&customers)
	for i := range customers {
		customers[i].CustomerPassword = nil
	}

	var propBasics []models.PropBasic
	database.DB.Where("delete_flag = 0").Select("prop_cd, prop_name").Find(&propBasics)

	c.JSON(http.StatusOK, gin.H{"customers": customers, "prop_basics": propBasics})
}

// PostInvoiceList handles POST /dashboard/invoice/search-list
func PostInvoiceList(c *gin.Context) {
	var req struct {
		CustomerCd   string `json:"customer_cd"`
		InvoiceMonth string `json:"invoice_month"`
		InvoiceType  *int16 `json:"invoice_type"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	db := database.DB.Where("delete_flag = 0")
	if req.CustomerCd != "" {
		db = db.Where("customer_cd = ?", req.CustomerCd)
	}
	if req.InvoiceMonth != "" {
		db = db.Where("invoice_month = ?", req.InvoiceMonth)
	}
	if req.InvoiceType != nil {
		db = db.Where("invoice_type = ?", *req.InvoiceType)
	}

	var invoices []models.CusInvoiceDetail
	db.Order("id DESC").Find(&invoices)
	c.JSON(http.StatusOK, gin.H{"invoices": invoices})
}

// GetInvoiceExportExcel handles GET /dashboard/invoice/export-excel
func GetInvoiceExportExcel(c *gin.Context) {
	var invoices []models.CusInvoiceDetail
	db := database.DB.Where("delete_flag = 0")
	if customerCd := c.Query("customer_cd"); customerCd != "" {
		db = db.Where("customer_cd = ?", customerCd)
	}
	if invoiceMonth := c.Query("invoice_month"); invoiceMonth != "" {
		db = db.Where("invoice_month = ?", invoiceMonth)
	}
	db.Order("id DESC").Find(&invoices)

	f := excelize.NewFile()
	sheet := "Sheet1"
	headers := []string{"ID", "顧客CD", "請求月", "請求種別", "金額", "税額"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}
	for row, inv := range invoices {
		r := row + 2
		vals := []interface{}{inv.ID, inv.CustomerCd, inv.InvoiceMonth, inv.InvoiceType, inv.TotalAmount, inv.TaxAmount}
		for col, v := range vals {
			cell, _ := excelize.CoordinatesToCellName(col+1, r)
			f.SetCellValue(sheet, cell, v)
		}
	}

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=invoice_list.xlsx")
	f.Write(c.Writer)
}

// PostInvoiceUpdateBiko handles POST /dashboard/invoice/update-biko
func PostInvoiceUpdateBiko(c *gin.Context) {
	var req struct {
		CustomerCd   string  `json:"customer_cd" binding:"required"`
		InvoiceMonth string  `json:"invoice_month" binding:"required"`
		BikoText     *string `json:"biko_text"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	var biko models.InvoiceBiko
	result := database.DB.Where("customer_cd = ? AND invoice_month = ?", req.CustomerCd, req.InvoiceMonth).First(&biko)
	if result.Error != nil {
		biko = models.InvoiceBiko{CustomerCd: req.CustomerCd, InvoiceMonth: req.InvoiceMonth, BikoText: req.BikoText}
		database.DB.Create(&biko)
	} else {
		database.DB.Model(&biko).Update("biko_text", req.BikoText)
	}
	c.JSON(http.StatusOK, gin.H{"message": "備考を更新しました"})
}

// PostInvoiceResubmitComment handles POST /dashboard/invoice/resubmit-comment
func PostInvoiceResubmitComment(c *gin.Context) {
	var req models.InvoiceResubmitComment
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Create(&req)
	c.JSON(http.StatusCreated, gin.H{"comment": req})
}
