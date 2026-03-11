package handlers

import (
	"kms_golang/database"
	"kms_golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCancellationList handles GET /dashboard/cancellation/list
// Returns list of work cancellations (orders with work_cancel_flg = 1)
func GetCancellationList(c *gin.Context) {
	type CancellationRow struct {
		PropCd       string  `json:"prop_cd"`
		CustomerCd   string  `json:"customer_cd"`
		CustomerName *string `json:"customer_name"`
		PropName     *string `json:"prop_name"`
	}

	db := database.DB.Table("order_prop_manage_work").
		Select("orders.prop_cd, orders.customer_cd").
		Joins("JOIN orders ON orders.order_id = order_prop_manage_work.order_id").
		Joins("LEFT JOIN prop_basics ON prop_basics.prop_cd = orders.prop_cd").
		Joins("LEFT JOIN customers ON customers.customer_cd = orders.customer_cd").
		Where("order_prop_manage_work.delete_flag = 0").
		Where("order_prop_manage_work.work_cancel_flg = 1")

	if customerCd := c.Query("customer_cd"); customerCd != "" {
		db = db.Where("orders.customer_cd = ?", customerCd)
	}
	if propCd := c.Query("prop_cd"); propCd != "" {
		db = db.Where("orders.prop_cd = ?", propCd)
	}
	if cusKana := c.Query("cus_kana"); cusKana != "" {
		db = db.Where("customers.customer_kana LIKE ?", "%"+cusKana+"%")
	}
	if propKana := c.Query("prop_kana"); propKana != "" {
		db = db.Where("prop_basics.prop_kana_name LIKE ?", "%"+propKana+"%")
	}

	db = db.Group("orders.customer_cd, orders.prop_cd")

	var rows []struct {
		PropCd     string `json:"prop_cd"`
		CustomerCd string `json:"customer_cd"`
	}
	if err := db.Find(&rows).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"cancellations": rows})
}

// GetCancellationDetail handles GET /dashboard/cancellation/detail?prop_cd=xxx&customer_cd=xxx
func GetCancellationDetail(c *gin.Context) {
	propCd := c.Query("prop_cd")
	customerCd := c.Query("customer_cd")
	if propCd == "" || customerCd == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "prop_cd and customer_cd are required"})
		return
	}

	var prop models.PropBasic
	database.DB.Where("prop_cd = ? AND delete_flag = 0", propCd).First(&prop)

	var customer models.Customer
	database.DB.Where("customer_cd = ? AND delete_flag = 0", customerCd).First(&customer)

	var cancelledWorks []models.OrderPropmanageWork
	database.DB.Table("order_prop_manage_work").
		Joins("JOIN orders ON orders.order_id = order_prop_manage_work.order_id").
		Where("order_prop_manage_work.delete_flag = 0").
		Where("order_prop_manage_work.work_cancel_flg = 1").
		Where("orders.prop_cd = ?", propCd).
		Where("orders.customer_cd = ?", customerCd).
		Find(&cancelledWorks)

	var cancels []models.WorkCancel
	database.DB.Table("work_cancel").
		Joins("JOIN orders ON orders.order_id = work_cancel.order_id").
		Where("orders.prop_cd = ?", propCd).
		Where("orders.customer_cd = ?", customerCd).
		Find(&cancels)

	c.JSON(http.StatusOK, gin.H{
		"property":        prop,
		"customer":        customer,
		"cancelled_works": cancelledWorks,
		"work_cancels":    cancels,
	})
}
