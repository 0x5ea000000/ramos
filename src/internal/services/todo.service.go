package services

import (
	"0x5ea000000/ramos/internal/entities"
	"0x5ea000000/ramos/internal/repositories"
)

type TodoService interface {
	GetAll() ([]*entities.Todo, error)
	Get(id string) (*entities.Todo, error)
	Create(todo *entities.Todo) (*entities.Todo, error)
	Update(todo *entities.Todo) (*entities.Todo, error)
	Delete(id string) error
}

type todoService struct {
	repo repositories.TodoRepository
}

func NewTodoService(repo repositories.TodoRepository) TodoService {
	return &todoService{repo: repo}
}

func (s *todoService) GetAll() ([]*entities.Todo, error) {
	return s.repo.GetAll()
}

func (s *todoService) Get(id string) (*entities.Todo, error) {
	return s.repo.Get(id)
}

func (s *todoService) Create(todo *entities.Todo) (*entities.Todo, error) {
	return s.repo.Create(todo)
}

func (s *todoService) Update(todo *entities.Todo) (*entities.Todo, error) {
	return s.repo.Update(todo)
}

func (s *todoService) Delete(id string) error {
	return s.repo.Delete(id)
}
