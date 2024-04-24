package repository

import "github.com/briannbig/afya-village/internal/domain/model"

type (
	UserRepository interface {
		Save(model *model.User) error
		FindById(id string) (*model.User, error)
		FindByIdentifier(identifier string) (*model.User, error)
	}
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
	MedicalRecordRepository interface {
		Save(model *model.MedicalRecord) error
		FindById(id string) (*model.MedicalRecord, error)
		FindByIdentifier(identifier string) (*model.MedicalRecord, error)
	}
	MedicRepository interface {
		Save(model *model.Medic) error
		FindById(id string) (*model.Medic, error)
		FindByIdentifier(identifier string) (*model.Medic, error)
	}
)
