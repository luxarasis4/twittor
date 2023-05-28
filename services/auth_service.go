package services

import (
	"sync"

	"github.com/luxarasis4/twittor/database"
	"github.com/luxarasis4/twittor/dtos"
)

var (
	authServiceOnce      sync.Once
	authServiceInterface *AuthService
)

type AuthService struct {
	userRepository database.IUserRepository
}

func NewUserService(userRepository database.IUserRepository) *AuthService {
	authServiceOnce.Do(func() {
		authServiceInterface = &AuthService{
			userRepository: userRepository,
		}
	})

	return authServiceInterface
}

func (service *AuthService) Registry(request dtos.AuthRegistryRequestDTO) (*dtos.AuthRegistryResponseDTO, bool, error) {
	user, found, _ := service.userRepository.FindByEmail(request.Email)

	response := &dtos.AuthRegistryResponseDTO{}

	if found {
		response.Id = user.ID.String()
		return response, false, nil
	}

	id, status, err := service.userRepository.Save(*request.ToUserEntity())

	if err != nil {
		return nil, false, err
	}

	response.Id = id

	return response, status, err
}
