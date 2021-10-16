package usecases

import "github.com/mkaiho/gin-sample/domain/models"

type UserRepository interface {
	FindById(id string) *models.User
}
