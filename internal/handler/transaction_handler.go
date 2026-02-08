package handler

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ahmadeko2017/backed-golang-tugas/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TransactionHandler struct {
	service service.TransactionService
}

func NewTransactionHandler(s service.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: s}
}

// formatValidationErrors converts validator errors to user-friendly messages
func formatValidationErrors(err error) []string {
	var errors []string
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			var msg string
			field := e.Field()

			// Make field names more readable
			switch {
			case strings.Contains(field, "ProductID"):
				field = "product_id"
			case strings.Contains(field, "Quantity"):
				field = "quantity"
			case strings.Contains(field, "Total"):
				field = "total"
			}

			switch e.Tag() {
			case "required":
				msg = fmt.Sprintf("Field '%s' wajib diisi", field)
			case "min":
				msg = fmt.Sprintf("Field '%s' minimal %s", field, e.Param())
			case "gt":
				msg = fmt.Sprintf("Field '%s' harus lebih dari %s", field, e.Param())
			default:
				msg = fmt.Sprintf("Field '%s' tidak valid", field)
			}
			errors = append(errors, msg)
		}
	}
	return errors
}

// Checkout godoc
// @Summary Create a transaction (checkout)
// @Description Create a new transaction with multiple items. All items are validated before any write (all-or-nothing).
// @Tags transactions
// @Accept json
// @Produce json
// @Param request body CheckoutRequest true "Checkout Request"
// @Success 201 {object} entity.Transaction
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /api/checkout [post]
func (h *TransactionHandler) Checkout(c *gin.Context) {
	var req CheckoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// Check if it's a validation error
		if validationErrors := formatValidationErrors(err); len(validationErrors) > 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Validasi gagal",
				"details": validationErrors,
			})
			return
		}
		// Other JSON parsing errors
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON tidak valid"})
		return
	}

	items := make([]service.CheckoutItem, 0, len(req.Items))
	for _, it := range req.Items {
		items = append(items, service.CheckoutItem{ProductID: it.ProductID, Quantity: it.Quantity})
	}

	tx, err := h.service.Checkout(items, req.Total)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, tx)
}

// ReportHariIni godoc
// @Summary Today's sales report
// @Description Returns total revenue, total transactions, and top product for today
// @Tags report
// @Produce json
// @Success 200 {object} ReportResponse
// @Failure 500 {object} map[string]string
// @Router /api/report/hari-ini [get]
func (h *TransactionHandler) ReportHariIni(c *gin.Context) {
	totalRevenue, totalTx, name, qty, err := h.service.ReportToday()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var top *TopProduct
	if name != "" && qty > 0 {
		top = &TopProduct{Nama: name, QtyTerjual: qty}
	}
	resp := ReportResponse{TotalRevenue: totalRevenue, TotalTransaksi: totalTx, ProdukTerlaris: top}
	c.JSON(http.StatusOK, resp)
}

// ReportRange godoc
// @Summary Sales report for a date range
// @Description Returns total revenue, total transactions, and top product between two dates
// @Tags report
// @Produce json
// @Param start_date query string true "Start date (YYYY-MM-DD)"
// @Param end_date query string true "End date (YYYY-MM-DD)"
// @Success 200 {object} ReportResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/report [get]
func (h *TransactionHandler) ReportRange(c *gin.Context) {
	startStr := c.Query("start_date")
	endStr := c.Query("end_date")
	if startStr == "" || endStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "start_date and end_date are required"})
		return
	}
	start, err := time.Parse("2006-01-02", startStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start_date format"})
		return
	}
	end, err := time.Parse("2006-01-02", endStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end_date format"})
		return
	}
	// set end to end of day
	end = end.Add(24*time.Hour - time.Nanosecond)

	totalRevenue, totalTx, name, qty, err := h.service.ReportRange(start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var top *TopProduct
	if name != "" && qty > 0 {
		top = &TopProduct{Nama: name, QtyTerjual: qty}
	}
	resp := ReportResponse{TotalRevenue: totalRevenue, TotalTransaksi: totalTx, ProdukTerlaris: top}
	c.JSON(http.StatusOK, resp)
}
