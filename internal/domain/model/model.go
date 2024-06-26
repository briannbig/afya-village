package model

import (
	"time"

	"github.com/rs/xid"
)

type AppointmentStatus string

const (
	StatusScheduled AppointmentStatus = "Scheduled"
	StatusCompleted AppointmentStatus = "Completed"
	StatusCancelled AppointmentStatus = "Canceled"
)

type (
	BaseModel struct {
		Id        string     `json:"id"`
		CreatedAt time.Time  `json:"createdAt" db:"created_at"`
		UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
		DeletedAt *time.Time `json:"deletedAt,omitempty"`
	}

	User struct {
		BaseModel
		Username  string `json:"username"`
		Email     string `json:"email"`
		Telephone string `json:"telephone"`
		Password  string `json:"-"`
	}

	Patient struct {
		User
		Name           string          `json:"name"`
		DateOfBirth    time.Time       `json:"dateOfBirth"`
		Gender         string          `json:"gender"`
		Location       string          `json:"location"`
		MedicalHistory []MedicalRecord `json:"medicalHistory"`
	}

	Medic struct {
		User
		Name           string   `json:"name"`
		Specialization string   `json:"specialization"`
		Qualifications []string `json:"qualifications"`
		Experience     int      `json:"experience"` // in years
	}

	Clinic struct {
		BaseModel
		Location          string        `json:"location"`
		AvailableServices []string      `json:"availableServices"`
		MedicalStaff      []Medic       `json:"medicalStaff"`
		Appointments      []Appointment `json:"appointments"`
	}

	Appointment struct {
		BaseModel
		Patient           Patient   `json:"patient"`
		MobileClinic      Clinic    `json:"mobileClinic"`
		ScheduledDateTime time.Time `json:"scheduledDateTime"`
		Status            string    `json:"status"` // AppointmentStatus
	}

	MedicalRecord struct {
		BaseModel
		Patient             Patient      `json:"patient"`
		Date                time.Time    `json:"date"`
		Diagnosis           string       `json:"diagnosis"`
		Treatment           string       `json:"treatment"`
		Prescription        []Medication `json:"prescription"`
		MedicalProfessional Medic        `json:"medicalProfessional"`
	}

	Medication struct {
		BaseModel
		Name         string `json:"name"`
		Dosage       string `json:"dosage"`
		Instructions string `json:"instructions"`
	}
)

func newBaseModel() BaseModel {
	now := time.Now()
	return BaseModel{
		Id:        xid.New().String(),
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: nil,
	}
}

func NewUser(username, email, telephone, password string) User {
	return User{
		BaseModel: newBaseModel(),
		Username:  username,
		Email:     email,
		Telephone: telephone,
		Password:  password,
	}
}

func NewPatient(username, email, telephone, name, gender, location, password string, dateOfBirth time.Time) Patient {
	user := NewUser(username, email, telephone, password)
	return Patient{
		User:           user,
		Name:           name,
		DateOfBirth:    dateOfBirth,
		Gender:         gender,
		Location:       location,
		MedicalHistory: nil,
	}
}

func NewMedicalProfessional(username, email, telephone, name, specialization, password string, qualifications []string, experience int) Medic {
	user := NewUser(username, email, telephone, password)
	return Medic{
		User:           user,
		Name:           name,
		Specialization: specialization,
		Qualifications: qualifications,
		Experience:     experience,
	}
}

func NewMobileClinic(location string, availableServices []string) Clinic {
	baseModel := newBaseModel()
	return Clinic{
		BaseModel:         baseModel,
		Location:          location,
		AvailableServices: availableServices,
		MedicalStaff:      nil,
		Appointments:      nil,
	}
}

func NewAppointment(patient Patient, mobileClinic Clinic, scheduledDateTime time.Time) Appointment {
	baseModel := newBaseModel()
	return Appointment{
		BaseModel:         baseModel,
		Patient:           patient,
		MobileClinic:      mobileClinic,
		ScheduledDateTime: scheduledDateTime,
		Status:            "Scheduled",
	}
}

func NewMedicalRecord(patient Patient, date time.Time, diagnosis, treatment string, prescription []Medication, medicalProfessional Medic) MedicalRecord {
	baseModel := newBaseModel()
	return MedicalRecord{
		BaseModel:           baseModel,
		Patient:             patient,
		Date:                date,
		Diagnosis:           diagnosis,
		Treatment:           treatment,
		Prescription:        prescription,
		MedicalProfessional: medicalProfessional,
	}
}

func NewMedication(name, dosage, instructions string) Medication {
	return Medication{
		Name:         name,
		Dosage:       dosage,
		Instructions: instructions,
	}
}
