package handlers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"hospital-middleware/internal/models"
	"hospital-middleware/internal/services"
	"net/http"
)

type StaffHandler struct {
	staffService *services.StaffService
	authService  *services.AuthService
}

func NewStaffHandler(staffService *services.StaffService, authService *services.AuthService) *StaffHandler {
	return &StaffHandler{
		staffService: staffService,
		authService: authService,
	}
}

func (h *StaffHandler) Create(c *gin.Context) {
	var input struct {
		Username   string `json:"username" binding:"required"`
		Password   string `json:"password" binding:"required"`
		HospitalID uint   `json:"hospital_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	staff := &models.Staff{
		Username:   input.Username,
		Password:   string(hashedPassword),
		HospitalID: input.HospitalID,
	}

	if err := h.staffService.Create(staff); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Staff created successfully"})
}

func (h *StaffHandler) Login(c *gin.Context) {
	var input struct {
		Username   string `json:"username" binding:"required"`
		Password   string `json:"password" binding:"required"`
		HospitalID uint   `json:"hospital_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	staff, err := h.staffService.GetByUsername(input.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if staff.HospitalID != input.HospitalID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid hospital"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(staff.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := h.authService.GenerateToken(staff.ID, staff.HospitalID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}