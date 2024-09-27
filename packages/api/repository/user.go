package repository

import (
	"context"
	"log"
	"time"

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

	newUser := models.User{
		ID:        id,
		Email:     email,
		Username:  username,
		Password:  password,
		Role:      "user",
		IsActive:  true,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	_, err := r.user.InsertOne(context.TODO(), newUser)

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

	updatedUser := models.User{}

	if user.Username != "" {
		updatedUser.Username = user.Username
	}

	if user.Email != "" {
		updatedUser.Email = user.Email
	}

	if user.Password != "" {
		updatedUser.Password = user.Password
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
