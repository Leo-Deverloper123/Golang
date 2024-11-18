package handlers

import (
	"github.com/gin-gonic/gin"
	"hospital-middleware/internal/services"
	"net/http"
)

type PatientHandler struct {
	patientService *services.PatientService
}

func NewPatientHandler(patientService *services.PatientService) *PatientHandler {
	return &PatientHandler{
		patientService: patientService,
	}
}

func (h *PatientHandler) Search(c *gin.Context) {
	hospitalID := c.GetUint("hospital_id")
	
	var query struct {
		NationalID  string `form:"national_id"`
		PassportID  string `form:"passport_id"`
		FirstName   string `form:"first_name"`
		MiddleName  string `form:"middle_name"`
		LastName    string `form:"last_name"`
		DateOfBirth string `form:"date_of_birth"`
		PhoneNumber string `form:"phone_number"`
		Email       string `form:"email"`
	}

	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	patients, err := h.patientService.Search(hospitalID, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, patients)
}