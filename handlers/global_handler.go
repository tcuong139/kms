package handlers

import (
	"kms_golang/database"
	"kms_golang/models"
	"kms_golang/utils"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// PostUploadFile handles POST /upload-files
func PostUploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ファイルが見つかりません"})
		return
	}

	if !utils.IsAllowedImageType(file.Filename) && !utils.IsAllowedDocumentType(file.Filename) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "許可されていないファイル形式です"})
		return
	}

	uploadDir := filepath.Join("public", "uploads")
	savedName, err := utils.SaveUploadedFile(file, uploadDir)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"filename": savedName,
		"path":     filepath.Join("uploads", savedName),
	})
}

// DeleteUploadedFile handles DELETE /upload-files/:filename
func DeleteUploadedFile(c *gin.Context) {
	filename := c.Param("filename")

	// Validate filename to prevent path traversal attacks
	if filepath.Base(filename) != filename {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid filename"})
		return
	}

	filePath := filepath.Join("public", "uploads", filename)
	if err := utils.DeleteFile(filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ファイルを削除しました"})
}

// GetPostSearch handles GET /postSearch?zip=xxx (zip code to address lookup)
func GetPostSearch(c *gin.Context) {
	zip := c.Query("zip")
	if zip == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "郵便番号を入力してください"})
		return
	}

	pref, city, town, err := utils.GetAddressByZip(zip)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "住所が見つかりません"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"pref_id":   pref.PrefectureID,
		"pref_name": pref.PrefectureName,
		"city_id":   city.CityID,
		"city_name": city.CityName,
		"town_id":   town.TownID,
		"town_name": town.TownName,
	})
}

// GetPrefectures handles GET /prefectures
func GetPrefectures(c *gin.Context) {
	prefs, err := utils.GetPrefectures()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"prefectures": prefs})
}

// GetCities handles GET /cities?pref_id=xxx
func GetCities(c *gin.Context) {
	prefID := c.Query("pref_id")
	if prefID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "pref_id is required"})
		return
	}

	cities, err := utils.GetCitiesByPref(prefID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"cities": cities})
}

// GetTowns handles GET /towns?city_id=xxx
func GetTowns(c *gin.Context) {
	cityID := c.Query("city_id")
	if cityID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "city_id is required"})
		return
	}

	towns, err := utils.GetTownsByCity(cityID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"towns": towns})
}

// GetWellKnownList handles GET /dashboard/well-known/list
func GetWellKnownList(c *gin.Context) {
	var wellKnowns []models.WellKnown
	if err := database.DB.Order("id ASC").Find(&wellKnowns).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"well_knowns": wellKnowns})
}

// PostWellKnownCreate handles POST /dashboard/well-known/create
func PostWellKnownCreate(c *gin.Context) {
	var req models.WellKnown
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	if err := database.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"well_known": req})
}

// PutWellKnownUpdate handles PUT /dashboard/well-known/:id
func PutWellKnownUpdate(c *gin.Context) {
	id := c.Param("id")
	var wk models.WellKnown
	if err := database.DB.First(&wk, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "見つかりません"})
		return
	}
	if err := c.ShouldBindJSON(&wk); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	if err := database.DB.Save(&wk).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"well_known": wk})
}

// DeleteWellKnown handles DELETE /dashboard/well-known/:id
func DeleteWellKnown(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.WellKnown{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "削除しました"})
}

// GetLoadCustomers handles GET /dashboard/loadCustomers
func GetLoadCustomers(c *gin.Context) {
	var customers []struct {
		CustomerCd   string  `json:"customer_cd"`
		CustomerName *string `json:"customer_name"`
	}
	database.DB.Table("customers").Select("customer_cd, customer_name").
		Where("delete_flag = 0").Order("customer_cd ASC").Find(&customers)
	c.JSON(http.StatusOK, gin.H{"customers": customers})
}

// GetDropdownSearch handles GET /dashboard/dropdown/search
// Returns customers, sekosakis, and props for Select2 dropdowns
func GetDropdownSearch(c *gin.Context) {
	search := c.Query("q")

	type DropdownItem struct {
		ID   string `json:"id"`
		Text string `json:"text"`
	}

	var customers []struct {
		CustomerCd   string  `gorm:"column:customer_cd"`
		CustomerName *string `gorm:"column:customer_name"`
	}
	db := database.DB.Table("customers").Select("customer_cd, customer_name").Where("delete_flag = 0")
	if search != "" {
		db = db.Where("customer_name LIKE ? OR customer_cd LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	db.Limit(50).Find(&customers)

	var sekosakis []struct {
		SekosakiCd   string  `gorm:"column:sekosaki_cd"`
		SekosakiName *string `gorm:"column:sekosaki_name"`
	}
	db2 := database.DB.Table("sekosakis").Select("sekosaki_cd, sekosaki_name").Where("delete_flag = 0")
	if search != "" {
		db2 = db2.Where("sekosaki_name LIKE ? OR sekosaki_cd LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	db2.Limit(50).Find(&sekosakis)

	var props []struct {
		PropCd   string  `gorm:"column:prop_cd"`
		PropName *string `gorm:"column:prop_name"`
	}
	db3 := database.DB.Table("prop_basics").Select("prop_cd, prop_name").Where("delete_flag = 0")
	if search != "" {
		db3 = db3.Where("prop_name LIKE ? OR prop_cd LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	db3.Limit(50).Find(&props)

	c.JSON(http.StatusOK, gin.H{
		"customers": customers,
		"sekosakis": sekosakis,
		"props":     props,
	})
}

// GetCompanyInfo handles GET /dashboard/company-info
func GetCompanyInfo(c *gin.Context) {
	var company []models.CompanyInfo
	database.DB.Order("co_code ASC").Find(&company)
	c.JSON(http.StatusOK, gin.H{"company_info": company})
}

// PostCompanyInfoRegister handles POST /dashboard/company-info/register
func PostCompanyInfoRegister(c *gin.Context) {
	var req models.CompanyInfo
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	if err := database.DB.Save(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"company_info": req})
}

// GetDashboardTop handles GET /dashboard (top page data)
func GetDashboardTop(c *gin.Context) {
	// Return summary counts for the dashboard
	var userCount, customerCount, propCount, receptionCount int64
	database.DB.Model(&models.User{}).Where("delete_flg = 0").Count(&userCount)
	database.DB.Model(&models.Customer{}).Where("delete_flag = 0").Count(&customerCount)
	database.DB.Model(&models.PropBasic{}).Where("delete_flag = 0").Count(&propCount)
	database.DB.Model(&models.Reception{}).Where("delete_flag = 0").Count(&receptionCount)

	c.JSON(http.StatusOK, gin.H{
		"user_count":      userCount,
		"customer_count":  customerCount,
		"property_count":  propCount,
		"reception_count": receptionCount,
	})
}
