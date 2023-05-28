package dtos

import (
	"time"

	"github.com/luxarasis4/twittor/database/entities"
)

type AuthRegistryRequestDTO struct {
	Name      string    `json:"name" validate:"required"`
	LastName  string    `json:"last_name" validate:"required"`
	BirthDate time.Time `json:"birth_date" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,min=8,max=32"`
	Avatar    string    `json:"avatar:omitempty"`
	Banner    string    `json:"banner:omitempty"`
	Biography string    `json:"biography:omitempty"`
	Location  string    `json:"location:omitempty"`
	WebSite   string    `json:"web_site:omitempty"`
}

func (dto *AuthRegistryRequestDTO) ToUserEntity() *entities.UserEntity {
	return &entities.UserEntity{
		Name:      dto.Name,
		LastName:  dto.LastName,
		BirthDate: dto.BirthDate,
		Email:     dto.Email,
		Password:  dto.Password,
		Avatar:    dto.Avatar,
		Banner:    dto.Banner,
		Biography: dto.Biography,
		Location:  dto.Location,
		WebSite:   dto.WebSite,
	}
}

type AuthRegistryResponseDTO struct {
	Id string `json:"id"`
}
