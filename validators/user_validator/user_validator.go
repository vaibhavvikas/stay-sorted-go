package user_validator

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/thedevsaddam/govalidator"
	"github.com/vaibhavvikas/stay-sorted/models"
	"github.com/vaibhavvikas/stay-sorted/providers"
	"github.com/vaibhavvikas/stay-sorted/validators"
	"gorm.io/gorm"
)

func ValidateCreateUserRequest(ctx *gin.Context) (models.CreateUserReq, error) {
	var createUserReq models.CreateUserReq

	err := ctx.ShouldBindJSON(&createUserReq)
	if err != nil {
		return createUserReq, errors.Wrap(err, "[ValidateUserRequest]")
	}

	opts := govalidator.Options{
		Data:  &createUserReq,
		Rules: createUserReqRules(),
	}

	errs := govalidator.New(opts).ValidateStruct()
	if len(errs) > 0 {
		return createUserReq, validators.InvalidRequestValidatorError(errs)
	}

	return createUserReq, errors.Wrap(err, "[ValidateUserRequest]")
}

func createUserReqRules() govalidator.MapData {
	return govalidator.MapData{
		"name":     []string{"required"},
		"email":    []string{"required", "email"},
		"gender":   []string{"required", "in:male,female,other"},
		"dob":      []string{"required", "date:yyyy-mm-dd"},
		"phone":    []string{"required", "digits:10"},
		"password": []string{"required", "min:6"},
	}
}

func ValidateGetUserByPidRequest(ctx *gin.Context) (string, error) {
	userPid := ctx.Param("user_pid")
	if userPid == "" {
		return "", validators.ParamMissingOrInvalidError("user_pid")
	}
	return userPid, nil
}

func ValidateLoginUserRequest(ctx *gin.Context) (models.UserLoginReq, error) {
	var userLoginReq models.UserLoginReq

	err := ctx.ShouldBindJSON(&userLoginReq)
	if err != nil {
		return userLoginReq, errors.Wrap(err, "[ValidateLoginUserRequest]")
	}

	rules := govalidator.MapData{
		"email":    []string{"required", "email"},
		"password": []string{"required"},
	}
	opts := govalidator.Options{
		Data:  &userLoginReq,
		Rules: rules,
	}

	errs := govalidator.New(opts).ValidateStruct()
	if len(errs) > 0 {
		return userLoginReq, validators.InvalidRequestValidatorError(errs)
	}

	user, err := providers.UserRepository.GetUserByEmail(ctx, userLoginReq.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userLoginReq, validators.InvalidRequestError("email", "User with following email does not exist in the system.")
		}

		return userLoginReq, errors.Wrap(err, "[ValidateLoginUserRequest]")
	}
	userLoginReq.User = user
	return userLoginReq, nil
}
