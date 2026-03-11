package handlers

import (
	"kms_golang/database"
	"kms_golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetOrderExportPDFType1 handles GET /dashboard/order-export/pdf-type1/:order_id
func GetOrderExportPDFType1(c *gin.Context) {
	orderID := c.Param("order_id")

	var order models.Order
	if err := database.DB.Where("order_id = ?", orderID).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "注文が見つかりません"})
		return
	}

	// Load related data for PDF generation
	var orderDetails []models.OrderPropmanageWork
	database.DB.Where("order_id = ?", orderID).Find(&orderDetails)

	var propBasic models.PropBasic
	if order.PropCd != nil {
		database.DB.Where("prop_cd = ?", *order.PropCd).First(&propBasic)
	}

	// Return data for client-side PDF generation
	c.JSON(http.StatusOK, gin.H{
		"order":         order,
		"order_details": orderDetails,
		"prop_basic":    propBasic,
	})
}

// GetOrderExportPDFType2 handles GET /dashboard/order-export/pdf-type2/:order_id
func GetOrderExportPDFType2(c *gin.Context) {
	orderID := c.Param("order_id")

	var order models.Order
	if err := database.DB.Where("order_id = ?", orderID).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "注文が見つかりません"})
		return
	}

	var orderConstructions []models.OrderConstructionWork
	database.DB.Where("order_id = ?", orderID).Find(&orderConstructions)

	var orderConstructionDetails []models.OrderConstructionWorkDetail
	for _, oc := range orderConstructions {
		var details []models.OrderConstructionWorkDetail
		database.DB.Where("order_construction_work_id = ?", oc.ID).Find(&details)
		orderConstructionDetails = append(orderConstructionDetails, details...)
	}

	var propBasic models.PropBasic
	if order.PropCd != nil {
		database.DB.Where("prop_cd = ?", *order.PropCd).First(&propBasic)
	}

	c.JSON(http.StatusOK, gin.H{
		"order":                      order,
		"order_constructions":        orderConstructions,
		"order_construction_details": orderConstructionDetails,
		"prop_basic":                 propBasic,
	})
}

// ---- PDF Controller (data endpoints for client-side PDF rendering) ----

// GetPdfExportData handles GET /dashboard/pdf/export-data
func GetPdfExportData(c *gin.Context) {
	pdfType := c.Query("type")
	id := c.Query("id")

	switch pdfType {
	case "property":
		var prop models.PropBasic
		if err := database.DB.Where("prop_cd = ? AND delete_flag = 0", id).First(&prop).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "物件が見つかりません"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": prop})

	case "task_request":
		var tasks []models.Task
		database.DB.Where("prop_cd = ? AND delete_flag <> 1", id).Find(&tasks)
		c.JSON(http.StatusOK, gin.H{"data": tasks})

	case "estimate":
		var estimate models.Estimate
		database.DB.Where("estimate_number = ? AND delete_flag = 0 AND activated = 1", id).First(&estimate)
		var estConstructions []models.EstConstruction
		database.DB.Where("estimate_number = ? AND subnumber = ?", estimate.EstimateNumber, estimate.Subnumber).Find(&estConstructions)
		c.JSON(http.StatusOK, gin.H{"estimate": estimate, "est_constructions": estConstructions})

	case "monthly_report":
		var report models.MonthlyReportNotes
		database.DB.Preload("Imgs").Where("id = ?", id).First(&report)
		c.JSON(http.StatusOK, gin.H{"data": report})

	case "notification":
		var notif models.NotificationList
		database.DB.Where("id = ?", id).First(&notif)
		c.JSON(http.StatusOK, gin.H{"data": notif})

	default:
		c.JSON(http.StatusBadRequest, gin.H{"message": "不明なPDFタイプ"})
	}
}

// GetPdfFormData handles GET /dashboard/pdf/form-data/:form_type
func GetPdfFormData(c *gin.Context) {
	formType := c.Param("form_type")
	orderID := c.Query("order_id")
	propCd := c.Query("prop_cd")

	switch formType {
	case "order":
		var order models.Order
		database.DB.Where("order_id = ?", orderID).First(&order)
		var propBasic models.PropBasic
		if propCd != "" {
			database.DB.Where("prop_cd = ?", propCd).First(&propBasic)
		}
		c.JSON(http.StatusOK, gin.H{"order": order, "prop_basic": propBasic})

	case "request":
		var reception models.Reception
		database.DB.Where("prop_cd = ? AND delete_flag = 0", propCd).Order("accept_number DESC").First(&reception)
		c.JSON(http.StatusOK, gin.H{"reception": reception})

	case "report":
		var reports []models.MonthlyReportNotes
		database.DB.Where("prop_cd = ?", propCd).Order("report_month DESC").Find(&reports)
		c.JSON(http.StatusOK, gin.H{"reports": reports})

	case "invoice":
		var invoiceDetails []models.CusInvoiceDetail
		database.DB.Where("customer_cd = ? AND delete_flag = 0", c.Query("customer_cd")).Find(&invoiceDetails)
		c.JSON(http.StatusOK, gin.H{"invoice_details": invoiceDetails})

	default:
		c.JSON(http.StatusBadRequest, gin.H{"message": "不明なフォームタイプ"})
	}
}

// ---- TempController (order form templates) ----

// GetTempOrderForm handles GET /dashboard/temp/order-form/:order_id
func GetTempOrderForm(c *gin.Context) {
	orderID := c.Param("order_id")

	var order models.Order
	if err := database.DB.Where("order_id = ?", orderID).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "注文が見つかりません"})
		return
	}

	var propBasic models.PropBasic
	if order.PropCd != nil {
		database.DB.Where("prop_cd = ?", *order.PropCd).First(&propBasic)
	}

	var customer models.Customer
	if order.CustomerCd != nil {
		database.DB.Where("customer_cd = ?", *order.CustomerCd).First(&customer)
	}

	var sekosaki models.Sekosaki
	var ocw models.OrderConstructionWork
	if err := database.DB.Where("order_id = ?", orderID).First(&ocw).Error; err == nil && ocw.SekosakiCd != nil {
		database.DB.Where("sekosaki_cd = ?", *ocw.SekosakiCd).First(&sekosaki)
	}

	// Load order details (type 1 or type 2)
	var orderPropmanageWorks []models.OrderPropmanageWork
	database.DB.Where("order_id = ?", orderID).Find(&orderPropmanageWorks)

	var orderConstructionWorks []models.OrderConstructionWork
	database.DB.Where("order_id = ?", orderID).Find(&orderConstructionWorks)

	var orderConstructionWorkDetails []models.OrderConstructionWorkDetail
	for _, ow := range orderConstructionWorks {
		var details []models.OrderConstructionWorkDetail
		database.DB.Where("order_construction_work_id = ?", ow.ID).Find(&details)
		orderConstructionWorkDetails = append(orderConstructionWorkDetails, details...)
	}

	c.JSON(http.StatusOK, gin.H{
		"order":                           order,
		"prop_basic":                      propBasic,
		"customer":                        customer,
		"sekosaki":                        sekosaki,
		"order_propmanage_works":          orderPropmanageWorks,
		"order_construction_works":        orderConstructionWorks,
		"order_construction_work_details": orderConstructionWorkDetails,
	})
}

// PostTempEditListOrder handles POST /dashboard/temp/edit-list-order
func PostTempEditListOrder(c *gin.Context) {
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
