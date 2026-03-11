package handlers

import (
	"kms_golang/database"
	"kms_golang/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// GetStaffDailyReportList handles GET /dashboard/staff-daily-report/list
func GetStaffDailyReportList(c *gin.Context) {
	date := c.DefaultQuery("date", time.Now().Format("2006-01"))
	userID := c.GetString("user_id") // from JWT

	type ReportRow struct {
		UserID          string  `json:"user_id"`
		UserName        string  `json:"user_name"`
		CrewName        string  `json:"crew_name"`
		SettingDate     string  `json:"setting_date"`
		Biko            string  `json:"biko"`
		TimeIn          string  `json:"time_in"`
		TimeOut         string  `json:"time_out"`
		OvertimeComment string  `json:"overtime_comment"`
		UserConfirm     string  `json:"user_confirm"`
		Title           string  `json:"title"`
		ActionStartDate string  `json:"action_start_date"`
		ActionStartTime string  `json:"action_start_time"`
		ActionEndedTime string  `json:"action_ended_time"`
		IsReport        *string `json:"is_report"`
	}

	year, _ := strconv.Atoi(date[:4])
	month, _ := strconv.Atoi(date[5:7])
	daysInMonth := time.Date(year, time.Month(month)+1, 0, 0, 0, 0, 0, time.UTC).Day()

	reportAll := make(map[string][]ReportRow)

	for i := 1; i <= daysInMonth; i++ {
		day := time.Date(year, time.Month(month), i, 0, 0, 0, 0, time.UTC).Format("2006-01-02")

		var rows []ReportRow
		db := database.DB.Table("user").
			Select(`user.user_id, user.user_name, crews.crew_name,
				user_dayoff_settings.biko, user_dayoff_settings.user_confirm,
				user_dayoff_settings.overtime_comment, user_dayoff_settings.user_id as is_report,
				action_plans.title, action_plans.action_start_date,
				action_plans.action_start_time, action_plans.action_ended_time`).
			Joins("LEFT JOIN crews ON user.user_id = crews.crew_code").
			Joins("LEFT JOIN user_dayoff_settings ON user.user_id = user_dayoff_settings.user_id AND user_dayoff_settings.setting_date LIKE ?", day+"%").
			Joins("LEFT JOIN action_plans ON user.user_id = action_plans.user_id AND action_plans.action_start_date LIKE ? AND action_plans.action_plan_type = '1' AND action_plans.task_finished_flg = '0' AND action_plans.delete_flag = '0'", day+"%").
			Where("user.delete_flg <> ?", "1")

		if userID != "" {
			db = db.Where("user.user_id = ?", userID)
		}

		db.Order("user.user_name ASC").Scan(&rows)
		reportAll[day] = rows
	}

	var users []struct {
		UserID   string `json:"user_id"`
		UserName string `json:"user_name"`
		LoginID  string `json:"login_id"`
	}
	database.DB.Table("user").Select("user_id, user_name, login_id").Where("delete_flg <> ?", "1").Scan(&users)

	c.JSON(http.StatusOK, gin.H{"data": reportAll, "user": users, "date": date})
}

// GetStaffDailyReportRegister handles GET /dashboard/staff-daily-report/register
func GetStaffDailyReportRegister(c *gin.Context) {
	userID := c.GetString("user_id")
	date := c.DefaultQuery("setting_date", time.Now().Format("2006-01-02"))

	var userName string
	database.DB.Table("user").Select("user_name").Where("user_id = ?", userID).Scan(&userName)

	var data models.UserDayoffSetting
	database.DB.Where("user_id = ? AND setting_date = ?", userID, date).First(&data)

	var actionPlans []models.ActionPlan
	database.DB.Where("(user_id = ? OR sub_user_id = ?) AND delete_flag = 0 AND action_start_date = ?", userID, userID, date).
		Order("action_plan_type DESC, action_start_time ASC").Find(&actionPlans)

	c.JSON(http.StatusOK, gin.H{
		"date":         date,
		"data":         data,
		"user_name":    userName,
		"user_id":      userID,
		"action_plans": actionPlans,
	})
}

// PostStaffDailyReportRegister handles POST /dashboard/staff-daily-report/register
func PostStaffDailyReportRegister(c *gin.Context) {
	var req struct {
		SettingDate     string `json:"setting_date" binding:"required"`
		UserID          string `json:"user_id" binding:"required"`
		Biko            string `json:"biko"`
		TimeIn          string `json:"time_in"`
		TimeOut         string `json:"time_out"`
		EndedHour       string `json:"ended_hour"`
		EndedMinute     string `json:"ended_minute"`
		StartedHour     string `json:"started_hour"`
		StartedMinute   string `json:"started_minute"`
		OvertimeReason  string `json:"overtime_reason"`
		OvertimeComment string `json:"overtime_comment"`
		ConfirmNotifi   string `json:"confirm_notifi"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	var existing models.UserDayoffSetting
	found := database.DB.Where("user_id = ? AND setting_date = ?", req.UserID, req.SettingDate).First(&existing).Error == nil

	updates := map[string]interface{}{
		"biko":                    req.Biko,
		"time_in":                 req.TimeIn,
		"time_out":                req.TimeOut,
		"overtime_ended_hour":     req.EndedHour,
		"overtime_ended_minute":   req.EndedMinute,
		"overtime_started_hour":   req.StartedHour,
		"overtime_started_minute": req.StartedMinute,
		"overtime_reason":         req.OvertimeReason,
		"overtime_comment":        req.OvertimeComment,
	}

	if found {
		if req.ConfirmNotifi == "true" {
			loggedUserID := c.GetString("user_id")
			uc := existing.UserID
			if uc == "" {
				uc = loggedUserID
			}
			updates["user_confirm"] = uc + "," + loggedUserID
		}
		database.DB.Model(&models.UserDayoffSetting{}).Where("user_id = ? AND setting_date = ?", req.UserID, req.SettingDate).Updates(updates)
	} else {
		newSetting := models.UserDayoffSetting{
			UserID: req.UserID,
		}
		database.DB.Create(&newSetting)
		database.DB.Model(&newSetting).Updates(updates)
	}

	c.JSON(http.StatusOK, gin.H{"message": "登録が完了しました"})
}

// GetStaffDailyReportManager handles GET /dashboard/staff-daily-report/manager
func GetStaffDailyReportManager(c *gin.Context) {
	currentUserID := c.GetString("user_id")
	dateType := c.DefaultQuery("setting_date_type", "1")
	userIDFilter := c.Query("user_id")

	var date string
	if dateType == "2" {
		date = c.DefaultQuery("setting_date_month", time.Now().Format("2006-01"))
	} else {
		date = c.DefaultQuery("setting_date_day", time.Now().Format("2006-01-02"))
	}

	type ManagerRow struct {
		UserID          string  `json:"user_id"`
		UserName        string  `json:"user_name"`
		CrewName        string  `json:"crew_name"`
		Biko            string  `json:"biko"`
		UserConfirm     string  `json:"user_confirm"`
		OvertimeComment string  `json:"overtime_comment"`
		IsReport        *string `json:"is_report"`
		Title           string  `json:"title"`
		ActionStartDate string  `json:"action_start_date"`
		ActionStartTime string  `json:"action_start_time"`
		ActionEndedTime string  `json:"action_ended_time"`
	}

	reportAll := make(map[string][]ManagerRow)

	if dateType == "2" {
		year, _ := strconv.Atoi(date[:4])
		month, _ := strconv.Atoi(date[5:7])
		daysInMonth := time.Date(year, time.Month(month)+1, 0, 0, 0, 0, 0, time.UTC).Day()

		for i := 1; i <= daysInMonth; i++ {
			day := time.Date(year, time.Month(month), i, 0, 0, 0, 0, time.UTC).Format("2006-01-02")
			var rows []ManagerRow
			db := database.DB.Table("user").
				Select(`user.user_id, user.user_name, crews.crew_name,
					user_dayoff_settings.biko, user_dayoff_settings.user_confirm,
					user_dayoff_settings.overtime_comment, user_dayoff_settings.user_id as is_report,
					action_plans.title, action_plans.action_start_date,
					action_plans.action_start_time, action_plans.action_ended_time`).
				Joins("JOIN crews ON user.user_id = crews.crew_code AND crews.retire_date IS NULL").
				Joins("LEFT JOIN user_dayoff_settings ON user.user_id = user_dayoff_settings.user_id AND user_dayoff_settings.setting_date LIKE ?", day+"%").
				Joins("LEFT JOIN action_plans ON user.user_id = action_plans.user_id AND action_plans.action_start_date LIKE ? AND action_plans.action_plan_type = '1' AND action_plans.task_finished_flg = '0' AND action_plans.delete_flag = '0'", day+"%").
				Where("(user.daily_report_auth_id LIKE ? OR user.daily_report_auth_id LIKE ? OR user.daily_report_auth_id LIKE ? OR user.daily_report_auth_id = ?)",
					currentUserID+",%", "%,"+currentUserID+",%", "%,"+currentUserID, currentUserID).
				Where("user.delete_flg <> ?", "1")

			if userIDFilter != "" {
				db = db.Where("user.user_id = ?", userIDFilter)
			}
			db.Order("user.user_name ASC").Scan(&rows)
			reportAll[day] = rows
		}
	} else {
		var rows []ManagerRow
		db := database.DB.Table("user").
			Select(`user.user_id, user.user_name, crews.crew_name,
				user_dayoff_settings.biko, user_dayoff_settings.user_confirm,
				user_dayoff_settings.overtime_comment, user_dayoff_settings.user_id as is_report,
				action_plans.title, action_plans.action_start_date,
				action_plans.action_start_time, action_plans.action_ended_time`).
			Joins("JOIN crews ON user.user_id = crews.crew_code AND crews.retire_date IS NULL").
			Joins("LEFT JOIN user_dayoff_settings ON user.user_id = user_dayoff_settings.user_id AND user_dayoff_settings.setting_date LIKE ?", date+"%").
			Joins("LEFT JOIN action_plans ON user.user_id = action_plans.user_id AND action_plans.action_start_date LIKE ? AND action_plans.action_plan_type = '1' AND action_plans.task_finished_flg = '0' AND action_plans.delete_flag = '0'", date+"%").
			Where("(user.daily_report_auth_id LIKE ? OR user.daily_report_auth_id LIKE ? OR user.daily_report_auth_id LIKE ? OR user.daily_report_auth_id = ?)",
				currentUserID+",%", "%,"+currentUserID+",%", "%,"+currentUserID, currentUserID).
			Where("user.delete_flg <> ?", "1")

		if userIDFilter != "" {
			db = db.Where("user.user_id = ?", userIDFilter)
		}
		db.Order("user.user_name ASC").Scan(&rows)
		reportAll[date] = rows
	}

	c.JSON(http.StatusOK, gin.H{"data": reportAll, "date": date, "setting_date_type": dateType})
}

// GetStaffDailyReportManagerDetail handles GET /dashboard/staff-daily-report/manager/detail
func GetStaffDailyReportManagerDetail(c *gin.Context) {
	userID := c.Query("user_id")
	date := c.Query("date")

	var data struct {
		models.UserDayoffSetting
		UserName string `json:"user_name"`
	}
	database.DB.Table("user_dayoff_settings").
		Select("user_dayoff_settings.*, user.user_name").
		Joins("LEFT JOIN user ON user_dayoff_settings.user_id = user.user_id").
		Where("user_dayoff_settings.user_id = ? AND user_dayoff_settings.setting_date = ?", userID, date).
		Scan(&data)

	var actionPlans []models.ActionPlan
	database.DB.Where("(user_id = ? OR sub_user_id = ?) AND delete_flag = 0 AND action_start_date = ?", userID, userID, date).
		Order("action_plan_type DESC, action_start_time ASC").Find(&actionPlans)

	c.JSON(http.StatusOK, gin.H{"data": data, "action_plans": actionPlans})
}

// PostStaffDailyReportSearch handles POST /dashboard/staff-daily-report/search
func PostStaffDailyReportSearch(c *gin.Context) {
	var req struct {
		SettingDate string `json:"setting_date"`
		UserID      string `json:"user_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	var data1 []models.UserDayoffSetting
	database.DB.Where("user_id = ? AND setting_date = ?", req.UserID, req.SettingDate).Find(&data1)

	var data2 []models.ActionPlan
	database.DB.Where("(user_id = ? OR sub_user_id = ?) AND delete_flag = 0 AND action_start_date = ?",
		req.UserID, req.UserID, req.SettingDate).
		Order("action_plan_type DESC, action_start_time ASC").
		Select("action_start_time, action_ended_time, prop_name, title").
		Find(&data2)

	c.JSON(http.StatusOK, gin.H{"data1": data1, "data2": data2})
}
