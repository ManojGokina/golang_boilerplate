package handler

import (
	"net/http"
	"strconv"

	"backend/internal/domain"
	"backend/pkg/logger"
	"backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService domain.UserService
	logger      logger.Logger
}

func NewUserHandler(userService domain.UserService, logger logger.Logger) *UserHandler {
	return &UserHandler{
		userService: userService,
		logger:      logger,
	}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param user body domain.CreateUserRequest true "User creation request"
// @Success 201 {object} response.Response{data=domain.UserResponse}
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req domain.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Invalid request body: ", err)
		response.Error(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	user, err := h.userService.CreateUser(&req)
	if err != nil {
		h.logger.Error("Failed to create user: ", err)
		response.Error(c, http.StatusBadRequest, "Failed to create user", err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "User created successfully", user)
}

// GetUser godoc
// @Summary Get user by ID
// @Description Get user information by user ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} response.Response{data=domain.UserResponse}
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /api/v1/users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "Invalid user ID", "User ID is required")
		return
	}

	user, err := h.userService.GetUser(id)
	if err != nil {
		if err.Error() == "user not found" {
			response.Error(c, http.StatusNotFound, "User not found", err.Error())
			return
		}
		h.logger.Error("Failed to get user: ", err)
		response.Error(c, http.StatusInternalServerError, "Failed to get user", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "User retrieved successfully", user)
}

// UpdateUser godoc
// @Summary Update user
// @Description Update user information
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body domain.UpdateUserRequest true "User update request"
// @Success 200 {object} response.Response{data=domain.UserResponse}
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /api/v1/users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "Invalid user ID", "User ID is required")
		return
	}

	var req domain.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Invalid request body: ", err)
		response.Error(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	user, err := h.userService.UpdateUser(id, &req)
	if err != nil {
		if err.Error() == "user not found" {
			response.Error(c, http.StatusNotFound, "User not found", err.Error())
			return
		}
		h.logger.Error("Failed to update user: ", err)
		response.Error(c, http.StatusInternalServerError, "Failed to update user", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "User updated successfully", user)
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete user by ID
// @Tags users
// @Param id path string true "User ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /api/v1/users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "Invalid user ID", "User ID is required")
		return
	}

	if err := h.userService.DeleteUser(id); err != nil {
		if err.Error() == "user not found" {
			response.Error(c, http.StatusNotFound, "User not found", err.Error())
			return
		}
		h.logger.Error("Failed to delete user: ", err)
		response.Error(c, http.StatusInternalServerError, "Failed to delete user", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "User deleted successfully", nil)
}

// ListUsers godoc
// @Summary List users
// @Description Get paginated list of users
// @Tags users
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} response.PaginatedResponse{data=[]domain.UserResponse}
// @Failure 500 {object} response.Response
// @Router /api/v1/users [get]
func (h *UserHandler) ListUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	users, total, err := h.userService.ListUsers(page, limit)
	if err != nil {
		h.logger.Error("Failed to list users: ", err)
		response.Error(c, http.StatusInternalServerError, "Failed to list users", err.Error())
		return
	}

	response.Paginated(c, http.StatusOK, "Users retrieved successfully", users, page, limit, int(total))
}