package repository

import "github.com/briannbig/afya-village/internal/domain/model"

type (
	PatientRepository interface {
		Save(model *model.Patient) error
		FindById(id string) (*model.Patient, error)
		FindByIdentifier(identifier string) (*model.Patient, error)
	}
)
