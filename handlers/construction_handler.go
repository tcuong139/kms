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

// GetSekosakiList handles GET /dashboard/sekosaki/list
func GetSekosakiList(c *gin.Context) {
	var sekosakis []models.Sekosaki
	db := database.DB.Where("delete_flag = 0")

	if name := c.Query("sekosaki_name"); name != "" {
		db = db.Where("sekosaki_name LIKE ?", "%"+name+"%")
	}
	if cd := c.Query("sekosaki_cd"); cd != "" {
		db = db.Where("sekosaki_cd LIKE ?", "%"+cd+"%")
	}

	if err := db.Order("sekosaki_cd ASC").Find(&sekosakis).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	for i := range sekosakis {
		sekosakis[i].SekosakiPassword = nil
	}

	c.JSON(http.StatusOK, gin.H{"sekosakis": sekosakis})
}

// GetSekosakiDetail handles GET /dashboard/sekosaki/:cd
func GetSekosakiDetail(c *gin.Context) {
	cd := c.Param("cd")
	var sekosaki models.Sekosaki
	if err := database.DB.Where("sekosaki_cd = ? AND delete_flag = 0", cd).First(&sekosaki).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "施工先が見つかりません"})
		return
	}
	sekosaki.SekosakiPassword = nil

	var personnel []models.SekosakiPersonnel
	database.DB.Where("sekosaki_cd = ?", cd).Find(&personnel)

	c.JSON(http.StatusOK, gin.H{"sekosaki": sekosaki, "personnel": personnel})
}

// PostSekosakiCreate handles POST /dashboard/sekosaki/create
func PostSekosakiCreate(c *gin.Context) {
	var req struct {
		SekosakiName string  `json:"sekosaki_name" binding:"required"`
		SekosakiKana string  `json:"sekosaki_kana"`
		PostCode     string  `json:"post_code"`
		PrefectureID *string `json:"prefecture_id"`
		CityID       *string `json:"city_id"`
		TownID       *string `json:"town_id"`
		BlockName    string  `json:"block_name"`
		Tel          string  `json:"tel"`
		LoginID      string  `json:"login_id"`
		Password     string  `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	cd, err := utils.GenerateSekosakiID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	sekosakiName := req.SekosakiName
	sekosakiKana := req.SekosakiKana
	postCode := req.PostCode
	blockName := req.BlockName
	tel := req.Tel
	loginID := req.LoginID
	delFlag := int16(0)

	sekosaki := models.Sekosaki{
		SekosakiCd:      cd,
		SekosakiName:    &sekosakiName,
		SekosakiKana:    &sekosakiKana,
		PostCode:        &postCode,
		BlockName:       &blockName,
		Tel:             &tel,
		PrefectureID:    req.PrefectureID,
		CityID:          req.CityID,
		TownID:          req.TownID,
		SekosakiLoginID: &loginID,
		DeleteFlag:      &delFlag,
	}

	if req.Password != "" {
		hashed, herr := services.HashPassword(req.Password)
		if herr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": herr.Error()})
			return
		}
		sekosaki.SekosakiPassword = &hashed
	}

	if err := database.DB.Create(&sekosaki).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	sekosaki.SekosakiPassword = nil
	c.JSON(http.StatusCreated, gin.H{"sekosaki": sekosaki})
}

// PutSekosakiUpdate handles PUT /dashboard/sekosaki/:cd
func PutSekosakiUpdate(c *gin.Context) {
	cd := c.Param("cd")
	var sekosaki models.Sekosaki
	if err := database.DB.Where("sekosaki_cd = ? AND delete_flag = 0", cd).First(&sekosaki).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "施工先が見つかりません"})
		return
	}

	if err := c.ShouldBindJSON(&sekosaki); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	sekosaki.SekosakiCd = cd
	if err := database.DB.Save(&sekosaki).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	sekosaki.SekosakiPassword = nil
	c.JSON(http.StatusOK, gin.H{"sekosaki": sekosaki})
}

// DeleteSekosaki handles DELETE /dashboard/sekosaki/:cd (soft delete)
func DeleteSekosaki(c *gin.Context) {
	cd := c.Param("cd")
	delFlag := int16(1)
	if err := database.DB.Model(&models.Sekosaki{}).Where("sekosaki_cd = ?", cd).Update("delete_flag", delFlag).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "施工先を削除しました"})
}

// GetSekosakiExportExcel handles GET /dashboard/sekosaki/export-excel
func GetSekosakiExportExcel(c *gin.Context) {
	var sekosakis []models.Sekosaki
	database.DB.Where("delete_flag = 0").Order("sekosaki_cd ASC").Find(&sekosakis)

	f := excelize.NewFile()
	sheet := "Sheet1"
	headers := []string{"施工先CD", "施工先名", "施工先名カナ", "郵便番号", "住所", "電話番号", "FAX"}
	for i, h := range headers {
		col := string(rune('A' + i))
		f.SetCellValue(sheet, col+"1", h)
	}
	for i, s := range sekosakis {
		row := strconv.Itoa(i + 2)
		f.SetCellValue(sheet, "A"+row, s.SekosakiCd)
		if s.SekosakiName != nil {
			f.SetCellValue(sheet, "B"+row, *s.SekosakiName)
		}
		if s.SekosakiKana != nil {
			f.SetCellValue(sheet, "C"+row, *s.SekosakiKana)
		}
		if s.PostCode != nil {
			f.SetCellValue(sheet, "D"+row, *s.PostCode)
		}
		if s.BlockName != nil {
			f.SetCellValue(sheet, "E"+row, *s.BlockName)
		}
		if s.Tel != nil {
			f.SetCellValue(sheet, "F"+row, *s.Tel)
		}
	}

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=construction_sites.xlsx")
	f.Write(c.Writer)
}

// GetSekosakiCustomer handles GET /dashboard/sekosaki/customer
func GetSekosakiCustomer(c *gin.Context) {
	sekosakiCd := c.Query("sekosaki_cd")
	var customers []models.Customer
	database.DB.Joins("JOIN prop_basic ON prop_basic.customer_cd = customer.customer_cd").
		Where("prop_basic.sekosaki_cd = ? AND customer.delete_flag = 0", sekosakiCd).
		Group("customer.customer_cd").Find(&customers)
	for i := range customers {
		customers[i].CustomerPassword = nil
	}
	c.JSON(http.StatusOK, gin.H{"customers": customers})
}

// PostSekosakiSearchCustomer handles POST /dashboard/sekosaki/search-customer
func PostSekosakiSearchCustomer(c *gin.Context) {
	var req struct {
		Keyword string `json:"keyword"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	var customers []models.Customer
	database.DB.Where("delete_flag = 0 AND (customer_name LIKE ? OR customer_cd LIKE ?)",
		"%"+req.Keyword+"%", "%"+req.Keyword+"%").Find(&customers)
	for i := range customers {
		customers[i].CustomerPassword = nil
	}
	c.JSON(http.StatusOK, gin.H{"customers": customers})
}

// PostSekosakiChangePassword handles POST /dashboard/sekosaki/:cd/change-password
func PostSekosakiChangePassword(c *gin.Context) {
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

	if err := database.DB.Model(&models.Sekosaki{}).Where("sekosaki_cd = ?", cd).Update("sekosaki_password", hashed).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "パスワードを変更しました"})
}

// GetSekosakiCheckLoginID handles GET /dashboard/sekosaki/check-login-id
func GetSekosakiCheckLoginID(c *gin.Context) {
	loginID := c.Query("login_id")
	sekosakiCd := c.Query("sekosaki_cd")
	var count int64
	db := database.DB.Model(&models.Sekosaki{}).Where("sekosaki_login_id = ? AND delete_flag = 0", loginID)
	if sekosakiCd != "" {
		db = db.Where("sekosaki_cd <> ?", sekosakiCd)
	}
	db.Count(&count)
	c.JSON(http.StatusOK, gin.H{"valid": count == 0})
}

// GetSekosakiCheckName handles GET /dashboard/sekosaki/check-name
func GetSekosakiCheckName(c *gin.Context) {
	name := c.Query("sekosaki_name")
	sekosakiCd := c.Query("sekosaki_cd")
	var count int64
	db := database.DB.Model(&models.Sekosaki{}).Where("sekosaki_name = ? AND delete_flag = 0", name)
	if sekosakiCd != "" {
		db = db.Where("sekosaki_cd <> ?", sekosakiCd)
	}
	db.Count(&count)
	c.JSON(http.StatusOK, gin.H{"valid": count == 0})
}

// GetSekosakiCheckAddress handles GET /dashboard/sekosaki/check-address
func GetSekosakiCheckAddress(c *gin.Context) {
	address := c.Query("block_name")
	sekosakiCd := c.Query("sekosaki_cd")
	var count int64
	db := database.DB.Model(&models.Sekosaki{}).Where("block_name = ? AND delete_flag = 0", address)
	if sekosakiCd != "" {
		db = db.Where("sekosaki_cd <> ?", sekosakiCd)
	}
	db.Count(&count)
	c.JSON(http.StatusOK, gin.H{"valid": count == 0})
}
