package handlers

import (
	"kms_golang/database"
	"kms_golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetEstimateList handles GET /dashboard/estimate/list
func GetEstimateList(c *gin.Context) {
	var estimates []models.Estimate
	db := database.DB.Where("delete_flag = 0")

	if propCd := c.Query("prop_cd"); propCd != "" {
		db = db.Where("prop_cd = ?", propCd)
	}
	if acceptNumber := c.Query("accept_number"); acceptNumber != "" {
		db = db.Where("accept_number = ?", acceptNumber)
	}
	if status := c.Query("status"); status != "" {
		db = db.Where("estimate_status = ?", status)
	}

	if err := db.Order("estimate_number DESC, subnumber ASC").Find(&estimates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"estimates": estimates})
}

// GetEstimateDetail handles GET /dashboard/estimate/:number/:subnumber
func GetEstimateDetail(c *gin.Context) {
	number := c.Param("number")
	subnumber := c.Param("subnumber")

	var estimate models.Estimate
	if err := database.DB.Where("estimate_number = ? AND subnumber = ? AND delete_flag = 0", number, subnumber).First(&estimate).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "見積もりが見つかりません"})
		return
	}

	var constructions []models.EstConstruction
	database.DB.Where("estimate_number = ? AND subnumber = ?", number, subnumber).Find(&constructions)

	var imgs []models.EstimateImg
	database.DB.Where("estimate_number = ? AND subnumber = ?", number, subnumber).Find(&imgs)

	var propManageDetails []models.EstPropManageDetail
	database.DB.Where("estimate_number = ? AND subnumber = ?", number, subnumber).Find(&propManageDetails)

	c.JSON(http.StatusOK, gin.H{
		"estimate":            estimate,
		"constructions":       constructions,
		"images":              imgs,
		"prop_manage_details": propManageDetails,
	})
}

// PostEstimateCreate handles POST /dashboard/estimate/create
func PostEstimateCreate(c *gin.Context) {
	var req models.Estimate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	// Generate estimate_number if not provided
	if req.EstimateNumber == "" {
		var maxNum struct{ MaxNum *string }
		database.DB.Raw("SELECT MAX(estimate_number) as max_num FROM estimate WHERE delete_flag = 0").Scan(&maxNum)
		if maxNum.MaxNum == nil {
			req.EstimateNumber = "E000001"
		} else {
			var seq int
			numStr := (*maxNum.MaxNum)[1:]
			seq = 0
			for _, ch := range numStr {
				seq = seq*10 + int(ch-'0')
			}
			seq++
			req.EstimateNumber = "E" + zeroPad(seq, 6)
		}
		req.Subnumber = "01"
	}

	delFlag := int16(0)
	req.DeleteFlag = &delFlag
	if err := database.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"estimate": req})
}

// PutEstimateUpdate handles PUT /dashboard/estimate/:number/:subnumber
func PutEstimateUpdate(c *gin.Context) {
	number := c.Param("number")
	subnumber := c.Param("subnumber")

	var estimate models.Estimate
	if err := database.DB.Where("estimate_number = ? AND subnumber = ? AND delete_flag = 0", number, subnumber).First(&estimate).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "見積もりが見つかりません"})
		return
	}

	if err := c.ShouldBindJSON(&estimate); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	estimate.EstimateNumber = number
	estimate.Subnumber = subnumber
	if err := database.DB.Save(&estimate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"estimate": estimate})
}

// DeleteEstimate handles DELETE /dashboard/estimate/:number/:subnumber (soft delete)
func DeleteEstimate(c *gin.Context) {
	number := c.Param("number")
	subnumber := c.Param("subnumber")
	if err := database.DB.Model(&models.Estimate{}).
		Where("estimate_number = ? AND subnumber = ?", number, subnumber).
		Update("delete_flag", int16(1)).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "見積もりを削除しました"})
}

// PutEstimateApprove handles PUT /dashboard/estimate/:number/:subnumber/approve
func PutEstimateApprove(c *gin.Context) {
	number := c.Param("number")
	subnumber := c.Param("subnumber")
	if err := database.DB.Model(&models.Estimate{}).
		Where("estimate_number = ? AND subnumber = ?", number, subnumber).
		Update("estimate_status", 2).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "見積もりを承認しました"})
}

// zeroPad pads an integer with leading zeros to width digits
func zeroPad(n, width int) string {
	s := ""
	for n > 0 {
		s = string(rune('0'+n%10)) + s
		n /= 10
	}
	for len(s) < width {
		s = "0" + s
	}
	return s
}

// GetEstimateRegisterForm handles GET /dashboard/estimate/register-form
func GetEstimateRegisterForm(c *gin.Context) {
	propCd := c.Query("prop_cd")
	acceptNumber := c.Query("accept_number")

	var propBasic models.PropBasic
	if propCd != "" {
		database.DB.Where("prop_cd = ? AND delete_flag = 0", propCd).First(&propBasic)
	}

	var reception models.Reception
	if acceptNumber != "" {
		database.DB.Where("accept_number = ? AND delete_flag = 0", acceptNumber).First(&reception)
	}

	var customers []models.Customer
	database.DB.Where("delete_flag = 0").Select("customer_cd, customer_name").Find(&customers)
	for i := range customers {
		customers[i].CustomerPassword = nil
	}

	var sekosakis []models.Sekosaki
	database.DB.Where("delete_flag = 0").Select("sekosaki_cd, sekosaki_name").Find(&sekosakis)
	for i := range sekosakis {
		sekosakis[i].SekosakiPassword = nil
	}

	c.JSON(http.StatusOK, gin.H{
		"prop_basic": propBasic,
		"reception":  reception,
		"customers":  customers,
		"sekosakis":  sekosakis,
	})
}

// PostEstimateSendCustomer handles POST /dashboard/estimate/send-customer
func PostEstimateSendCustomer(c *gin.Context) {
	var req struct {
		EstimateNumber string `json:"estimate_number" binding:"required"`
		Subnumber      string `json:"subnumber" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	database.DB.Model(&models.Estimate{}).
		Where("estimate_number = ? AND subnumber = ?", req.EstimateNumber, req.Subnumber).
		Update("send_customer_flg", 1)
	c.JSON(http.StatusOK, gin.H{"message": "顧客に送信しました"})
}

// PostEstimateUpdateApproveState handles POST /dashboard/estimate/update-approve-state
func PostEstimateUpdateApproveState(c *gin.Context) {
	var req struct {
		EstimateNumber string `json:"estimate_number" binding:"required"`
		Subnumber      string `json:"subnumber" binding:"required"`
		EstimateStatus int16  `json:"estimate_status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	database.DB.Model(&models.Estimate{}).
		Where("estimate_number = ? AND subnumber = ?", req.EstimateNumber, req.Subnumber).
		Update("estimate_status", req.EstimateStatus)
	c.JSON(http.StatusOK, gin.H{"message": "承認状態を更新しました"})
}

// PostEstimateCreateBK handles POST /dashboard/estimate/create-bk
func PostEstimateCreateBK(c *gin.Context) {
	var req struct {
		EstimateNumber string `json:"estimate_number" binding:"required"`
		Subnumber      string `json:"subnumber" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	// Copy existing estimate as BK (new subnumber)
	var original models.Estimate
	if err := database.DB.Where("estimate_number = ? AND subnumber = ?", req.EstimateNumber, req.Subnumber).First(&original).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "見積もりが見つかりません"})
		return
	}

	// Find next subnumber
	var maxSub struct{ MaxSub *string }
	database.DB.Raw("SELECT MAX(subnumber) as max_sub FROM estimate WHERE estimate_number = ?", req.EstimateNumber).Scan(&maxSub)
	nextSub := "02"
	if maxSub.MaxSub != nil {
		n := 0
		for _, ch := range *maxSub.MaxSub {
			n = n*10 + int(ch-'0')
		}
		n++
		nextSub = zeroPad(n, 2)
	}

	original.Subnumber = nextSub
	activated := int16(0)
	original.Activated = &activated
	database.DB.Create(&original)

	c.JSON(http.StatusCreated, gin.H{"estimate": original})
}

// GetEstimatePreviewPDF handles GET /dashboard/estimate/preview-pdf
func GetEstimatePreviewPDF(c *gin.Context) {
	estimateNumber := c.Query("estimate_number")
	subnumber := c.Query("subnumber")

	var estimate models.Estimate
	database.DB.Where("estimate_number = ? AND subnumber = ?", estimateNumber, subnumber).First(&estimate)

	var constructions []models.EstConstruction
	database.DB.Where("estimate_number = ? AND subnumber = ?", estimateNumber, subnumber).Find(&constructions)

	var propBasic models.PropBasic
	if estimate.PropCd != nil {
		database.DB.Where("prop_cd = ?", *estimate.PropCd).First(&propBasic)
	}

	c.JSON(http.StatusOK, gin.H{
		"estimate":      estimate,
		"constructions": constructions,
		"prop_basic":    propBasic,
	})
}

// PostEstimateExportPDF handles POST /dashboard/estimate/export-pdf
func PostEstimateExportPDF(c *gin.Context) {
	var req struct {
		EstimateNumber string `json:"estimate_number" binding:"required"`
		Subnumber      string `json:"subnumber" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	var estimate models.Estimate
	database.DB.Where("estimate_number = ? AND subnumber = ?", req.EstimateNumber, req.Subnumber).First(&estimate)

	var constructions []models.EstConstruction
	database.DB.Where("estimate_number = ? AND subnumber = ?", req.EstimateNumber, req.Subnumber).Find(&constructions)

	c.JSON(http.StatusOK, gin.H{"estimate": estimate, "constructions": constructions})
}

// GetEstimateCheckUnusedCustomer handles GET /dashboard/estimate/check-unused-customer
func GetEstimateCheckUnusedCustomer(c *gin.Context) {
	customerCd := c.Query("customer_cd")
	var count int64
	database.DB.Model(&models.Estimate{}).Where("customer_cd = ? AND delete_flag = 0", customerCd).Count(&count)
	c.JSON(http.StatusOK, gin.H{"unused": count == 0})
}

// GetEstimateCheckIssetOrder handles GET /dashboard/estimate/check-isset-order
func GetEstimateCheckIssetOrder(c *gin.Context) {
	estimateNumber := c.Query("estimate_number")
	var count int64
	database.DB.Model(&models.Order{}).Where("estimate_number = ?", estimateNumber).Count(&count)
	c.JSON(http.StatusOK, gin.H{"has_order": count > 0})
}

// PostEstimateDeleteOrder handles POST /dashboard/estimate/delete-order
func PostEstimateDeleteOrder(c *gin.Context) {
	var req struct {
		EstimateNumber string `json:"estimate_number" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Delete(&models.Order{}, "estimate_number = ?", req.EstimateNumber)
	c.JSON(http.StatusOK, gin.H{"message": "発注を削除しました"})
}

// PostEstimateChangeAmount handles POST /dashboard/estimate/change-amount
func PostEstimateChangeAmount(c *gin.Context) {
	var req struct {
		EstimateNumber string   `json:"estimate_number" binding:"required"`
		Subnumber      string   `json:"subnumber" binding:"required"`
		TotalAmount    *float64 `json:"total_amount"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Model(&models.Estimate{}).
		Where("estimate_number = ? AND subnumber = ?", req.EstimateNumber, req.Subnumber).
		Update("total_amount", req.TotalAmount)
	c.JSON(http.StatusOK, gin.H{"message": "金額を変更しました"})
}

// PostEstimateSaveConstructionDetail handles POST /dashboard/estimate/save-construction-detail
func PostEstimateSaveConstructionDetail(c *gin.Context) {
	var req models.EstConstruction
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	if req.ID > 0 {
		database.DB.Save(&req)
	} else {
		database.DB.Create(&req)
	}
	c.JSON(http.StatusOK, gin.H{"construction": req})
}

// PostEstimateSaveOther2 handles POST /dashboard/estimate/save-other2
func PostEstimateSaveOther2(c *gin.Context) {
	var req struct {
		EstimateNumber string                 `json:"estimate_number" binding:"required"`
		Subnumber      string                 `json:"subnumber" binding:"required"`
		Data           map[string]interface{} `json:"data"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Model(&models.Estimate{}).
		Where("estimate_number = ? AND subnumber = ?", req.EstimateNumber, req.Subnumber).
		Updates(req.Data)
	c.JSON(http.StatusOK, gin.H{"message": "更新しました"})
}

// PostEstimateSavePropManageDetail handles POST /dashboard/estimate/save-prop-manage-detail
func PostEstimateSavePropManageDetail(c *gin.Context) {
	var req models.EstPropManageDetail
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	if req.ID > 0 {
		database.DB.Save(&req)
	} else {
		database.DB.Create(&req)
	}
	c.JSON(http.StatusOK, gin.H{"detail": req})
}

// PostEstimateSaveFreeInputDetail handles POST /dashboard/estimate/save-free-input-detail
func PostEstimateSaveFreeInputDetail(c *gin.Context) {
	var req struct {
		EstimateNumber string                 `json:"estimate_number" binding:"required"`
		Subnumber      string                 `json:"subnumber" binding:"required"`
		Data           map[string]interface{} `json:"data"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Model(&models.Estimate{}).
		Where("estimate_number = ? AND subnumber = ?", req.EstimateNumber, req.Subnumber).
		Updates(req.Data)
	c.JSON(http.StatusOK, gin.H{"message": "更新しました"})
}
