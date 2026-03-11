package handlers

import (
	"fmt"
	"kms_golang/database"
	"kms_golang/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetActionPlanList handles GET /dashboard/action-plan/list
func GetActionPlanList(c *gin.Context) {
	var plans []models.ActionPlan
	db := database.DB

	if userID := c.Query("user_id"); userID != "" {
		db = db.Where("user_id = ?", userID)
	}
	if propCd := c.Query("prop_cd"); propCd != "" {
		db = db.Where("prop_cd = ?", propCd)
	}
	if planDate := c.Query("plan_date"); planDate != "" {
		db = db.Where("plan_date = ?", planDate)
	}

	if err := db.Order("plan_date DESC, user_id ASC").Find(&plans).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"action_plans": plans})
}

// GetActionPlanDetail handles GET /dashboard/action-plan/:user_id/:serial
func GetActionPlanDetail(c *gin.Context) {
	userID := c.Param("user_id")
	serial := c.Param("serial")

	var plan models.ActionPlan
	if err := database.DB.Where("user_id = ? AND action_serial_num = ?", userID, serial).First(&plan).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "アクションプランが見つかりません"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"action_plan": plan})
}

// PostActionPlanCreate handles POST /dashboard/action-plan/create
func PostActionPlanCreate(c *gin.Context) {
	var req models.ActionPlan
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	// Generate serial number
	var maxSerial struct{ MaxSerial *int }
	database.DB.Raw("SELECT MAX(CAST(action_serial_num AS UNSIGNED)) as max_serial FROM action_plan WHERE user_id = ?", req.UserID).Scan(&maxSerial)
	if maxSerial.MaxSerial == nil {
		req.ActionSerialNum = "1"
	} else {
		next := *maxSerial.MaxSerial + 1
		req.ActionSerialNum = fmt.Sprintf("%d", next)
	}

	if err := database.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"action_plan": req})
}

// PutActionPlanUpdate handles PUT /dashboard/action-plan/:user_id/:serial
func PutActionPlanUpdate(c *gin.Context) {
	userID := c.Param("user_id")
	serial := c.Param("serial")

	var plan models.ActionPlan
	if err := database.DB.Where("user_id = ? AND action_serial_num = ?", userID, serial).First(&plan).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "アクションプランが見つかりません"})
		return
	}

	if err := c.ShouldBindJSON(&plan); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	plan.UserID = userID
	plan.ActionSerialNum = serial
	if err := database.DB.Save(&plan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"action_plan": plan})
}

// DeleteActionPlan handles DELETE /dashboard/action-plan/:user_id/:serial
func DeleteActionPlan(c *gin.Context) {
	userID := c.Param("user_id")
	serial := c.Param("serial")

	if err := database.DB.Where("user_id = ? AND action_serial_num = ?", userID, serial).Delete(&models.ActionPlan{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "アクションプランを削除しました"})
}

// GetWorkScheduleList handles GET /dashboard/work-schedule/list
func GetWorkScheduleList(c *gin.Context) {
	var schedules []models.WorkSchedule
	db := database.DB

	if userID := c.Query("user_id"); userID != "" {
		db = db.Where("user_id = ?", userID)
	}
	if scheduleDate := c.Query("schedule_date"); scheduleDate != "" {
		db = db.Where("schedule_date = ?", scheduleDate)
	}

	if err := db.Order("schedule_date DESC").Find(&schedules).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"work_schedules": schedules})
}

// PostWorkScheduleCreate handles POST /dashboard/work-schedule/create
func PostWorkScheduleCreate(c *gin.Context) {
	var req models.WorkSchedule
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	if err := database.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"work_schedule": req})
}

// GetTaskList handles GET /dashboard/task/list
func GetTaskList(c *gin.Context) {
	var tasks []models.Task
	db := database.DB.Where("delete_flag = 0")

	if userID := c.Query("user_id"); userID != "" {
		db = db.Where("user_id = ?", userID)
	}
	if propCd := c.Query("prop_cd"); propCd != "" {
		db = db.Where("prop_cd = ?", propCd)
	}
	if taskStatus := c.Query("task_status"); taskStatus != "" {
		db = db.Where("task_status = ?", taskStatus)
	}

	if err := db.Order("deadline ASC").Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// PostTaskCreate handles POST /dashboard/task/create
func PostTaskCreate(c *gin.Context) {
	var req models.Task
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	delFlagTask := int16(0)
	req.DeleteFlag = &delFlagTask
	if err := database.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"task": req})
}

// PutTaskUpdate handles PUT /dashboard/task/:id
func PutTaskUpdate(c *gin.Context) {
	idStr := c.Param("id")
	tid, terr := strconv.ParseUint(idStr, 10, 64)
	if terr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "無効なID"})
		return
	}
	var task models.Task
	if err := database.DB.Where("id = ? AND delete_flag = 0", tid).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "タスクが見つかりません"})
		return
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	task.ID = uint(tid)
	if err := database.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task": task})
}

// DeleteTask handles DELETE /dashboard/task/:id (soft delete)
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Model(&models.Task{}).Where("id = ?", id).Update("delete_flag", 1).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "タスクを削除しました"})
}

// GetWorkBukkenInfo handles GET /dashboard/work/bukken-info
func GetWorkBukkenInfo(c *gin.Context) {
	propCd := c.Query("prop_cd")
	var propBasic models.PropBasic
	if err := database.DB.Where("prop_cd = ? AND delete_flag = 0", propCd).First(&propBasic).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "物件が見つかりません"})
		return
	}

	var orders []models.Order
	database.DB.Where("prop_cd = ?", propCd).Find(&orders)

	var receptions []models.Reception
	database.DB.Where("prop_cd = ? AND delete_flag = 0", propCd).Find(&receptions)

	c.JSON(http.StatusOK, gin.H{"prop_basic": propBasic, "orders": orders, "receptions": receptions})
}

// GetWorkDropdown handles GET /dashboard/work/dropdown
func GetWorkDropdown(c *gin.Context) {
	var users []models.User
	database.DB.Where("delete_flg = 0").Select("user_id, user_name").Find(&users)
	for i := range users {
		users[i].Password = nil
	}

	var props []models.PropBasic
	database.DB.Where("delete_flag = 0").Select("prop_cd, prop_name").Find(&props)

	var sekosakis []models.Sekosaki
	database.DB.Where("delete_flag = 0").Select("sekosaki_cd, sekosaki_name").Find(&sekosakis)
	for i := range sekosakis {
		sekosakis[i].SekosakiPassword = nil
	}

	c.JSON(http.StatusOK, gin.H{"users": users, "prop_basics": props, "sekosakis": sekosakis})
}

// GetWorkWeekNav handles GET /dashboard/work/week-nav
func GetWorkWeekNav(c *gin.Context) {
	date := c.Query("date")
	userID := c.Query("user_id")

	var schedules []models.WorkSchedule
	db := database.DB.Where("delete_flag = 0")
	if date != "" {
		db = db.Where("work_date = ?", date)
	}
	if userID != "" {
		db = db.Where("user_id = ?", userID)
	}
	db.Order("work_date ASC").Find(&schedules)

	var results []models.WorkResult
	db2 := database.DB.Where("delete_flag = 0")
	if date != "" {
		db2 = db2.Where("work_date = ?", date)
	}
	if userID != "" {
		db2 = db2.Where("user_id = ?", userID)
	}
	db2.Order("work_date ASC").Find(&results)

	c.JSON(http.StatusOK, gin.H{"work_schedules": schedules, "work_results": results})
}

// PostWorkSearch handles POST /dashboard/work/search
func PostWorkSearch(c *gin.Context) {
	var req struct {
		UserID   string `json:"user_id"`
		PropCd   string `json:"prop_cd"`
		DateFrom string `json:"date_from"`
		DateTo   string `json:"date_to"`
		WorkType string `json:"work_type"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	db := database.DB.Where("delete_flag = 0")
	if req.UserID != "" {
		db = db.Where("user_id = ?", req.UserID)
	}
	if req.PropCd != "" {
		db = db.Where("prop_cd = ?", req.PropCd)
	}
	if req.DateFrom != "" {
		db = db.Where("work_date >= ?", req.DateFrom)
	}
	if req.DateTo != "" {
		db = db.Where("work_date <= ?", req.DateTo)
	}
	if req.WorkType != "" {
		db = db.Where("work_type = ?", req.WorkType)
	}

	var schedules []models.WorkSchedule
	db.Order("work_date ASC").Find(&schedules)
	c.JSON(http.StatusOK, gin.H{"work_schedules": schedules})
}

// PostWorkRegularWork handles POST /dashboard/work/regular-work
func PostWorkRegularWork(c *gin.Context) {
	var req struct {
		UserID   string `json:"user_id" binding:"required"`
		PropCd   string `json:"prop_cd"`
		WorkDate string `json:"work_date" binding:"required"`
		WorkType *int16 `json:"work_type"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	delFlag := int16(0)
	userID := req.UserID
	ws := models.WorkSchedule{
		PropCd:     &req.PropCd,
		WorkDate:   &req.WorkDate,
		WorkType:   req.WorkType,
		DeleteFlag: &delFlag,
		RegistUser: &userID,
	}
	database.DB.Create(&ws)
	c.JSON(http.StatusCreated, gin.H{"work_schedule": ws})
}

// GetWorkResultList handles GET /dashboard/work-result/list
func GetWorkResultList(c *gin.Context) {
	var results []models.WorkResult
	db := database.DB.Where("delete_flag = 0")
	if userID := c.Query("user_id"); userID != "" {
		db = db.Where("user_id = ?", userID)
	}
	if propCd := c.Query("prop_cd"); propCd != "" {
		db = db.Where("prop_cd = ?", propCd)
	}
	if workDate := c.Query("work_date"); workDate != "" {
		db = db.Where("work_date = ?", workDate)
	}
	db.Order("work_date DESC").Find(&results)
	c.JSON(http.StatusOK, gin.H{"work_results": results})
}

// PostWorkResultCreate handles POST /dashboard/work-result/create
func PostWorkResultCreate(c *gin.Context) {
	var req models.WorkResult
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
	c.JSON(http.StatusCreated, gin.H{"work_result": req})
}

// PutWorkResultUpdate handles PUT /dashboard/work-result/:id
func PutWorkResultUpdate(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "無効なID"})
		return
	}
	var wr models.WorkResult
	if err := database.DB.Where("id = ? AND delete_flag = 0", id).First(&wr).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "見つかりません"})
		return
	}
	if err := c.ShouldBindJSON(&wr); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	wr.ID = uint(id)
	database.DB.Save(&wr)
	c.JSON(http.StatusOK, gin.H{"work_result": wr})
}

// DeleteWorkResult handles DELETE /dashboard/work-result/:id (soft delete)
func DeleteWorkResult(c *gin.Context) {
	id := c.Param("id")
	database.DB.Model(&models.WorkResult{}).Where("id = ?", id).Update("delete_flag", 1)
	c.JSON(http.StatusOK, gin.H{"message": "削除しました"})
}

// PutWorkScheduleUpdate handles PUT /dashboard/work-schedule/:id
func PutWorkScheduleUpdate(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "無効なID"})
		return
	}
	var ws models.WorkSchedule
	if err := database.DB.Where("id = ? AND delete_flag = 0", id).First(&ws).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "見つかりません"})
		return
	}
	if err := c.ShouldBindJSON(&ws); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	ws.ID = uint(id)
	database.DB.Save(&ws)
	c.JSON(http.StatusOK, gin.H{"work_schedule": ws})
}

// DeleteWorkSchedule handles DELETE /dashboard/work-schedule/:id (soft delete)
func DeleteWorkSchedule(c *gin.Context) {
	id := c.Param("id")
	database.DB.Model(&models.WorkSchedule{}).Where("id = ?", id).Update("delete_flag", 1)
	c.JSON(http.StatusOK, gin.H{"message": "削除しました"})
}

// GetWorkCancelList handles GET /dashboard/work-cancel/list
func GetWorkCancelList(c *gin.Context) {
	var cancels []models.WorkCancel
	db := database.DB
	if propCd := c.Query("prop_cd"); propCd != "" {
		db = db.Where("prop_cd = ?", propCd)
	}
	db.Order("id DESC").Find(&cancels)
	c.JSON(http.StatusOK, gin.H{"work_cancels": cancels})
}

// PostWorkCancelCreate handles POST /dashboard/work-cancel/create
func PostWorkCancelCreate(c *gin.Context) {
	var req models.WorkCancel
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Create(&req)
	c.JSON(http.StatusCreated, gin.H{"work_cancel": req})
}

// GetTaskDetail handles GET /dashboard/task/:id
func GetTaskDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "無効なID"})
		return
	}
	var task models.Task
	if err := database.DB.Where("id = ? AND delete_flag = 0", id).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "タスクが見つかりません"})
		return
	}
	var imgs []models.TaskImg
	database.DB.Where("task_id = ?", id).Find(&imgs)
	c.JSON(http.StatusOK, gin.H{"task": task, "task_images": imgs})
}
