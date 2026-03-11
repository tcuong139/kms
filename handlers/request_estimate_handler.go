package handlers

import (
	"kms_golang/database"
	"kms_golang/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

// GetRequestEstimateList handles GET /dashboard/request-estimate/list
func GetRequestEstimateList(c *gin.Context) {
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

	db := database.DB.Table("request_quotation").
		Select("request_quotation.*, reception.prop_name, reception.customer_cd, reception.room_number").
		Joins("JOIN reception ON reception.accept_number = request_quotation.accept_number").
		Joins("JOIN prop_basics ON prop_basics.prop_cd = request_quotation.prop_cd").
		Where("(request_quotation.delete_flg <> 1 OR request_quotation.delete_flg IS NULL)").
		Order("request_quotation.regist_datetime DESC")

	if reqName := c.Query("request_name"); reqName != "" {
		db = db.Where("request_quotation.request_name LIKE ?", "%"+reqName+"%")
	}
	if recoveryDate := c.Query("recovery_date"); recoveryDate != "" {
		db = db.Where("DATE(request_quotation.submission_deadline) = ?", recoveryDate)
	}
	if propName := c.Query("prop_name"); propName != "" {
		db = db.Where("prop_basics.prop_name LIKE ?", "%"+propName+"%")
	}
	if roomNumber := c.Query("room_number"); roomNumber != "" {
		db = db.Where("reception.room_number LIKE ?", "%"+roomNumber+"%")
	}
	if propKana := c.Query("prop_kana"); propKana != "" {
		db = db.Where("prop_basics.prop_kana_name LIKE ?", "%"+propKana+"%")
	}

	var total int64
	db.Count(&total)

	var results []map[string]interface{}
	offset := (page - 1) * perPage
	db.Offset(offset).Limit(perPage).Scan(&results)

	c.JSON(http.StatusOK, gin.H{"request_quotations": results, "total": total, "page": page, "per_page": perPage})
}

// GetRequestEstimateExportExcel handles GET /dashboard/request-estimate/export-excel
func GetRequestEstimateExportExcel(c *gin.Context) {
	db := database.DB.Table("request_quotation").
		Select("request_quotation.*, reception.prop_name, reception.room_number").
		Joins("JOIN reception ON reception.accept_number = request_quotation.accept_number").
		Joins("JOIN prop_basics ON prop_basics.prop_cd = request_quotation.prop_cd").
		Where("(request_quotation.delete_flg <> 1 OR request_quotation.delete_flg IS NULL)").
		Order("request_quotation.regist_datetime DESC")

	if reqName := c.Query("request_name"); reqName != "" {
		db = db.Where("request_quotation.request_name LIKE ?", "%"+reqName+"%")
	}
	if propName := c.Query("prop_name"); propName != "" {
		db = db.Where("prop_basics.prop_name LIKE ?", "%"+propName+"%")
	}

	var results []map[string]interface{}
	db.Scan(&results)

	f := excelize.NewFile()
	sheet := "Sheet1"
	headers := []string{"見積依頼名", "物件名", "部屋番号", "提出期限", "登録日"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}
	for i, r := range results {
		row := i + 2
		setCellValRE(f, sheet, 1, row, r["request_name"])
		setCellValRE(f, sheet, 2, row, r["prop_name"])
		setCellValRE(f, sheet, 3, row, r["room_number"])
		setCellValRE(f, sheet, 4, row, r["submission_deadline"])
		setCellValRE(f, sheet, 5, row, r["regist_datetime"])
	}

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=requestEstimate.xlsx")
	f.Write(c.Writer)
}

func setCellValRE(f *excelize.File, sheet string, col, row int, val interface{}) {
	cell, _ := excelize.CoordinatesToCellName(col, row)
	if val != nil {
		f.SetCellValue(sheet, cell, val)
	}
}

// ---- Seko Request Estimate ----

// GetSekoRequestEstimateList handles GET /dashboard/seko-request-estimate/list
func GetSekoRequestEstimateList(c *gin.Context) {
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

	sekosakiCd := c.GetString("sekosaki_cd") // from JWT for sekosaki guard

	db := database.DB.Table("request_quotation").
		Select("request_quotation.*, prop_basics.prop_name, reception.customer_cd, reception.room_number").
		Joins("JOIN reception ON reception.accept_number = request_quotation.accept_number").
		Joins("JOIN prop_basics ON prop_basics.prop_cd = request_quotation.prop_cd").
		Joins("LEFT JOIN user ON user.user_id = request_quotation.regist_user").
		Where("(request_quotation.delete_flg = 0 OR request_quotation.delete_flg IS NULL)")

	if sekosakiCd != "" {
		db = db.Where("request_quotation.sekosaki_cd = ?", sekosakiCd)
	}
	if keyword := c.Query("keyword"); keyword != "" {
		db = db.Where("request_quotation.request_name LIKE ?", "%"+keyword+"%")
	}
	if recoveryDate := c.Query("recovery_date"); recoveryDate != "" {
		db = db.Where("DATE(request_quotation.submission_deadline) = ?", recoveryDate)
	}
	if propName := c.Query("prop_name"); propName != "" {
		db = db.Where("prop_basics.prop_name LIKE ?", "%"+propName+"%")
	}
	if userName := c.Query("user_name"); userName != "" {
		db = db.Where("user.user_name LIKE ?", "%"+userName+"%")
	}
	if status := c.Query("status"); status != "" {
		if status == "1" {
			db = db.Where("NOT EXISTS (SELECT 1 FROM request_quotation_pdfs WHERE request_quotation_pdfs.request_quotation_id = request_quotation.id)")
		} else {
			db = db.Where("EXISTS (SELECT 1 FROM request_quotation_pdfs WHERE request_quotation_pdfs.request_quotation_id = request_quotation.id)")
		}
	}

	var total int64
	db.Count(&total)

	var results []map[string]interface{}
	offset := (page - 1) * perPage
	db.Offset(offset).Limit(perPage).Order("request_quotation.id ASC").Scan(&results)

	c.JSON(http.StatusOK, gin.H{"request_quotations": results, "total": total, "page": page, "per_page": perPage})
}

// GetSekoRequestEstimateDetail handles GET /dashboard/seko-request-estimate/:id
func GetSekoRequestEstimateDetail(c *gin.Context) {
	id := c.Param("id")

	var rq models.RequestQuotation
	if err := database.DB.Preload("Imgs").Preload("PDFs").Where("id = ?", id).First(&rq).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "見積依頼が見つかりません"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"request_quotation": rq})
}

// PostSekoRequestEstimateRegister handles POST /dashboard/seko-request-estimate/register
func PostSekoRequestEstimateRegister(c *gin.Context) {
	var req struct {
		RQID         uint   `json:"rq_id" binding:"required"`
		SekosakiBiko string `json:"sekosaki_biko"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	database.DB.Model(&models.RequestQuotation{}).Where("id = ?", req.RQID).Updates(map[string]interface{}{
		"sekosaki_biko": req.SekosakiBiko,
		"reply_flg":     1,
	})

	c.JSON(http.StatusOK, gin.H{"message": "登録しました"})
}

// GetSekoRequestEstimateExportExcel handles GET /dashboard/seko-request-estimate/export-excel
func GetSekoRequestEstimateExportExcel(c *gin.Context) {
	sekosakiCd := c.GetString("sekosaki_cd")

	db := database.DB.Table("request_quotation").
		Select("request_quotation.*, prop_basics.prop_name, reception.room_number").
		Joins("JOIN reception ON reception.accept_number = request_quotation.accept_number").
		Joins("JOIN prop_basics ON prop_basics.prop_cd = request_quotation.prop_cd").
		Where("(request_quotation.delete_flg = 0 OR request_quotation.delete_flg IS NULL)")

	if sekosakiCd != "" {
		db = db.Where("request_quotation.sekosaki_cd = ?", sekosakiCd)
	}

	var results []map[string]interface{}
	db.Scan(&results)

	f := excelize.NewFile()
	sheet := "Sheet1"
	headers := []string{"見積依頼名", "物件名", "部屋番号", "提出期限", "回答状況"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}
	for i, r := range results {
		row := i + 2
		setCellValRE(f, sheet, 1, row, r["request_name"])
		setCellValRE(f, sheet, 2, row, r["prop_name"])
		setCellValRE(f, sheet, 3, row, r["room_number"])
		setCellValRE(f, sheet, 4, row, r["submission_deadline"])
		setCellValRE(f, sheet, 5, row, r["reply_flg"])
	}

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=seko_request_estimate.xlsx")
	f.Write(c.Writer)
}
