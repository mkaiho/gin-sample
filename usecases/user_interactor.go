package usecases

type UserInteractor struct {
	UserRepository UserRepository
}

type FindUserRequest struct {
	ID string
}

type FindUserResponse struct {
	ID   string
	Name string
}

func (ui *UserInteractor) FindUser(req FindUserRequest) *FindUserResponse {
	user := ui.UserRepository.FindById(req.ID)

	if user == nil {
		return nil
	}

	response := &FindUserResponse{}
	response.ID = user.ID
	response.Name = user.Name

	return response
}
