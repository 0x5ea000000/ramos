package database

import (
	"gorm.io/gorm"
)

type DatabaseRepository[T any] struct {
	db *gorm.DB
}

// NewDatabaseRepository returns a new instance of DatabaseRepository
func NewDatabaseRepository[T any](db *gorm.DB) *DatabaseRepository[T] {
	return &DatabaseRepository[T]{db}
}

func (r *DatabaseRepository[T]) Create(entity *T) (*T, error) {
	err := r.db.Create(entity).Error
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (r *DatabaseRepository[T]) GetAll() ([]*T, error) {
	var entities []*T
	err := r.db.Find(&entities).Error
	return entities, err
}

func (r *DatabaseRepository[T]) Get(id string) (*T, error) {
	var entity T
	err := r.db.First(&entity, id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *DatabaseRepository[T]) Update(entity *T) (*T, error) {
	err := r.db.Save(entity).Error
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (r *DatabaseRepository[T]) Delete(id string) error {
	var entity T
	return r.db.Delete(&entity, id).Error
}
