package providers

import (
	"github.com/vaibhavvikas/stay-sorted/repository/user_repository"
)

// Repositories
var UserRepository = user_repository.UserRepositoryImpl{}
