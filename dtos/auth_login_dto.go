package dtos

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthLoginRequestDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type AuthLoginResponseDTO struct {
	ID        primitive.ObjectID `json:"id"`
	Name      string             `json:"name,omitempty"`
	LastName  string             `json:"last_name,omitempty"`
	BirthDate time.Time          `json:"birth_date,omitempty"`
	Email     string             `json:"email"`
	Avatar    string             `json:"avatar,omitempty"`
	Banner    string             `json:"banner,omitempty"`
	Biography string             `json:"biography,omitempty"`
	Location  string             `json:"location,omitempty"`
	WebSite   string             `json:"web_site,omitempty"`
}
