package handlers

import (
	"fmt"
	"kms_golang/database"
	"kms_golang/models"
	"kms_golang/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

// GetUserList handles GET /dashboard/user/list
func GetUserList(c *gin.Context) {
	var users []models.User
	db := database.DB.Where("delete_flg = 0")

	if name := c.Query("name"); name != "" {
		db = db.Where("user_name LIKE ?", "%"+name+"%")
	}

	if err := db.Order("user_id ASC").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	for i := range users {
		users[i].Password = nil
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

// GetUserDetail handles GET /dashboard/user/:id
func GetUserDetail(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := database.DB.Where("user_id = ? AND delete_flg = 0", id).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "ユーザーが見つかりません"})
		return
	}
	user.Password = nil
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// PostUserCreate handles POST /dashboard/user/create
func PostUserCreate(c *gin.Context) {
	var req struct {
		UserName string `json:"user_name" binding:"required"`
		LoginID  string `json:"login_id" binding:"required"`
		Password string `json:"password" binding:"required,min=6"`
		Auth     *int16 `json:"auth"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	// Check duplicate login_id
	var exists models.User
	if err := database.DB.Where("login_id = ? AND delete_flg = 0", req.LoginID).First(&exists).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"message": "このログインIDは既に使用されています"})
		return
	}

	hashed, err := services.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	var maxID struct{ MaxID *int }
	database.DB.Raw("SELECT MAX(CAST(user_id AS UNSIGNED)) as max_id FROM user WHERE delete_flg = 0").Scan(&maxID)
	nextID := 1
	if maxID.MaxID != nil {
		nextID = *maxID.MaxID + 1
	}

	userID := fmt.Sprintf("%d", nextID)
	userName := req.UserName
	loginID := req.LoginID
	var auth int16 = 0
	if req.Auth != nil {
		auth = *req.Auth
	}
	delFlg := int16(0)

	user := models.User{
		UserID:    userID,
		UserName:  &userName,
		LoginID:   &loginID,
		Password:  &hashed,
		Auth:      &auth,
		DeleteFlg: &delFlg,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	user.Password = nil
	c.JSON(http.StatusCreated, gin.H{"user": user})
}

// PutUserUpdate handles PUT /dashboard/user/:id
func PutUserUpdate(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := database.DB.Where("user_id = ? AND delete_flg = 0", id).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "ユーザーが見つかりません"})
		return
	}

	var req struct {
		UserName *string `json:"user_name"`
		LoginID  *string `json:"login_id"`
		Password *string `json:"password"`
		Auth     *int16  `json:"auth"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	if req.UserName != nil {
		user.UserName = req.UserName
	}
	if req.LoginID != nil {
		user.LoginID = req.LoginID
	}
	if req.Password != nil {
		hashed, err := services.HashPassword(*req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		user.Password = &hashed
	}
	if req.Auth != nil {
		user.Auth = req.Auth
	}

	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	user.Password = nil
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// DeleteUser handles DELETE /dashboard/user/:id (soft delete)
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	delFlg := int16(1)
	if err := database.DB.Model(&models.User{}).Where("user_id = ?", id).Update("delete_flg", delFlg).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ユーザーを削除しました"})
}

// GetUserRegisterForm handles GET /dashboard/user/register-form
func GetUserRegisterForm(c *gin.Context) {
	// Return dropdown data needed for user creation form
	var crews []models.Crew
	database.DB.Where("delete_flag = 0").Select("crew_code, crew_name").Find(&crews)
	c.JSON(http.StatusOK, gin.H{"crews": crews})
}

// GetUserCrewWorkplace handles GET /dashboard/user/:id/crew-workplace
func GetUserCrewWorkplace(c *gin.Context) {
	id := c.Param("id")
	var workplaces []models.CrewWorkplace
	database.DB.Where("user_id = ?", id).Find(&workplaces)
	c.JSON(http.StatusOK, gin.H{"crew_workplaces": workplaces})
}

// PostUserCrewWorkplaceCreate handles POST /dashboard/user/crew-workplace/create
func PostUserCrewWorkplaceCreate(c *gin.Context) {
	var req models.CrewWorkplace
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	if err := database.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"crew_workplace": req})
}

// PutUserCrewWorkplaceUpdate handles PUT /dashboard/user/crew-workplace/:id
func PutUserCrewWorkplaceUpdate(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "無効なID"})
		return
	}
	var wp models.CrewWorkplace
	if err := database.DB.First(&wp, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "見つかりません"})
		return
	}
	if err := c.ShouldBindJSON(&wp); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	wp.ID = uint(id)
	database.DB.Save(&wp)
	c.JSON(http.StatusOK, gin.H{"crew_workplace": wp})
}

// PostUserUpdateWorkEndDate handles POST /dashboard/user/update-work-end-date
func PostUserUpdateWorkEndDate(c *gin.Context) {
	var req struct {
		UserID      string  `json:"user_id" binding:"required"`
		WorkEndDate *string `json:"work_end_date"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Model(&models.User{}).Where("user_id = ?", req.UserID).Update("work_end_date", req.WorkEndDate)
	c.JSON(http.StatusOK, gin.H{"message": "更新しました"})
}

// GetUserListByCategory handles GET /dashboard/user/list-by-category
func GetUserListByCategory(c *gin.Context) {
	auth := c.Query("auth")
	var users []models.User
	db := database.DB.Where("delete_flg = 0")
	if auth != "" {
		db = db.Where("auth = ?", auth)
	}
	db.Order("user_id ASC").Find(&users)
	for i := range users {
		users[i].Password = nil
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

// GetUserExportExcel handles GET /dashboard/user/export-excel
func GetUserExportExcel(c *gin.Context) {
	var users []models.User
	database.DB.Where("delete_flg = 0").Order("user_id ASC").Find(&users)

	f := excelize.NewFile()
	sheet := "Sheet1"
	headers := []string{"ユーザーID", "ユーザー名", "ログインID", "権限"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}
	for row, u := range users {
		r := row + 2
		vals := []interface{}{u.UserID, ptrStr(u.UserName), ptrStr(u.LoginID), u.Auth}
		for col, v := range vals {
			cell, _ := excelize.CoordinatesToCellName(col+1, r)
			f.SetCellValue(sheet, cell, v)
		}
	}

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=user_list.xlsx")
	f.Write(c.Writer)
}

// GetUserAuthorityList handles GET /dashboard/user/authority-list
func GetUserAuthorityList(c *gin.Context) {
	var users []models.User
	database.DB.Where("delete_flg = 0").Select("user_id, user_name, auth").Order("user_id ASC").Find(&users)
	for i := range users {
		users[i].Password = nil
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

// PostUserUpdateAuthority handles POST /dashboard/user/update-authority
func PostUserUpdateAuthority(c *gin.Context) {
	var req struct {
		UserID string `json:"user_id" binding:"required"`
		Auth   int16  `json:"auth"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Model(&models.User{}).Where("user_id = ?", req.UserID).Update("auth", req.Auth)
	c.JSON(http.StatusOK, gin.H{"message": "権限を更新しました"})
}

// GetUserCheckLoginID handles GET /dashboard/user/check-login-id
func GetUserCheckLoginID(c *gin.Context) {
	loginID := c.Query("login_id")
	excludeID := c.Query("exclude_id")

	db := database.DB.Model(&models.User{}).Where("login_id = ? AND delete_flg = 0", loginID)
	if excludeID != "" {
		db = db.Where("user_id != ?", excludeID)
	}

	var count int64
	db.Count(&count)
	c.JSON(http.StatusOK, gin.H{"exists": count > 0})
}
