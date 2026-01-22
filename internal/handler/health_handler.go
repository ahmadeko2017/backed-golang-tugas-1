package handler

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
)

type HealthHandler struct {
	StartTime time.Time
}

func NewHealthHandler(startTime time.Time) *HealthHandler {
	return &HealthHandler{StartTime: startTime}
}

// HealthCheck godoc
// @Summary Get application health and system metrics
// @Description Get current status, uptime, CPU, and RAM usage
// @Tags health
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func (h *HealthHandler) HealthCheck(c *gin.Context) {
	// Uptime
	uptime := time.Since(h.StartTime).String()

	// Memory
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	// CPU
	// Get CPU usage percentage over a 1 second interval
	// Note: This blocks for 1 second. For non-blocking, use a background goroutine or shorter interval/instant calculation if acceptable.
	// For this simple implementation, we'll try to get it instantly (0 interval) which might return last stored value or need short wait.
	// Common practice for instant response is using 0, but it might be 0 on first call.
	// Let's use 200ms to be responsive enough but get actual data.
	cpuPercent, _ := cpu.Percent(200*time.Millisecond, false)

	// Init default value if retrieval fails
	var cpuUsage float64
	if len(cpuPercent) > 0 {
		cpuUsage = cpuPercent[0]
	}

	numCPU := runtime.NumCPU()

	c.JSON(http.StatusOK, gin.H{
		"status":    "UP",
		"uptime":    uptime,
		"timestamp": time.Now(),
		"system": gin.H{
			"num_cpu":           numCPU,
			"cpu_usage_percent": fmt.Sprintf("%.2f%%", cpuUsage),
			"memory_mb": gin.H{
				"allocated":       bToMb(memStats.Alloc),
				"total_allocated": bToMb(memStats.TotalAlloc),
				"sys_memory":      bToMb(memStats.Sys),
				"num_gc":          memStats.NumGC,
			},
		},
	})
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
