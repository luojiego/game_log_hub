package controllers

import (
	"game_log_hub/database"
	"game_log_hub/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateLoginError creates a new login error record
func CreateLoginError(c *gin.Context) {
	var loginError models.LoginError
	if err := c.ShouldBindJSON(&loginError); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loginError.CreatedAt = time.Now()
	result := database.DB.Create(&loginError)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create login error record"})
		return
	}

	c.JSON(http.StatusCreated, loginError)
}

// GetLoginErrors retrieves login error records with pagination
func GetLoginErrors(c *gin.Context) {
	var loginErrors []models.LoginError

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	offset := (page - 1) * pageSize

	// Get total count
	var total int64
	database.DB.Model(&models.LoginError{}).Count(&total)

	// Get records with pagination
	result := database.DB.Order("created_at desc").Offset(offset).Limit(pageSize).Find(&loginErrors)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve login errors"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  loginErrors,
		"total": total,
		"page":  page,
		"size":  pageSize,
	})
}

// GetLoginErrorByID retrieves a single login error record by ID
func GetLoginErrorByID(c *gin.Context) {
	id := c.Param("id")
	var loginError models.LoginError

	result := database.DB.First(&loginError, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Login error record not found"})
		return
	}

	c.JSON(http.StatusOK, loginError)
}

// DeleteLoginError deletes a login error record
func DeleteLoginError(c *gin.Context) {
	id := c.Param("id")
	result := database.DB.Delete(&models.LoginError{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete login error record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login error record deleted successfully"})
}
