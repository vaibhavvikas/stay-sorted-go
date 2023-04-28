package user_service

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/vaibhavvikas/stay-sorted/models"
	"github.com/vaibhavvikas/stay-sorted/models/entities"
	"github.com/vaibhavvikas/stay-sorted/repository/user_repository"
	"github.com/vaibhavvikas/stay-sorted/utils"
)

type UserServiceImpl struct {
	UserRepository user_repository.UserRepositoryImpl
}

func (u UserServiceImpl) CreateUser(ctx *gin.Context, userReq models.CreateUserReq) (entities.User, error) {
	var user entities.User

	salt, err := utils.GenerateSalt()
	if err != nil {
		return user, errors.Wrap(err, "[CreateUser]")
	}

	password := utils.GenerateHashedPassword(salt, userReq.Password)

	user = entities.User{
		Name:     userReq.Name,
		Email:    userReq.Email,
		Gender:   userReq.Gender,
		Dob:      userReq.Dob,
		Phone:    userReq.Phone,
		Rating:   0,
		IsActive: true,
		Salt:     hex.EncodeToString(salt),
		Password: hex.EncodeToString(password),
	}

	user, err = u.UserRepository.CreateUser(ctx, user)
	if err != nil {
		return user, errors.Wrap(err, "[CreateUser]")
	}

	return user, errors.Wrap(err, "[CreateUser]")
}

func (u UserServiceImpl) GetUserByPid(ctx *gin.Context, userPid string) (entities.User, error) {
	user, err := u.UserRepository.GetUserByPid(ctx, userPid)
	if err != nil {
		return user, errors.Wrap(err, "[GetUserByPid]")
	}
	return user, errors.Wrap(err, "[GetUserByPid]")
}

func (u UserServiceImpl) AuthenticateUser(ctx *gin.Context, userLoginReq models.UserLoginReq) (string, error) {
	salt, err := hex.DecodeString(userLoginReq.User.Salt)
	if err != nil {
		return "", errors.Wrap(err, "[AuthenticateUser]")
	}

	saltedPassword := append(salt, []byte(userLoginReq.Password)...)

	// Hash the salted password using SHA-256
	enteredPasswordHash := sha256.Sum256(saltedPassword)

	// Compare the entered password hash with the stored hash value
	storedPasswordHash, err := hex.DecodeString(userLoginReq.User.Password)
	if err != nil {
		return "", errors.Wrap(err, "[AuthenticateUser]")
	}

	if hex.EncodeToString(enteredPasswordHash[:]) != hex.EncodeToString(storedPasswordHash) {
		return "", errors.Wrap(getInvalidPasswordError(), "[AuthenticateUser]")
	}

	token, err := utils.GenerateJwtToken(userLoginReq.User.Pid)
	if err != nil {
		return "", errors.Wrap(err, "[AuthenticateUser]")
	}

	return token, errors.Wrap(err, "[AuthenticateUser]")
}

func getInvalidPasswordError() error {
	return &models.AppError{
		Message: "Invalid Password",
	}
}
