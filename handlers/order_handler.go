package handlers

import (
	"kms_golang/database"
	"kms_golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetOrderList handles GET /dashboard/order/list
func GetOrderList(c *gin.Context) {
	var orders []models.Order
	db := database.DB

	if propCd := c.Query("prop_cd"); propCd != "" {
		db = db.Where("prop_cd = ?", propCd)
	}
	if estimateNumber := c.Query("estimate_number"); estimateNumber != "" {
		db = db.Where("estimate_number = ?", estimateNumber)
	}
	if orderStatus := c.Query("order_status"); orderStatus != "" {
		db = db.Where("order_status = ?", orderStatus)
	}

	if err := db.Order("order_id DESC").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

// GetOrderDetail handles GET /dashboard/order/:id
func GetOrderDetail(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := database.DB.Where("order_id = ?", id).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "発注が見つかりません"})
		return
	}

	var propWork []models.OrderPropmanageWork
	database.DB.Where("order_id = ?", id).Find(&propWork)

	var constrWork []models.OrderConstructionWork
	database.DB.Where("order_id = ?", id).Find(&constrWork)

	c.JSON(http.StatusOK, gin.H{
		"order":              order,
		"prop_manage_works":  propWork,
		"construction_works": constrWork,
	})
}

// PostOrderCreate handles POST /dashboard/order/create
func PostOrderCreate(c *gin.Context) {
	var req models.Order
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	if err := database.DB.Create(&req).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"order": req})
}

// PutOrderUpdate handles PUT /dashboard/order/:id
func PutOrderUpdate(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := database.DB.Where("order_id = ?", id).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "発注が見つかりません"})
		return
	}

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	order.OrderID = id
	if err := database.DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

// DeleteOrder handles DELETE /dashboard/order/:id (soft delete)
func DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Order{}, "order_id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "発注を削除しました"})
}

// PutOrderStatus handles PUT /dashboard/order/:id/status
func PutOrderStatus(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Status int16 `json:"order_status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	if err := database.DB.Model(&models.Order{}).Where("order_id = ?", id).Update("order_status", req.Status).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ステータスを更新しました"})
}

// GetOrderType1 handles GET /dashboard/order/:id/type1
func GetOrderType1(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := database.DB.Where("order_id = ?", id).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "発注が見つかりません"})
		return
	}
	var propWork []models.OrderPropmanageWork
	database.DB.Where("order_id = ?", id).Find(&propWork)

	var propBasic models.PropBasic
	if order.PropCd != nil {
		database.DB.Where("prop_cd = ?", *order.PropCd).First(&propBasic)
	}

	c.JSON(http.StatusOK, gin.H{"order": order, "prop_manage_works": propWork, "prop_basic": propBasic})
}

// GetOrderType2 handles GET /dashboard/order/:id/type2
func GetOrderType2(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := database.DB.Where("order_id = ?", id).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "発注が見つかりません"})
		return
	}
	var constrWork []models.OrderConstructionWork
	database.DB.Where("order_id = ?", id).Find(&constrWork)

	var constrDetails []models.OrderConstructionWorkDetail
	for _, cw := range constrWork {
		var details []models.OrderConstructionWorkDetail
		database.DB.Where("order_construction_work_id = ?", cw.ID).Find(&details)
		constrDetails = append(constrDetails, details...)
	}

	var propBasic models.PropBasic
	if order.PropCd != nil {
		database.DB.Where("prop_cd = ?", *order.PropCd).First(&propBasic)
	}

	c.JSON(http.StatusOK, gin.H{
		"order":                     order,
		"construction_works":        constrWork,
		"construction_work_details": constrDetails,
		"prop_basic":                propBasic,
	})
}

// GetOrderSekosakiListType1 handles GET /dashboard/order/:id/sekosaki-list-type1
func GetOrderSekosakiListType1(c *gin.Context) {
	id := c.Param("id")
	var propWorks []models.OrderPropmanageWork
	database.DB.Where("order_id = ?", id).Find(&propWorks)

	sekosakiCds := map[string]bool{}
	for _, pw := range propWorks {
		if pw.SekosakiCd != nil {
			sekosakiCds[*pw.SekosakiCd] = true
		}
	}

	var sekosakis []models.Sekosaki
	if len(sekosakiCds) > 0 {
		cds := make([]string, 0, len(sekosakiCds))
		for cd := range sekosakiCds {
			cds = append(cds, cd)
		}
		database.DB.Where("sekosaki_cd IN ? AND delete_flag = 0", cds).Find(&sekosakis)
		for i := range sekosakis {
			sekosakis[i].SekosakiPassword = nil
		}
	}
	c.JSON(http.StatusOK, gin.H{"sekosakis": sekosakis, "prop_manage_works": propWorks})
}

// GetOrderSekosakiListType2 handles GET /dashboard/order/:id/sekosaki-list-type2
func GetOrderSekosakiListType2(c *gin.Context) {
	id := c.Param("id")
	var constrWorks []models.OrderConstructionWork
	database.DB.Where("order_id = ?", id).Find(&constrWorks)

	sekosakiCds := map[string]bool{}
	for _, cw := range constrWorks {
		if cw.SekosakiCd != nil {
			sekosakiCds[*cw.SekosakiCd] = true
		}
	}

	var sekosakis []models.Sekosaki
	if len(sekosakiCds) > 0 {
		cds := make([]string, 0, len(sekosakiCds))
		for cd := range sekosakiCds {
			cds = append(cds, cd)
		}
		database.DB.Where("sekosaki_cd IN ? AND delete_flag = 0", cds).Find(&sekosakis)
		for i := range sekosakis {
			sekosakis[i].SekosakiPassword = nil
		}
	}
	c.JSON(http.StatusOK, gin.H{"sekosakis": sekosakis, "construction_works": constrWorks})
}

// GetOrderFinancialType1 handles GET /dashboard/order/:id/financial-type1
func GetOrderFinancialType1(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	database.DB.Where("order_id = ?", id).First(&order)

	var propWorks []models.OrderPropmanageWork
	database.DB.Where("order_id = ?", id).Find(&propWorks)

	c.JSON(http.StatusOK, gin.H{"order": order, "prop_manage_works": propWorks})
}

// GetOrderFinancialType2 handles GET /dashboard/order/:id/financial-type2
func GetOrderFinancialType2(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	database.DB.Where("order_id = ?", id).First(&order)

	var constrWorks []models.OrderConstructionWork
	database.DB.Where("order_id = ?", id).Find(&constrWorks)

	c.JSON(http.StatusOK, gin.H{"order": order, "construction_works": constrWorks})
}

// GetOrderWorkListType1 handles GET /dashboard/order/:id/work-list-type1
func GetOrderWorkListType1(c *gin.Context) {
	id := c.Param("id")
	var propWorks []models.OrderPropmanageWork
	database.DB.Where("order_id = ?", id).Find(&propWorks)

	var workSchedules []models.WorkSchedule
	database.DB.Where("order_id = ? AND delete_flag = 0", id).Find(&workSchedules)

	c.JSON(http.StatusOK, gin.H{"prop_manage_works": propWorks, "work_schedules": workSchedules})
}

// GetOrderWorkDetailType1 handles GET /dashboard/order/:id/work-detail-type1
func GetOrderWorkDetailType1(c *gin.Context) {
	id := c.Param("id")
	pmwID := c.Query("pmw_id")

	var detail models.OrderPropmanageWork
	if pmwID != "" {
		database.DB.Where("id = ? AND order_id = ?", pmwID, id).First(&detail)
	}

	c.JSON(http.StatusOK, gin.H{"detail": detail})
}

// PostOrderSelectedSekosakiType1 handles POST /dashboard/order/:id/selected-sekosaki-type1
func PostOrderSelectedSekosakiType1(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		SekosakiCd string `json:"sekosaki_cd" binding:"required"`
		PmwIDs     []uint `json:"pmw_ids"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	if len(req.PmwIDs) > 0 {
		database.DB.Model(&models.OrderPropmanageWork{}).Where("id IN ? AND order_id = ?", req.PmwIDs, id).
			Update("sekosaki_cd", req.SekosakiCd)
	}
	c.JSON(http.StatusOK, gin.H{"message": "施工先を選択しました"})
}

// PostOrderChangeContractTermStart handles POST /dashboard/order/:id/change-contract-term-start
func PostOrderChangeContractTermStart(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		ContractTermStart *string `json:"contract_term_start"`
	}
	c.ShouldBindJSON(&req)
	database.DB.Model(&models.OrderPropmanageWork{}).Where("order_id = ?", id).
		Update("contract_term_start", req.ContractTermStart)
	c.JSON(http.StatusOK, gin.H{"message": "更新しました"})
}

// PostOrderChangeWorkCancelDate handles POST /dashboard/order/:id/change-work-cancel-date
func PostOrderChangeWorkCancelDate(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		WorkCancelDate *string `json:"work_cancel_date"`
		PmwID          uint    `json:"pmw_id"`
	}
	c.ShouldBindJSON(&req)
	if req.PmwID > 0 {
		database.DB.Model(&models.OrderPropmanageWork{}).Where("id = ? AND order_id = ?", req.PmwID, id).
			Update("work_cancel_flg", req.WorkCancelDate)
	}
	c.JSON(http.StatusOK, gin.H{"message": "更新しました"})
}

// PostOrderChangeBillingCustomer handles POST /dashboard/order/:id/change-billing-customer
func PostOrderChangeBillingCustomer(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		CustomerCd string `json:"customer_cd" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Model(&models.Order{}).Where("order_id = ?", id).Update("customer_cd", req.CustomerCd)
	c.JSON(http.StatusOK, gin.H{"message": "請求先を変更しました"})
}

// PostOrderChangeSekosaki handles POST /dashboard/order/:id/change-sekosaki
func PostOrderChangeSekosaki(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		SekosakiCd string `json:"sekosaki_cd" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Model(&models.Order{}).Where("order_id = ?", id).Update("sekosaki_cd", req.SekosakiCd)
	c.JSON(http.StatusOK, gin.H{"message": "施工先を変更しました"})
}

// PostOrderUpdateBikoOrders handles POST /dashboard/order/:id/update-biko
func PostOrderUpdateBikoOrders(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Biko *string `json:"biko"`
	}
	c.ShouldBindJSON(&req)
	database.DB.Model(&models.Order{}).Where("order_id = ?", id).Update("biko", req.Biko)
	c.JSON(http.StatusOK, gin.H{"message": "備考を更新しました"})
}

// PostOrderDeleteType1OfSekosaki handles POST /dashboard/order/:id/delete-type1-sekosaki
func PostOrderDeleteType1OfSekosaki(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		SekosakiCd string `json:"sekosaki_cd" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Model(&models.OrderPropmanageWork{}).Where("order_id = ? AND sekosaki_cd = ?", id, req.SekosakiCd).
		Update("sekosaki_cd", nil)
	c.JSON(http.StatusOK, gin.H{"message": "施工先を削除しました"})
}

// PostOrderDeleteType2OfSekosaki handles POST /dashboard/order/:id/delete-type2-sekosaki
func PostOrderDeleteType2OfSekosaki(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		SekosakiCd string `json:"sekosaki_cd" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Model(&models.OrderConstructionWork{}).Where("order_id = ? AND sekosaki_cd = ?", id, req.SekosakiCd).
		Update("sekosaki_cd", nil)
	c.JSON(http.StatusOK, gin.H{"message": "施工先を削除しました"})
}

// PostOrderSendToSekosakiType1 handles POST /dashboard/order/:id/send-sekosaki-type1
func PostOrderSendToSekosakiType1(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		SekosakiCd string `json:"sekosaki_cd" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Model(&models.Order{}).Where("order_id = ?", id).
		Updates(map[string]interface{}{"sekosaki_cd": req.SekosakiCd, "seko_send_flg": 1})
	c.JSON(http.StatusOK, gin.H{"message": "施工先に送信しました"})
}

// PostOrderSendToSekosakiType2 handles POST /dashboard/order/:id/send-sekosaki-type2
func PostOrderSendToSekosakiType2(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		SekosakiCd string `json:"sekosaki_cd" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Model(&models.Order{}).Where("order_id = ?", id).
		Updates(map[string]interface{}{"sekosaki_cd": req.SekosakiCd, "seko_send_flg": 1})
	c.JSON(http.StatusOK, gin.H{"message": "施工先に送信しました"})
}

// PostOrderChangeAdjustmentTaxAmount handles POST /dashboard/order/:id/change-adjustment-tax
func PostOrderChangeAdjustmentTaxAmount(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		AdjustmentTaxAmount *float64 `json:"adjustment_tax_amount"`
	}
	c.ShouldBindJSON(&req)
	database.DB.Model(&models.Order{}).Where("order_id = ?", id).
		Update("adjustment_tax_amount", req.AdjustmentTaxAmount)
	c.JSON(http.StatusOK, gin.H{"message": "更新しました"})
}

// PostOrderLoadDataConstructionWorkDetail handles POST /dashboard/order/:id/load-construction-work-detail
func PostOrderLoadDataConstructionWorkDetail(c *gin.Context) {
	var req struct {
		OrderConstructionWorkID uint `json:"order_construction_work_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	var details []models.OrderConstructionWorkDetail
	database.DB.Where("order_construction_work_id = ?", req.OrderConstructionWorkID).Find(&details)
	c.JSON(http.StatusOK, gin.H{"details": details})
}

// PostOrderUpdateSekosakiConstructionWorkDetail handles POST /dashboard/order/update-sekosaki-construction-work-detail
func PostOrderUpdateSekosakiConstructionWorkDetail(c *gin.Context) {
	var req struct {
		ID         uint   `json:"id" binding:"required"`
		SekosakiCd string `json:"sekosaki_cd" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Model(&models.OrderConstructionWork{}).Where("id = ?", req.ID).
		Update("sekosaki_cd", req.SekosakiCd)
	c.JSON(http.StatusOK, gin.H{"message": "更新しました"})
}

// PostOrderUpdateDateConstructionDetail handles POST /dashboard/order/update-date-construction-detail
func PostOrderUpdateDateConstructionDetail(c *gin.Context) {
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

// PostOrderUpdateConstructionDetail handles POST /dashboard/order/update-construction-detail
func PostOrderUpdateConstructionDetail(c *gin.Context) {
	var req struct {
		ID   uint                   `json:"id" binding:"required"`
		Data map[string]interface{} `json:"data"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Model(&models.OrderConstructionWork{}).Where("id = ?", req.ID).Updates(req.Data)
	c.JSON(http.StatusOK, gin.H{"message": "更新しました"})
}

// PostOrderRegisterActionPlan handles POST /dashboard/order/:id/register-action-plan
func PostOrderRegisterActionPlan(c *gin.Context) {
	var req models.ActionPlan
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Create(&req)
	c.JSON(http.StatusCreated, gin.H{"action_plan": req})
}

// PostOrderReportWork handles POST /dashboard/order/:id/report-work
func PostOrderReportWork(c *gin.Context) {
	var req models.PropWorkReport
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}
	database.DB.Create(&req)
	c.JSON(http.StatusCreated, gin.H{"work_report": req})
}

// GetOrderCustomerNameReport handles GET /dashboard/order/customer-name-report
func GetOrderCustomerNameReport(c *gin.Context) {
	customerCd := c.Query("customer_cd")
	var customer models.Customer
	database.DB.Where("customer_cd = ?", customerCd).Select("customer_cd, customer_name").First(&customer)
	customer.CustomerPassword = nil
	c.JSON(http.StatusOK, gin.H{"customer": customer})
}

// PostOrderPurchaseOrder handles POST /dashboard/order/:id/purchase-order
func PostOrderPurchaseOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := database.DB.Where("order_id = ?", id).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "発注が見つかりません"})
		return
	}
	// Return data for purchase order PDF
	var propBasic models.PropBasic
	if order.PropCd != nil {
		database.DB.Where("prop_cd = ?", *order.PropCd).First(&propBasic)
	}
	var customer models.Customer
	if order.CustomerCd != nil {
		database.DB.Where("customer_cd = ?", *order.CustomerCd).First(&customer)
		customer.CustomerPassword = nil
	}
	var sekosaki models.Sekosaki
	// SekosakiCd is on OrderConstructionWork, not Order directly
	var ocw models.OrderConstructionWork
	if err := database.DB.Where("order_id = ?", order.OrderID).First(&ocw).Error; err == nil && ocw.SekosakiCd != nil {
		database.DB.Where("sekosaki_cd = ?", *ocw.SekosakiCd).First(&sekosaki)
		sekosaki.SekosakiPassword = nil
	}
	c.JSON(http.StatusOK, gin.H{
		"order":      order,
		"prop_basic": propBasic,
		"customer":   customer,
		"sekosaki":   sekosaki,
	})
}

// PostOrderSearchWorkDetailType1 handles POST /dashboard/order/search-work-detail-type1
func PostOrderSearchWorkDetailType1(c *gin.Context) {
	var req struct {
		OrderID string `json:"order_id"`
		Keyword string `json:"keyword"`
	}
	c.ShouldBindJSON(&req)
	var details []models.OrderPropmanageWork
	db := database.DB.Table("order_propmanage_work")
	if req.OrderID != "" {
		db = db.Where("order_id = ?", req.OrderID)
	}
	if req.Keyword != "" {
		db = db.Where("work_content LIKE ?", "%"+req.Keyword+"%")
	}
	db.Find(&details)
	c.JSON(http.StatusOK, gin.H{"details": details})
}
