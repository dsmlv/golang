package controllers

import (
	"final/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get total sales per product
func GetSalesReport(c *gin.Context) {
	var report []struct {
		ProductName string
		TotalSales  float64
	}

	config.DB.Table("order_items").
		Select("products.name AS product_name, SUM(order_items.price * order_items.quantity) AS total_sales").
		Joins("JOIN products ON order_items.product_id = products.product_id").
		Group("products.name").
		Scan(&report)

	c.JSON(http.StatusOK, report)
}
