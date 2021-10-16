package usecases_test

import (
	"testing"

	"github.com/mkaiho/gin-sample/domain/models"
	"github.com/mkaiho/gin-sample/usecases"
)

const (
	dummyUserId   = "8d73e15f-be3d-454a-925f-3a7a96acfdb3"
	dummyUserName = "John Smith"
)

type DummyUserRepository struct{}

func (ur DummyUserRepository) FindById(id string) *models.User {
	if id != dummyUserId {
		return nil
	}

	return &models.User{
		ID:   dummyUserId,
		Name: dummyUserName,
	}
}

func TestFindUser(t *testing.T) {
	t.Run("Non null value is Returned when there is user which has id equivalent with passed.", func(t *testing.T) {
		expectedId := dummyUserId
		param := usecases.FindUserRequest{ID: expectedId}
		sut := usecases.UserInteractor{
			UserRepository: DummyUserRepository{},
		}

		actual := sut.FindUser(param)

		if actual == nil {
			t.Errorf("Returned null value; want non null value.")
		}
	})
}
