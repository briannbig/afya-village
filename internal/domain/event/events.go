package event

import "github.com/briannbig/afya-village/internal/domain/model"

type (
	UserRegisteredEvent struct {
		User model.User
	}
	
	PatientRegisteredEvent struct {
		Patient model.Patient
	}

	PatientUpdatedEvent struct {
		Patient model.Patient
	}

	MedicalProfessionalRegisteredEvent struct {
		MedicalProfessional model.Medic
	}

	MedicalProfessionalUpdatedEvent struct {
		MedicalProfessional model.Medic
	}

	MobileClinicCreatedEvent struct {
		MobileClinic model.Clinic
	}

	MobileClinicUpdatedEvent struct {
		MobileClinic model.Clinic
	}

	AppointmentScheduledEvent struct {
		Appointment model.Appointment
	}

	AppointmentCompletedEvent struct {
		Appointment model.Appointment
	}

	AppointmentCanceledEvent struct {
		Appointment model.Appointment
	}

	MedicalRecordCreatedEvent struct {
		MedicalRecord model.MedicalRecord
	}

	MedicalRecordUpdatedEvent struct {
		MedicalRecord model.MedicalRecord
	}
)
