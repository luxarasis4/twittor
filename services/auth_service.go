package services

import (
	"sync"

	"github.com/luxarasis4/twittor/database"
	"github.com/luxarasis4/twittor/dtos"
	"github.com/luxarasis4/twittor/util/crypto"
)

var (
	authServiceOnce      sync.Once
	authServiceInterface *AuthService
)

type AuthService struct {
	userRepository database.IUserRepository
	bcryptUtil     *crypto.BcryptUtil
}

func NewUserService(userRepository database.IUserRepository) *AuthService {
	authServiceOnce.Do(func() {
		authServiceInterface = &AuthService{
			userRepository: userRepository,
			bcryptUtil:     &crypto.BcryptUtil{},
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

func (service *AuthService) Login(request dtos.AuthLoginRequestDTO) (*dtos.AuthLoginResponseDTO, bool) {
	user, found, _ := service.userRepository.FindByEmail(request.Email)
	if !found {
		return nil, false
	}

	valid, _ := service.bcryptUtil.ValidatePassword(request.Password, user.Password)
	if !valid {
		return nil, false
	}

	response := &dtos.AuthLoginResponseDTO{
		ID:        user.ID,
		Name:      user.Name,
		LastName:  user.LastName,
		BirthDate: user.BirthDate,
		Email:     user.Email,
		Avatar:    user.Avatar,
		Banner:    user.Banner,
		Biography: user.Biography,
		Location:  user.Location,
		WebSite:   user.WebSite,
	}

	return response, true
}
