package services

import (
	"encoding/json"
	"fmt"
	"hospital-middleware/internal/models"
	"net/http"
)

type PatientService struct {
	db *gorm.DB
}

func NewPatientService(db *gorm.DB) *PatientService {
	return &PatientService{db: db}
}

func (s *PatientService) Search(hospitalID uint, query interface{}) ([]models.Patient, error) {
	var hospital models.Hospital
	if err := s.db.First(&hospital, hospitalID).Error; err != nil {
		return nil, err
	}

	// First search in local database
	var patients []models.Patient
	db := s.db.Where("hospital_id = ?", hospitalID)

	if q, ok := query.(struct{
		NationalID  string
		PassportID  string
		FirstName   string
		MiddleName  string
		LastName    string
		DateOfBirth string
		PhoneNumber string
		Email       string
	}); ok {
		if q.NationalID != "" {
			db = db.Where("national_id = ?", q.NationalID)
		}
		if q.PassportID != "" {
			db = db.Where("passport_id = ?", q.PassportID)
		}
		// Add other conditions...
	}

	if err := db.Find(&patients).Error; err != nil {
		return nil, err
	}

	// If no results found, search in external HIS
	if len(patients) == 0 {
		externalPatients, err := s.searchExternalHIS(hospital, query)
		if err != nil {
			return nil, err
		}
		return externalPatients, nil
	}

	return patients, nil
}

func (s *PatientService) searchExternalHIS(hospital models.Hospital, query interface{}) ([]models.Patient, error) {
	// Implementation for external HIS API call
	client := &http.Client{}
	req, err := http.NewRequest("GET", hospital.APIUrl+"/patient/search", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+hospital.APIToken)
	
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var patients []models.Patient
	if err := json.NewDecoder(resp.Body).Decode(&patients); err != nil {
		return nil, err
	}

	return patients, nil
}