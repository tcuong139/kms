package handlers

import (
	"kms_golang/database"
	"kms_golang/models"
	"kms_golang/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

// GetPropertyList handles GET /dashboard/property/list
func GetPropertyList(c *gin.Context) {
	var properties []models.PropBasic
	db := database.DB.Where("delete_flag = 0")

	if propCd := c.Query("prop_cd"); propCd != "" {
		db = db.Where("prop_cd LIKE ?", "%"+propCd+"%")
	}
	if name := c.Query("prop_name"); name != "" {
		db = db.Where("prop_name LIKE ?", "%"+name+"%")
	}
	if customerCd := c.Query("customer_cd"); customerCd != "" {
		db = db.Where("customer_cd = ?", customerCd)
	}

	if err := db.Order("prop_cd ASC").Find(&properties).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"properties": properties})
}

// GetPropertyDetail handles GET /dashboard/property/:cd
func GetPropertyDetail(c *gin.Context) {
	cd := c.Param("cd")
	var property models.PropBasic
	if err := database.DB.Where("prop_cd = ? AND delete_flag = 0", cd).First(&property).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "物件が見つかりません"})
		return
	}

	// Load related data
	var propOther models.PropBasicOther
	database.DB.Where("prop_cd = ?", cd).First(&propOther)

	var facilities models.PropFacilities
	database.DB.Where("prop_cd = ?", cd).First(&facilities)

	var propCustomers []models.PropCustomer
	database.DB.Where("prop_cd = ?", cd).Find(&propCustomers)

	var propImgs []models.PropImg
	database.DB.Where("prop_cd = ?", cd).Find(&propImgs)

	c.JSON(http.StatusOK, gin.H{
		"property":       property,
		"prop_other":     propOther,
		"facilities":     facilities,
		"prop_customers": propCustomers,
		"prop_imgs":      propImgs,
	})
}

// PostPropertyCreate handles POST /dashboard/property/create
func PostPropertyCreate(c *gin.Context) {
	var req models.PropBasic
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	cd, err := utils.GeneratePropertyID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	req.PropCd = cd
	delFlagProp := int16(0)
	req.DeleteFlag = &delFlagProp

	if err := database.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"property": req})
}

// PutPropertyUpdate handles PUT /dashboard/property/:cd
func PutPropertyUpdate(c *gin.Context) {
	cd := c.Param("cd")
	var property models.PropBasic
	if err := database.DB.Where("prop_cd = ? AND delete_flag = 0", cd).First(&property).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "物件が見つかりません"})
		return
	}

	if err := c.ShouldBindJSON(&property); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	property.PropCd = cd
	if err := database.DB.Save(&property).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"property": property})
}

// DeleteProperty handles DELETE /dashboard/property/:cd (soft delete)
func DeleteProperty(c *gin.Context) {
	cd := c.Param("cd")
	if err := database.DB.Model(&models.PropBasic{}).Where("prop_cd = ?", cd).Update("delete_flag", 1).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "物件を削除しました"})
}

// GetPropertyImages handles GET /dashboard/property/:cd/images
func GetPropertyImages(c *gin.Context) {
	cd := c.Param("cd")
	var imgs []models.PropImg
	if err := database.DB.Where("prop_cd = ?", cd).Find(&imgs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"images": imgs})
}

// GetPropertyWorkReports handles GET /dashboard/property/:cd/work-reports
func GetPropertyWorkReports(c *gin.Context) {
	cd := c.Param("cd")
	var reports []models.PropWorkReport
	if err := database.DB.Where("prop_cd = ?", cd).Order("regist_datetime DESC").Find(&reports).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"work_reports": reports})
}

// GetPropertyDropdownList handles GET /dashboard/property/dropdown
func GetPropertyDropdownList(c *gin.Context) {
	var props []models.PropBasic
	database.DB.Where("delete_flag = 0").Select("prop_cd, prop_name, customer_cd").Order("prop_cd ASC").Find(&props)
	c.JSON(http.StatusOK, gin.H{"properties": props})
}

// GetPropertyDropdownReception handles GET /dashboard/property/dropdown-reception
func GetPropertyDropdownReception(c *gin.Context) {
	propCd := c.Query("prop_cd")
	var receptions []models.Reception
	database.DB.Where("prop_cd = ? AND delete_flag = 0", propCd).Order("accept_number DESC").Find(&receptions)
	c.JSON(http.StatusOK, gin.H{"receptions": receptions})
}

// GetPropertyRegisterForm handles GET /dashboard/property/register-form
func GetPropertyRegisterForm(c *gin.Context) {
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

	c.JSON(http.StatusOK, gin.H{"customers": customers, "sekosakis": sekosakis})
}

// PostPropertyCustomerRegister handles POST /dashboard/property/:cd/customer-register
func PostPropertyCustomerRegister(c *gin.Context) {
	cd := c.Param("cd")
	var req struct {
		CustomerCd string `json:"customer_cd" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	propCust := models.PropCustomer{PropCd: cd, CustomerCd: req.CustomerCd}
	database.DB.Create(&propCust)
	c.JSON(http.StatusCreated, gin.H{"prop_customer": propCust})
}

// PostPropertySekosakiRegister handles POST /dashboard/property/:cd/sekosaki-register
func PostPropertySekosakiRegister(c *gin.Context) {
	cd := c.Param("cd")
	var req struct {
		SekosakiCd string `json:"sekosaki_cd" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Model(&models.PropBasic{}).Where("prop_cd = ?", cd).Update("sekosaki_cd", req.SekosakiCd)
	c.JSON(http.StatusOK, gin.H{"message": "施工先を登録しました"})
}

// GetPropertyExportExcel handles GET /dashboard/property/export-excel
func GetPropertyExportExcel(c *gin.Context) {
	var props []models.PropBasic
	database.DB.Where("delete_flag = 0").Order("prop_cd ASC").Find(&props)

	f := excelize.NewFile()
	sheet := "Sheet1"
	headers := []string{"物件CD", "物件名", "顧客CD", "郵便番号", "住所", "TEL"}
	for i, h := range headers {
		col := string(rune('A' + i))
		f.SetCellValue(sheet, col+"1", h)
	}
	for i, p := range props {
		row := strconv.Itoa(i + 2)
		f.SetCellValue(sheet, "A"+row, p.PropCd)
		if p.PropName != nil {
			f.SetCellValue(sheet, "B"+row, *p.PropName)
		}
		if p.CustomerCd != nil {
			f.SetCellValue(sheet, "C"+row, *p.CustomerCd)
		}
		if p.PostCode != nil {
			f.SetCellValue(sheet, "D"+row, *p.PostCode)
		}
		if p.BlockName != nil {
			f.SetCellValue(sheet, "E"+row, *p.BlockName)
		}
		if p.Tel != nil {
			f.SetCellValue(sheet, "F"+row, *p.Tel)
		}
	}

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=properties.xlsx")
	f.Write(c.Writer)
}

// GetPropertyNotifyInforExportExcel handles GET /dashboard/property/notify-infor-export-excel
func GetPropertyNotifyInforExportExcel(c *gin.Context) {
	propCd := c.Query("prop_cd")
	var notifications []models.NotificationList
	db := database.DB.Where("delete_flag = 0")
	if propCd != "" {
		db = db.Where("prop_cd = ?", propCd)
	}
	db.Find(&notifications)

	f := excelize.NewFile()
	sheet := "Sheet1"
	f.SetCellValue(sheet, "A1", "ID")
	f.SetCellValue(sheet, "B1", "タイトル")
	for i, n := range notifications {
		row := strconv.Itoa(i + 2)
		f.SetCellValue(sheet, "A"+row, n.ID)
		if n.NotifyContent != nil {
			f.SetCellValue(sheet, "B"+row, *n.NotifyContent)
		}
	}

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=property_notifications.xlsx")
	f.Write(c.Writer)
}

// PostPropertyCopy handles POST /dashboard/property/:cd/copy
func PostPropertyCopy(c *gin.Context) {
	cd := c.Param("cd")
	var original models.PropBasic
	if err := database.DB.Where("prop_cd = ? AND delete_flag = 0", cd).First(&original).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "物件が見つかりません"})
		return
	}

	newCd, err := utils.GeneratePropertyID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	original.PropCd = newCd
	delFlag := int16(0)
	original.DeleteFlag = &delFlag
	if err := database.DB.Create(&original).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"property": original})
}

// PostPropertyBrowseQuotes handles POST /dashboard/property/:cd/browse-quotes
func PostPropertyBrowseQuotes(c *gin.Context) {
	cd := c.Param("cd")
	var estimates []models.Estimate
	database.DB.Where("prop_cd = ? AND delete_flag = 0 AND activated = 1", cd).
		Order("estimate_number DESC, subnumber ASC").Find(&estimates)
	c.JSON(http.StatusOK, gin.H{"estimates": estimates})
}

// PostPropertyApproveConfirm handles POST /dashboard/property/:cd/approve-confirm
func PostPropertyApproveConfirm(c *gin.Context) {
	cd := c.Param("cd")
	var req struct {
		EstimateNumber string `json:"estimate_number" binding:"required"`
		Subnumber      string `json:"subnumber" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	database.DB.Model(&models.Estimate{}).
		Where("estimate_number = ? AND subnumber = ? AND prop_cd = ?", req.EstimateNumber, req.Subnumber, cd).
		Update("estimate_status", 2)
	c.JSON(http.StatusOK, gin.H{"message": "承認しました"})
}

// GetPropertyWorkDetailRegister handles GET /dashboard/property/:cd/work-detail-register
func GetPropertyWorkDetailRegister(c *gin.Context) {
	cd := c.Param("cd")
	var orders []models.Order
	database.DB.Where("prop_cd = ?", cd).Find(&orders)

	var sekosakis []models.Sekosaki
	database.DB.Where("delete_flag = 0").Select("sekosaki_cd, sekosaki_name").Find(&sekosakis)
	for i := range sekosakis {
		sekosakis[i].SekosakiPassword = nil
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders, "sekosakis": sekosakis})
}

// PostPropertyEditOrderDetail handles POST /dashboard/property/:cd/edit-order-detail
func PostPropertyEditOrderDetail(c *gin.Context) {
	var req struct {
		OrderID string                 `json:"order_id" binding:"required"`
		Data    map[string]interface{} `json:"data"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Model(&models.Order{}).Where("order_id = ?", req.OrderID).Updates(req.Data)
	c.JSON(http.StatusOK, gin.H{"message": "更新しました"})
}

// PostPropertyWorkReport handles POST /dashboard/property/:cd/work-report
func PostPropertyWorkReport(c *gin.Context) {
	cd := c.Param("cd")
	var req models.PropWorkReport
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	req.PropCd = cd
	database.DB.Create(&req)
	c.JSON(http.StatusCreated, gin.H{"work_report": req})
}

// GetPropertyTask handles GET /dashboard/property/:cd/tasks
func GetPropertyTask(c *gin.Context) {
	cd := c.Param("cd")
	var tasks []models.Task
	database.DB.Where("prop_cd = ? AND delete_flag = 0", cd).Order("deadline ASC").Find(&tasks)
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// PostPropertySearchTask handles POST /dashboard/property/search-task
func PostPropertySearchTask(c *gin.Context) {
	var req struct {
		PropCd  string `json:"prop_cd"`
		Keyword string `json:"keyword"`
	}
	c.ShouldBindJSON(&req)
	var tasks []models.Task
	db := database.DB.Where("delete_flag = 0")
	if req.PropCd != "" {
		db = db.Where("prop_cd = ?", req.PropCd)
	}
	if req.Keyword != "" {
		db = db.Where("task_name LIKE ?", "%"+req.Keyword+"%")
	}
	db.Find(&tasks)
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// GetPropertyPersonelInfo handles GET /dashboard/property/:cd/personnel-info
func GetPropertyPersonelInfo(c *gin.Context) {
	cd := c.Param("cd")
	var prop models.PropBasic
	database.DB.Where("prop_cd = ?", cd).First(&prop)
	if prop.CustomerCd == nil {
		c.JSON(http.StatusOK, gin.H{"personnel": []interface{}{}})
		return
	}
	var personnel []models.CustomerPersonnel
	database.DB.Where("customer_cd = ?", *prop.CustomerCd).Find(&personnel)
	c.JSON(http.StatusOK, gin.H{"personnel": personnel})
}

// PostPropertyUpdateOrderConstructionWorkSekosaki handles POST /dashboard/property/update-order-construction-work-sekosaki
func PostPropertyUpdateOrderConstructionWorkSekosaki(c *gin.Context) {
	var req struct {
		OrderConstructionWorkID uint   `json:"order_construction_work_id" binding:"required"`
		SekosakiCd              string `json:"sekosaki_cd" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Model(&models.OrderConstructionWork{}).Where("id = ?", req.OrderConstructionWorkID).
		Update("sekosaki_cd", req.SekosakiCd)
	c.JSON(http.StatusOK, gin.H{"message": "更新しました"})
}

// PostPropertyUpdateOrderConstructionWorkDetail handles POST /dashboard/property/update-order-construction-work-detail
func PostPropertyUpdateOrderConstructionWorkDetail(c *gin.Context) {
	var req struct {
		ID   uint                   `json:"id" binding:"required"`
		Data map[string]interface{} `json:"data"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Model(&models.OrderConstructionWorkDetail{}).Where("id = ?", req.ID).Updates(req.Data)
	c.JSON(http.StatusOK, gin.H{"message": "更新しました"})
}

// PostPropertySendToSekosaki handles POST /dashboard/property/send-to-sekosaki
func PostPropertySendToSekosaki(c *gin.Context) {
	var req struct {
		OrderID    string `json:"order_id" binding:"required"`
		SekosakiCd string `json:"sekosaki_cd" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Model(&models.Order{}).Where("order_id = ?", req.OrderID).
		Updates(map[string]interface{}{"sekosaki_cd": req.SekosakiCd, "seko_send_flg": 1})
	c.JSON(http.StatusOK, gin.H{"message": "施工先に送信しました"})
}

// PostPropertyDeleteOrder handles POST /dashboard/property/delete-order
func PostPropertyDeleteOrder(c *gin.Context) {
	var req struct {
		OrderID string `json:"order_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Delete(&models.Order{}, "order_id = ?", req.OrderID)
	c.JSON(http.StatusOK, gin.H{"message": "削除しました"})
}

// PostPropertySearchNotification handles POST /dashboard/property/search-notification
func PostPropertySearchNotification(c *gin.Context) {
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

// GetPropertyCommissionedWork handles GET /dashboard/property/:cd/commissioned-work
func GetPropertyCommissionedWork(c *gin.Context) {
	cd := c.Param("cd")
	var orders []models.Order
	database.DB.Where("prop_cd = ? AND order_status >= 1", cd).Find(&orders)
	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

// PostPropertyCreateInvoice handles POST /dashboard/property/:cd/create-invoice
func PostPropertyCreateInvoice(c *gin.Context) {
	_ = c.Param("cd")
	var req struct {
		InvoiceMonth string `json:"invoice_month" binding:"required"`
		CustomerCd   string `json:"customer_cd" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	invoice := models.CusInvoiceDetail{
		CustomerCd:   req.CustomerCd,
		InvoiceMonth: req.InvoiceMonth,
	}
	delFlag := int16(0)
	invoice.DeleteFlag = &delFlag
	database.DB.Create(&invoice)
	c.JSON(http.StatusCreated, gin.H{"invoice": invoice})
}

// PostPropertyActionPlan handles POST /dashboard/property/:cd/action-plan
func PostPropertyActionPlan(c *gin.Context) {
	var req models.ActionPlan
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Create(&req)
	c.JSON(http.StatusCreated, gin.H{"action_plan": req})
}

// GetPropertyWorkDetailList handles GET /dashboard/property/:cd/work-detail-list
func GetPropertyWorkDetailList(c *gin.Context) {
	cd := c.Param("cd")
	var orders []models.Order
	database.DB.Where("prop_cd = ?", cd).Find(&orders)

	var orderDetails []models.OrderPropmanageWork
	for _, o := range orders {
		var details []models.OrderPropmanageWork
		database.DB.Where("order_id = ?", o.OrderID).Find(&details)
		orderDetails = append(orderDetails, details...)
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders, "order_details": orderDetails})
}

// GetPropertyOrderPropManageWorkDetail handles GET /dashboard/property/order-prop-manage-work-detail
func GetPropertyOrderPropManageWorkDetail(c *gin.Context) {
	orderID := c.Query("order_id")
	var details []models.OrderPropmanageWork
	database.DB.Where("order_id = ?", orderID).Find(&details)
	c.JSON(http.StatusOK, gin.H{"details": details})
}

// PostPropertyEditOrderPropManageWorkDetail handles POST /dashboard/property/edit-order-prop-manage-work-detail
func PostPropertyEditOrderPropManageWorkDetail(c *gin.Context) {
	var req struct {
		ID   uint                   `json:"id" binding:"required"`
		Data map[string]interface{} `json:"data"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Model(&models.OrderPropmanageWork{}).Where("id = ?", req.ID).Updates(req.Data)
	c.JSON(http.StatusOK, gin.H{"message": "更新しました"})
}

// GetPropertyNotifyInformationRegister handles GET /dashboard/property/:cd/notify-information-register
func GetPropertyNotifyInformationRegister(c *gin.Context) {
	cd := c.Param("cd")
	var prop models.PropBasic
	database.DB.Where("prop_cd = ?", cd).First(&prop)
	var styles []models.NotificationStyles
	database.DB.Find(&styles)
	c.JSON(http.StatusOK, gin.H{"property": prop, "notification_styles": styles})
}

// PostPropertyNotifyInformationRegister handles POST /dashboard/property/:cd/notify-information-register
func PostPropertyNotifyInformationRegister(c *gin.Context) {
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

// GetPropertyNotifyInformation handles GET /dashboard/property/:cd/notify-information
func GetPropertyNotifyInformation(c *gin.Context) {
	cd := c.Param("cd")
	var notifications []models.NotificationList
	database.DB.Where("prop_cd = ? AND delete_flag = 0", cd).Order("regist_datetime DESC").Find(&notifications)
	c.JSON(http.StatusOK, gin.H{"notifications": notifications})
}

// GetPropertyUnitList handles GET /dashboard/property/unit-list
func GetPropertyUnitList(c *gin.Context) {
	var units []map[string]interface{}
	database.DB.Table("unit").Find(&units)
	c.JSON(http.StatusOK, gin.H{"units": units})
}

// GetPropertyCustomer handles GET /dashboard/property/customer
func GetPropertyCustomer(c *gin.Context) {
	customerCd := c.Query("customer_cd")
	var customer models.Customer
	if err := database.DB.Where("customer_cd = ? AND delete_flag = 0", customerCd).First(&customer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "顧客が見つかりません"})
		return
	}
	customer.CustomerPassword = nil
	c.JSON(http.StatusOK, gin.H{"customer": customer})
}

// PostPropertySearchEstimate handles POST /dashboard/property/search-estimate
func PostPropertySearchEstimate(c *gin.Context) {
	var req struct {
		PropCd string `json:"prop_cd"`
	}
	c.ShouldBindJSON(&req)
	var estimates []models.Estimate
	database.DB.Where("prop_cd = ? AND delete_flag = 0 AND activated = 1", req.PropCd).Find(&estimates)
	c.JSON(http.StatusOK, gin.H{"estimates": estimates})
}

// PostPropertyUpdateBikoOrders handles POST /dashboard/property/update-biko-orders
func PostPropertyUpdateBikoOrders(c *gin.Context) {
	var req struct {
		OrderID string  `json:"order_id" binding:"required"`
		Biko    *string `json:"biko"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Model(&models.Order{}).Where("order_id = ?", req.OrderID).Update("biko", req.Biko)
	c.JSON(http.StatusOK, gin.H{"message": "更新しました"})
}

// GetPropertyWorkList handles GET /dashboard/property/:cd/work-list
func GetPropertyWorkList(c *gin.Context) {
	cd := c.Param("cd")
	var schedules []models.WorkSchedule
	database.DB.Where("prop_cd = ? AND delete_flag = 0", cd).Order("work_date ASC").Find(&schedules)
	c.JSON(http.StatusOK, gin.H{"work_schedules": schedules})
}

// PostPropertyChangeMonthWork handles POST /dashboard/property/change-month-work
func PostPropertyChangeMonthWork(c *gin.Context) {
	var req struct {
		OrderID   string `json:"order_id" binding:"required"`
		WorkMonth string `json:"work_month" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Model(&models.Order{}).Where("order_id = ?", req.OrderID).Update("work_month", req.WorkMonth)
	c.JSON(http.StatusOK, gin.H{"message": "更新しました"})
}

// PostPropertyExportInvoice handles POST /dashboard/property/:cd/export-invoice
func PostPropertyExportInvoice(c *gin.Context) {
	cd := c.Param("cd")
	var invoices []models.CusInvoiceDetail
	database.DB.Where("prop_cd = ? AND delete_flag = 0", cd).Find(&invoices)
	c.JSON(http.StatusOK, gin.H{"invoices": invoices})
}

// PostPropertyDeleteFileOtherDetail handles POST /dashboard/property/delete-file-other-detail
func PostPropertyDeleteFileOtherDetail(c *gin.Context) {
	var req struct {
		ID uint `json:"id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Delete(&models.PropImg{}, "id = ?", req.ID)
	c.JSON(http.StatusOK, gin.H{"message": "削除しました"})
}

// PostPropertyAddAdjustmentTaxAmount handles POST /dashboard/property/add-adjustment-tax-amount
func PostPropertyAddAdjustmentTaxAmount(c *gin.Context) {
	var req struct {
		OrderID             string   `json:"order_id" binding:"required"`
		AdjustmentTaxAmount *float64 `json:"adjustment_tax_amount"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Model(&models.Order{}).Where("order_id = ?", req.OrderID).
		Update("adjustment_tax_amount", req.AdjustmentTaxAmount)
	c.JSON(http.StatusOK, gin.H{"message": "更新しました"})
}

// GetPropertyCheckPropName handles GET /dashboard/property/check-prop-name
func GetPropertyCheckPropName(c *gin.Context) {
	propName := c.Query("prop_name")
	propCd := c.Query("prop_cd")
	var count int64
	db := database.DB.Model(&models.PropBasic{}).Where("prop_name = ? AND delete_flag = 0", propName)
	if propCd != "" {
		db = db.Where("prop_cd <> ?", propCd)
	}
	db.Count(&count)
	c.JSON(http.StatusOK, gin.H{"exists": count > 0})
}

// PostPropertyConfidential handles POST /dashboard/property/:cd/confidential
func PostPropertyConfidential(c *gin.Context) {
	cd := c.Param("cd")
	var req struct {
		ConfidentialFlg *int16 `json:"confidential_flg"`
	}
	c.ShouldBindJSON(&req)
	database.DB.Model(&models.PropBasic{}).Where("prop_cd = ?", cd).
		Update("confidential_flg", req.ConfidentialFlg)
	c.JSON(http.StatusOK, gin.H{"message": "更新しました"})
}

// PostPropertyPropCustomer handles POST /dashboard/property/prop-customer
func PostPropertyPropCustomer(c *gin.Context) {
	var req struct {
		PropCd     string `json:"prop_cd" binding:"required"`
		CustomerCd string `json:"customer_cd" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Model(&models.PropBasic{}).Where("prop_cd = ?", req.PropCd).
		Update("customer_cd", req.CustomerCd)
	c.JSON(http.StatusOK, gin.H{"message": "顧客を設定しました"})
}

// PostPropertyInsertWorkSchedule handles POST /dashboard/property/insert-work-schedule
func PostPropertyInsertWorkSchedule(c *gin.Context) {
	var req models.WorkSchedule
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	delFlag := int16(0)
	req.DeleteFlag = &delFlag
	database.DB.Create(&req)
	c.JSON(http.StatusCreated, gin.H{"work_schedule": req})
}
