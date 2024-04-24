package usecase

import (
	"errors"
	"time"

	"github.com/briannbig/afya-village/internal/domain/event"
	"github.com/briannbig/afya-village/internal/domain/model"
	"github.com/briannbig/afya-village/internal/domain/queue"
	"github.com/briannbig/afya-village/internal/domain/repository"
)

type ScheduleAppointmentUseCase struct {
	patientRepository      repository.PatientRepository
	mobileClinicRepository repository.MobileClinicRepository
	appointmentRepository  repository.AppointmentRepository
	producer               queue.Producer
}

func NewAppointmentScheduler(
	patientRepository repository.PatientRepository,
	mobileClinicRepository repository.MobileClinicRepository,
	appointmentRepository repository.AppointmentRepository,
	producer queue.Producer) *ScheduleAppointmentUseCase {
	return &ScheduleAppointmentUseCase{
		patientRepository:      patientRepository,
		mobileClinicRepository: mobileClinicRepository,
	}
}

func (s *ScheduleAppointmentUseCase) ScheduleAppointment(patientID, mobileClinicID, appointmentDateTime string) (*model.Appointment, error) {
	// Retrieve patient and mobile clinic information
	patient, err := s.patientRepository.FindById(patientID)
	if err != nil {
		return nil, err
	}

	mobileClinic, err := s.mobileClinicRepository.FindById(mobileClinicID)
	if err != nil {
		return nil, err
	}

	// Check if the patient and mobile clinic are valid
	if patient == nil || mobileClinic == nil {
		return nil, errors.New("invalid patient or mobile clinic")
	}

	// Check if the appointment date and time are valid
	appointmentDate, err := time.Parse(time.RFC3339, appointmentDateTime)
	if err != nil {
		return nil, err
	}

	// todo:  Check if the mobile clinic is available at the specified date and time depending on the specific business rules for mobile clinic availability

	var appointment model.Appointment = model.Appointment{
		Patient:           *patient,
		MobileClinic:      *mobileClinic,
		ScheduledDateTime: appointmentDate,
		Status:            string(model.StatusScheduled),
	}

	if err := s.appointmentRepository.Save(&appointment); err != nil {
		return nil, err
	}

	if err := s.producer.Publish(
		event.AppointmentScheduledEvent{Appointment: appointment},
	); err != nil {
		return nil, err
	}

	return &appointment, nil
}
