package repository

import "github.com/briannbig/afya-village/internal/domain/model"

type (
	PatientRepository interface {
		Save(model *model.Patient) error
		FindById(id string) (*model.Patient, error)
		FindByIdentifier(identifier string) (*model.Patient, error)
	}
	MobileClinicRepository interface {
		Save(model *model.Clinic) error
		FindById(id string) (*model.Clinic, error)
		FindByIdentifier(identifier string) (*model.Clinic, error)
	}
	AppointmentRepository interface {
		Save(model *model.Appointment) error
		FindById(id string) (*model.Appointment, error)
		FindByIdentifier(identifier string) (*model.Appointment, error)
	}
)
