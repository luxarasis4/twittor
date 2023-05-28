package handlers

import (
	"github.com/luxarasis4/twittor/controllers"
	"github.com/luxarasis4/twittor/database"
	"github.com/luxarasis4/twittor/database/repositories"
	"github.com/luxarasis4/twittor/services"
)

func newUserController() controllers.IAuthController {
	return controllers.NewAuthController(
		newAuthService(),
	)
}

func newAuthService() services.IAuthService {
	return services.NewUserService(
		newUserRepository(),
	)
}

func newUserRepository() database.IUserRepository {
	return repositories.NewUserRepository()
}
