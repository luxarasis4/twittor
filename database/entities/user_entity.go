package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserEntity struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	LastName  string             `bson:"last_name"`
	BirthDate time.Time          `bson:"birth_date"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	Avatar    string             `bson:"avatar"`
	Banner    string             `bson:"banner"`
	Biography string             `bson:"biography"`
	Location  string             `bson:"location"`
	WebSite   string             `bson:"web_site"`
}
