package handlers

import (
	"kms_golang/database"
	"kms_golang/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

// GetPaymentList handles GET /dashboard/payment/list
func GetPaymentList(c *gin.Context) {
	var slips []models.PaymentSlip
	db := database.DB.Where("delete_flag = 0")

	if propCd := c.Query("prop_cd"); propCd != "" {
		db = db.Where("prop_cd = ?", propCd)
	}
	if customerCd := c.Query("customer_cd"); customerCd != "" {
		db = db.Where("customer_cd = ?", customerCd)
	}
	if yearMonth := c.Query("year_month"); yearMonth != "" {
		db = db.Where("payment_month = ?", yearMonth)
	}

	if err := db.Order("id DESC").Find(&slips).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"payment_slips": slips})
}

// GetPaymentDetail handles GET /dashboard/payment/:number
func GetPaymentDetail(c *gin.Context) {
	number := c.Param("number")
	var slip models.PaymentSlip
	if err := database.DB.Where("payment_number = ? AND delete_flag = 0", number).First(&slip).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "支払伝票が見つかりません"})
		return
	}

	var details []models.PaymentDetail
	database.DB.Where("payment_number = ?", number).Find(&details)

	c.JSON(http.StatusOK, gin.H{"payment_slip": slip, "payment_details": details})
}

// PostPaymentCreate handles POST /dashboard/payment/create
func PostPaymentCreate(c *gin.Context) {
	var req models.PaymentSlip
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	delFlagPay := int16(0)
	req.DeleteFlag = &delFlagPay
	if err := database.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"payment_slip": req})
}

// PutPaymentUpdate handles PUT /dashboard/payment/:number
func PutPaymentUpdate(c *gin.Context) {
	number := c.Param("number")
	nid, nerr := strconv.ParseUint(number, 10, 64)
	if nerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "無効なID"})
		return
	}
	var slip models.PaymentSlip
	if err := database.DB.Where("id = ? AND delete_flag = 0", nid).First(&slip).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "支払伝票が見つかりません"})
		return
	}

	if err := c.ShouldBindJSON(&slip); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	slip.ID = uint(nid)
	if err := database.DB.Save(&slip).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"payment_slip": slip})
}

// DeletePayment handles DELETE /dashboard/payment/:number (soft delete)
func DeletePayment(c *gin.Context) {
	number := c.Param("number")
	if err := database.DB.Model(&models.PaymentSlip{}).Where("payment_number = ?", number).Update("delete_flag", 1).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "支払伝票を削除しました"})
}

// GetDepositList handles GET /dashboard/deposit/list
func GetDepositList(c *gin.Context) {
	var slips []models.DepositSlip
	db := database.DB.Where("delete_flag = 0")

	if propCd := c.Query("prop_cd"); propCd != "" {
		db = db.Where("prop_cd = ?", propCd)
	}
	if customerCd := c.Query("customer_cd"); customerCd != "" {
		db = db.Where("customer_cd = ?", customerCd)
	}

	if err := db.Order("deposit_number DESC").Find(&slips).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deposit_slips": slips})
}

// GetDepositDetail handles GET /dashboard/deposit/:number
func GetDepositDetail(c *gin.Context) {
	number := c.Param("number")
	var slip models.DepositSlip
	if err := database.DB.Where("deposit_number = ? AND delete_flag = 0", number).First(&slip).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "入金伝票が見つかりません"})
		return
	}

	var details []models.DepositDetail
	database.DB.Where("deposit_number = ?", number).Find(&details)

	c.JSON(http.StatusOK, gin.H{"deposit_slip": slip, "deposit_details": details})
}

// PostDepositCreate handles POST /dashboard/deposit/create
func PostDepositCreate(c *gin.Context) {
	var req models.DepositSlip
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	delFlagDep := int16(0)
	req.DeleteFlag = &delFlagDep
	if err := database.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"deposit_slip": req})
}

// GetCashBackList handles GET /dashboard/cashback/list
func GetCashBackList(c *gin.Context) {
	var cashbacks []models.CashBack
	db := database.DB.Where("delete_flag = 0")

	if propCd := c.Query("prop_cd"); propCd != "" {
		db = db.Where("prop_cd = ?", propCd)
	}

	if err := db.Order("id DESC").Find(&cashbacks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"cashbacks": cashbacks})
}

// PutDepositUpdate handles PUT /dashboard/deposit/:number
func PutDepositUpdate(c *gin.Context) {
	number := c.Param("number")
	nid, nerr := strconv.ParseUint(number, 10, 64)
	if nerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "無効なID"})
		return
	}
	var slip models.DepositSlip
	if err := database.DB.Where("id = ? AND delete_flag = 0", nid).First(&slip).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "入金伝票が見つかりません"})
		return
	}
	if err := c.ShouldBindJSON(&slip); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	slip.ID = uint(nid)
	if err := database.DB.Save(&slip).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deposit_slip": slip})
}

// DeleteDeposit handles DELETE /dashboard/deposit/:number (soft delete)
func DeleteDeposit(c *gin.Context) {
	number := c.Param("number")
	if err := database.DB.Model(&models.DepositSlip{}).Where("deposit_number = ?", number).Update("delete_flag", 1).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "入金伝票を削除しました"})
}

// GetDepositFreeInputList handles GET /dashboard/deposit-free-input/list
func GetDepositFreeInputList(c *gin.Context) {
	var slips []models.DepositSlipFreeInput
	db := database.DB.Where("delete_flag = 0")
	if customerCd := c.Query("customer_cd"); customerCd != "" {
		db = db.Where("customer_cd = ?", customerCd)
	}
	if err := db.Order("id DESC").Find(&slips).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deposit_free_inputs": slips})
}

// PostDepositFreeInputCreate handles POST /dashboard/deposit-free-input/create
func PostDepositFreeInputCreate(c *gin.Context) {
	var req models.DepositSlipFreeInput
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	if err := database.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"deposit_free_input": req})
}

// GetPaymentDropdownCustomers handles GET /dashboard/payment/dropdown-customers
func GetPaymentDropdownCustomers(c *gin.Context) {
	var customers []models.Customer
	database.DB.Where("delete_flag = 0").Select("customer_cd, customer_name").Find(&customers)
	for i := range customers {
		customers[i].CustomerPassword = nil
	}
	c.JSON(http.StatusOK, gin.H{"customers": customers})
}

// GetPaymentTermList handles GET /dashboard/payment/term-list
func GetPaymentTermList(c *gin.Context) {
	var closings []models.ClosingDayReq
	database.DB.Find(&closings)
	c.JSON(http.StatusOK, gin.H{"closing_days": closings})
}

// GetPaymentDetailFull handles GET /dashboard/payment/:number/full
func GetPaymentDetailFull(c *gin.Context) {
	number := c.Param("number")
	var slip models.PaymentSlip
	if err := database.DB.Where("payment_number = ? AND delete_flag = 0", number).First(&slip).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "支払伝票が見つかりません"})
		return
	}

	var details []models.PaymentDetail
	database.DB.Where("payment_number = ?", number).Find(&details)

	var propBasic models.PropBasic
	if slip.PropCd != nil {
		database.DB.Where("prop_cd = ?", *slip.PropCd).First(&propBasic)
	}

	var customer models.Customer
	if slip.CustomerCd != nil {
		database.DB.Where("customer_cd = ?", *slip.CustomerCd).First(&customer)
		customer.CustomerPassword = nil
	}

	c.JSON(http.StatusOK, gin.H{
		"payment_slip":    slip,
		"payment_details": details,
		"prop_basic":      propBasic,
		"customer":        customer,
	})
}

// PostPaymentConfirm handles POST /dashboard/payment/confirm
func PostPaymentConfirm(c *gin.Context) {
	var req struct {
		PaymentNumber string `json:"payment_number" binding:"required"`
		ConfirmFlg    int16  `json:"confirm_flg"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Model(&models.PaymentSlip{}).Where("payment_number = ?", req.PaymentNumber).Update("confirm_flg", req.ConfirmFlg)
	c.JSON(http.StatusOK, gin.H{"message": "確認しました"})
}

// PostPaymentApproval handles POST /dashboard/payment/approval
func PostPaymentApproval(c *gin.Context) {
	var req struct {
		PaymentNumber string `json:"payment_number" binding:"required"`
		ApprovalFlg   int16  `json:"approval_flg"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Model(&models.PaymentSlip{}).Where("payment_number = ?", req.PaymentNumber).Update("approval_flg", req.ApprovalFlg)
	c.JSON(http.StatusOK, gin.H{"message": "承認しました"})
}

// PostPaymentUpdateBiko handles POST /dashboard/payment/update-biko
func PostPaymentUpdateBiko(c *gin.Context) {
	var req struct {
		PaymentNumber string  `json:"payment_number" binding:"required"`
		Biko          *string `json:"biko"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Model(&models.PaymentSlip{}).Where("payment_number = ?", req.PaymentNumber).Update("biko", req.Biko)
	c.JSON(http.StatusOK, gin.H{"message": "備考を更新しました"})
}

// PostPaymentSendInvoice handles POST /dashboard/payment/send-invoice
func PostPaymentSendInvoice(c *gin.Context) {
	var req struct {
		PaymentNumber string `json:"payment_number" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Model(&models.PaymentSlip{}).Where("payment_number = ?", req.PaymentNumber).Update("send_invoice_flg", 1)
	c.JSON(http.StatusOK, gin.H{"message": "請求書を送信しました"})
}

// GetPaymentExportExcel handles GET /dashboard/payment/export-excel
func GetPaymentExportExcel(c *gin.Context) {
	var slips []models.PaymentSlip
	db := database.DB.Where("delete_flag = 0")
	if propCd := c.Query("prop_cd"); propCd != "" {
		db = db.Where("prop_cd = ?", propCd)
	}
	if yearMonth := c.Query("year_month"); yearMonth != "" {
		db = db.Where("payment_month = ?", yearMonth)
	}
	db.Order("id DESC").Find(&slips)

	f := excelize.NewFile()
	sheet := "Sheet1"
	headers := []string{"支払番号", "物件CD", "顧客CD", "年月", "合計金額", "備考"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}

	for row, s := range slips {
		r := row + 2
		vals := []interface{}{
			s.ID, ptrStr(s.PropCd), ptrStr(s.CustomerCd),
			ptrStr(s.PaymentMonth), s.TotalAmount, s.TaxAmount,
		}
		for col, v := range vals {
			cell, _ := excelize.CoordinatesToCellName(col+1, r)
			f.SetCellValue(sheet, cell, v)
		}
	}

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=payment_list.xlsx")
	f.Write(c.Writer)
}

// GetPaymentSekosakiTimeline handles GET /dashboard/payment/sekosaki-timeline
func GetPaymentSekosakiTimeline(c *gin.Context) {
	sekosakiCd := c.Query("sekosaki_cd")
	var slips []models.PaymentSlip
	database.DB.Where("sekosaki_cd = ? AND delete_flag = 0", sekosakiCd).Order("payment_month DESC").Find(&slips)
	c.JSON(http.StatusOK, gin.H{"payment_slips": slips})
}

// GetPaymentUserPayment handles GET /dashboard/payment/user-payment
func GetPaymentUserPayment(c *gin.Context) {
	userID := c.GetString("user_id")
	var slips []models.PaymentSlip
	database.DB.Where("user_id = ? AND delete_flag = 0", userID).Order("payment_number DESC").Find(&slips)
	c.JSON(http.StatusOK, gin.H{"payment_slips": slips})
}

// GetPaymentSekosakiPayment handles GET /dashboard/payment/sekosaki-payment
func GetPaymentSekosakiPayment(c *gin.Context) {
	sekosakiCd := c.Query("sekosaki_cd")
	var slips []models.PaymentSlip
	database.DB.Where("sekosaki_cd = ? AND delete_flag = 0", sekosakiCd).Order("payment_number DESC").Find(&slips)
	c.JSON(http.StatusOK, gin.H{"payment_slips": slips})
}

// GetPaymentDepositManagement handles GET /dashboard/payment/deposit-management
func GetPaymentDepositManagement(c *gin.Context) {
	customerCd := c.Query("customer_cd")
	yearMonth := c.Query("year_month")

	var deposits []models.DepositSlip
	db := database.DB.Where("delete_flag = 0")
	if customerCd != "" {
		db = db.Where("customer_cd = ?", customerCd)
	}
	if yearMonth != "" {
		db = db.Where("deposit_month = ?", yearMonth)
	}
	db.Order("id DESC").Find(&deposits)

	var details []models.DepositDetail
	for _, d := range deposits {
		var dd []models.DepositDetail
		database.DB.Where("deposit_slip_id = ?", d.ID).Find(&dd)
		details = append(details, dd...)
	}

	c.JSON(http.StatusOK, gin.H{"deposit_slips": deposits, "deposit_details": details})
}

// PostPaymentDetailCreate handles POST /dashboard/payment/detail/create
func PostPaymentDetailCreate(c *gin.Context) {
	var req models.PaymentDetail
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	if err := database.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"payment_detail": req})
}

// PostPaymentDetailUpdate handles POST /dashboard/payment/detail/update
func PostPaymentDetailUpdate(c *gin.Context) {
	var req models.PaymentDetail
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	if req.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "IDが必要です"})
		return
	}
	if err := database.DB.Save(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"payment_detail": req})
}

// PostDepositDetailCreate handles POST /dashboard/deposit/detail/create
func PostDepositDetailCreate(c *gin.Context) {
	var req models.DepositDetail
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	if err := database.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"deposit_detail": req})
}

// PostDepositDetailUpdate handles POST /dashboard/deposit/detail/update
func PostDepositDetailUpdate(c *gin.Context) {
	var req models.DepositDetail
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	if req.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "IDが必要です"})
		return
	}
	if err := database.DB.Save(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deposit_detail": req})
}

// GetCashBackDetail handles GET /dashboard/cashback/:id
func GetCashBackDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "無効なID"})
		return
	}
	var cb models.CashBack
	if err := database.DB.Where("id = ? AND delete_flag = 0", id).First(&cb).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "見つかりません"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"cashback": cb})
}

// PostCashBackCreate handles POST /dashboard/cashback/create
func PostCashBackCreate(c *gin.Context) {
	var req models.CashBack
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
	c.JSON(http.StatusCreated, gin.H{"cashback": req})
}

// PostCashBackUpdate handles POST /dashboard/cashback/update
func PostCashBackUpdate(c *gin.Context) {
	var req models.CashBack
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	if req.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "IDが必要です"})
		return
	}
	if err := database.DB.Save(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"cashback": req})
}

// GetTaxRateSettings handles GET /dashboard/payment/tax-rate
func GetTaxRateSettings(c *gin.Context) {
	var settings []models.TaxrateSetting
	database.DB.Find(&settings)
	c.JSON(http.StatusOK, gin.H{"tax_rate_settings": settings})
}

// GetClosingDayPayments handles GET /dashboard/payment/closing-day
func GetClosingDayPayments(c *gin.Context) {
	yearMonth := c.Query("year_month")
	customerCd := c.Query("customer_cd")

	var payments []models.ClosingDayPayments
	db := database.DB
	if yearMonth != "" {
		db = db.Where("year_month = ?", yearMonth)
	}
	if customerCd != "" {
		db = db.Where("customer_cd = ?", customerCd)
	}
	db.Find(&payments)

	var details []models.ClosingDayPaymentDetail
	for _, p := range payments {
		var dd []models.ClosingDayPaymentDetail
		database.DB.Where("closing_day_payment_id = ?", p.ID).Find(&dd)
		details = append(details, dd...)
	}
	c.JSON(http.StatusOK, gin.H{"closing_day_payments": payments, "closing_day_details": details})
}

// GetInvoiceAuthorityList handles GET /dashboard/invoice/authority-list
func GetInvoiceAuthorityList(c *gin.Context) {
	var auths []models.InvoiceAuthority
	database.DB.Find(&auths)
	c.JSON(http.StatusOK, gin.H{"invoice_authorities": auths})
}

// PostInvoiceUpdateAuthority handles POST /dashboard/invoice/update-authority
func PostInvoiceUpdateAuthority(c *gin.Context) {
	var req models.InvoiceAuthority
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	if req.ID > 0 {
		database.DB.Save(&req)
	} else {
		database.DB.Create(&req)
	}
	c.JSON(http.StatusOK, gin.H{"invoice_authority": req})
}

// GetWithHoldingList handles GET /dashboard/payment/withholding
func GetWithHoldingList(c *gin.Context) {
	var whs []models.WithHolding
	database.DB.Find(&whs)
	c.JSON(http.StatusOK, gin.H{"withholdings": whs})
}
