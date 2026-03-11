package handlers

import (
	"kms_golang/database"
	"kms_golang/models"
	"kms_golang/services"
	"kms_golang/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

// GetCustomerList handles GET /dashboard/customer/list
func GetCustomerList(c *gin.Context) {
	var customers []models.Customer
	db := database.DB.Where("delete_flag = 0")

	if name := c.Query("customer_name"); name != "" {
		db = db.Where("customer_name LIKE ?", "%"+name+"%")
	}
	if cd := c.Query("customer_cd"); cd != "" {
		db = db.Where("customer_cd LIKE ?", "%"+cd+"%")
	}
	if phone := c.Query("tel"); phone != "" {
		db = db.Where("tel LIKE ?", "%"+phone+"%")
	}

	if err := db.Order("customer_cd ASC").Find(&customers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Mask passwords
	for i := range customers {
		customers[i].CustomerPassword = nil
	}

	c.JSON(http.StatusOK, gin.H{"customers": customers})
}

// GetCustomerDetail handles GET /dashboard/customer/:cd
func GetCustomerDetail(c *gin.Context) {
	cd := c.Param("cd")
	var customer models.Customer
	if err := database.DB.Where("customer_cd = ? AND delete_flag = 0", cd).First(&customer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "顧客が見つかりません"})
		return
	}
	customer.CustomerPassword = nil

	// Load related personnel
	var personnel []models.CustomerPersonnel
	database.DB.Where("customer_cd = ?", cd).Find(&personnel)

	c.JSON(http.StatusOK, gin.H{"customer": customer, "personnel": personnel})
}

// PostCustomerCreate handles POST /dashboard/customer/create
func PostCustomerCreate(c *gin.Context) {
	var req struct {
		CustomerName string  `json:"customer_name" binding:"required"`
		CustomerKana string  `json:"customer_kana"`
		PostCode     string  `json:"post_code"`
		PrefectureID *string `json:"prefecture_id"`
		CityID       *string `json:"city_id"`
		TownID       *string `json:"town_id"`
		BlockName    string  `json:"block_name"`
		Tel          string  `json:"tel"`
		Fax          string  `json:"fax"`
		CustomerType *int16  `json:"customer_type"`
		LoginID      string  `json:"login_id"`
		Password     string  `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	cd, err := utils.GenerateCustomerID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	customerName := req.CustomerName
	customerKana := req.CustomerKana
	postCode := req.PostCode
	blockName := req.BlockName
	tel := req.Tel
	fax := req.Fax
	loginID := req.LoginID
	delFlag := int16(0)

	customer := models.Customer{
		CustomerCd:      cd,
		CustomerName:    &customerName,
		CustomerKana:    &customerKana,
		PostCode:        &postCode,
		BlockName:       &blockName,
		Tel:             &tel,
		Fax:             &fax,
		PrefectureID:    req.PrefectureID,
		CityID:          req.CityID,
		TownID:          req.TownID,
		CustomerType:    req.CustomerType,
		CustomerLoginID: &loginID,
		DeleteFlag:      &delFlag,
	}

	if req.Password != "" {
		hashed, herr := services.HashPassword(req.Password)
		if herr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": herr.Error()})
			return
		}
		customer.CustomerPassword = &hashed
	}

	if err := database.DB.Create(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	customer.CustomerPassword = nil
	c.JSON(http.StatusCreated, gin.H{"customer": customer})
}

// PutCustomerUpdate handles PUT /dashboard/customer/:cd
func PutCustomerUpdate(c *gin.Context) {
	cd := c.Param("cd")
	var customer models.Customer
	if err := database.DB.Where("customer_cd = ? AND delete_flag = 0", cd).First(&customer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "顧客が見つかりません"})
		return
	}

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	customer.CustomerCd = cd
	if err := database.DB.Save(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	customer.CustomerPassword = nil
	c.JSON(http.StatusOK, gin.H{"customer": customer})
}

// DeleteCustomer handles DELETE /dashboard/customer/:cd (soft delete)
func DeleteCustomer(c *gin.Context) {
	cd := c.Param("cd")
	delFlag := int16(1)
	if err := database.DB.Model(&models.Customer{}).Where("customer_cd = ?", cd).Update("delete_flag", delFlag).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "顧客を削除しました"})
}

// GetCustomerExportExcel handles GET /dashboard/customer/export-excel
func GetCustomerExportExcel(c *gin.Context) {
	var customers []models.Customer
	database.DB.Where("delete_flag = 0").Order("customer_cd ASC").Find(&customers)

	f := excelize.NewFile()
	sheet := "Sheet1"
	headers := []string{"顧客CD", "顧客名", "顧客名カナ", "郵便番号", "住所", "電話番号", "FAX"}
	for i, h := range headers {
		col := string(rune('A' + i))
		f.SetCellValue(sheet, col+"1", h)
	}
	for i, cust := range customers {
		row := strconv.Itoa(i + 2)
		f.SetCellValue(sheet, "A"+row, cust.CustomerCd)
		if cust.CustomerName != nil {
			f.SetCellValue(sheet, "B"+row, *cust.CustomerName)
		}
		if cust.CustomerKana != nil {
			f.SetCellValue(sheet, "C"+row, *cust.CustomerKana)
		}
		if cust.PostCode != nil {
			f.SetCellValue(sheet, "D"+row, *cust.PostCode)
		}
		if cust.BlockName != nil {
			f.SetCellValue(sheet, "E"+row, *cust.BlockName)
		}
		if cust.Tel != nil {
			f.SetCellValue(sheet, "F"+row, *cust.Tel)
		}
		if cust.Fax != nil {
			f.SetCellValue(sheet, "G"+row, *cust.Fax)
		}
	}

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=customers.xlsx")
	f.Write(c.Writer)
}

// GetCustomerEstimateList handles GET /dashboard/customer/:cd/estimates
func GetCustomerEstimateList(c *gin.Context) {
	cd := c.Param("cd")
	var estimates []models.Estimate
	database.DB.Where("customer_cd = ? AND delete_flag = 0", cd).Order("estimate_number DESC").Find(&estimates)
	c.JSON(http.StatusOK, gin.H{"estimates": estimates})
}

// GetCustomerInvoiceList handles GET /dashboard/customer/:cd/invoices
func GetCustomerInvoiceList(c *gin.Context) {
	cd := c.Param("cd")
	var invoices []models.CusInvoiceDetail
	database.DB.Where("customer_cd = ? AND delete_flag = 0", cd).Order("id DESC").Find(&invoices)
	c.JSON(http.StatusOK, gin.H{"invoices": invoices})
}

// GetCustomerContractList handles GET /dashboard/customer/:cd/contracts
func GetCustomerContractList(c *gin.Context) {
	cd := c.Param("cd")
	var props []models.PropBasic
	database.DB.Where("customer_cd = ? AND delete_flag = 0", cd).Find(&props)
	c.JSON(http.StatusOK, gin.H{"contracts": props})
}

// GetCustomerContractDetail handles GET /dashboard/customer/:cd/contract/:prop_cd
func GetCustomerContractDetail(c *gin.Context) {
	propCd := c.Param("prop_cd")
	var prop models.PropBasic
	if err := database.DB.Where("prop_cd = ? AND delete_flag = 0", propCd).First(&prop).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "契約が見つかりません"})
		return
	}
	var orders []models.Order
	database.DB.Where("prop_cd = ?", propCd).Find(&orders)
	c.JSON(http.StatusOK, gin.H{"property": prop, "orders": orders})
}

// PostCustomerChangePassword handles POST /dashboard/customer/:cd/change-password
func PostCustomerChangePassword(c *gin.Context) {
	cd := c.Param("cd")
	var req struct {
		Password string `json:"password" binding:"required,min=6"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	hashed, err := services.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if err := database.DB.Model(&models.Customer{}).Where("customer_cd = ?", cd).Update("customer_password", hashed).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "パスワードを変更しました"})
}

// GetCustomerCheckLoginID handles GET /dashboard/customer/check-login-id
func GetCustomerCheckLoginID(c *gin.Context) {
	loginID := c.Query("login_id")
	customerCd := c.Query("customer_cd")
	var count int64
	db := database.DB.Model(&models.Customer{}).Where("customer_login_id = ? AND delete_flag = 0", loginID)
	if customerCd != "" {
		db = db.Where("customer_cd <> ?", customerCd)
	}
	db.Count(&count)
	c.JSON(http.StatusOK, gin.H{"valid": count == 0})
}

// PostCustomerContractConfirm handles POST /dashboard/customer/:cd/contract-confirm
func PostCustomerContractConfirm(c *gin.Context) {
	cd := c.Param("cd")
	var req struct {
		PropCd string `json:"prop_cd" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	// Update confirmation status
	database.DB.Model(&models.PropBasic{}).Where("prop_cd = ? AND customer_cd = ?", req.PropCd, cd).
		Update("contract_confirm", 1)
	c.JSON(http.StatusOK, gin.H{"message": "契約確認しました"})
}

// PostCustomerSetDatePayment handles POST /dashboard/customer/set-date-payment
func PostCustomerSetDatePayment(c *gin.Context) {
	var req struct {
		CustomerCd string `json:"customer_cd" binding:"required"`
		PaymentDay *int   `json:"payment_day"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	database.DB.Model(&models.Customer{}).Where("customer_cd = ?", req.CustomerCd).
		Update("payment_day", req.PaymentDay)
	c.JSON(http.StatusOK, gin.H{"message": "支払日を設定しました"})
}

// GetCustomerPreviewInvoice handles GET /dashboard/customer/:cd/preview-invoice
func GetCustomerPreviewInvoice(c *gin.Context) {
	cd := c.Param("cd")
	invoiceMonth := c.Query("invoice_month")
	var invoiceDetails []models.CusInvoiceDetail
	db := database.DB.Where("customer_cd = ? AND delete_flag = 0", cd)
	if invoiceMonth != "" {
		db = db.Where("invoice_month = ?", invoiceMonth)
	}
	db.Find(&invoiceDetails)

	var customer models.Customer
	database.DB.Where("customer_cd = ?", cd).First(&customer)
	customer.CustomerPassword = nil

	c.JSON(http.StatusOK, gin.H{"customer": customer, "invoice_details": invoiceDetails})
}

// GetCustomerExportInvoiceExcel handles GET /dashboard/customer/:cd/export-invoice-excel
func GetCustomerExportInvoiceExcel(c *gin.Context) {
	cd := c.Param("cd")
	invoiceMonth := c.Query("invoice_month")
	var invoiceDetails []models.CusInvoiceDetail
	db := database.DB.Where("customer_cd = ? AND delete_flag = 0", cd)
	if invoiceMonth != "" {
		db = db.Where("invoice_month = ?", invoiceMonth)
	}
	db.Find(&invoiceDetails)

	f := excelize.NewFile()
	sheet := "Sheet1"
	f.SetCellValue(sheet, "A1", "請求明細")
	for i, d := range invoiceDetails {
		row := strconv.Itoa(i + 2)
		f.SetCellValue(sheet, "A"+row, d.ID)
		f.SetCellValue(sheet, "B"+row, d.InvoiceMonth)
	}

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=customer_invoice.xlsx")
	f.Write(c.Writer)
}
