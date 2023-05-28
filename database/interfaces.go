package database

import "github.com/luxarasis4/twittor/database/entities"

type IUserRepository interface {
	Save(entity entities.UserEntity) (string, bool, error)
	FindByEmail(email string) (entities.UserEntity, bool, error)
}
