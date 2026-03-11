package handlers

import (
	"kms_golang/database"
	"kms_golang/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

// GetPaymentKeshikomiList handles GET /dashboard/payment-keshikomi/list
func GetPaymentKeshikomiList(c *gin.Context) {
	type KeshikomiRow struct {
		CustomerCd   string `json:"customer_cd"`
		CustomerName string `json:"customer_name"`
		CustomerKana string `json:"customer_kana"`
	}

	db := database.DB.Table("payment_keshikomis").
		Select("DISTINCT payment_keshikomis.customer_cd, customers.customer_name, customers.customer_kana").
		Joins("JOIN customers ON customers.customer_cd = payment_keshikomis.customer_cd").
		Where("payment_keshikomis.delete_flag <> ?", "1")

	if customerCd := c.Query("customer_cd"); customerCd != "" {
		db = db.Where("payment_keshikomis.customer_cd = ?", customerCd)
	}
	if checkKeshikomi := c.Query("check_keshikomi"); checkKeshikomi == "1" {
		db = db.Where("payment_keshikomis.keshikomi_day IS NULL")
	}
	if monthStart := c.Query("month_start"); monthStart != "" {
		db = db.Where("payment_keshikomis.invoice_month >= ?", monthStart)
	}
	if monthEnd := c.Query("month_end"); monthEnd != "" {
		db = db.Where("payment_keshikomis.invoice_month <= ?", monthEnd)
	}
	if cusKana := c.Query("cus_kana"); cusKana != "" {
		db = db.Where("customers.customer_kana LIKE ?", "%"+cusKana+"%")
	}

	var rows []KeshikomiRow
	db.Order("customers.customer_kana ASC").Scan(&rows)

	c.JSON(http.StatusOK, gin.H{"data": rows})
}

// PostSearchPaymentKeshikomi handles POST /dashboard/payment-keshikomi/search
func PostSearchPaymentKeshikomi(c *gin.Context) {
	var req struct {
		Page           int      `json:"page"`
		PerPage        int      `json:"per_page"`
		CustomerCd     string   `json:"customer_cd"`
		CheckKeshikomi int      `json:"check_keshikomi"`
		MonthStart     string   `json:"month_start"`
		MonthEnd       string   `json:"month_end"`
		CusKana        string   `json:"cus_kana"`
		CoCode         []string `json:"co_code"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	if req.Page < 1 {
		req.Page = 1
	}
	if req.PerPage < 1 {
		req.PerPage = 10
	}

	db := database.DB.Table("payment_keshikomis").
		Select(`prop_basics.prop_name, payment_keshikomis.fee, customers.customer_name,
			payment_keshikomis.biko, payment_keshikomis.`+"`index`"+`, payment_keshikomis.renban,
			payment_keshikomis.prop_cd, payment_keshikomis.deposit, payment_keshikomis.work_name,
			payment_keshikomis.customer_cd, payment_keshikomis.work_amount,
			payment_keshikomis.billing_date, payment_keshikomis.work_devision,
			payment_keshikomis.invoice_month, payment_keshikomis.keshikomi_day,
			payment_keshikomis.payment_amount, payment_keshikomis.estimate_number,
			payment_keshikomis.with_holding_id`).
		Joins("JOIN customers ON customers.customer_cd = payment_keshikomis.customer_cd").
		Joins("LEFT JOIN prop_basics ON prop_basics.prop_cd = payment_keshikomis.prop_cd").
		Where("payment_keshikomis.delete_flag <> ?", "1")

	if req.CustomerCd != "" {
		db = db.Where("payment_keshikomis.customer_cd = ?", req.CustomerCd)
	}
	if req.CheckKeshikomi == 1 {
		db = db.Where("payment_keshikomis.keshikomi_day IS NULL")
	}
	if req.MonthStart != "" {
		db = db.Where("payment_keshikomis.invoice_month >= ?", req.MonthStart)
	}
	if req.MonthEnd != "" {
		db = db.Where("payment_keshikomis.invoice_month <= ?", req.MonthEnd)
	}
	if req.CusKana != "" {
		db = db.Where("customers.customer_kana LIKE ?", "%"+req.CusKana+"%")
	}

	var total int64
	db.Count(&total)

	offset := (req.Page - 1) * req.PerPage
	var results []map[string]interface{}
	db.Offset(offset).Limit(req.PerPage).Scan(&results)

	c.JSON(http.StatusOK, gin.H{"results": results, "countResults": total})
}

// PostUpdatePaymentKeshikomi handles POST /dashboard/payment-keshikomi/update
func PostUpdatePaymentKeshikomi(c *gin.Context) {
	var req struct {
		Data []struct {
			CustomerCd     string  `json:"customer_cd"`
			InvoiceMonth   string  `json:"invoice_month"`
			Index          int     `json:"index"`
			Renban         int     `json:"renban"`
			Biko           string  `json:"biko"`
			Fee            float64 `json:"fee"`
			Deposit        float64 `json:"deposit"`
			PaymentAmount  float64 `json:"payment_amount"`
			KeshikomiDay   *string `json:"keshikomi_day"`
			EstimateNumber string  `json:"estimate_number"`
		} `json:"data"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	loggedUserID := c.GetString("user_id")

	tx := database.DB.Begin()
	for _, item := range req.Data {
		updates := map[string]interface{}{
			"biko":            item.Biko,
			"with_holding_id": nil,
			"keshikomi_day":   item.KeshikomiDay,
			"update_user":     loggedUserID,
			"fee":             item.Fee,
			"deposit":         item.Deposit,
			"payment_amount":  item.PaymentAmount,
		}

		tx.Table("payment_keshikomis").
			Where("customer_cd = ? AND invoice_month = ? AND `index` = ? AND renban = ?",
				item.CustomerCd, item.InvoiceMonth, item.Index, item.Renban).
			Updates(updates)
	}
	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"message": "更新しました"})
}

// GetPaymentKeshikomiExportExcel handles GET /dashboard/payment-keshikomi/export-excel
func GetPaymentKeshikomiExportExcel(c *gin.Context) {
	customerCd := c.Query("customer_cd")
	monthStart := c.Query("month_start")
	monthEnd := c.Query("month_end")

	db := database.DB.Table("payment_keshikomis").
		Select("payment_keshikomis.*, customers.customer_name, prop_basics.prop_name").
		Joins("JOIN customers ON customers.customer_cd = payment_keshikomis.customer_cd").
		Joins("LEFT JOIN prop_basics ON prop_basics.prop_cd = payment_keshikomis.prop_cd").
		Where("payment_keshikomis.delete_flag <> ?", "1")

	if customerCd != "" {
		db = db.Where("payment_keshikomis.customer_cd = ?", customerCd)
	}
	if monthStart != "" {
		db = db.Where("payment_keshikomis.invoice_month >= ?", monthStart)
	}
	if monthEnd != "" {
		db = db.Where("payment_keshikomis.invoice_month <= ?", monthEnd)
	}

	var results []map[string]interface{}
	db.Scan(&results)

	f := excelize.NewFile()
	sheet := "Sheet1"
	headers := []string{"顧客コード", "顧客名", "物件名", "請求月", "消込日", "金額", "入金額", "支払額", "備考"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}
	for i, r := range results {
		row := i + 2
		setCellVal(f, sheet, 1, row, r["customer_cd"])
		setCellVal(f, sheet, 2, row, r["customer_name"])
		setCellVal(f, sheet, 3, row, r["prop_name"])
		setCellVal(f, sheet, 4, row, r["invoice_month"])
		setCellVal(f, sheet, 5, row, r["keshikomi_day"])
		setCellVal(f, sheet, 6, row, r["fee"])
		setCellVal(f, sheet, 7, row, r["deposit"])
		setCellVal(f, sheet, 8, row, r["payment_amount"])
		setCellVal(f, sheet, 9, row, r["biko"])
	}

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=payment_keshikomi.xlsx")
	f.Write(c.Writer)
}

func setCellVal(f *excelize.File, sheet string, col, row int, val interface{}) {
	cell, _ := excelize.CoordinatesToCellName(col, row)
	if val != nil {
		f.SetCellValue(sheet, cell, val)
	}
}

// GetPaymentSearchList handles GET /dashboard/payment-search/list
func GetPaymentSearchList(c *gin.Context) {
	var customers []models.Customer
	database.DB.Where("delete_flag = 0").Select("customer_cd, customer_name").Find(&customers)
	c.JSON(http.StatusOK, gin.H{"customers": customers})
}

// ---- SekosakiPayment handlers ----

// GetSekosakiPaymentList handles GET /dashboard/sekosaki-payment/list
func GetSekosakiPaymentList(c *gin.Context) {
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

	db := database.DB.Table("payment_slip").Where("delete_flag = 0")

	if sekosakiCd := c.Query("sekosaki_cd"); sekosakiCd != "" {
		db = db.Where("sekosaki_cd = ?", sekosakiCd)
	}
	if yearMonth := c.Query("year_month"); yearMonth != "" {
		db = db.Where("payment_month = ?", yearMonth)
	}

	var total int64
	db.Count(&total)

	var slips []models.PaymentSlip
	offset := (page - 1) * perPage
	db.Offset(offset).Limit(perPage).Order("id DESC").Find(&slips)

	c.JSON(http.StatusOK, gin.H{"payment_slips": slips, "total": total, "page": page, "per_page": perPage})
}

// PostSekosakiPaymentRegister handles POST /dashboard/sekosaki-payment/register
func PostSekosakiPaymentRegister(c *gin.Context) {
	var req models.PaymentSlip
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
	c.JSON(http.StatusCreated, gin.H{"payment_slip": req})
}

// GetSekosakiPaymentDetail handles GET /dashboard/sekosaki-payment/:id
func GetSekosakiPaymentDetail(c *gin.Context) {
	id := c.Param("id")
	var slip models.PaymentSlip
	if err := database.DB.Where("id = ? AND delete_flag = 0", id).First(&slip).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "支払伝票が見つかりません"})
		return
	}

	var details []models.PaymentDetail
	database.DB.Where("payment_slip_id = ?", id).Find(&details)

	c.JSON(http.StatusOK, gin.H{"payment_slip": slip, "payment_details": details})
}

// GetSekosakiPaymentTerm handles GET /dashboard/sekosaki-payment/term
func GetSekosakiPaymentTerm(c *gin.Context) {
	sekosakiCd := c.Query("sekosaki_cd")
	month := c.Query("month")

	var slips []models.PaymentSlip
	db := database.DB.Where("delete_flag = 0")
	if sekosakiCd != "" {
		db = db.Where("sekosaki_cd = ?", sekosakiCd)
	}
	if month != "" {
		db = db.Where("payment_month = ?", month)
	}
	db.Find(&slips)

	c.JSON(http.StatusOK, gin.H{"payment_slips": slips})
}
