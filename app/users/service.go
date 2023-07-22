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
	}

	response, err := u.core.CreateUser(ctx, user)
	if err != nil {
		return nil, errorclass.NewError(errorclass.InternalServerError).Wrap(err.Error())
	}

	return &structs.UserSingupRequestResponse{
		Id: response.ID,
	}, nil
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
