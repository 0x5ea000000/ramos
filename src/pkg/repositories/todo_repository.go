package repositories

import "0x5ea000000/ramos/pkg/models"

type TodoRepository interface {
	Repository[models.Todo]
}
