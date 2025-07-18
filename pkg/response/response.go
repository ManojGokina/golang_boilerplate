package response

import (
	"math"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type PaginatedResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

type Pagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

func Success(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Error(c *gin.Context, statusCode int, message string, error string) {
	c.JSON(statusCode, Response{
		Success: false,
		Message: message,
		Error:   error,
	})
}

func Paginated(c *gin.Context, statusCode int, message string, data interface{}, page, limit, total int) {
	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	c.JSON(statusCode, PaginatedResponse{
		Success: true,
		Message: message,
		Data:    data,
		Pagination: Pagination{
			Page:       page,
			Limit:      limit,
			Total:      total,
			TotalPages: totalPages,
		},
	})
}