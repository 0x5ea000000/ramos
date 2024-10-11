package repositories

type Repository[T any] interface {
	Create(todo *T) (*T, error)
	GetAll() ([]*T, error)
	Get(id string) (*T, error)
	Update(todo *T) (*T, error)
	Delete(id string) error
}
