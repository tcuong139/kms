package handlers

import (
	"kms_golang/database"
	"kms_golang/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetNotificationList handles GET /dashboard/notification/list
func GetNotificationList(c *gin.Context) {
	var notifications []models.NotificationList
	db := database.DB.Where("delete_flag = 0")

	if guardType, exists := c.Get("guard_type"); exists {
		switch guardType {
		case "customer":
			db = db.Where("target_type IN (0, 2)") // 0=all, 2=customers
		case "sekosaki":
			db = db.Where("target_type IN (0, 3)") // 0=all, 3=sekosakis
		default:
			// web users see all
		}
	}

	if err := db.Order("regist_datetime DESC").Find(&notifications).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"notifications": notifications})
}

// GetNotificationDetail handles GET /dashboard/notification/:id
func GetNotificationDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "無効なID"})
		return
	}
	var notification models.NotificationList
	if err := database.DB.Where("id = ? AND delete_flag = 0", id).First(&notification).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "お知らせが見つかりません"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"notification": notification})
}

// PostNotificationCreate handles POST /dashboard/notification/create
func PostNotificationCreate(c *gin.Context) {
	var req models.NotificationList
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

	c.JSON(http.StatusCreated, gin.H{"notification": req})
}

// PutNotificationUpdate handles PUT /dashboard/notification/:id
func PutNotificationUpdate(c *gin.Context) {
	idStr := c.Param("id")
	nid, nerr := strconv.ParseUint(idStr, 10, 64)
	if nerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "無効なID"})
		return
	}
	var notification models.NotificationList
	if err := database.DB.Where("id = ? AND delete_flag = 0", nid).First(&notification).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "お知らせが見つかりません"})
		return
	}

	if err := c.ShouldBindJSON(&notification); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	notification.ID = uint(nid)
	if err := database.DB.Save(&notification).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"notification": notification})
}

// DeleteNotification handles DELETE /dashboard/notification/:id (soft delete)
func DeleteNotification(c *gin.Context) {
	idStr := c.Param("id")
	did, derr := strconv.ParseUint(idStr, 10, 64)
	if derr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "無効なID"})
		return
	}
	delFlag2 := int16(1)
	if err := database.DB.Model(&models.NotificationList{}).Where("id = ?", did).Update("delete_flag", delFlag2).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "お知らせを削除しました"})
}

// GetUserNotifications handles GET /dashboard/notification/user/:user_id
func GetUserNotifications(c *gin.Context) {
	userID := c.Param("user_id")
	var userNotifications []models.UserNotifi
	database.DB.Where("user_id = ?", userID).Order("regist_datetime DESC").Find(&userNotifications)
	c.JSON(http.StatusOK, gin.H{"user_notifications": userNotifications})
}

// GetNotificationCompany handles GET /dashboard/notification/company
func GetNotificationCompany(c *gin.Context) {
	var companies []models.CompanyInfo
	database.DB.Find(&companies)
	c.JSON(http.StatusOK, gin.H{"companies": companies})
}

// PostNotificationExportPDF handles POST /dashboard/notification/export-pdf
func PostNotificationExportPDF(c *gin.Context) {
	idStr := c.Query("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)
	var notification models.NotificationList
	database.DB.Where("id = ?", id).First(&notification)
	c.JSON(http.StatusOK, gin.H{"notification": notification})
}

// PostNotificationUpdatePrintFlgAndGiveout handles POST /dashboard/notification/update-print-giveout
func PostNotificationUpdatePrintFlgAndGiveout(c *gin.Context) {
	var req struct {
		IDs      []uint `json:"ids" binding:"required"`
		PrintFlg *int16 `json:"print_flg"`
		Giveout  *int16 `json:"giveout"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	updates := map[string]interface{}{}
	if req.PrintFlg != nil {
		updates["print_flg"] = *req.PrintFlg
	}
	if req.Giveout != nil {
		updates["giveout"] = *req.Giveout
	}
	database.DB.Model(&models.NotificationList{}).Where("id IN ?", req.IDs).Updates(updates)
	c.JSON(http.StatusOK, gin.H{"message": "更新しました"})
}

// PostNotificationUpdateGiveoutSingleRow handles POST /dashboard/notification/update-giveout-single
func PostNotificationUpdateGiveoutSingleRow(c *gin.Context) {
	var req struct {
		ID      uint  `json:"id" binding:"required"`
		Giveout int16 `json:"giveout"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Model(&models.NotificationList{}).Where("id = ?", req.ID).Update("giveout", req.Giveout)
	c.JSON(http.StatusOK, gin.H{"message": "更新しました"})
}

// PostNotificationUpdatePrintFlgSingleRow handles POST /dashboard/notification/update-print-single
func PostNotificationUpdatePrintFlgSingleRow(c *gin.Context) {
	var req struct {
		ID       uint  `json:"id" binding:"required"`
		PrintFlg int16 `json:"print_flg"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Model(&models.NotificationList{}).Where("id = ?", req.ID).Update("print_flg", req.PrintFlg)
	c.JSON(http.StatusOK, gin.H{"message": "更新しました"})
}

// GetNotificationRegisterForm handles GET /dashboard/notification/register-form
func GetNotificationRegisterForm(c *gin.Context) {
	var styles []models.NotificationStyles
	database.DB.Find(&styles)

	var customers []models.Customer
	database.DB.Where("delete_flag = 0").Select("customer_cd, customer_name").Find(&customers)
	for i := range customers {
		customers[i].CustomerPassword = nil
	}

	c.JSON(http.StatusOK, gin.H{"styles": styles, "customers": customers})
}

// PostNotificationStylesRegister handles POST /dashboard/notification/styles-register
func PostNotificationStylesRegister(c *gin.Context) {
	var req models.NotificationStyles
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Create(&req)
	c.JSON(http.StatusCreated, gin.H{"notification_style": req})
}

// PostNotificationStylesUpdate handles POST /dashboard/notification/styles-update
func PostNotificationStylesUpdate(c *gin.Context) {
	var req models.NotificationStyles
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Save(&req)
	c.JSON(http.StatusOK, gin.H{"notification_style": req})
}

// PostNotificationCheckUpdate handles POST /dashboard/notification/check-update
func PostNotificationCheckUpdate(c *gin.Context) {
	var req struct {
		ID uint `json:"id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	var notification models.NotificationList
	database.DB.Where("id = ?", req.ID).First(&notification)
	c.JSON(http.StatusOK, gin.H{"notification": notification})
}

// GetNotificationSekosaki handles GET /dashboard/notification/sekosaki
func GetNotificationSekosaki(c *gin.Context) {
	var sekosakis []models.Sekosaki
	database.DB.Where("delete_flag = 0").Select("sekosaki_cd, sekosaki_name").Find(&sekosakis)
	for i := range sekosakis {
		sekosakis[i].SekosakiPassword = nil
	}
	c.JSON(http.StatusOK, gin.H{"sekosakis": sekosakis})
}

// PostNotificationSearchSekosaki handles POST /dashboard/notification/search-sekosaki
func PostNotificationSearchSekosaki(c *gin.Context) {
	var req struct {
		Keyword string `json:"keyword"`
	}
	c.ShouldBindJSON(&req)
	var sekosakis []models.Sekosaki
	database.DB.Where("delete_flag = 0 AND (sekosaki_name LIKE ? OR sekosaki_cd LIKE ?)",
		"%"+req.Keyword+"%", "%"+req.Keyword+"%").Find(&sekosakis)
	for i := range sekosakis {
		sekosakis[i].SekosakiPassword = nil
	}
	c.JSON(http.StatusOK, gin.H{"sekosakis": sekosakis})
}

// PostNotificationSearchCompany handles POST /dashboard/notification/search-company
func PostNotificationSearchCompany(c *gin.Context) {
	var req struct {
		Keyword string `json:"keyword"`
	}
	c.ShouldBindJSON(&req)
	var companies []models.CompanyInfo
	database.DB.Where("company_name LIKE ?", "%"+req.Keyword+"%").Find(&companies)
	c.JSON(http.StatusOK, gin.H{"companies": companies})
}

// GetNotificationShowPreview handles GET /dashboard/notification/preview/:id
func GetNotificationShowPreview(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)
	var notification models.NotificationList
	database.DB.Where("id = ?", id).First(&notification)
	c.JSON(http.StatusOK, gin.H{"notification": notification})
}

// GetPropertyNotificationList handles GET /dashboard/notification/property-list
func GetPropertyNotificationList(c *gin.Context) {
	propCd := c.Query("prop_cd")
	var notifications []models.NotificationList
	db := database.DB.Where("delete_flag = 0")
	if propCd != "" {
		db = db.Where("prop_cd = ?", propCd)
	}
	db.Order("regist_datetime DESC").Find(&notifications)
	c.JSON(http.StatusOK, gin.H{"notifications": notifications})
}

// GetPropertyNotificationRegister handles GET /dashboard/notification/property-register
func GetPropertyNotificationRegister(c *gin.Context) {
	var styles []models.NotificationStyles
	database.DB.Find(&styles)
	c.JSON(http.StatusOK, gin.H{"styles": styles})
}

// PostPropertyNotificationRegister handles POST /dashboard/notification/property-register
func PostPropertyNotificationRegister(c *gin.Context) {
	var req models.NotificationList
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	delFlag := int16(0)
	req.DeleteFlag = &delFlag
	database.DB.Create(&req)
	c.JSON(http.StatusCreated, gin.H{"notification": req})
}

// PostPropertyNotificationUpdate handles POST /dashboard/notification/property-update
func PostPropertyNotificationUpdate(c *gin.Context) {
	var req models.NotificationList
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Save(&req)
	c.JSON(http.StatusOK, gin.H{"notification": req})
}

// PostSearchPropertyNotification handles POST /dashboard/notification/search-property
func PostSearchPropertyNotification(c *gin.Context) {
	var req struct {
		PropCd  string `json:"prop_cd"`
		Keyword string `json:"keyword"`
	}
	c.ShouldBindJSON(&req)
	var notifications []models.NotificationList
	db := database.DB.Where("delete_flag = 0")
	if req.PropCd != "" {
		db = db.Where("prop_cd = ?", req.PropCd)
	}
	if req.Keyword != "" {
		db = db.Where("notification_title LIKE ?", "%"+req.Keyword+"%")
	}
	db.Find(&notifications)
	c.JSON(http.StatusOK, gin.H{"notifications": notifications})
}
