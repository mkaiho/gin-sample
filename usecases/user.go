package usecases

import (
	"github.com/mkaiho/gin-sample/domain/models"
	"github.com/mkaiho/gin-sample/interfaces/gateways"
)

type UserInteractor struct {
	UserRepository gateways.UserRepository
}

func (ui *UserInteractor) FindUser(id string) *models.User {
	return ui.UserRepository.FindById(id)
}
