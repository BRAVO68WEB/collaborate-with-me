package repository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/BRAVO68WEB/collaborate-with-me/packages/api/helpers"
	"github.com/BRAVO68WEB/collaborate-with-me/packages/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UpdateUser struct {
	Username string
	Email    string
	Password string
	Role     string
}

type UserRepository interface {
	CreateUser(
		email string,
		password string,
		username string,
	) (models.User, error)

	UpdateUserByID(
		ID string,
		user UpdateUser,
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

type userRepository struct {
	user *mongo.Collection
}

func NewUserRepository(
	col *mongo.Collection,
) UserRepository {
	return &userRepository{
		user: col,
	}
}

func (r *userRepository) CreateUser(
	email string,
	password string,
	username string,
) (models.User, error) {
	id := primitive.NewObjectID()

	passwordHash, err := helpers.HashPassword(password)

	if err != nil {
		return models.User{}, err
	}

	newUser := models.User{
		ID:        id,
		Email:     email,
		Username:  username,
		Password:  passwordHash,
		Role:      "user",
		IsActive:  true,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	_, err = r.user.InsertOne(context.TODO(), newUser)

	if err != nil {
		log.Fatal(err)
	}

	return newUser, nil
}

func (r *userRepository) UpdateUserByID(
	ID string,
	user UpdateUser,
) (models.User, error) {
	_ID, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return models.User{}, err
	}

	user_to_update := r.user.FindOne(context.TODO(), bson.M{
		"_id": _ID,
	})

	var user_data models.User

	err = user_to_update.Decode(&user_data)

	if err != nil {
		return models.User{}, err
	}

	updatedUser := models.User{}

	if user.Username != "" {
		updatedUser.Username = user.Username
	}

	if user.Email != "" {
		updatedUser.Email = user.Email
	}

	if user.Password != "" {
		passwordHash, err := helpers.HashPassword(user.Password)

		if err != nil {
			return models.User{}, err
		}

		updatedUser.Password = passwordHash
	}

	if user.Role != "" {
		updatedUser.Role = user.Role
	}

	updatedUser.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	_, err = r.user.UpdateOne(context.TODO(), bson.M{
		"_id": _ID,
	}, bson.M{
		"$set": updatedUser,
	})

	if err != nil {
		return models.User{}, err
	}

	return updatedUser, nil
}

func (r *userRepository) DisableUserByID(
	ID string,
) (models.User, error) {
	_ID, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return models.User{}, err
	}

	user_to_disable := r.user.FindOne(context.TODO(), bson.M{
		"_id": _ID,
	})

	var user_data models.User

	err = user_to_disable.Decode(&user_data)

	if err != nil {
		return models.User{}, err
	}

	disabledUser := models.User{
		IsActive: false,
	}

	_, err = r.user.UpdateOne(context.TODO(), bson.M{
		"_id": _ID,
	}, bson.M{
		"$set": disabledUser,
	})

	if err != nil {
		return models.User{}, err
	}

	return disabledUser, nil
}

func (r *userRepository) GetUsers(
	page int,
) ([]models.User, error) {
	users := []models.User{}

	cursor, err := r.user.Find(context.TODO(), bson.M{})
	if err != nil {
		return users, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user models.User

		err := cursor.Decode(&user)

		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) GetUserByEmail(
	email string,
) (models.User, error) {
	result := r.user.FindOne(context.TODO(), bson.M{
		"email": email,
	})

	var user models.User

	err := result.Decode(&user)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *userRepository) GetUserByID(
	ID string,
) (models.User, error) {
	_ID, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return models.User{}, err
	}

	result := r.user.FindOne(context.TODO(), bson.M{
		"_id": _ID,
	})

	var user models.User

	err = result.Decode(&user)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *userRepository) GetUserByUserName(
	username string,
) (models.User, error) {
	result := r.user.FindOne(context.TODO(), bson.M{
		"username": username,
	})

	var user models.User

	err := result.Decode(&user)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *userRepository) Login(
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

func (r *userRepository) CheckIfUserIsActiveByID(
	ID string,
) bool {
	_ID, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return false
	}

	result := r.user.FindOne(context.TODO(), bson.M{
		"_id": _ID,
	})

	var user models.User

	err = result.Decode(&user)

	if err != nil {
		return false
	}

	return user.IsActive
}

func (r *userRepository) CheckIfUserIsActiveByEmail(
	email string,
) bool {
	result := r.user.FindOne(context.TODO(), bson.M{
		"email": email,
	})

	var user models.User

	err := result.Decode(&user)

	if err != nil {
		return false
	}

	return user.IsActive
}

func (r *userRepository) CheckIfUserIsAdmin(
	ID string,
) bool {
	_ID, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return false
	}

	result := r.user.FindOne(context.TODO(), bson.M{
		"_id": _ID,
	})

	var user models.User

	err = result.Decode(&user)

	if err != nil {
		return false
	}

	return user.Role == "admin"
}
