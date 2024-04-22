package usecase

import (
	"time"

	"github.com/briannbig/afya-village/internal/domain/event"
	"github.com/briannbig/afya-village/internal/domain/model"
	"github.com/briannbig/afya-village/internal/domain/queue"
	"github.com/briannbig/afya-village/internal/domain/repository"
)

type PatientRegistrationUseCase struct {
	repo     repository.PatientRepository
	producer queue.Producer
}

func NewPatientRegistrationUseCase(repo repository.PatientRepository, producer queue.Producer) *PatientRegistrationUseCase {
	return &PatientRegistrationUseCase{repo: repo, producer: producer}
}

func (u *PatientRegistrationUseCase) RegisterNewPatient(name string, dateOfBirth time.Time, gender, location string) (*model.Patient, error) {
	patient := model.NewPatient(name, dateOfBirth, gender, location)
	if err := u.repo.Save(&patient); err != nil {
		return nil, err
	}
	err := u.producer.Produce(event.PatientRegisteredEvent{Patient: patient})
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

func (u *PatientRegistrationUseCase) RegisterExistingPatient(identifier string) (*model.Patient, error) {
	patient, err := u.repo.FindByIdentifier(identifier)
	if err != nil {
		return nil, err
	}

	err = u.producer.Produce(event.PatientRegisteredEvent{Patient: *patient})
	if err != nil {
		return nil, err
	}
	return patient, nil
}
