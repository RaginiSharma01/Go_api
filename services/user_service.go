package services

import (
	"api/models"
	"api/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		Repo: repo,
	}
}
func (s *UserService) CreateUser(user models.User) error {
	return s.Repo.CreateUser(user)
}
func (s *UserService) GetUsers() ([]models.User, error) {

	return s.Repo.GetUsers()
}
func (s *UserService) GetUserByID(id int) (models.User, error) {
	return s.Repo.GetUserByID(id)
}
func (s *UserService) DeleteUser(id int) error {
	return s.Repo.DeleteUser(id)
}
