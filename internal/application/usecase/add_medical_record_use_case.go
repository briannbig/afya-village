package usecase

import (
	"errors"
	"time"

	"github.com/briannbig/afya-village/internal/domain/event"
	"github.com/briannbig/afya-village/internal/domain/model"
	"github.com/briannbig/afya-village/internal/domain/queue"
	"github.com/briannbig/afya-village/internal/domain/repository"
)

type AddMedicalRecordUseCase struct {
	patientRepository       repository.PatientRepository
	medicalRecordRepository repository.MedicalRecordRepository
	medicRepository         repository.MedicRepository
	producer                queue.Producer
}

func NewAddMedicalRecordUseCase(
	patientRepository repository.PatientRepository,
	medicalRecordRepository repository.MedicalRecordRepository,
	medicRepositorty repository.MedicRepository,
	producer queue.Producer) *AddMedicalRecordUseCase {
	return &AddMedicalRecordUseCase{
		patientRepository:       patientRepository,
		medicalRecordRepository: medicalRecordRepository,
	}
}

func (r *AddMedicalRecordUseCase) AddMedicalRecord(
	patientID, dateTime, diagnosis, treatment, medicId string, prescriptionIds []string) (*model.MedicalRecord, error) {

	recordDate, err := time.Parse(time.RFC3339, dateTime)
	if err != nil {
		return nil, err
	}
	// Retrieve patient and medical professional information
	patient, err := r.patientRepository.FindById(patientID)
	if err != nil {
		return nil, err
	}

	medic, err := r.medicRepository.FindById(medicId)
	if err != nil {
		return nil, err
	}

	if patient == nil || medic == nil {
		return nil, errors.New("invalid patient or medical professional")
	}

	medicalRecord := model.MedicalRecord{
		Patient:   *patient,
		Date:      recordDate,
		Diagnosis: diagnosis,
		Treatment: treatment,
		// todo: fetch prescriptions from db and add them
		MedicalProfessional: *medic,
	}

	if err := r.medicalRecordRepository.Save(&medicalRecord); err != nil {
		return nil, err
	}

	if err := r.producer.Produce(
		event.MedicalRecordCreatedEvent{
			MedicalRecord: medicalRecord},
	); err != nil {
		return nil, err
	}

	return &medicalRecord, nil
}
