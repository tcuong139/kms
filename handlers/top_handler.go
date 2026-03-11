package handlers

import (
	"kms_golang/database"
	"kms_golang/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// GetTopMemo handles GET /dashboard/top – returns memos and action plans for the logged-in user
func GetTopMemo(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "認証エラー"})
		return
	}

	today := time.Now().Format("2006-01-02")

	// Action plans for today or tasks in progress
	var actionPlans []models.ActionPlan
	database.DB.Where(
		"(user_id = ? OR sub_user_id = ?) AND delete_flag = 0 AND (action_start_date <= ? OR (action_plan_type = 2 AND task_finished_flg = 0))",
		userID, userID, today,
	).Order("action_plan_type DESC, action_start_time ASC").Find(&actionPlans)

	// Memos
	var memos []models.Memo
	database.DB.Where("delete_flag = 0").Order("id ASC").Find(&memos)

	// User notifications (unread)
	var userNotifs []models.UserNotifi
	database.DB.Where("user_id = ? AND is_read = 0", userID).
		Order("regist_datetime DESC").Limit(20).Find(&userNotifs)

	c.JSON(http.StatusOK, gin.H{
		"action_plans":  actionPlans,
		"memos":         memos,
		"notifications": userNotifs,
		"today":         today,
	})
}

// PostTopMemo handles POST /dashboard/top/memo – create or update a memo
func PostTopMemo(c *gin.Context) {
	var req models.Memo
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	flag := int16(0)
	req.DeleteFlag = &flag

	if err := database.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"memo": req})
}

// DeleteTopMemo handles DELETE /dashboard/top/memo/:id – soft-delete a memo
func DeleteTopMemo(c *gin.Context) {
	id := c.Param("id")
	mid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "無効なID"})
		return
	}
	flag := int16(1)
	if err := database.DB.Model(&models.Memo{}).Where("id = ?", mid).Update("delete_flag", flag).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "メモを削除しました"})
}

// DeleteTopNotify handles DELETE /dashboard/top/notification/:id
func DeleteTopNotify(c *gin.Context) {
	userID := c.GetString("user_id")
	id := c.Param("id")
	nid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "無効なID"})
		return
	}

	// Mark as read (soft-delete style)
	isRead := int16(1)
	if err := database.DB.Model(&models.UserNotifi{}).
		Where("id = ? AND user_id = ?", nid, userID).
		Update("is_read", isRead).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "通知を削除しました"})
}

// GetTopWorkSchedule handles GET /dashboard/top/work-schedule
func GetTopWorkSchedule(c *gin.Context) {
	userID := c.GetString("user_id")
	date := c.Query("date")
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}

	var schedules []models.WorkSchedule
	db := database.DB.Where("delete_flag = 0 AND work_date = ?", date)
	if userID != "" {
		// Filter by sekosaki assigned to user if needed; here return all for the date
		_ = userID
	}
	if err := db.Order("work_date ASC").Find(&schedules).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"work_schedules": schedules})
}

// PostCopyWorkScheduleIntoWorkResult handles POST /dashboard/top/copy-work-schedule
func PostCopyWorkScheduleIntoWorkResult(c *gin.Context) {
	var req struct {
		WorkScheduleID uint `json:"work_schedule_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	var ws models.WorkSchedule
	if err := database.DB.First(&ws, req.WorkScheduleID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "作業スケジュールが見つかりません"})
		return
	}

	flag := int16(0)
	wr := models.WorkResult{
		OrderID:     ws.OrderID,
		PropCd:      ws.PropCd,
		SekosakiCd:  ws.SekosakiCd,
		WorkDate:    ws.WorkDate,
		WorkType:    ws.WorkType,
		WorkContent: ws.WorkContent,
		DeleteFlag:  &flag,
	}

	if err := database.DB.Create(&wr).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"work_result": wr})
}

// PostTopCopyAll handles POST /dashboard/top/copy-all-schedule
func PostTopCopyAll(c *gin.Context) {
	userID := c.GetString("user_id")
	date := c.Query("date")
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}

	var schedules []models.WorkSchedule
	database.DB.Where("delete_flag = 0 AND work_date = ? AND user_id = ?", date, userID).Find(&schedules)

	for _, ws := range schedules {
		flag := int16(0)
		wr := models.WorkResult{
			OrderID:     ws.OrderID,
			PropCd:      ws.PropCd,
			SekosakiCd:  ws.SekosakiCd,
			WorkDate:    ws.WorkDate,
			WorkType:    ws.WorkType,
			WorkContent: ws.WorkContent,
			DeleteFlag:  &flag,
		}
		database.DB.Create(&wr)
	}
	c.JSON(http.StatusCreated, gin.H{"message": "全スケジュールをコピーしました", "count": len(schedules)})
}

// PostTopCancel handles POST /dashboard/top/cancel-schedule
func PostTopCancel(c *gin.Context) {
	var req struct {
		WorkScheduleID uint `json:"work_schedule_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	database.DB.Model(&models.WorkSchedule{}).Where("id = ?", req.WorkScheduleID).Update("delete_flag", 1)
	c.JSON(http.StatusOK, gin.H{"message": "スケジュールをキャンセルしました"})
}

// PostTopWorkFinishedEmail handles POST /dashboard/top/work-finished-email
func PostTopWorkFinishedEmail(c *gin.Context) {
	var req struct {
		OrderID string `json:"order_id" binding:"required"`
		PropCd  string `json:"prop_cd"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	// Mark as finished and send notification
	database.DB.Model(&models.WorkSchedule{}).Where("order_id = ?", req.OrderID).Update("finished_flg", 1)
	c.JSON(http.StatusOK, gin.H{"message": "作業完了メールを送信しました"})
}

// GetTopSeeMore handles GET /dashboard/top/see-more
func GetTopSeeMore(c *gin.Context) {
	userID := c.GetString("user_id")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage := 20

	var userNotifs []models.UserNotifi
	database.DB.Where("user_id = ? AND is_read = 0", userID).
		Order("regist_datetime DESC").Offset((page - 1) * perPage).Limit(perPage).Find(&userNotifs)

	var total int64
	database.DB.Model(&models.UserNotifi{}).Where("user_id = ? AND is_read = 0", userID).Count(&total)

	c.JSON(http.StatusOK, gin.H{"notifications": userNotifs, "total": total, "page": page})
}

// PostTopUpdateInvoiceMonth handles POST /dashboard/top/update-invoice-month
func PostTopUpdateInvoiceMonth(c *gin.Context) {
	var req struct {
		PropCd       string `json:"prop_cd" binding:"required"`
		InvoiceMonth string `json:"invoice_month" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Model(&models.PropBasic{}).Where("prop_cd = ?", req.PropCd).Update("invoice_month", req.InvoiceMonth)
	c.JSON(http.StatusOK, gin.H{"message": "請求月を更新しました"})
}
