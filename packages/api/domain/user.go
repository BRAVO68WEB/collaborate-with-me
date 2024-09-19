package domain

import (
	"errors"

	"github.com/BRAVO68WEB/collaborate-with-me/packages/api/db"
	"github.com/BRAVO68WEB/collaborate-with-me/packages/api/helpers"
	"github.com/BRAVO68WEB/collaborate-with-me/packages/api/models"
	"github.com/BRAVO68WEB/collaborate-with-me/packages/api/repository"
)

type UserDomain interface {
	CreateUser(
		email string,
		password string,
		username string,
	) (models.User, error)

	UpdateUserByID(
		ID string,
		user repository.UpdateUser,
	) (models.User, error)

	DisableUserByID(
		ID string,
	) (models.User, error)

	GetUsers(
		page int,
	) ([]models.User, error)

	GetUserByEmail(
		email string,
	) (models.User, error)

	GetUserByID(
		ID string,
	) (models.User, error)

	GetUserByUserName(
		username string,
	) (models.User, error)

	Login(
		email string,
		password string,
	) (string, error)

	CheckIfUserIsActiveByID(
		ID string,
	) bool

	CheckIfUserIsActiveByEmail(
		email string,
	) bool

	CheckIfUserIsAdmin(
		ID string,
	) bool
}

type userDomain struct {
	user_repo repository.UserRepository
}

func NewUserDomain(
	db db.Connection,
) UserDomain {
	return &userDomain{
		user_repo: repository.NewUserRepository(
			db.GetCollection("user"),
		),
	}
}

func (u *userDomain) CreateUser(email string, password string, username string) (models.User, error) {
	if _, err := u.user_repo.GetUserByEmail(email); err == nil {
		return models.User{}, errors.New("email already exists")
	}

	if _, err := u.user_repo.GetUserByUserName(username); err == nil {
		return models.User{}, errors.New("username already exists")
	}

	hashedPassword, err := helpers.HashPassword(password)

	if err != nil {
		return models.User{}, err
	}

	return u.user_repo.CreateUser(email, hashedPassword, username)
}

func (u *userDomain) UpdateUserByID(ID string, user repository.UpdateUser) (models.User, error) {
	if _, err := u.user_repo.GetUserByID(ID); err != nil {
		return models.User{}, errors.New("user not found")
	}

	if user.Password != "" {
		hashedPassword, err := helpers.HashPassword(user.Password)

		if err != nil {
			return models.User{}, err
		}

		user.Password = hashedPassword
	}

	return u.user_repo.UpdateUserByID(ID, user)
}

func (u *userDomain) DisableUserByID(ID string) (models.User, error) {
	if _, err := u.user_repo.GetUserByID(ID); err != nil {
		return models.User{}, errors.New("user not found")
	}

	return u.user_repo.DisableUserByID(ID)
}

func (u *userDomain) GetUsers(page int) ([]models.User, error) {
	return u.user_repo.GetUsers(page)
}

func (u *userDomain) GetUserByEmail(email string) (models.User, error) {
	return u.user_repo.GetUserByEmail(email)
}

func (u *userDomain) GetUserByID(ID string) (models.User, error) {
	return u.user_repo.GetUserByID(ID)
}

func (u *userDomain) GetUserByUserName(username string) (models.User, error) {
	return u.user_repo.GetUserByUserName(username)
}

func (r *userDomain) Login(
	email string,
	password string,
) (string, error) {
	user, err := r.GetUserByEmail(email)

	if err != nil {
		return "", errors.New("user not found")
	}

	if !helpers.ValidatePassword(password, user.Password) {
		return "", errors.New("invalid password")
	}

	token, err := helpers.CreateJWT(user.ID.Hex(), user.Username, user.Role)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *userDomain) CheckIfUserIsActiveByID(
	ID string,
) bool {
	result, err := r.user_repo.GetUserByID(ID)

	if err != nil {
		return false
	}

	return result.IsActive
}

func (r *userDomain) CheckIfUserIsActiveByEmail(
	email string,
) bool {
	result, err := r.user_repo.GetUserByEmail(email)

	if err != nil {
		return false
	}

	return result.IsActive
}

func (r *userDomain) CheckIfUserIsAdmin(
	ID string,
) bool {
	user, err := r.user_repo.GetUserByID(ID)

	if err != nil {
		return false
	}

	return user.Role == "admin"
}
