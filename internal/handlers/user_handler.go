package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/klmunoz2/grusolpac_api/internal/models"
	"gorm.io/gorm"
)

// UserHandler holds the database connection
type UserHandler struct {
	DB *gorm.DB
}

// NewUserHandler creates a new UserHandler instance
func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{DB: db}
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Get a list of all registered users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User
// @Failure 500 {object} gin.H{"error": string}
// @Router /users [get]
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	var users []models.User
	// Preload Role to include role information
	if err := h.DB.Preload("Role").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving users: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Get details of a specific user by their ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {object} gin.H{"error": string} "Invalid ID format"
// @Failure 404 {object} gin.H{"error": string} "User not found"
// @Failure 500 {object} gin.H{"error": string}
// @Router /users/{id} [get]
func (h *UserHandler) GetUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	var user models.User
	// Preload Role to include role information
	if err := h.DB.Preload("Role").First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user: " + err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, user)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Add a new user to the database
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.User true "User object to create"
// @Success 201 {object} models.User
// @Failure 400 {object} gin.H{"error": string} "Invalid input data"
// @Failure 500 {object} gin.H{"error": string}
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data: " + err.Error()})
		return
	}

	// TODO: Add validation logic here (e.g., check required fields, email format, password complexity)
	// TODO: Hash the password before saving
	// Example: newUser.Password = hashPasswordFunction(newUser.Password)

	if err := h.DB.Create(&newUser).Error; err != nil {
		// TODO: Handle specific DB errors like unique constraint violations
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user: " + err.Error()})
		return
	}

	// Preload Role for the response
	h.DB.Preload("Role").First(&newUser, newUser.ID)
	c.JSON(http.StatusCreated, newUser)
}

// UpdateUser godoc
// @Summary Update an existing user
// @Description Update user details by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Param user body models.User true "User object with updated fields"
// @Success 200 {object} models.User
// @Failure 400 {object} gin.H{"error": string} "Invalid ID format or input data"
// @Failure 404 {object} gin.H{"error": string} "User not found"
// @Failure 500 {object} gin.H{"error": string}
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	var user models.User
	if err := h.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding user: " + err.Error()})
		}
		return
	}

	var updatedData models.User
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data: " + err.Error()})
		return
	}

	// TODO: Add validation for updatedData
	// TODO: Be careful about updating sensitive fields like Password.
	// You might want a separate endpoint or specific logic for password changes.
	// Avoid overwriting the password unintentionally if it's not provided in the update request.
	// If updatedData.Password is not empty, hash it before saving.

	// Update user fields (excluding ID, CreatedAt, UpdatedAt)
	// Use Model & Updates to only update non-zero fields from updatedData
	// or explicitly map fields you want to allow updating.
	// For simplicity here, we'll use Save which updates all fields.
	// Consider using h.DB.Model(&user).Updates(updatedData) for partial updates.

	// Manually update fields to avoid overwriting with zero values from the request
	// if they weren't provided.
	user.Username = updatedData.Username
	user.Email = updatedData.Email
	user.FirstName = updatedData.FirstName
	user.LastName = updatedData.LastName
	user.IsActive = updatedData.IsActive
	user.RoleID = updatedData.RoleID
	// Handle password update explicitly if needed
	// if updatedData.Password != "" {
	//     user.Password = hashPasswordFunction(updatedData.Password)
	// }

	if err := h.DB.Save(&user).Error; err != nil {
		// TODO: Handle specific DB errors
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user: " + err.Error()})
		return
	}

	// Preload Role for the response
	h.DB.Preload("Role").First(&user, user.ID)
	c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user by their ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 204 "No Content"
// @Failure 400 {object} gin.H{"error": string} "Invalid ID format"
// @Failure 404 {object} gin.H{"error": string} "User not found"
// @Failure 500 {object} gin.H{"error": string}
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	// Use Exec for soft delete (if gorm.DeletedAt is in your model)
	// or Delete for permanent delete. gorm.Model includes gorm.DeletedAt.
	result := h.DB.Delete(&models.User{}, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user: " + result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.Status(http.StatusNoContent)
}
