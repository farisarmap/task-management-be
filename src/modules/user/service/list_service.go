package services

import (
	"context"
	"task-management-be/src/domain"
	"task-management-be/src/helper"
)

func (us *UserService) List(ctx context.Context) ([]domain.CreateUserResponse, error) {
	resp, err := us.Repository.List(ctx)

	if err != nil {
		return []domain.CreateUserResponse{}, helper.NewAppError(err.Error())
	}

	response := make([]domain.CreateUserResponse, len(resp))
	for i := 0; i < len(resp); i++ {
		response[i].Email = resp[i].Email
		response[i].Name = resp[i].Name
		response[i].Created_At = resp[i].Created_At
		response[i].Updated_At = resp[i].Updated_At
	}

	return response, nil
}
