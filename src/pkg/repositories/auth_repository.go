package repositories

import "0x5ea000000/ramos/pkg/models"

type AuthRepository interface {
	Repository[models.User]
}
