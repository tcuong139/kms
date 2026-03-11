package handlers

import (
	"kms_golang/database"
	"kms_golang/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

// GetWorkManagementList handles GET /dashboard/work-management/list
func GetWorkManagementList(c *gin.Context) {
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

	db := database.DB.Model(&models.Task{}).Where("delete_flag <> 1")

	if code := c.Query("work_code"); code != "" {
		db = db.Where("work_code LIKE ?", "%"+code+"%")
	}
	if name := c.Query("work_name"); name != "" {
		db = db.Where("work_name LIKE ?", "%"+name+"%")
	}
	if omitted := c.Query("task_omitted"); omitted != "" {
		db = db.Where("task_omitted LIKE ?", "%"+omitted+"%")
	}
	if kana := c.Query("task_kana_name"); kana != "" {
		db = db.Where("task_kana_name LIKE ?", "%"+kana+"%")
	}
	if ft := c.Query("fulltime"); ft == "1" {
		if pt := c.Query("parttime"); pt == "1" {
			db = db.Where("type IN ?", []int{1, 2})
		} else {
			db = db.Where("type = ?", 1)
		}
	} else if pt := c.Query("parttime"); pt == "1" {
		db = db.Where("type = ?", 2)
	}

	var total int64
	db.Count(&total)

	var tasks []models.Task
	offset := (page - 1) * perPage
	if err := db.Offset(offset).Limit(perPage).Order("work_code ASC").Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks, "total": total, "page": page, "per_page": perPage})
}

// GetWorkManagementRegister handles GET /dashboard/work-management/register
func GetWorkManagementRegister(c *gin.Context) {
	var units []models.Unit
	database.DB.Where("is_active = 1").Order("id").Find(&units)

	workCD := c.Query("work_cd")
	if workCD != "" {
		var task models.Task
		if err := database.DB.Where("work_code = ?", workCD).First(&task).Error; err == nil {
			c.JSON(http.StatusOK, gin.H{"units": units, "task": task, "work_cd": workCD})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"units": units, "work_cd": workCD})
}

// PostWorkManagementRegister handles POST /dashboard/work-management/register
func PostWorkManagementRegister(c *gin.Context) {
	var req struct {
		WorkCode     string  `json:"work_code" binding:"required"`
		WorkName     string  `json:"work_name"`
		TaskOmitted  string  `json:"task_omitted"`
		TaskKanaName string  `json:"task_kana_name"`
		UnitPrice    float64 `json:"unit_price"`
		Type         string  `json:"type"`
		Unit         int     `json:"unit"`
		Comment      string  `json:"comment"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	var task models.Task
	existing := database.DB.Where("work_code = ?", req.WorkCode).First(&task).Error == nil

	taskType := int16(1)
	if req.Type == "parttime" {
		taskType = 2
	}
	unitVal := 1
	if taskType == 2 {
		unitVal = req.Unit
	}

	if existing {
		database.DB.Model(&task).Updates(map[string]interface{}{
			"work_name":      req.WorkName,
			"task_omitted":   req.TaskOmitted,
			"task_kana_name": req.TaskKanaName,
			"unit_price":     req.UnitPrice,
			"type":           taskType,
			"unit":           unitVal,
			"comment":        req.Comment,
		})
	} else {
		task = models.Task{}
		delFlag := int16(0)
		task.DeleteFlag = &delFlag
		database.DB.Exec("INSERT INTO tasks (work_code, work_name, task_omitted, task_kana_name, unit_price, type, unit, comment, delete_flag) VALUES (?, ?, ?, ?, ?, ?, ?, ?, 0)",
			req.WorkCode, req.WorkName, req.TaskOmitted, req.TaskKanaName, req.UnitPrice, taskType, unitVal, req.Comment)
	}

	c.JSON(http.StatusOK, gin.H{"message": "登録を完了しました"})
}

// GetWorkManagementExportExcel handles GET /dashboard/work-management/export-excel
func GetWorkManagementExportExcel(c *gin.Context) {
	db := database.DB.Model(&models.Task{}).Where("delete_flag <> 1")

	if name := c.Query("work_name"); name != "" {
		db = db.Where("work_name LIKE ?", "%"+name+"%")
	}
	if kana := c.Query("task_kana_name"); kana != "" {
		db = db.Where("task_kana_name LIKE ?", "%"+kana+"%")
	}
	if omitted := c.Query("task_omitted"); omitted != "" {
		db = db.Where("task_omitted LIKE ?", "%"+omitted+"%")
	}

	var tasks []models.Task
	db.Order("work_code ASC").Find(&tasks)

	f := excelize.NewFile()
	sheet := "Sheet1"
	headers := []string{"作業コード", "作業名", "略称", "カナ名", "単価", "タイプ", "単位"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}
	for i, t := range tasks {
		row := i + 2
		f.SetCellValue(sheet, cellName(1, row), ptrStr(t.AcceptNumber))
		f.SetCellValue(sheet, cellName(2, row), ptrStr(t.TaskContent))
		f.SetCellValue(sheet, cellName(3, row), ptrStr(t.TaskContent))
	}

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=work_management_list.xlsx")
	f.Write(c.Writer)
}

func cellName(col, row int) string {
	name, _ := excelize.CoordinatesToCellName(col, row)
	return name
}

func ptrStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
