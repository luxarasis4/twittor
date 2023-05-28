package repositories

import (
	"context"
	"sync"
	"time"

	"github.com/luxarasis4/twittor/database"
	"github.com/luxarasis4/twittor/database/entities"
	"github.com/luxarasis4/twittor/util/crypto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	databaseTwittor = "twittor"

	collectionUser = "user"

	timeout = 15 * time.Second
)

var (
	userRepositoryOnce      sync.Once
	userRepositoryInterface *UserRepository
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository() *UserRepository {
	userRepositoryOnce.Do(func() {
		userRepositoryInterface = &UserRepository{
			collection: database.MongoCN.Database(databaseTwittor).Collection(collectionUser),
		}
	})

	return userRepositoryInterface
}

func (repository *UserRepository) Save(entity entities.UserEntity) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	entity.Password, _ = crypto.EncryptPasswordWithBcrypt(entity.Password)

	result, err := repository.collection.InsertOne(ctx, entity)

	if err != nil {
		return "", false, err
	}

	id, _ := result.InsertedID.(primitive.ObjectID)

	return id.String(), true, nil
}

func (repository *UserRepository) FindByEmail(email string) (entities.UserEntity, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	query := bson.M{"email": email}

	var user entities.UserEntity

	err := repository.collection.FindOne(ctx, query).Decode(&user)

	return user, err == nil, err
}
