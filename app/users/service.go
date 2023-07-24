package users

import (
	"context"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	errorclass "interview/app/error"
	"interview/app/models"
	"interview/app/structs"
	"io"
)

var service userservice

type userservice struct {
	core IUserCore
}

func InitializeUserService(core IUserCore) IUserService {
	service = userservice{
		core: core,
	}
	return &service
}

func (u userservice) Signup(ctx context.Context, data *structs.UserSingupRequest) (*structs.UserSingupRequestResponse, *errorclass.Error) {
	// Check If User Exist
	filter := map[string]interface{}{
		"email": data.Email,
	}

	existinguser, err := u.core.GetUser(ctx, filter)
	if err != nil {
		return nil, errorclass.NewError(errorclass.InternalServerError).Wrap("Error in fetching data")
	}

	if existinguser != nil {
		return nil, errorclass.NewError(errorclass.RecordAlreadyExist).Wrap("user with this email already exist")
	}

	// If user does not exist create new user

	salt := generateRandomSalt(10)
	hash := hashPassword(data.Password, salt)

	user := &models.User{
		Email:         data.Email,
		Hash:          hash,
		Salt:          salt,
		ActiveAccount: true,
		CountryCode:   data.CountryCode,
		Type:          data.Type,
	}

	response, err := u.core.CreateUser(ctx, user)
	if err != nil {
		return nil, errorclass.NewError(errorclass.InternalServerError).Wrap(err.Error())
	}

	return &structs.UserSingupRequestResponse{
		Id: response.ID,
	}, nil
}

func (u userservice) Login(ctx context.Context, data *structs.UserLoginRequest) (*structs.UserLoginRequestResponse, *errorclass.Error) {
	filter := map[string]interface{}{
		"email": data.Email,
	}

	user, err := u.core.GetUser(ctx, filter)
	if err != nil {
		return nil, errorclass.NewError(errorclass.InternalServerError).Wrap("Error in fetching data")
	}

	if user == nil {
		return nil, errorclass.NewError(errorclass.RecordNotFound).Wrap("user with this email does not exist")
	}

	isValidPassword := isValidPassword(user.Hash, data.Password, user.Salt)
	if !isValidPassword {
		return nil, errorclass.NewError(errorclass.BadRequestValidationError).Wrap("credentials does not match")
	}

	return &structs.UserLoginRequestResponse{
		Id:                  user.ID,
		Name:                user.Name,
		Email:               user.Email,
		PhoneNumber:         user.PhoneNumber,
		EmailVerified:       user.EmailVerified,
		PhoneNumberVerified: user.PhoneNumberVerified,
		Type:                user.Type,
		Admin:               user.Admin,
		CountryCode:         user.CountryCode,
		ActiveAccount:       user.ActiveAccount,
		CreatedAt:           user.CreatedAt,
		UpdatedAt:           user.UpdatedAt,
	}, nil
}

func isValidPassword(hash string, password string, salt string) bool {
	newHash := hashPassword(password, salt)
	if newHash == hash {
		return true
	}
	return false
}

func hashPassword(password string, salt string) string {
	var passwordBytes = []byte(password)

	var sha512Hasher = sha512.New()

	passwordBytes = append(passwordBytes, salt...)

	sha512Hasher.Write(passwordBytes)

	var hashedPasswordBytes = sha512Hasher.Sum(nil)

	var hashedPasswordHex = hex.EncodeToString(hashedPasswordBytes)

	return hashedPasswordHex
}

func generateRandomSalt(saltSize int) string {
	var salt = make([]byte, saltSize)

	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(salt)
}
