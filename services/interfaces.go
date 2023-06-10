package services

import "github.com/luxarasis4/twittor/dtos"

type IAuthService interface {
	Registry(request dtos.AuthRegistryRequestDTO) (*dtos.AuthRegistryResponseDTO, bool, error)
	Login(request dtos.AuthLoginRequestDTO) (*dtos.AuthLoginResponseDTO, bool)
}
