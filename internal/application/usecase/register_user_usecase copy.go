package usecase

import (
	"github.com/briannbig/afya-village/internal/application/dto"
	"github.com/briannbig/afya-village/internal/domain/event"
	"github.com/briannbig/afya-village/internal/domain/model"
	"github.com/briannbig/afya-village/internal/domain/queue"
	"github.com/briannbig/afya-village/internal/domain/repository"
)

type UserRegistrationUseCase struct {
	repo     repository.UserRepository
	producer queue.Producer
}

func NewUserRegistrationUseCase(repo repository.UserRepository, producer queue.Producer) *UserRegistrationUseCase {
	return &UserRegistrationUseCase{repo: repo, producer: producer}
}

func (u *UserRegistrationUseCase) RegisterNewUser(req dto.RequestCreateUser) (*model.User, error) {
	user := model.NewUser(req.Username, req.Email, req.Telephone, req.Password)
	if err := u.repo.Save(&user); err != nil {
		return nil, err
	}
	err := u.producer.Publish(event.UserRegisteredEvent{User: user})
	if err != nil {
		return nil, err
	}
	return &user, nil
}
