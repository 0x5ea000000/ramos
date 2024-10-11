package services

import (
	"0x5ea000000/ramos/pkg/models"
	"0x5ea000000/ramos/pkg/repositories"
)

type TodoService interface {
	GetAll() ([]*models.Todo, error)
	Get(id string) (*models.Todo, error)
	Create(todo *models.Todo) (*models.Todo, error)
	Update(todo *models.Todo) (*models.Todo, error)
	Delete(id string) error
}

type todoService struct {
	repo repositories.TodoRepository
}

func NewTodoService(repo repositories.TodoRepository) TodoService {
	return &todoService{repo: repo}
}

func (s *todoService) GetAll() ([]*models.Todo, error) {
	return s.repo.GetAll()
}

func (s *todoService) Get(id string) (*models.Todo, error) {
	return s.repo.Get(id)
}

func (s *todoService) Create(todo *models.Todo) (*models.Todo, error) {
	return s.repo.Create(todo)
}

func (s *todoService) Update(todo *models.Todo) (*models.Todo, error) {
	return s.repo.Update(todo)
}

func (s *todoService) Delete(id string) error {
	return s.repo.Delete(id)
}
