package exports

import (
	"fmt"
	"kms_golang/database"
	"kms_golang/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func safeStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func safeFloat(f *float64) float64 {
	if f == nil {
		return 0
	}
	return *f
}

// ExportCustomers handles GET /dashboard/export/customers
func ExportCustomers(c *gin.Context) {
	var customers []models.Customer
	db := database.DB.Where("delete_flag = 0")

	if name := c.Query("customer_name"); name != "" {
		db = db.Where("customer_name LIKE ?", "%"+name+"%")
	}

	if err := db.Order("customer_cd ASC").Find(&customers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	f := excelize.NewFile()
	defer f.Close()

	sheet := "顧客一覧"
	f.SetSheetName("Sheet1", sheet)

	// Header row
	headers := []string{"顧客CD", "顧客名", "顧客名（カナ）", "郵便番号", "住所", "電話番号", "FAX", "メールアドレス"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}

	// Style headers
	style, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{Bold: true},
		Fill: excelize.Fill{Type: "pattern", Color: []string{"#D9E1F2"}, Pattern: 1},
	})
	f.SetCellStyle(sheet, "A1", string(rune('A'+len(headers)-1))+"1", style)

	// Data rows
	for i, customer := range customers {
		row := i + 2
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), customer.CustomerCd)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), safeStr(customer.CustomerName))
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), safeStr(customer.CustomerKana))
		f.SetCellValue(sheet, fmt.Sprintf("D%d", row), safeStr(customer.PostCode))
		f.SetCellValue(sheet, fmt.Sprintf("E%d", row), safeStr(customer.BlockName))
		f.SetCellValue(sheet, fmt.Sprintf("F%d", row), safeStr(customer.Tel))
		f.SetCellValue(sheet, fmt.Sprintf("G%d", row), safeStr(customer.Fax))
		f.SetCellValue(sheet, fmt.Sprintf("H%d", row), "")
	}

	// Set column widths
	f.SetColWidth(sheet, "A", "A", 12)
	f.SetColWidth(sheet, "B", "C", 20)
	f.SetColWidth(sheet, "D", "D", 12)
	f.SetColWidth(sheet, "E", "E", 30)
	f.SetColWidth(sheet, "F", "G", 15)
	f.SetColWidth(sheet, "H", "H", 25)

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=customers.xlsx")

	if err := f.Write(c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
}

// ExportProperties handles GET /dashboard/export/properties
func ExportProperties(c *gin.Context) {
	var properties []models.PropBasic
	db := database.DB.Where("delete_flag = 0")

	if propCd := c.Query("prop_cd"); propCd != "" {
		db = db.Where("prop_cd LIKE ?", "%"+propCd+"%")
	}
	if customerCd := c.Query("customer_cd"); customerCd != "" {
		db = db.Where("customer_cd = ?", customerCd)
	}

	if err := db.Order("prop_cd ASC").Find(&properties).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	f := excelize.NewFile()
	defer f.Close()

	sheet := "物件一覧"
	f.SetSheetName("Sheet1", sheet)

	headers := []string{"物件CD", "物件名", "郵便番号", "住所", "顧客CD"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}

	for i, prop := range properties {
		row := i + 2
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), prop.PropCd)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), safeStr(prop.PropName))
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), safeStr(prop.PostCode))
		f.SetCellValue(sheet, fmt.Sprintf("D%d", row), safeStr(prop.BlockName))
		f.SetCellValue(sheet, fmt.Sprintf("E%d", row), safeStr(prop.CustomerCd))
	}

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=properties.xlsx")

	if err := f.Write(c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
}

// ExportInvoices handles GET /dashboard/export/invoices
func ExportInvoices(c *gin.Context) {
	var invoices []models.CusInvoiceDetail
	db := database.DB.Where("delete_flag = 0")

	if customerCd := c.Query("customer_cd"); customerCd != "" {
		db = db.Where("customer_cd = ?", customerCd)
	}
	if yearMonth := c.Query("year_month"); yearMonth != "" {
		db = db.Where("invoice_month = ?", yearMonth)
	}

	if err := db.Order("invoice_number DESC").Find(&invoices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	f := excelize.NewFile()
	defer f.Close()

	sheet := "請求書一覧"
	f.SetSheetName("Sheet1", sheet)

	headers := []string{"請求書番号", "顧客CD", "物件CD", "年月", "請求金額", "消費税", "合計"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}

	for i, inv := range invoices {
		row := i + 2
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), inv.ID)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), inv.CustomerCd)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), "")
		f.SetCellValue(sheet, fmt.Sprintf("D%d", row), inv.InvoiceMonth)
		f.SetCellValue(sheet, fmt.Sprintf("E%d", row), safeFloat(inv.TotalAmount))
		f.SetCellValue(sheet, fmt.Sprintf("F%d", row), safeFloat(inv.TaxAmount))
		f.SetCellValue(sheet, fmt.Sprintf("G%d", row), safeFloat(inv.TotalAmount))
	}

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", "attachment; filename=invoices.xlsx")

	if err := f.Write(c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
}
